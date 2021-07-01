package accounting_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/accounting"
	accountingtest "github.com/nspcc-dev/neofs-api-go/pkg/accounting/test"
	"github.com/stretchr/testify/require"
)

func TestDecimal_Value(t *testing.T) {
	var d accounting.Decimal

	v := int64(3)
	d.SetValue(v)

	require.Equal(t, v, d.Value())
}

func TestDecimal_Precision(t *testing.T) {
	var d accounting.Decimal

	p := uint32(3)
	d.SetPrecision(p)

	require.Equal(t, p, d.Precision())
}

func TestDecimalEncoding(t *testing.T) {
	d := accountingtest.Generate()

	t.Run("binary", func(t *testing.T) {
		data, err := d.Marshal()
		require.NoError(t, err)

		var d2 accounting.Decimal
		require.NoError(t, d2.Unmarshal(data))

		require.Equal(t, d, d2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := d.MarshalJSON()
		require.NoError(t, err)

		var d2 accounting.Decimal
		require.NoError(t, d2.UnmarshalJSON(data))

		require.Equal(t, d, d2)
	})
}
