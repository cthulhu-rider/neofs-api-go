package accounting

import (
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
)

// Decimal represents NeoFS API V2-compatible decimal number.
//
// Can be created using var declaration syntax.
type Decimal accounting.Decimal

// FromV2 restores Decimal from accounting.Decimal message.
func (d *Decimal) FromV2(d2 accounting.Decimal) {
	*d = Decimal(d2)
}

// WriteToV2 writes Decimal to accounting.Decimal message.
//
// Message must not be nil.
func (d Decimal) WriteToV2(d2 *accounting.Decimal) {
	v := (accounting.Decimal)(d)

	d2.SetValue(v.GetValue())
	d2.SetPrecision(v.GetPrecision())
}

// Value returns value of the decimal number.
func (d Decimal) Value() int64 {
	return (*accounting.Decimal)(&d).
		GetValue()
}

// SetValue sets value of the decimal number.
func (d *Decimal) SetValue(v int64) {
	(*accounting.Decimal)(d).
		SetValue(v)
}

// Precision returns precision of the decimal number.
func (d Decimal) Precision() uint32 {
	return (*accounting.Decimal)(&d).
		GetPrecision()
}

// SetPrecision sets precision of the decimal number.
func (d *Decimal) SetPrecision(p uint32) {
	(*accounting.Decimal)(d).
		SetPrecision(p)
}

// Marshal marshals Decimal into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (d Decimal) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*accounting.Decimal)(&d).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Decimal.
func (d *Decimal) Unmarshal(data []byte) error {
	return (*accounting.Decimal)(d).
		Unmarshal(data)
}

// MarshalJSON encodes Decimal to protobuf JSON format.
func (d Decimal) MarshalJSON() ([]byte, error) {
	return (*accounting.Decimal)(&d).
		MarshalJSON()
}

// UnmarshalJSON decodes Decimal from protobuf JSON format.
func (d *Decimal) UnmarshalJSON(data []byte) error {
	return (*accounting.Decimal)(d).
		UnmarshalJSON(data)
}
