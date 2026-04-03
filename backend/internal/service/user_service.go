package service

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(req dto.CreateUserReq) error {
	exists, _ := s.userRepo.FindByUsername(req.Username)
	if exists != nil {
		return errors.New("用户名已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   req.Status,
	}

	if user.Status == 0 {
		user.Status = 1
	}

	return s.userRepo.Create(user)
}

func (s *UserService) Update(req dto.UpdateUserReq) error {
	user, err := s.userRepo.Detail(req.ID)
	if err != nil {
		return errors.New("用户不存在")
	}

	if req.Nickname != nil {
		user.Nickname = req.Nickname
	}
	if req.Email != nil {
		user.Email = req.Email
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.Status != nil {
		user.Status = *req.Status
	}

	return s.userRepo.Update(user)
}

func (s *UserService) Delete(id int64) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) Detail(id int64) (*dto.UserResp, error) {
	user, err := s.userRepo.Detail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	roles, _ := s.userRepo.GetUserRoles(id)
	roleResps := make([]dto.RoleResp, len(roles))
	for i, role := range roles {
		roleResps[i] = dto.RoleResp{
			ID:   role.ID,
			Name: role.Name,
			Code: role.Code,
		}
	}

	return &dto.UserResp{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Roles:     roleResps,
	}, nil
}

func (s *UserService) List(req dto.QueryUserReq) (*dto.PageResp, error) {
	users, total, err := s.userRepo.List(req)
	if err != nil {
		return nil, err
	}

	list := make([]dto.UserResp, len(users))
	for i, user := range users {
		roles, _ := s.userRepo.GetUserRoles(user.ID)
		roleResps := make([]dto.RoleResp, len(roles))
		for j, role := range roles {
			roleResps[j] = dto.RoleResp{
				ID:   role.ID,
				Name: role.Name,
				Code: role.Code,
			}
		}
		list[i] = dto.UserResp{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Roles:     roleResps,
		}
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}

	return &dto.PageResp{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		List:     list,
	}, nil
}

func (s *UserService) UpdatePassword(req dto.UpdatePasswordReq) error {
	user, err := s.userRepo.Detail(req.ID)
	if err != nil {
		return errors.New("用户不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("原密码错误")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Update(user)
}

func (s *UserService) AssignRoles(req dto.AssignRolesReq) error {
	return s.userRepo.AssignRoles(req.UserID, req.RoleIDs)
}
