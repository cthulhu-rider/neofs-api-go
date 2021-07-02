package client

import (
	"context"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	oid "github.com/nspcc-dev/neofs-api-go/pkg/object/id"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	v2refs "github.com/nspcc-dev/neofs-api-go/v2/refs"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
)

// PutObjectPrm groups the parameters of Client.PutObject operation.
type PutObjectPrm struct {
	commonPrmWithTokens
}

// PutObjectRes groups the results of Client.PutObject operation.
type PutObjectRes struct {
	stream PutObjectStream
}

// PutObjectStream is a tool to stream the object to the NeoFS network.
type PutObjectStream struct {
	called bool

	prm commonPrmWithTokens

	v2stream rpcapi.PutObjectStream

	reqInit v2object.PutObjectPartInit
	reqBody v2object.PutRequestBody
}

// WriteObjectHeaderPrm groups the parameters of PutObjectStream.WriteHeader operation.
type WriteObjectHeaderPrm struct {
	hdrSet bool
	hdr    object.HeaderWithIDAndSignature
}

// SetHeader sets object information except payload.
//
// Required parameter.
func (x *WriteObjectHeaderPrm) SetHeader(h object.HeaderWithIDAndSignature) {
	x.hdr = h
	x.hdrSet = true
}

// WriteObjectHeaderRes groups the results of PutObjectStream.WriteHeader operation.
type WriteObjectHeaderRes struct {
	payloadStream PayloadStream
}

// PayloadStream returns payload stream instance.
func (x WriteObjectHeaderRes) PayloadStream() PayloadStream {
	return x.payloadStream
}

// WriteHeader initializes stream with writing the object header. Must be called once.
// To write the payload res.PayloadStream() can be used.
//
// All required parameters must be set. Result must not be nil.
func (x *PutObjectStream) WriteHeader(prm WriteObjectHeaderPrm, res *WriteObjectHeaderRes) error {
	switch {
	case !prm.hdrSet:
		panic("header required")
	case x.called:
		panic("re-call detected")
	}

	x.called = true

	x.reqBody.SetObjectPart(&x.reqInit)

	{ // construct init part
		var (
			idv2 *v2refs.ObjectID
			hv2  *v2object.Header
			sv2  *v2refs.Signature
		)

		if prm.hdr.WithID() {
			idv2 = new(v2refs.ObjectID)
		}

		if prm.hdr.WithHeader() {
			hv2 = new(v2object.Header)
		}

		if prm.hdr.WithSignature() {
			sv2 = new(v2refs.Signature)
		}

		prm.hdr.WriteToV2(hv2, sv2, idv2)

		x.reqInit.SetObjectID(idv2)
		x.reqInit.SetHeader(hv2)
		x.reqInit.SetSignature(sv2)
	}

	err := sendObjectPutRequest(&x.reqBody, x.prm, x.v2stream)
	if err != nil {
		return err
	}

	res.payloadStream.prm = x.prm
	res.payloadStream.v2stream = x.v2stream
	res.payloadStream.reqBody.SetObjectPart(&res.payloadStream.reqChunk)

	return nil
}

// PayloadStream provides the interface to stream the payload chunks.
//
// Implements io.WriterCloser.
type PayloadStream struct {
	prm commonPrmWithTokens

	v2stream rpcapi.PutObjectStream

	reqBody  v2object.PutRequestBody
	reqChunk v2object.PutObjectPartChunk

	id oid.ID
}

func (x *PayloadStream) Write(buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, nil
	}

	var (
		err error

		written, chunkLen int
	)

	for {
		const maxChunkLen = 3 * 1 << 20 // 3MB

		if chunkLen = len(buf); chunkLen > maxChunkLen {
			chunkLen = maxChunkLen
		}

		{ // construct chunk part
			x.reqChunk.SetChunk(buf[:chunkLen])
		}

		err = sendObjectPutRequest(&x.reqBody, x.prm, x.v2stream)
		if err != nil {
			return written, err
		}

		written += chunkLen

		// update the buffer
		buf = buf[chunkLen:]
		if len(buf) == 0 {
			break
		}
	}

	return written, nil
}

// Close closes the stream. Should be called once after which there should be no write operations.
// ID of the saved object can be received using ID().
func (x *PayloadStream) Close() error {
	if err := x.v2stream.CloseSend(); err != nil {
		return err
	}

	resp := x.v2stream.Response()

	body := resp.GetBody()
	if body == nil {
		// some sort of selfishness because NeoFS API does not tell "MUST NOT be null" and perhaps it would be worth
		return errMalformedResponse
	}

	idv2 := body.GetObjectID()
	if idv2 == nil {
		// some sort of selfishness because NeoFS API does not tell "MUST NOT be null" and perhaps it would be worth
		return errMalformedResponse
	}

	x.id.FromV2(*idv2)

	return nil
}

// ID returns identifier of the saved object. Should be called after the successful Close.
func (x PayloadStream) ID() oid.ID {
	return x.id
}

func sendObjectPutRequest(
	reqBody *v2object.PutRequestBody,
	prm commonPrmWithTokens,
	v2stream rpcapi.PutObjectStream,
) error {
	var (
		err error
		req v2object.PutRequest
	)

	{ // construct the request
		if err = prepareRequest(&req, prm, func(r requestInterface) {
			r.(*v2object.PutRequest).SetBody(reqBody)
		}); err != nil {
			return err
		}
	}

	{ // send the request
		if err = v2stream.Write(req); err != nil {
			return fmt.Errorf("rpc error: %w", err)
		}
	}

	return nil
}

// PutObject requests to store the object in NeoFS network. After the successful call object stream is opened
// and can be used to save the object in parts.
//
// All required parameters must be set. Result must not be nil.
//
// Context is used for network communication. To set the timeout, use context.WithTimeout or context.WithDeadline.
// It must not be nil.
func (x Client) PutObject(ctx context.Context, prm PutObjectPrm, res *PutObjectRes) error {
	// prelim checks
	prm.checkInputs(ctx, res)

	var (
		err    error
		rpcRes rpcapi.PutObjectRes
	)

	{ // exec RPC
		var rpcPrm rpcapi.PutObjectPrm

		err = rpcapi.PutObject(ctx, x.c, rpcPrm, &rpcRes)
		if err != nil {
			return fmt.Errorf("rpc error: %w", err)
		}
	}

	{ // set results
		res.stream.v2stream = rpcRes.Stream()
		res.stream.prm = prm.commonPrmWithTokens
	}

	return nil
}

// check checks if all required parameters are set and panics if not.
func (x PutObjectPrm) checkInputs(ctx context.Context, res *PutObjectRes) {
	x.commonPrm.checkInputs(ctx)

	if res == nil {
		panic("nil result")
	}
}

// Stream returns initialized PutObjectStream.
func (x PutObjectRes) Stream() PutObjectStream {
	return x.stream
}
