namespace go api

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

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct Comment {
    1: i64 id
    2: User user
    3: string content
    4: string create_date
}

struct FriendUser {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: optional string message
    8: required i64 msgType
}

struct Message{
    1: i64 id
    2: i64 to_user_id
    3: i64 from_user_id
    4: string content
    5: string create_time
}

struct MessageSendEvent{
    1: i64 user_id
    2: i64 to_user_id
    3: i64 from_user_id
    4: string msg_content
}

struct MessagePushEvent{
    1: i64 from_user_id
    2: string msg_content
}

/*
* 视频流接口的请求与响应
* */
struct VideoStreamReq {
    1: optional i64 latest_time     (api.query = "latest_time")
    2: string token                 (api.query = "token")
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
    1: required binary data     (api.form = "data")
    2: required string token    (api.form = "token")
    3: required string title    (api.form = "title")
}

struct VideoUploadResp {
    1: required BaseResp base_resp
}


/*
* 发布列表接口的请求与响应
*/
struct VideoListReq {
    1: required string token    (api.query = "token")
    2: required i64 user_id     (api.query = "user_id")
}

struct VideoListResp {
    1: required BaseResp base_resp
    2: required list<Video> video_list
}


//注册
struct RegisterReq {
    1: required string username   (api.query = "username", api.vd = "len($) < 32")
    2: required string password   (api.query = "password", api.vd = "len($) < 32")
}

struct RegisterResp {
    1: required BaseResp base_resp
    2: required i64 user_id
    3: required string token
}


//登录
struct LoginReq {
    1: required string username (api.query = "username", api.vd = "len($) < 32")
    2: required string password (api.query = "password", api.vd = "len($) < 32")
}

struct LoginResp {
    1: required BaseResp base_resp
    2: required i64 user_id
    3: required string token
}


//用户信息
struct UserInfoReq {
    1: required i64 user_id     (api.query = "user_id")
    2: required string token    (api.query = "token")
}

struct UserInfoResp {
    1: required BaseResp base_resp
    2: required User user
}

/*
* 赞操作
* */
struct LikeActionReq {
    1: required string token    (api.query = "token")
    2: required i64 video_id    (api.query = "video_id")
    3: required i64 action_type (api.query = "action_type")
}

struct LikeActionResp {
    1: required BaseResp base_resp
}

/*
* 喜欢列表
*/
struct LikeListReq {
    1: required i64 user_id     (api.query = "user_id")
    2: required string token    (api.query = "token")
}

struct LikeListResp {
    1: required BaseResp base_resp
    2: required list<Video> video_list
}


/*
* 评论相关
* */
struct CommentActionReq{
    1: required string token    (api.query = "token")
    2: required i64 video_id    (api.query = "video_id")
    3: required i64 action_type (api.query = "action_type")
    4: string comment_text      (api.query = "comment_text")
    5: i64 comment_id           (api.query = "comment_id")
}

struct CommentActionResp {
    1: required BaseResp base_resp
    2: required Comment comment
}

struct CommentListReq {
    1: required string token  (api.query = "token")
    2: required i64 video_id  (api.query = "video_id")
}

struct CommentListResp {
    1: required BaseResp base_resp
    2: required list<Comment> comment_list
}

/*
* 关注操作
* */

struct FollowActionReq {
    1: required string token    (api.query = "token")
    2: required i64 to_user_id  (api.query = "to_user_id")
    3: required i64 action_type (api.query = "action_type")
}

struct FollowActionResp {
    1: required BaseResp base_resp
}

/*
* 关注列表
* */

struct FollowListReq {
    1: required string token    (api.query = "token")
    2: required i64 user_id     (api.query = "user_id")
}

struct FollowListResp {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

/*
* 粉丝列表
* */

struct FollowerListReq {
    1: required string token    (api.query = "token")
    2: required i64 user_id     (api.query = "user_id")
}

struct FollowerListResp {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

/*
* 粉丝列表
* */
struct FriendListReq{
    1: required i64 user_id     (api.query = "user_id")
    2: required string token    (api.query = "token")
}

struct FriendListResp{
    1: required BaseResp base_resp
    2: required list<FriendUser> user_list
}

/*
* 聊天记录
* */
struct MessageChatReq{
    1: required string token    (api.query = "token")
    2: required i64 to_user_id  (api.query = "to_user_id")
}

struct MessageChatResp{
    1: required BaseResp base_resp
    2: required list<Message> message_list
}

/*
* 消息操作
* */
struct MessageActionReq{
    1: required string token        (api.query = "token")
    2: required i64 to_user_id      (api.query = "to_user_id")
    3: required i64 action_type     (api.query = "action_type")
    4: required string content      (api.query = "content")
}

struct MessageActionResp{
    1: required BaseResp base_resp
}

service ApiService{
        //user
        RegisterResp Register(1: RegisterReq req)   (api.post = "/douyin/user/register/")
        LoginResp Login(1: LoginReq req)        (api.post = "/douyin/user/login/")
        UserInfoResp UserInfo(1: UserInfoReq req)   (api.get = "/douyin/user/")

        //video
        VideoStreamResp VideoStream(1: VideoStreamReq req) (api.get = "/douyin/feed/")
        VideoUploadResp VideoUpload(1: VideoUploadReq req)  (api.post = "/douyin/publish/action/")
        VideoListResp VideoList(1: VideoListReq req)    (api.get = "/douyin/publish/list/")

        //like
        LikeActionResp LikeAction(1: LikeActionReq req) (api.post = "/douyin/favorite/action/")
        LikeListResp LikeList(1: LikeListReq req)   (api.get = "/douyin/favorite/list/")

        //comment
        CommentActionResp CommentAction(1: CommentActionReq req)    (api.post = "/douyin/comment/action/")
        CommentListResp CommentList(1: CommentListReq req)    (api.get = "/douyin/comment/list/")

        //follow
        FollowActionResp FollowAction(1: FollowActionReq req)   (api.post = "/douyin/relation/action/")
        FollowListResp FollowList(1: FollowListReq req)   (api.get = "/douyin/relation/follow/list/")
        FollowerListResp FollowerList(1: FollowerListReq req)   (api.get = "/douyin/relation/follower/list/")

        //friend
        FriendListResp FriendList(1:FriendListReq req)  (api.get = "/douyin/relation/friend/list/")
        MessageChatResp MessageChat(1:MessageChatReq req)   (api.get = "/douyin/message/chat/")
        MessageActionResp MessageAction(1:MessageActionReq req) (api.post = "/douyin/message/action/")
}