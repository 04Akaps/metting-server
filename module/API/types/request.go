package types

type RegisterUser struct {
	UserName    string   `json:"userName" binding:"required"`
	Description string   `json:"description" binding:"required,max=300"`
	Hobby       []string `json:"hobby" binding:"required"`
	Latitude    float64  `json:"latitude" binding:"required,min=-90,max=90"`   // 위도
	Hardness    float64  `json:"hardness" binding:"required,min=-180,max=180"` // 경도
}

type FindAroundFriendsReq struct {
	User  string `form:"user" binding:"required"`
	Range int64  `form:"range" binding:"required"`
	Limit int64  `form:"limit"`
}

type LikeRequest struct {
	FromUser string `json:"fromUser" binding:"required"`
	ToUser   string `json:"toUser" binding:"required"`
}

type RefuseRequest struct {
	FromUser string `json:"fromUser" binding:"required"`
	ToUser   string `json:"toUser" binding:"required"`
}

type AcceptedRequest struct {
	FromUser string `json:"fromUser" binding:"required"`
	ToUser   string `json:"toUser" binding:"required"`
}

type GetLikedListRequest struct {
	User string `form:"user" binding:"required"`
}
