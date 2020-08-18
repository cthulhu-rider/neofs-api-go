package service

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	signatureKeyField   = 1
	signatureValueField = 2

	versionMajorField = 1
	versionMinorField = 2

	xheaderKeyField   = 1
	xheaderValueField = 2

	lifetimeExpirationField     = 1
	lifetimeNotValidBeforeField = 2
	lifetimeIssuedAtField       = 3

	objectCtxVerbField    = 1
	objectCtxAddressField = 2

	sessionTokenBodyIDField        = 1
	sessionTokenBodyOwnerField     = 2
	sessionTokenBodyLifetimeField  = 3
	sessionTokenBodyKeyField       = 4
	sessionTokenBodyObjectCtxField = 5

	sessionTokenBodyField      = 1
	sessionTokenSignatureField = 2

	bearerTokenBodyACLField      = 1
	bearerTokenBodyOwnerField    = 2
	bearerTokenBodyLifetimeField = 3

	bearerTokenBodyField      = 1
	bearerTokenSignatureField = 2

	reqMetaHeaderVersionField      = 1
	reqMetaHeaderEpochField        = 2
	reqMetaHeaderTTLField          = 3
	reqMetaHeaderXHeadersField     = 4
	reqMetaHeaderSessionTokenField = 5
	reqMetaHeaderBearerTokenField  = 6
	reqMetaHeaderOriginField       = 7

	reqVerifHeaderBodySignatureField   = 1
	reqVerifHeaderMetaSignatureField   = 2
	reqVerifHeaderOriginSignatureField = 3
	reqVerifHeaderOriginField          = 4

	respMetaHeaderVersionField  = 1
	respMetaHeaderEpochField    = 2
	respMetaHeaderTTLField      = 3
	respMetaHeaderXHeadersField = 4
	respMetaHeaderOriginField   = 5

	respVerifHeaderBodySignatureField   = 1
	respVerifHeaderMetaSignatureField   = 2
	respVerifHeaderOriginSignatureField = 3
	respVerifHeaderOriginField          = 4
)

func (s *Signature) StableMarshal(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.BytesMarshal(signatureKeyField, buf[offset:], s.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.BytesMarshal(signatureValueField, buf[offset:], s.sign)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Signature) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.BytesSize(signatureKeyField, s.key)
	size += proto.BytesSize(signatureValueField, s.sign)

	return size
}

func (v *Version) StableMarshal(buf []byte) ([]byte, error) {
	if v == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, v.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.UInt32Marshal(versionMajorField, buf[offset:], v.major)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.UInt32Marshal(versionMinorField, buf[offset:], v.minor)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (v *Version) StableSize() (size int) {
	if v == nil {
		return 0
	}

	size += proto.UInt32Size(versionMajorField, v.major)
	size += proto.UInt32Size(versionMinorField, v.minor)

	return size
}

func (x *XHeader) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.StringMarshal(xheaderKeyField, buf[offset:], x.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.StringMarshal(xheaderValueField, buf[offset:], x.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *XHeader) StableSize() (size int) {
	if x == nil {
		return 0
	}

	size += proto.StringSize(xheaderKeyField, x.key)
	size += proto.StringSize(xheaderValueField, x.val)

	return size
}

func (l *TokenLifetime) StableMarshal(buf []byte) ([]byte, error) {
	if l == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, l.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.UInt64Marshal(lifetimeExpirationField, buf[offset:], l.exp)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(lifetimeNotValidBeforeField, buf[offset:], l.nbf)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.UInt64Marshal(lifetimeIssuedAtField, buf[offset:], l.iat)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (l *TokenLifetime) StableSize() (size int) {
	if l == nil {
		return 0
	}

	size += proto.UInt64Size(lifetimeExpirationField, l.exp)
	size += proto.UInt64Size(lifetimeNotValidBeforeField, l.nbf)
	size += proto.UInt64Size(lifetimeIssuedAtField, l.iat)

	return size
}

func (c *ObjectSessionContext) StableMarshal(buf []byte) ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, c.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.EnumMarshal(objectCtxVerbField, buf[offset:], int32(c.verb))
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(objectCtxAddressField, buf[offset:], c.addr)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *ObjectSessionContext) StableSize() (size int) {
	if c == nil {
		return 0
	}

	size += proto.EnumSize(objectCtxVerbField, int32(c.verb))
	size += proto.NestedStructureSize(objectCtxAddressField, c.addr)

	return size
}

func (t *SessionTokenBody) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.BytesMarshal(sessionTokenBodyIDField, buf[offset:], t.id)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(sessionTokenBodyOwnerField, buf[offset:], t.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(sessionTokenBodyLifetimeField, buf[offset:], t.lifetime)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(sessionTokenBodyKeyField, buf[offset:], t.sessionKey)
	if err != nil {
		return nil, err
	}

	offset += n

	if t.ctx != nil {
		switch v := t.ctx.(type) {
		case *ObjectSessionContext:
			_, err = proto.NestedStructureMarshal(sessionTokenBodyObjectCtxField, buf[offset:], v)
			if err != nil {
				return nil, err
			}
		default:
			panic("cannot marshal unknown session token context")
		}
	}

	return buf, nil
}

func (t *SessionTokenBody) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += proto.BytesSize(sessionTokenBodyIDField, t.id)
	size += proto.NestedStructureSize(sessionTokenBodyOwnerField, t.ownerID)
	size += proto.NestedStructureSize(sessionTokenBodyLifetimeField, t.lifetime)
	size += proto.BytesSize(sessionTokenBodyKeyField, t.sessionKey)

	if t.ctx != nil {
		switch v := t.ctx.(type) {
		case *ObjectSessionContext:
			size += proto.NestedStructureSize(sessionTokenBodyObjectCtxField, v)
		default:
			panic("cannot marshal unknown session token context")
		}
	}

	return size
}

func (t *SessionToken) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(sessionTokenBodyField, buf[offset:], t.body)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(sessionTokenSignatureField, buf[offset:], t.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (t *SessionToken) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += proto.NestedStructureSize(sessionTokenBodyField, t.body)
	size += proto.NestedStructureSize(sessionTokenSignatureField, t.sig)

	return size
}

func (bt *BearerTokenBody) StableMarshal(buf []byte) ([]byte, error) {
	if bt == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, bt.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(bearerTokenBodyACLField, buf[offset:], bt.eacl)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(bearerTokenBodyOwnerField, buf[offset:], bt.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(bearerTokenBodyLifetimeField, buf[offset:], bt.lifetime)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (bt *BearerTokenBody) StableSize() (size int) {
	if bt == nil {
		return 0
	}

	size += proto.NestedStructureSize(bearerTokenBodyACLField, bt.eacl)
	size += proto.NestedStructureSize(bearerTokenBodyOwnerField, bt.ownerID)
	size += proto.NestedStructureSize(bearerTokenBodyLifetimeField, bt.lifetime)

	return size
}

func (bt *BearerToken) StableMarshal(buf []byte) ([]byte, error) {
	if bt == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, bt.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(bearerTokenBodyField, buf[offset:], bt.body)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(bearerTokenSignatureField, buf[offset:], bt.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (bt *BearerToken) StableSize() (size int) {
	if bt == nil {
		return 0
	}

	size += proto.NestedStructureSize(bearerTokenBodyField, bt.body)
	size += proto.NestedStructureSize(bearerTokenSignatureField, bt.sig)

	return size
}

func (r *RequestMetaHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(reqMetaHeaderVersionField, buf[offset:], r.version)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(reqMetaHeaderEpochField, buf[offset:], r.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(reqMetaHeaderTTLField, buf[offset:], r.ttl)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range r.xHeaders {
		n, err = proto.NestedStructureMarshal(reqMetaHeaderXHeadersField, buf[offset:], r.xHeaders[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	n, err = proto.NestedStructureMarshal(reqMetaHeaderSessionTokenField, buf[offset:], r.sessionToken)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(reqMetaHeaderBearerTokenField, buf[offset:], r.bearerToken)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(reqMetaHeaderOriginField, buf[offset:], r.origin)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *RequestMetaHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	if r.version != nil {
		size += proto.NestedStructureSize(reqMetaHeaderVersionField, r.version)
	}

	size += proto.UInt64Size(reqMetaHeaderEpochField, r.epoch)
	size += proto.UInt32Size(reqMetaHeaderTTLField, r.ttl)

	for i := range r.xHeaders {
		size += proto.NestedStructureSize(reqMetaHeaderXHeadersField, r.xHeaders[i])
	}

	size += proto.NestedStructureSize(reqMetaHeaderSessionTokenField, r.sessionToken)
	size += proto.NestedStructureSize(reqMetaHeaderBearerTokenField, r.bearerToken)
	size += proto.NestedStructureSize(reqMetaHeaderOriginField, r.origin)

	return size
}

func (r *RequestVerificationHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(reqVerifHeaderBodySignatureField, buf[offset:], r.bodySig)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(reqVerifHeaderMetaSignatureField, buf[offset:], r.metaSig)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(reqVerifHeaderOriginSignatureField, buf[offset:], r.originSig)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(reqVerifHeaderOriginField, buf[offset:], r.origin)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *RequestVerificationHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(reqVerifHeaderBodySignatureField, r.bodySig)
	size += proto.NestedStructureSize(reqVerifHeaderMetaSignatureField, r.metaSig)
	size += proto.NestedStructureSize(reqVerifHeaderOriginSignatureField, r.originSig)
	size += proto.NestedStructureSize(reqVerifHeaderOriginField, r.origin)

	return size
}

func (r *ResponseMetaHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(respMetaHeaderVersionField, buf[offset:], r.version)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(respMetaHeaderEpochField, buf[offset:], r.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(respMetaHeaderTTLField, buf[offset:], r.ttl)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range r.xHeaders {
		n, err = proto.NestedStructureMarshal(respMetaHeaderXHeadersField, buf[offset:], r.xHeaders[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	_, err = proto.NestedStructureMarshal(respMetaHeaderOriginField, buf[offset:], r.origin)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *ResponseMetaHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	if r.version != nil {
		size += proto.NestedStructureSize(respMetaHeaderVersionField, r.version)
	}

	size += proto.UInt64Size(respMetaHeaderEpochField, r.epoch)
	size += proto.UInt32Size(respMetaHeaderTTLField, r.ttl)

	for i := range r.xHeaders {
		size += proto.NestedStructureSize(respMetaHeaderXHeadersField, r.xHeaders[i])
	}

	size += proto.NestedStructureSize(respMetaHeaderOriginField, r.origin)

	return size
}

func (r *ResponseVerificationHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(respVerifHeaderBodySignatureField, buf[offset:], r.bodySig)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(respVerifHeaderMetaSignatureField, buf[offset:], r.metaSig)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(respVerifHeaderOriginSignatureField, buf[offset:], r.originSig)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(respVerifHeaderOriginField, buf[offset:], r.origin)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *ResponseVerificationHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(respVerifHeaderBodySignatureField, r.bodySig)
	size += proto.NestedStructureSize(respVerifHeaderMetaSignatureField, r.metaSig)
	size += proto.NestedStructureSize(respVerifHeaderOriginSignatureField, r.originSig)
	size += proto.NestedStructureSize(respVerifHeaderOriginField, r.origin)

	return size
}
