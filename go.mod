module github.com/cthulhu-rider/neofs-api-go/v2

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/kr/pretty v0.2.0 // indirect
	github.com/nspcc-dev/neofs-crypto v0.3.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

// Used for debug reasons
// replace github.com/nspcc-dev/neofs-crypto => ../neofs-crypto
