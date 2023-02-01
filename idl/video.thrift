namespace go video
/*
* 定义该idl下三个接口涉及的错误码
* */
enum StatusCode{
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    TokenInvalid               = 10003
}

/*
* 定义结构体
* */

struct BaseResp {
    1: i64 status_code
    2: string status_msg
}

struct Video {
    1: i64 id
    2: Author author
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
}

struct Author {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

/*
* 定义请求与响应
*/


/*
* 视频流接口的请求与响应
* */
struct VideoStreamReq {
    1: optional i64 latest_time
    2: string token
}

struct VideoStreamResp {
    1: required BaseResp base_resp
    2: required i64 next_time
    3: required list<Video> video_list
}


/*
* 投稿接口的请求与响应
*/
struct VideoUploadReq {
    1: required binary data
    2: required string token
    3: required string title
}

struct VideoUploadResp {
    1: required BaseResp base_resp
}


/*
* 发布列表接口的请求与响应
*/
struct VideoListReq {
    1: required string token
    2: required i64 user_id
}

struct VideoListResp {
    1: required BaseResp base_resp
    2: required list<Video> video_list
}

/*
* 定义服务
* */
service VideoService {
    VideoStreamResp VideoStream(1: VideoStreamReq req)
    VideoUploadResp VideoUpload(1: VideoUploadReq req)
    VideoListResp VideoList(1: VideoListReq req)
}






