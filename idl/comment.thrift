namespace go comment

struct Comment {
    1:required i64 id
    2:required User user
    3:required string content
    4:required string create_date
}

struct DouyinCommentListRequest {
    1:required string token
    2:required i64 video_id
}

struct DouyinCommentListReponse {
    
}