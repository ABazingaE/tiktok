namespace go comment

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

struct Comment {
    1: i64 id
    2: User user
    3: string content
    4: string create_date
}

struct CommentActionReq{
    1: required string token
    2: required i64 video_id
    3: required i64 action_type (vt.in = "1", vt.in = "2")
    4: string comment_text
    5: i64 comment_id
}

struct CommentActionResp {
    1: required BaseResp base_resp
    2: required Comment comment
}

struct CommentListReq {
    1: required string token
    2: required i64 video_id
}

struct CommentListResp {
    1: required BaseResp base_resp
    2: required list<Comment> comment_list
}

service CommentService {
    CommentActionResp CommentAction(1: CommentActionReq req)
    CommentListResp CommentList(1: CommentListReq req)
}


