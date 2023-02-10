package service

import (
	"context"
	"strconv"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/rpc"
	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/video"
	"time"
)

type VideoStreamService struct {
	ctx context.Context
}

func NewVideoFeedService(ctx context.Context) *VideoStreamService {
	return &VideoStreamService{ctx: ctx}
}

func (s *VideoStreamService) VideoFeed(req *video.VideoStreamReq) (resp *video.VideoStreamResp, err error) {
	//判断latest_time参数，若为0则设定为当下时间
	latestTime := req.LatestTime

	if *latestTime == int64(0) || req.LatestTime == nil {
		latestTime = new(int64)
		*latestTime = time.Now().Unix()
	}

	//取出请求者id
	var requestId int
	if req.Token != "" {
		requestId, err = strconv.Atoi(req.Token)
		if err != nil {
			return nil, err
		}
	}

	//查询视频基本信息
	videoInfo, nextTime, err := db.GetVideoStream(s.ctx, *latestTime)
	if err != nil {
		return nil, err
	}
	var videoList []*video.Video

	//遍历videoInfo,每次遍历取出authorId，根据authorId查询用户信息
	for _, v := range videoInfo {
		authorId := v.AuthorId
		//构造请求，rpc调用user中的接口
		authorInfoReq := &user.UserInfoReq{
			UserId: authorId,
		}
		if req.Token != "" {
			authorInfoReq.Token = req.Token
		}
		//查询视频作者信息
		authorInfoResp, err := rpc.GetUserInfo(s.ctx, authorInfoReq)
		if err != nil {
			return nil, err
		}

		//查询当前用户是否喜欢该视频
		var isFavorite bool
		if req.Token != "" {
			isFavorite, err = db.IsFavorite(s.ctx, int64(requestId), int64(v.Id))
			if err != nil {
				return nil, err
			}
		} else {
			isFavorite = false
		}

		//完善videoList
		videoList = append(videoList, &video.Video{
			Id:            int64(v.Id),
			Author:        (*video.Author)(authorInfoResp.User),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: int64(v.FavoriteCount),
			CommentCount:  int64(v.CommentCount),
			IsFavorite:    isFavorite,
			Title:         v.Title,
		})

	}

	resp = &video.VideoStreamResp{
		VideoList: videoList,
		NextTime:  nextTime,
	}

	return resp, nil

}
