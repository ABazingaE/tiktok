package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tiktok/pkg/errno"
)

type Response struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}
