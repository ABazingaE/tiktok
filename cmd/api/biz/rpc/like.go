package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"tiktok/kitex_gen/like"
	"tiktok/kitex_gen/like/likeservice"
	"tiktok/pkg/consts"
	"tiktok/pkg/mw"
)

var likeClient likeservice.Client

func initLike() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := likeservice.NewClient(
		consts.LikeServiceName,
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
	likeClient = c
}

// like action
func LikeAction(ctx context.Context, req *like.LikeActionReq) (r *like.LikeActionResp, err error) {
	resp, err := likeClient.LikeAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
