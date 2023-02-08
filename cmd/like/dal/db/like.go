package db

import (
	"context"
)

type Like struct {
	Id      int `gorm:"column:id" json:"id"`             //type:*int   comment:          version:2023-01-07 18:05
	UserId  int `gorm:"column:user_id" json:"user_id"`   //type:*int   comment:          version:2023-01-07 18:05
	VideoId int `gorm:"column:video_id" json:"video_id"` //type:*int   comment:视频id    version:2023-01-07 18:05
}

// TableName 表名:like，。
// 说明:
func (l *Like) TableName() string {
	return "like"
}

// 更新点赞信息——点赞或取消点赞
func LikeAction(ctx context.Context, user_id int, video_id int, action_type int) error {
	/*
		判断action_type,如果是1则点赞，如果是0则取消点赞
		点赞即为向like表插入一条数据，取消点赞即删除一条数据
	*/
	if action_type == 1 {
		like := &Like{
			UserId:  user_id,
			VideoId: video_id,
		}
		if err := DB.WithContext(ctx).Create(like).Error; err != nil {
			return err
		}
	} else {
		if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", user_id, video_id).Delete(&Like{}).Error; err != nil {
			return err
		}
	}
	return nil
}
