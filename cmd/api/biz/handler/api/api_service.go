// Code generated by hertz generator.

package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "tiktok/cmd/api/biz/model/api"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RegisterResp)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.LoginResp)

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserInfoResp)

	c.JSON(consts.StatusOK, resp)
}

// VideoStream .
// @router /douyin/feed/ [GET]
func VideoStream(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.VideoStreamReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.VideoStreamResp)

	c.JSON(consts.StatusOK, resp)
}

// VideoUpload .
// @router /douyin/publish/action/ [POST]
func VideoUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.VideoUploadReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.VideoUploadResp)

	c.JSON(consts.StatusOK, resp)
}

// VideoList .
// @router /douyin/publish/list/ [GET]
func VideoList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.VideoListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.VideoListResp)

	c.JSON(consts.StatusOK, resp)
}

// LikeAction .
// @router /douyin/favorite/action/ [POST]
func LikeAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LikeActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.LikeActionResp)

	c.JSON(consts.StatusOK, resp)
}

// LikeList .
// @router /douyin/favorite/list/ [GET]
func LikeList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LikeListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.LikeListResp)

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentActionResp)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentListResp)

	c.JSON(consts.StatusOK, resp)
}

// FollowAction .
// @router /douyin/relation/action/ [POST]
func FollowAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FollowActionResp)

	c.JSON(consts.StatusOK, resp)
}

// FollowList .
// @router /douyin/relation/follow/list/ [GET]
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FollowListResp)

	c.JSON(consts.StatusOK, resp)
}

// FollowerList .
// @router /douyin/relation/follower/list/ [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowerListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FollowerListResp)

	c.JSON(consts.StatusOK, resp)
}

// FriendList .
// @router /douyin/relation/friend/list/ [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FriendListResp)

	c.JSON(consts.StatusOK, resp)
}

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageChatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.MessageChatResp)

	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.MessageActionResp)

	c.JSON(consts.StatusOK, resp)
}
