package service

import (
	"context"
	"strconv"
	apiUtil "tiktok/cmd/api/biz/util"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/rpc"
	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/video"
)

type VideoListService struct {
	ctx context.Context
}

func NewVideoListService(ctx context.Context) *VideoListService {
	return &VideoListService{ctx: ctx}
}

func (s *VideoListService) VideoList(req *video.VideoListReq) (resp *video.VideoListResp, err error) {
	//1.查询用户信息
	//解析token，获取请求者的id
	tokenString := req.Token
	requestUserId, err := apiUtil.GetUserIdFromToken(tokenString)
	userInfo, err := rpc.GetUserInfo(s.ctx, &user.UserInfoReq{
		UserId: req.UserId,
		Token:  requestUserId,
	})
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(requestUserId)
	if err != nil {
		return nil, err
	}

	//2.查询用户id下的视频信息
	videoInfo, err := db.GetVideoByAuthorId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	//3.将用户信息和视频信息组装成返回值
	var videoList []*video.Video

	for _, v := range videoInfo {

		//查询当前用户是否喜欢该视频
		isFavorite, err := db.IsFavorite(s.ctx, int64(userId), int64(v.Id))
		if err != nil {
			return nil, err
		}
		videoList = append(videoList, &video.Video{
			Id:            int64(v.Id),
			Author:        (*video.Author)(userInfo.User),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: int64(v.FavoriteCount),
			CommentCount:  int64(v.CommentCount),
			IsFavorite:    isFavorite,
			Title:         v.Title,
		})
	}
	resp = &video.VideoListResp{
		VideoList: videoList,
	}
	return resp, nil
}
