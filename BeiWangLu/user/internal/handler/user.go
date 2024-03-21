package handler

import (
	"context"
	"fmt"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

//
//var UserSrvIns *UserSrv
//var UserSrvOnce sync.Once
//
//type UserSrv struct {
//	service.UnimplementedUserServiceServer
//}
//
//func GetUserSrv() *UserSrv {
//	UserSrvOnce.Do(func() {
//		UserSrvIns = &UserSrv{}
//	})
//	return UserSrvIns
//}

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// UserLogin 用户登录 带token
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

// 用户注册
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	fmt.Println("Before UserCreat, req:", req)
	user, err = user.UserCreat(req)
	fmt.Println("After UserCreat, user:", user)
	fmt.Println(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	fmt.Println("ppp", user)
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}
