package main

import (
	"context"
	like "tiktok/kitex_gen/like"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// LikeAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeAction(ctx context.Context, req *like.LikeActionReq) (resp *like.LikeActionResp, err error) {
	//  Your code here...
	return
}

// LikeList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeList(ctx context.Context, req *like.LikeListReq) (resp *like.LikeListResp, err error) {
	//  Your code here...
	return
}
