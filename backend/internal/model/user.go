package model

import (
	"time"
)

type User struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string     `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Password  string     `json:"-" gorm:"type:varchar(255);not null"`
	Nickname  *string    `json:"nickname" gorm:"type:varchar(50)"`
	Email     *string    `json:"email" gorm:"type:varchar(100)"`
	Phone     *string    `json:"phone" gorm:"type:varchar(20)"`
	Avatar    *string    `json:"avatar" gorm:"type:varchar(255)"`
	Status    int        `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}

type Role struct {
	ID          int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" gorm:"type:varchar(50);uniqueIndex;not null"`
	Code        string     `json:"code" gorm:"type:varchar(50);uniqueIndex;not null"`
	Description *string    `json:"description" gorm:"type:varchar(255)"`
	Status      int        `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
}

func (Role) TableName() string {
	return "roles"
}

type Permission struct {
	ID          int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" gorm:"type:varchar(50);not null"`
	Code        string     `json:"code" gorm:"type:varchar(100);uniqueIndex;not null"`
	Description *string    `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
}

func (Permission) TableName() string {
	return "permissions"
}

type UserRole struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int64     `json:"user_id" gorm:"index;not null"`
	RoleID    int64     `json:"role_id" gorm:"index;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

type RolePermission struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID       int64     `json:"role_id" gorm:"index;not null"`
	PermissionID int64     `json:"permission_id" gorm:"index;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
