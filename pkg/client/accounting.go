package client

import (
	"context"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/pkg/accounting"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	v2accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
)

// AccountBalancePrm groups the parameters of Client.AccountBalance operation.
type AccountBalancePrm struct {
	commonPrm

	ownerSet bool
	owner    owner.ID
}

// AccountBalanceRes groups the results of Client.AccountBalance operation.
type AccountBalanceRes struct {
	fundNum accounting.Decimal
}

// AccountBalance requests the current balance of the NeoFS account.
//
// All required parameters must be set. Result must not be nil.
func (x Client) AccountBalance(ctx context.Context, prm AccountBalancePrm, res *AccountBalanceRes) error {
	// prelim checks
	prm.check()

	if res == nil {
		panic("nil result argument")
	}

	// construct the request body
	var reqBody v2accounting.BalanceRequestBody

	reqBody.SetOwnerID(prm.owner.ToV2())

	// construct the request meta header
	var hMeta v2session.RequestMetaHeader

	prm.writeToRequestMetaHeader(&hMeta)

	// construct the request
	var (
		err error
		req v2accounting.BalanceRequest
	)

	if err = prm.prepareRequest(&req, func(r requestInterface) {
		r.(*v2accounting.BalanceRequest).SetBody(&reqBody)
	}); err != nil {
		return err
	}

	// exec RPC
	var rpcPrm rpcapi.BalancePrm

	rpcPrm.SetRequest(req)

	var rpcRes rpcapi.BalanceRes

	err = rpcapi.Balance(ctx, x.c, rpcPrm, &rpcRes)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}

	// verify the response
	resp := rpcRes.Response()

	body := resp.GetBody()
	if body == nil {
		return errMalformedResponse
	}

	bal := body.GetBalance()
	if bal == nil {
		return errMalformedResponse
	}

	if err = verifyResponseSignature(&resp); err != nil {
		return err
	}

	// set results
	res.fundNum.FromV2(*bal)

	return nil
}

// check checks if all required parameters are set and panics if not.
func (x AccountBalancePrm) check() {
	x.commonPrm.check()

	switch {
	case !x.ownerSet:
		panic("account ID is required")
	}
}

// SetOwner sets account identifier to get the balance.
//
// Required parameter.
func (x *AccountBalancePrm) SetOwner(id owner.ID) {
	x.ownerSet = true
	x.owner = id
}

// NumberOfFunds returns current number of funds on the account.
func (x AccountBalanceRes) NumberOfFunds() accounting.Decimal {
	return x.fundNum
}
