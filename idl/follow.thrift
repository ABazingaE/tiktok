namespace go follow

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

/*
* 关注操作
* */

struct FollowActionReq {
    1: required string token
    2: required i64 to_user_id
    3: required i64 action_type (vt.in = "1", vt.in = "2")
}

struct FollowActionResp {
    1: required BaseResp base_resp
}

/*
* 关注列表
* */

struct FollowListReq {
    1: required string token
    2: required i64 user_id
}

struct FollowListResp {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

/*
* 粉丝列表
* */

struct FollowerListReq {
    1: required string token
    2: required i64 user_id
}

struct FollowerListResp {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

service FollowService {
    FollowActionResp FollowAction(1: FollowActionReq req)
    FollowListResp FollowList(1: FollowListReq req)
    FollowerListResp FollowerList(1: FollowerListReq req)
}