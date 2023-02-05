package service

import (
	"context"
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

	if *latestTime == int64(0) {
		latestTime = new(int64)
		*latestTime = time.Now().Unix()
	}

	videoInfo, nextTime, err := db.GetVideoStream(s.ctx, *latestTime)
	if err != nil {
		return nil, err
	}
	videoList := make([]*video.Video, len(videoInfo))
	//去除videoList的第一个空数据
	videoList = videoList[1:]
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
		authorInfoResp, err := rpc.GetUserInfo(s.ctx, authorInfoReq)
		if err != nil {
			return nil, err
		}
		//完善videoList
		videoList = append(videoList, &video.Video{
			Id:            int64(v.Id),
			Author:        (*video.Author)(authorInfoResp.User),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: int64(v.FavoriteCount),
			CommentCount:  int64(v.CommentCount),
			//TODO: 未完善 后续完成喜欢操作时，填充is_favorite字段,暂时写死
			IsFavorite: false,
			Title:      v.Title,
		})

	}

	resp = &video.VideoStreamResp{
		VideoList: videoList,
		NextTime:  nextTime,
	}

	return resp, nil

}
