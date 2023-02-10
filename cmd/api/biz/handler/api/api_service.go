// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"io"
	"mime/multipart"
	"net/http"
	api "tiktok/cmd/api/biz/model/api"
	"tiktok/cmd/api/biz/mw"
	"tiktok/cmd/api/biz/rpc"
	apiUtil "tiktok/cmd/api/biz/util"
	"tiktok/kitex_gen/comment"
	"tiktok/kitex_gen/follow"
	"tiktok/kitex_gen/friend"
	"tiktok/kitex_gen/like"
	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	//注册，获取id
	_, err = rpc.CreateUser(context.Background(), &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	//登录
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	//解析token，获取请求者的id
	tokenString := req.Token
	requestUserId, err := apiUtil.GetUserIdFromToken(tokenString)
	var resp *user.UserInfoResp
	resp, err = rpc.GetUserInfo(context.Background(), &user.UserInfoReq{
		UserId: req.UserID,
		Token:  requestUserId,
	})

	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user":        resp.User,
	})
}

// VideoStream .
// @router /douyin/feed/ [GET]
func VideoStream(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.VideoStreamReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	tokenString := req.Token
	videoSteamReq := &video.VideoStreamReq{}

	videoSteamReq.LatestTime = req.LatestTime
	if tokenString != "" {
		requestUserId, err := apiUtil.GetUserIdFromToken(tokenString)
		if err != nil {
			SendResponse(c, errno.ConvertErr(err))
			return
		}
		videoSteamReq.Token = requestUserId
	}

	var resp *video.VideoStreamResp
	resp, err = rpc.VideoStream(context.Background(), videoSteamReq)

	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"next_time":   resp.NextTime,
		"video_list":  resp.VideoList,
	})
}

// VideoUpload .
// @router /douyin/publish/action/ [POST]
func VideoUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	//var req api.VideoUploadReq
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	SendResponse(c, errno.ConvertErr(err))
	//	return
	//}
	var form *multipart.Form
	form, err = c.MultipartForm()
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	//获取视频
	var videoData *multipart.FileHeader
	videoData = form.File["data"][0]
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	videoReader, err := videoData.Open()
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	videoByte, err := io.ReadAll(videoReader)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	//解析token，获取请求者的id
	tokenString := form.Value["token"][0]
	requestUserId, err := apiUtil.GetUserIdFromToken(tokenString)

	//获取title
	title := form.Value["title"][0]

	_, err = rpc.VideoUpload(context.Background(), &video.VideoUploadReq{
		Data:  videoByte,
		Title: title,
		Token: requestUserId,
	})

	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
	})
}

// VideoList .
// @router /douyin/publish/list/ [GET]
func VideoList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.VideoListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	resp, err := rpc.VideoList(context.Background(), &video.VideoListReq{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"video_list":  resp.VideoList,
	})
}

// LikeAction .
// @router /douyin/favorite/action/ [POST]
func LikeAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LikeActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	_, err = rpc.LikeAction(context.Background(), &like.LikeActionReq{
		Token:      userId,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
	})

}

// LikeList .
// @router /douyin/favorite/list/ [GET]
func LikeList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LikeListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	resp, err := rpc.LikeList(context.Background(), &like.LikeListReq{
		UserId: req.UserID,
		Token:  userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"video_list":  resp.VideoList,
	})
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	resp, err := rpc.CommentAction(context.Background(), &comment.CommentActionReq{
		Token:       userId,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"comment":     resp.Comment,
	})
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	resp, err := rpc.CommentList(context.Background(), &comment.CommentListReq{
		Token:   userId,
		VideoId: req.VideoID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code":  0,
		"status_msg":   "success",
		"comment_list": resp.CommentList,
	})
}

// FollowAction .
// @router /douyin/relation/action/ [POST]
func FollowAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	_, err = rpc.FollowAction(context.Background(), &follow.FollowActionReq{
		Token:      userId,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
	})
}

// FollowList .
// @router /douyin/relation/follow/list/ [GET]
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	resp := new(follow.FollowListResp)
	resp, err = rpc.FollowList(context.Background(), &follow.FollowListReq{
		Token:  userId,
		UserId: req.UserID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   resp.UserList,
	})
}

// FollowerList .
// @router /douyin/relation/follower/list/ [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowerListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	resp := new(follow.FollowerListResp)
	resp, err = rpc.FollowerList(context.Background(), &follow.FollowerListReq{
		Token:  userId,
		UserId: req.UserID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   resp.UserList,
	})
}

// FriendList .
// @router /douyin/relation/friend/list/ [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	token := req.Token
	userId, err := apiUtil.GetUserIdFromToken(token)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	resp := new(friend.FriendListResp)
	resp, err = rpc.FriendList(context.Background(), &friend.FriendListReq{
		Token:  userId,
		UserId: req.UserID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   resp.UserList,
	})
}

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageChatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	resp := new(api.MessageChatResp)

	c.JSON(0, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	resp := new(api.MessageActionResp)

	c.JSON(0, resp)
}
