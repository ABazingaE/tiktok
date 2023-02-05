package service

import (
	"context"
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

	//2.查询用户id下的视频信息
	videoInfo, err := db.GetVideoByAuthorId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	//3.将用户信息和视频信息组装成返回值
	videoList := make([]*video.Video, len(videoInfo))
	//去除videoList的第一个空数据
	videoList = videoList[1:]
	for _, v := range videoInfo {
		videoList = append(videoList, &video.Video{
			Id:            int64(v.Id),
			Author:        (*video.Author)(userInfo.User),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: int64(v.FavoriteCount),
			CommentCount:  int64(v.CommentCount),
			//TODO: 未完善 后续完成喜欢操作时，填充is_favorite字段,暂时写死
			IsFavorite: false,
			Title:      v.Title,
		})
	}
	resp = &video.VideoListResp{
		VideoList: videoList,
	}
	return resp, nil
}
