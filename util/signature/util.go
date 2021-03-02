package signature

import (
	"github.com/pkg/errors"
)

func dataForSignature(src DataSource) ([]byte, error) {
	if src == nil {
		return nil, errors.New("nil source")
	}

	size := src.SignedDataSize()
	if size < 0 {
		return nil, errors.New("negative length")
	}

	return make([]byte, size), nil
}
