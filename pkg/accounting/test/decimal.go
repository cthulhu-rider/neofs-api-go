package accountingtest

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/accounting"
)

// Generate returns random accounting.Decimal.
func Generate() (r accounting.Decimal) {
	r.SetValue(1)
	r.SetPrecision(2)

	return
}
