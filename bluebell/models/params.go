package models

// 定义请求的参数结构体

type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID string `json:"post_id" binding:"required"` // 帖子id
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"` // 增长票(1)还是反对票（-1）取消票0
}