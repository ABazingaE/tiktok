package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"tiktok/kitex_gen/follow"
	"tiktok/kitex_gen/follow/followservice"
	"tiktok/pkg/consts"
	"tiktok/pkg/mw"
)

var followClient followservice.Client

func initFollow() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := followservice.NewClient(
		consts.FollowServiceName,
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
	followClient = c
}

// follow action
func FollowAction(ctx context.Context, req *follow.FollowActionReq) (r *follow.FollowActionResp, err error) {
	resp, err := followClient.FollowAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// follow list
func FollowList(ctx context.Context, req *follow.FollowListReq) (r *follow.FollowListResp, err error) {
	resp, err := followClient.FollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// follower list
func FollowerList(ctx context.Context, req *follow.FollowerListReq) (r *follow.FollowerListResp, err error) {
	resp, err := followClient.FollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
