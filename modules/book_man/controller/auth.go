package controller

import (
	"context"
	"devinggo/modules/book_man/api"
	"devinggo/modules/book_man/model/res"
	userService "devinggo/modules/book_man/service"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/request"
	service2 "devinggo/modules/system/service"
	"os/exec"
)

type cAuthController struct {
	base.BaseController
}

var AuthController = &cAuthController{}

func (c *cAuthController) Login(ctx context.Context, in *api.LoginReq) (out *api.LoginRes, err error) {
	token, expire, err := userService.Auth().Login(ctx, &in.LoginBody)
	if utils.IsError(err) {
		return &api.LoginRes{
			Message: err.Error(),
			Success: false,
		}, nil
	}
	out = &api.LoginRes{
		Token:     token,
		ExpiresIn: expire,
		Success:   true,
	}
	return
}

func (c *cAuthController) Register(ctx context.Context, in *api.RegisterReq) (out *api.RegisterRes, err error) {
	result, err := userService.Auth().Register(ctx, &in.RegisterBody)
	return &api.RegisterRes{
		Result: result,
	}, err
}

func (c *cAuthController) GetUserInfo(ctx context.Context, in *api.GetUserInfoReq) (out *api.GetUserInfoRes, err error) {
	user, err := userService.Auth().GetUserInfo(ctx)
	if utils.IsError(err) {
		return nil, err
	}
	return &api.GetUserInfoRes{
		UserInfo: res.AppUser{
			Username: user.Username,
			Nickname: user.Nickname,
			Id:       user.Id,
		},
	}, nil
}

func (c *cAuthController) Logout(ctx context.Context, in *api.LogoutReq) (out *api.LogoutRes, err error) {
	err = service2.Token().Logout(request.GetHttpRequest(ctx))
	if utils.IsError(err) {
		return
	}
	return &api.LogoutRes{
		Result: true,
	}, nil
}

func (c *cAuthController) StartScript(ctx context.Context, in *api.StartScriptReq) (out *api.StartScriptRes, err error) {
	err = exec.Command("python", "/Users/zeratul/Downloads/文本摘要.py").Start()
	if err != nil {
		return nil, err
	}
	return
}
