namespace go like


struct BaseResp {
    1: i64 status_code
    2: string status_msg
}

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
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

struct Comment {
    1: i64 id
    2: User user
    3: string content
    4: string create_date
}

/*
* 赞操作
* */
struct LikeActionReq {
    1: required string token
    2: required i64 video_id
    3: required i64 action_type (vt.in = "1", vt.in = "2")
}

struct LikeActionResp {
    1: required BaseResp base_resp
}

/*
* 喜欢列表
*/
struct LikeListReq {
    1: required i64 user_id
    2: required string token
}

struct LikeListResp {
    1: required BaseResp base_resp
    2: required list<Video> video_list
}


service LikeService {
    LikeActionResp LikeAction(1: LikeActionReq req)
    LikeListResp LikeList(1: LikeListReq req)
}