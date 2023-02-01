namespace go friend

struct BaseResp {
    1: i64 status_code
    2: string status_msg
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
* 粉丝列表
* */
struct FriendListReq{
    1: required i64 user_id
    2: required string token
}

struct FriendListResp{
    1: required BaseResp base_resp
    2: required list<FriendUser> user_list
}

/*
* 聊天记录
* */
struct MessageChatReq{
    1: required string token
    2: required i64 to_user_id
}

struct MessageChatResp{
    1: required BaseResp base_resp
    2: required list<Message> message_list
}

/*
* 消息操作
* */
struct MessageActionReq{
    1: required string token
    2: required i64 to_user_id
    3: required i64 action_type
    4: required string content
}

struct MessageActionResp{
    1: required BaseResp base_resp
}

service FriendService{
    FriendListResp FriendList(1:FriendListReq req)
    MessageChatResp MessageChat(1:MessageChatReq req)
    MessageActionResp MessageAction(1:MessageActionReq req)
}
