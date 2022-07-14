package tests

import (
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	
	"github.com/cs3org/reva/pkg/sdk"
	"github.com/cs3org/reva/pkg/sdk/common/net"
)

func PurgeRecycleBin(session *sdk.Session) error {
	// TODO: Use SDK once it has been updated
	req := &provider.GetHomeRequest{}
	res, err := session.Client().GetHome(session.Context(), req)
	if err := net.CheckRPCInvocation("querying home directory", res, err); err != nil {
		return err
	}
	homePath := res.Path
	ref := &provider.Reference{Path: homePath}
	req2 := &provider.PurgeRecycleRequest{Ref: ref}
	res2, err := session.Client().PurgeRecycle(session.Context(), req2)
	return net.CheckRPCInvocation("purging recycle bin", res2, err)
}
