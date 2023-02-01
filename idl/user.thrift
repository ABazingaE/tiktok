namespace go user
/*
* 定义状态码
* */
enum StatusCode {
    UserNotExist                    = 10004
    AuthorizationFailedErrCode      = 10005
    UserExist                       = 10006
}


/*
* 定义结构体
* */

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
* 定义请求与响应
* */

//注册
struct RegisterReq {
    1: required string username (vt.min_size = "1",vt.max_size = "32")
    2: required string password (vt.min_size = "1",vt.max_size = "32")
}

struct RegisterResp {
    1: required BaseResp base_resp
    2: required i64 user_id
    3: required string token
}


//登录
struct LoginReq {
    1: required string username (vt.min_size = "1",vt.max_size = "32")
    2: required string password (vt.min_size = "1",vt.max_size = "32")
}

struct LoginResp {
    1: required BaseResp base_resp
    2: required i64 user_id
    3: required string token
}


//用户信息
struct UserInfoReq {
    1: required i64 user_id
    2: required string token
}

struct UserInfoResp {
    1: required BaseResp base_resp
    2: required User user
}

/*
* 定义服务
* */

service UserService {
    RegisterResp Register(1: RegisterReq req)
    LoginResp Login(1: LoginReq req)
    UserInfoResp UserInfo(1: UserInfoReq req)
}