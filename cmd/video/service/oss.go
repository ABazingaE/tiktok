package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/tidwall/gjson"
	"os"
	"sync"
	"tiktok/pkg/consts"
)

func InitVodClient(accessKeyId string, accessKeySecret string) (client *vod.Client, err error) {
	// 点播服务接入区域
	regionId := "cn-shanghai"
	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		accessKeyId,
		accessKeySecret,
	}
	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

func MyCreateUploadVideo(client *vod.Client) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = "Sample Video Title"
	request.Description = "Sample Description"
	request.FileName = "/opt/video/sample/video_file.mp4"
	//request.CateId = "-1"
	//Cover URL示例：http://example.alicdn.com/tps/TB1qnJ1PVXXXXXCXXXXXXXXXXXX-700-****.png
	request.CoverURL = "<your CoverURL>"
	request.Tags = "tag1,tag2"
	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

func InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}

func UploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, localFile string) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传本地文件。
	err = bucket.PutObjectFromFile(uploadAddressDTO.FileName, localFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

type UploadAuthDTO struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}
type UploadAddressDTO struct {
	Endpoint string
	Bucket   string
	FileName string
}

// 整合服务
func UploadAndGetVideoInfo(path string) (playUrl string, coverUrl string) {
	//1.上传视频
	videoId := UploadVideo(path)

	//2.获取AI封面模板id
	coverTemplateId := GetAIImageTemplateId()

	//3.提交AI封面任务，生成封面
	SubmitAIImageJob(videoId, coverTemplateId)

	//4.获取视频和封面地址
	playUrl, coverUrl = GetPlayInfo(videoId)
	return playUrl, coverUrl
}

/*上传视频*/
func UploadVideo(path string) (Id string) {
	var accessKeyId string = consts.AccessKeyId         // 您的AccessKeyId
	var accessKeySecret string = consts.AccessKeySecret // 您的AccessKeySecret
	var localFile string = path                         // 需要上传到VOD的本地视频文件的完整路径
	// 初始化VOD客户端并获取上传地址和凭证
	var vodClient, initVodClientErr = InitVodClient(accessKeyId, accessKeySecret)
	if initVodClientErr != nil {
		fmt.Println("Error:", initVodClientErr)
		return
	}
	// 获取上传地址和凭证
	var response, createUploadVideoErr = MyCreateUploadVideo(vodClient)
	if createUploadVideoErr != nil {
		fmt.Println("Error:", createUploadVideoErr)
		return
	}
	// 执行成功会返回VideoId、UploadAddress和UploadAuth
	var videoId = response.VideoId
	var uploadAuthDTO UploadAuthDTO
	var uploadAddressDTO UploadAddressDTO
	var uploadAuthDecode, _ = base64.StdEncoding.DecodeString(response.UploadAuth)
	var uploadAddressDecode, _ = base64.StdEncoding.DecodeString(response.UploadAddress)
	json.Unmarshal(uploadAuthDecode, &uploadAuthDTO)
	json.Unmarshal(uploadAddressDecode, &uploadAddressDTO)
	// 使用UploadAuth和UploadAddress初始化OSS客户端
	var ossClient, _ = InitOssClient(uploadAuthDTO, uploadAddressDTO)
	// 上传文件，注意是同步上传会阻塞等待，耗时与文件大小和网络上行带宽有关
	UploadLocalFile(ossClient, uploadAddressDTO, localFile)
	//MultipartUploadFile(ossClient, uploadAddressDTO, localFile)
	fmt.Println("Succeed, VideoId:", videoId)
	return videoId
}

/*查询AI封面模板id*/
func GetAIImageTemplateId() (TemplateId string) {
	var accessKeyId string = consts.AccessKeyId         // 您的AccessKeyId
	var accessKeySecret string = consts.AccessKeySecret // 您的AccessKeySecret

	// 初始化VOD客户端并获取上传地址和凭证
	var vodClient, initVodClientErr = InitVodClient(accessKeyId, accessKeySecret)
	if initVodClientErr != nil {
		fmt.Println("Error:", initVodClientErr)
		return
	}
	request := vod.CreateListAITemplateRequest()
	request.TemplateType = "AIImage"
	resp, err := vodClient.ListAITemplate(request)
	if err != nil {
		fmt.Print(err.Error())
		return AddAIImageTemplate()
	}
	jsonByte, _ := json.Marshal(resp)
	jsonString := string(jsonByte[:])
	TemplateId = gjson.Get(jsonString, "TemplateInfoList.0.TemplateId").String()
	return TemplateId
}

/*添加AI封面模板*/

type TemplateConfigInfo struct {
	Format          string `json:"Format"`
	SetDefaultCover string `json:"SetDefaultCover"`
}

func AddAIImageTemplate() (TemplateId string) {
	var accessKeyId string = consts.AccessKeyId         // 您的AccessKeyId
	var accessKeySecret string = consts.AccessKeySecret // 您的AccessKeySecret

	// 初始化VOD客户端并获取上传地址和凭证
	var vodClient, initVodClientErr = InitVodClient(accessKeyId, accessKeySecret)
	if initVodClientErr != nil {
		fmt.Println("Error:", initVodClientErr)
		return
	}
	configInfo := &TemplateConfigInfo{
		Format:          "png",
		SetDefaultCover: "true",
	}
	req := vod.CreateAddAITemplateRequest()
	req.TemplateType = "AIImage"
	req.TemplateName = "GenerateCover"
	req.TemplateConfig = *util.ToJSONString(tea.ToMap(configInfo))

	resp, err := vodClient.AddAITemplate(req)
	if err != nil {
		fmt.Println("Error:", err)

	}

	jsonByte, _ := json.Marshal(resp)
	jsonString := string(jsonByte[:])
	TemplateId = gjson.Get(jsonString, "TemplateId").String()
	return TemplateId
}

/*提交AI图片任务*/
func SubmitAIImageJob(VideoId string, AITemplateId string) {
	var accessKeyId string = consts.AccessKeyId         // 您的AccessKeyId
	var accessKeySecret string = consts.AccessKeySecret // 您的AccessKeySecret

	// 初始化VOD客户端并获取上传地址和凭证
	var vodClient, initVodClientErr = InitVodClient(accessKeyId, accessKeySecret)
	if initVodClientErr != nil {
		fmt.Println("Error:", initVodClientErr)
		return
	}
	request := vod.CreateSubmitAIImageJobRequest()
	request.VideoId = VideoId
	request.AITemplateId = AITemplateId
	resp, err := vodClient.SubmitAIImageJob(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Succeed, JobId:", resp.JobId)

}

var wg = sync.WaitGroup{}

/*获取视频信息*/
func GetPlayInfo(videoId string) (playUrl string, coverUrl string) {
	var accessKeyId string = consts.AccessKeyId         // 您的AccessKeyId
	var accessKeySecret string = consts.AccessKeySecret // 您的AccessKeySecret

	// 初始化VOD客户端并获取上传地址和凭证
	var vodClient, initVodClientErr = InitVodClient(accessKeyId, accessKeySecret)
	if initVodClientErr != nil {
		fmt.Println("Error:", initVodClientErr)
		return
	}
	urlResult := make(chan UrlResult, 1)
	go GetUrlResult(vodClient, videoId, urlResult)
	wg.Wait()
	//获取channel中的值
	result := <-urlResult
	playUrl = result.playUrl
	coverUrl = result.coverUrl
	//打印两个url
	fmt.Println("playUrl:", playUrl)
	fmt.Println("coverUrl:", coverUrl)
	return playUrl, coverUrl
}

type UrlResult struct {
	playUrl  string
	coverUrl string
}

func GetUrlResult(vodClient *vod.Client, videoId string, result chan UrlResult) {
	wg.Add(1)
	for {
		request := vod.CreateGetPlayInfoRequest()
		request.VideoId = videoId
		resp, err := vodClient.GetPlayInfo(request)
		if err != nil {
			fmt.Println("Error:", err)

		}
		jsonByte, _ := json.Marshal(resp)
		jsonString := string(jsonByte[:])
		playUrl := gjson.Get(jsonString, "PlayInfoList.PlayInfo.0.PlayURL").String()
		coverUrl := gjson.Get(jsonString, "VideoBase.CoverURL").String()
		if playUrl != "" && coverUrl != "<your CoverURL>" {
			resp := UrlResult{
				playUrl:  playUrl,
				coverUrl: coverUrl,
			}
			result <- resp
			break
		}
	}
	wg.Done()

}
