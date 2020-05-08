package service

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TTL constants.
const (
	// ZeroTTL is an upper bound of invalid TTL values.
	ZeroTTL = iota

	// NonForwardingTTL is a TTL value that does not imply a request forwarding.
	NonForwardingTTL

	// SingleForwardingTTL is a TTL value that imply potential forwarding with NonForwardingTTL.
	SingleForwardingTTL
)

// SetTTL is a TTL field setter.
func (m *RequestMetaHeader) SetTTL(v uint32) {
	m.TTL = v
}

// IRNonForwarding condition that allows NonForwardingTTL only for IR.
func IRNonForwarding(role NodeRole) TTLCondition {
	return func(ttl uint32) error {
		if ttl == NonForwardingTTL && role != InnerRingNode {
			return ErrInvalidTTL
		}

		return nil
	}
}

// ProcessRequestTTL validates and updates requests with TTL.
func ProcessRequestTTL(req TTLContainer, cond ...TTLCondition) error {
	ttl := req.GetTTL()

	if ttl == ZeroTTL {
		return status.New(codes.InvalidArgument, ErrInvalidTTL.Error()).Err()
	}

	for i := range cond {
		if cond[i] == nil {
			continue
		}

		// check specific condition:
		if err := cond[i](ttl); err != nil {
			if st, ok := status.FromError(errors.Cause(err)); ok {
				return st.Err()
			}

			return status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	req.SetTTL(ttl - 1)

	return nil
}