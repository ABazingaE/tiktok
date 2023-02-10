package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"tiktok/kitex_gen/friend"
	"tiktok/kitex_gen/friend/friendservice"
	"tiktok/pkg/consts"
	"tiktok/pkg/mw"
)

var friendClient friendservice.Client

func initFriend() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := friendservice.NewClient(
		consts.FriendServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	friendClient = c
}

// friend list
func FriendList(ctx context.Context, req *friend.FriendListReq) (r *friend.FriendListResp, err error) {
	resp, err := friendClient.FriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
