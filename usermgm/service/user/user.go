package user

import (
	"context"
	userpb "quiz/gunk/v1/user"
	"quiz/usermgm/storage"
)

type CoreUser interface {
	RegisterUser(u storage.User) (*storage.User, error)
	RegisterAdmin(storage.User) (*storage.User, error)
	Login(storage.Login) (*storage.User, error)
	UserList() ([]storage.User, error)
	DeleteUser(id int32) error
	GetUserIdForEdit(id int) (*storage.User, error)
	UpdateUser(u storage.User) (*storage.User, error)
}

type UserSvc struct {
	userpb.UnimplementedUserServiceServer
	core CoreUser
}

func NewUserSvc(cu CoreUser) *UserSvc {
	return &UserSvc{
		core: cu,
	}
}

func (us UserSvc) RegisterUser(ctx context.Context, r *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	user := storage.User{
		ID:        0,
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
	}

	if err := user.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.core.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterUserResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsAdmin:   u.IsAdmin,
			IsActive:  u.IsActive,
		},
	}, nil
}
func (us UserSvc) RegisterAdmin(ctx context.Context, r *userpb.RegisterAdminRequest) (*userpb.RegisterAdminResponse, error) {
	user := storage.User{

		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
	}

	if err := user.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.core.RegisterAdmin(user)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterAdminResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsAdmin:   true,
			IsActive:  true,
		},
	}, nil
}

func (us UserSvc) UserDelete(ctx context.Context, r *userpb.UserDeleteRequest) (*userpb.UserDeleteResponse, error) {
	err := us.core.DeleteUser(r.ID)
	if err != nil {
		return nil, err
	}
	return &userpb.UserDeleteResponse{}, nil
}
func (cu UserSvc) UserEdit(ctx context.Context, r *userpb.UserEditRequest) (*userpb.UserEditResponse, error) {
	ru, err := cu.core.GetUserIdForEdit(int(r.Id))
	if err != nil {
		return nil, err
	}

	return &userpb.UserEditResponse{
		User: &userpb.User{
			ID:        int32(ru.ID),
			FirstName: ru.FirstName,
			LastName:  ru.LastName,
			Username:  ru.Username,
			Email:     ru.Email,
			IsAdmin:   ru.IsAdmin,
			IsActive:  ru.IsActive,
		},
	}, err
}

func (cu UserSvc) UserUpdate(ctx context.Context, r *userpb.UserUpdateRequest) (*userpb.UserUpdateResponse, error) {
	user := storage.User{
		ID:        int(r.ID),
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Username:  r.Username,
		IsAdmin:   r.IsAdmin,
		IsActive:  r.IsActive,
	}

	_, err := cu.core.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &userpb.UserUpdateResponse{}, nil
}
func (us UserSvc) UserList(ctx context.Context, r *userpb.UserListRequest) (*userpb.UserListResponse, error) {
	userList, err := us.core.UserList()
	if err != nil {
		return nil, err
	}

	users := make([]*userpb.User, len(userList))
	for i, user := range userList {
		users[i] = &userpb.User{
			ID:        int32(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Email:     user.Email,
			IsAdmin:   user.IsAdmin,
			IsActive:  user.IsActive,
		}
	}

	return &userpb.UserListResponse{
		User: users,
	}, nil
}

func (us UserSvc) Login(ctx context.Context, r *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	login := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		return nil, err
	}

	u, err := us.core.Login(login)
	if err != nil {
		return nil, err
	}

	return &userpb.LoginResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsAdmin:   u.IsAdmin,
			IsActive:  u.IsActive,
		},
	}, nil
}
