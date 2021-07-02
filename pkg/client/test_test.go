package client_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	"github.com/nspcc-dev/neofs-api-go/pkg/client"
	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	oid "github.com/nspcc-dev/neofs-api-go/pkg/object/id"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	"github.com/stretchr/testify/require"
)

func TestIk(t *testing.T) {
	d, err := os.ReadFile("../../../devenv/wallets/wallet.key")
	require.NoError(t, err)

	k, err := keys.NewPrivateKeyFromBytes(d)
	require.NoError(t, err)

	key := k.PrivateKey

	var dialPrm client.DialPrm

	dialPrm.SetAddress("s01.neofs.devenv:8080")
	dialPrm.SetTimeout(5 * time.Second)

	var dialRes client.DialRes

	err = dialPrm.Dial(&dialRes)
	require.NoError(t, err)

	cl := dialRes.Client()
	defer cl.Close()

	ctx := context.Background()

	const cidTxt = "E7wHSZBTmVUV78ifUiihekEAKqaSDnnWYh1LbzXGHUmL"

	var cnrID cid.ID

	require.NoError(t, cnrID.UnmarshalText([]byte(cidTxt)))

	var owID owner.ID

	require.NoError(t, owID.UnmarshalText([]byte("NbUgTSFvPmsRxmGeWpuuGeJUoRoi6PErcM")))

	var (
		sesPrm client.CreateSessionKeyPrm
		sesRes client.CreateSessionKeyRes
	)

	sesPrm.SetECDSAPrivateKey(key)
	sesPrm.SetOwner(owID)

	err = cl.CreateSessionKey(ctx, sesPrm, &sesRes)
	require.NoError(t, err)

	var tok session.Token

	tok.SetID(sesRes.ID())
	tok.SetOwnerID(owID)
	tok.SetSessionKey(sesRes.PublicKey())

	var octx session.ObjectContext

	var a oid.Address
	a.SetContainer(cnrID)

	octx.ForPut()
	octx.SetObject(a)

	tok.ForObject(octx)

	err = tok.SignECDSA(key)
	require.NoError(t, err)

	var (
		putPrm client.PutObjectPrm
		putRes client.PutObjectRes
	)

	putPrm.SetECDSAPrivateKey(key)
	putPrm.SetSessionToken(tok)

	err = cl.PutObject(ctx, putPrm, &putRes)
	require.NoError(t, err)

	stream := putRes.Stream()

	var hdr object.HeaderWithIDAndSignature

	hdr.SetContainer(cnrID)
	hdr.SetOwner(owID)

	var (
		hdrPrm client.WriteObjectHeaderPrm
		hdrRes client.WriteObjectHeaderRes
	)

	hdrPrm.SetHeader(hdr)

	err = stream.WriteHeader(hdrPrm, &hdrRes)
	require.NoError(t, err)

	pStream := hdrRes.PayloadStream()

	f, err := os.Open("../../Makefile")
	require.NoError(t, err)

	_, err = io.Copy(&pStream, f)
	require.NoError(t, err)

	require.NoError(t, pStream.Close())

	id := pStream.ID()

	txt, err := id.MarshalText()
	require.NoError(t, err)

	fmt.Println(string(txt))
}
