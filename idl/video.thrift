namespace go video
include "user.thrift"

struct Video{
    1:required i64 id
    2:required user.User author
    3:required string play_url
    4:required string cover_url
    5:required i64 favorite_count
    6:required i64 comment_count
    7:required bool is_favorite
    8:required string title
}

struct DouyinFeedRequest {
    1:optional i64 latest_time
    2:optional string token
}

struct DouyinFeedResponse {
    1:user.BaseResp base_resp
    2:list<Video> video_list
    3:optional i64 next_time
}

struct DouyinPublishActionRequest {
    1:required string token
    2:Video video
}

struct DouyinPublishActionResponse {
    1:user.BaseResp base_resp
}

struct DouyinPublishListRequest {
    1:required i64 user_id
    2:required string token
}

struct DouyinPublishListResponse {
    1:user.BaseResp base_resp
    2:list<Video> video_list
}

struct DouyinFavoriteActionRequest {
    1:required i64 user_id
    2:required string token
    3:required i64 video_id
    4:required i32 action_type
}

struct DouyinFavoriteListRequest {
    1:required i64 user_id
    2:required string token
}

struct DouyinFavoriteListResponse {
    1:user.BaseResp base_resp
    2:list<Video> video_list
}

service VideoService {
    DouyinFeedResponse DouyinFeed(1:DouyinFeedRequest req) # 视频流服务
    DouyinPublishActionResponse DouyinPublishAction(1:DouyinPublishActionRequest req) # 视频投稿服务
    DouyinPublishListResponse DouyinPublishList(1:DouyinPublishListRequest req) # 发布列表服务
    DouyinFavoriteListResponse DouyinFavoriteList(1:DouyinFavoriteListRequest req) # 视频点赞列表
}