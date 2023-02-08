package main

import (
	"context"
	"tiktok/cmd/like/service"
	like "tiktok/kitex_gen/like"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// LikeAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeAction(ctx context.Context, req *like.LikeActionReq) (resp *like.LikeActionResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	_, err = service.NewLikeActionService(ctx).LikeAction(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// LikeList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeList(ctx context.Context, req *like.LikeListReq) (resp *like.LikeListResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	resp, err = service.NewLikeListService(ctx).LikeList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
