package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRequest
	ginCtx.ShouldBind(&userReq)
	fmt.Println("pp", userReq)
	//gin.key中去除服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	fmt.Println(userResp)
	PanicTfUserError(err)
	r := res.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)

}

// UserLogin 用户登录
func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicTfUserError(ginCtx.Bind(&userReq))
	//gin.key中去获取服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicTfUserError(err)
	token, err := util.GenerateToken(uint(userResp.UserDetail.UserID))
	r := res.Response{
		Data: res.TokenData{
			User:  userResp.UserDetail,
			Token: token,
		},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
		Error:  err.Error(),
	}
	ginCtx.JSON(http.StatusOK, r)
}
