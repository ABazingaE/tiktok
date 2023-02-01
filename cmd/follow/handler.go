package main

import (
	"context"
	follow "tiktok/kitex_gen/follow"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// FollowAction implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowAction(ctx context.Context, req *follow.FollowActionReq) (resp *follow.FollowActionResp, err error) {
	//  Your code here...
	return
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListReq) (resp *follow.FollowListResp, err error) {
	//  Your code here...
	return
}

// FollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.FollowerListReq) (resp *follow.FollowerListResp, err error) {
	// Your code here...
	return
}
