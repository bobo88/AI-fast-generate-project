package repository

import (
	"backend/internal/dto"
	"backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id int64) error {
	return r.db.Delete(&model.User{}, "id = ?", id).Error
}

func (r *UserRepository) Detail(id int64) (*model.User, error) {
	var user model.User
	err := r.db.Preload("UserRoles.Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) List(req dto.QueryUserReq) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.Model(&model.User{})

	if req.Username != nil && *req.Username != "" {
		query = query.Where("username LIKE ?", "%"+*req.Username+"%")
	}
	if req.Nickname != nil && *req.Nickname != "" {
		query = query.Where("nickname LIKE ?", "%"+*req.Nickname+"%")
	}
	if req.Email != nil && *req.Email != "" {
		query = query.Where("email = ?", *req.Email)
	}
	if req.Phone != nil && *req.Phone != "" {
		query = query.Where("phone = ?", *req.Phone)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	query.Count(&total)

	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&users).Error

	return users, total, err
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) AssignRoles(userID int64, roleIDs []int64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		for _, roleID := range roleIDs {
			if err := tx.Create(&model.UserRole{UserID: userID, RoleID: roleID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *UserRepository) GetUserRoles(userID int64) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}
