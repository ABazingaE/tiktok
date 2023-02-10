package main

import (
	"context"
	"tiktok/cmd/follow/service"
	follow "tiktok/kitex_gen/follow"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// FollowAction implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowAction(ctx context.Context, req *follow.FollowActionReq) (resp *follow.FollowActionResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	_, err = service.NewFollowActionService(ctx).FollowAction(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListReq) (resp *follow.FollowListResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	resp, err = service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.FollowerListReq) (resp *follow.FollowerListResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	resp, err = service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
