package accounting

import (
	"github.com/cthulhu-rider/neofs-api-go/v2/refs"
	"github.com/cthulhu-rider/neofs-api-go/v2/session"
)

type BalanceRequestBody struct {
	ownerID *refs.OwnerID
}

type BalanceResponseBody struct {
	bal *Decimal
}

type Decimal struct {
	val int64

	prec uint32
}

func (b *BalanceRequestBody) GetOwnerID() *refs.OwnerID {
	if b != nil {
		return b.ownerID
	}

	return nil
}

func (b *BalanceRequestBody) SetOwnerID(v *refs.OwnerID) {
	if b != nil {
		b.ownerID = v
	}
}

func (b *BalanceRequest) GetBody() *BalanceRequestBody {
	if b != nil {
		return b.body
	}

	return nil
}

func (b *BalanceRequest) SetBody(v *BalanceRequestBody) {
	if b != nil {
		b.body = v
	}
}

func (b *BalanceRequest) GetMetaHeader() *session.RequestMetaHeader {
	if b != nil {
		return b.metaHeader
	}

	return nil
}

func (b *BalanceRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if b != nil {
		b.metaHeader = v
	}
}

func (b *BalanceRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if b != nil {
		return b.verifyHeader
	}

	return nil
}

func (b *BalanceRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if b != nil {
		b.verifyHeader = v
	}
}

func (d *Decimal) GetValue() int64 {
	if d != nil {
		return d.val
	}

	return 0
}

func (d *Decimal) SetValue(v int64) {
	if d != nil {
		d.val = v
	}
}

func (d *Decimal) GetPrecision() uint32 {
	if d != nil {
		return d.prec
	}

	return 0
}

func (d *Decimal) SetPrecision(v uint32) {
	if d != nil {
		d.prec = v
	}
}

func (br *BalanceResponseBody) GetBalance() *Decimal {
	if br != nil {
		return br.bal
	}

	return nil
}

func (br *BalanceResponseBody) SetBalance(v *Decimal) {
	if br != nil {
		br.bal = v
	}
}

func (br *BalanceResponse) GetBody() *BalanceResponseBody {
	if br != nil {
		return br.body
	}

	return nil
}

func (br *BalanceResponse) SetBody(v *BalanceResponseBody) {
	if br != nil {
		br.body = v
	}
}

func (br *BalanceResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if br != nil {
		return br.metaHeader
	}

	return nil
}

func (br *BalanceResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if br != nil {
		br.metaHeader = v
	}
}

func (br *BalanceResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if br != nil {
		return br.verifyHeader
	}

	return nil
}

func (br *BalanceResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if br != nil {
		br.verifyHeader = v
	}
}
