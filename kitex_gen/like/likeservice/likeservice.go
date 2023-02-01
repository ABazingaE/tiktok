// Code generated by Kitex v0.4.4. DO NOT EDIT.

package likeservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	like "tiktok/kitex_gen/like"
)

func serviceInfo() *kitex.ServiceInfo {
	return likeServiceServiceInfo
}

var likeServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "LikeService"
	handlerType := (*like.LikeService)(nil)
	methods := map[string]kitex.MethodInfo{
		"LikeAction": kitex.NewMethodInfo(likeActionHandler, newLikeServiceLikeActionArgs, newLikeServiceLikeActionResult, false),
		"LikeList":   kitex.NewMethodInfo(likeListHandler, newLikeServiceLikeListArgs, newLikeServiceLikeListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "like",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func likeActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceLikeActionArgs)
	realResult := result.(*like.LikeServiceLikeActionResult)
	success, err := handler.(like.LikeService).LikeAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceLikeActionArgs() interface{} {
	return like.NewLikeServiceLikeActionArgs()
}

func newLikeServiceLikeActionResult() interface{} {
	return like.NewLikeServiceLikeActionResult()
}

func likeListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceLikeListArgs)
	realResult := result.(*like.LikeServiceLikeListResult)
	success, err := handler.(like.LikeService).LikeList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceLikeListArgs() interface{} {
	return like.NewLikeServiceLikeListArgs()
}

func newLikeServiceLikeListResult() interface{} {
	return like.NewLikeServiceLikeListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) LikeAction(ctx context.Context, req *like.LikeActionReq) (r *like.LikeActionResp, err error) {
	var _args like.LikeServiceLikeActionArgs
	_args.Req = req
	var _result like.LikeServiceLikeActionResult
	if err = p.c.Call(ctx, "LikeAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LikeList(ctx context.Context, req *like.LikeListReq) (r *like.LikeListResp, err error) {
	var _args like.LikeServiceLikeListArgs
	_args.Req = req
	var _result like.LikeServiceLikeListResult
	if err = p.c.Call(ctx, "LikeList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}