package client

import (
	"crypto/ecdsa"
	"fmt"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// commonPrm groups the parameters inherited by all Client operations.
type commonPrm struct {
	keySet bool
	key    ecdsa.PrivateKey

	withSes bool
	sesTok  session.Token

	withBrr bool
	brrTok  token.BearerToken
}

// check checks if all required parameters are set and panics if not.
func (x commonPrm) check() {
	switch {
	case !x.keySet:
		panic("no private key provided")
	}
}

// writeToRequestMetaHeader writes the parameters that should be reflected
// in request to RequestMetaHeader.
func (x commonPrm) writeToRequestMetaHeader(h *v2session.RequestMetaHeader) {
	if x.withSes {
		h.SetSessionToken(x.sesTok.ToV2())
	}

	if x.withBrr {
		h.SetBearerToken(x.brrTok.ToV2())
	}
}

type requestInterface interface {
	GetMetaHeader() *v2session.RequestMetaHeader
	SetMetaHeader(*v2session.RequestMetaHeader)
	GetVerificationHeader() *v2session.RequestVerificationHeader
	SetVerificationHeader(*v2session.RequestVerificationHeader)
}

// prepareRequest forms meta header, sets body using callback and signs the request with private key.
func (x commonPrm) prepareRequest(r requestInterface, bodyCallback func(requestInterface)) error {
	// construct the request meta header
	var hMeta v2session.RequestMetaHeader

	x.writeToRequestMetaHeader(&hMeta)

	bodyCallback(r)

	r.SetMetaHeader(&hMeta)

	// arg interface is copy-pasted from signature pkg, may it is worth to share
	err := signature.SignServiceMessage(neofsecdsa.Signer(&x.key), r)
	if err != nil {
		return fmt.Errorf("could not sign the request: %w", err)
	}

	return nil
}

func verifyResponseSignature(r interface {
	GetMetaHeader() *v2session.ResponseMetaHeader
	GetVerificationHeader() *v2session.ResponseVerificationHeader
	SetVerificationHeader(*v2session.ResponseVerificationHeader)
}) error {
	err := signature.VerifyServiceMessage(r)
	if err != nil {
		return fmt.Errorf("invalid response signature: %w", err)
	}

	return err
}

// SetECDSAPrivateKey sets ecdsa.PrivateKey to sign the request(s).
//
// Required parameter.
func (x *commonPrm) SetECDSAPrivateKey(key ecdsa.PrivateKey) {
	x.keySet = true
	x.key = key
}

// SetSessionToken sets token of the session within which the operation should be carried out.
//
// Token should be signed by its owner.
func (x *commonPrm) SetSessionToken(sesTok session.Token) {
	x.withSes = true
	x.sesTok = sesTok
}

// SetBearerToken sets bearer token that will be attached to all requests within the operation.
//
// Token should be signed by its owner.
func (x *commonPrm) SetBearerToken(sesTok session.Token) {
	x.withBrr = true
	x.sesTok = sesTok
}
