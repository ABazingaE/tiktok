package service

import (
	"context"
	"go.etcd.io/etcd/pkg/v3/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"tiktok/cmd/video/dal/db"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/consts"
	"time"
)

type VideoUploadService struct {
	ctx context.Context
}

func NewVideoUploadService(ctx context.Context) *VideoUploadService {
	return &VideoUploadService{ctx: ctx}
}

func (s *VideoUploadService) VideoUpload(req *video.VideoUploadReq) (resp *video.VideoUploadResp, err error) {
	//将request中的视频信息写入临时文件夹,根据时间戳生成临时视频的文件名

	//获取当下时间戳
	now := time.Now().Unix()
	//将时间戳转换为字符串
	timeString := strconv.FormatInt(now, 10)
	filePath := consts.TempPath + timeString + ".mp4"
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ioutil.WriteAndSyncFile(filePath, req.Data, 0644)
		wg.Done()
	}()
	wg.Wait()

	playUrl, coverUrl := UploadAndGetVideoInfo(filePath)
	//修改Url, 读取字符串，直到遇到“?”为止,保留“?”之前的字符串
	playUrl = playUrl[:strings.Index(playUrl, "?")]
	coverUrl = coverUrl[:strings.Index(coverUrl, "?")]
	//删除临时文件
	os.Remove(filePath)

	//将视频信息写入数据库
	authorId, _ := strconv.ParseInt(req.Token, 10, 64)
	err = db.CreateVideo(s.ctx, []*db.Video{{
		PlayUrl:     playUrl,
		CoverUrl:    coverUrl,
		PublishTime: int(now),
		AuthorId:    authorId,
		Title:       req.Title,
	}})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
