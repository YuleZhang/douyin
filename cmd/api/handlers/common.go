package handlers

// type NoteParam struct {
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
