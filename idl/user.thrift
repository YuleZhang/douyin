namespace go user

struct User {
    1:required i64 id
    2:required string name
    3:optional i64 follow_count # 关注总数
    4:optional i64 follower_count # 粉丝总数
    5:required bool is_follow
}

/*
考虑再三还是决定加上BaseResp，作为基础消息体，它的作用有两个
1. 作为基础类型内嵌到其他结构中
2. 当没有合适返回消息体时，作为默认选项（例如发生Error）
*/
struct BaseResp {
    1:required i32 status_code
    2:required string status_msg
}

struct DouyinUserRegisterRequest {
    1:required string username
    2:required string password
}

struct DouyinUserRegisterResponse {
    1:BaseResp base_resp
    2:required i64 user_id
    3:required string token
}

struct DouyinUserRequest {
    1:required i64 user_id
    2:required string token
}

struct DouyinUserResponse {
    1:BaseResp base_resp
    2:list<User> user
}

struct DouyinUserLoginRequest {
    1:required string username
    2:required string password
}

struct DouyinUserLoginResponse {
    1:BaseResp base_resp
    2:required i64 user_id
    3:required string token
}

service UserService {
    DouyinUserRegisterResponse DouyinUserRegister(1:DouyinUserRegisterRequest req)
    DouyinUserResponse DouyinUser(1:DouyinUserRequest req)
    DouyinUserLoginResponse DouyinUserLogin(1:DouyinUserLoginRequest req)
}