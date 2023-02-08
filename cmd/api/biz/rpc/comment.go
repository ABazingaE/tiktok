package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"tiktok/kitex_gen/comment"
	"tiktok/kitex_gen/comment/commentservice"
	"tiktok/pkg/consts"
	"tiktok/pkg/mw"
)

var commentClient commentservice.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := commentservice.NewClient(
		consts.CommentServiceName,
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
	commentClient = c
}

// comment action
func CommentAction(ctx context.Context, req *comment.CommentActionReq) (r *comment.CommentActionResp, err error) {
	resp, err := commentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// comment list
func CommentList(ctx context.Context, req *comment.CommentListReq) (r *comment.CommentListResp, err error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
