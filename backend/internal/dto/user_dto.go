package dto

import "time"

type CreateUserReq struct {
	Username string  `json:"username" binding:"required,max=50"`
	Password string  `json:"password" binding:"required,min=6,max=100"`
	Nickname *string `json:"nickname"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Phone    *string `json:"phone" binding:"omitempty,max=20"`
	Status   int     `json:"status"`
}

type UpdateUserReq struct {
	ID       int64   `json:"id" binding:"required"`
	Nickname *string `json:"nickname"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Phone    *string `json:"phone" binding:"omitempty,max=20"`
	Status   *int    `json:"status"`
}

type UpdatePasswordReq struct {
	ID          int64  `json:"id" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=100"`
}

type QueryUserReq struct {
	Username *string `form:"username"`
	Nickname *string `form:"nickname"`
	Email    *string `form:"email"`
	Phone    *string `form:"phone"`
	Status   *int    `form:"status"`
	Page     int     `form:"page" binding:"omitempty,min=1"`
	PageSize int     `form:"page_size" binding:"omitempty,min=1,max=100"`
}

type UserResp struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Nickname  *string    `json:"nickname"`
	Email     *string    `json:"email"`
	Phone     *string    `json:"phone"`
	Avatar    *string    `json:"avatar"`
	Status    int        `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Roles     []RoleResp `json:"roles"`
}

type RoleResp struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type PageResp struct {
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"page_size"`
	List     []UserResp `json:"list"`
}

type AssignRolesReq struct {
	UserID  int64   `json:"user_id" binding:"required"`
	RoleIDs []int64 `json:"role_ids" binding:"required"`
}
