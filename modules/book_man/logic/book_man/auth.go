package book_man

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	"devinggo/modules/book_man/model/req"
	userService "devinggo/modules/book_man/service"
	consts2 "devinggo/modules/system/consts"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/secure"
	"devinggo/modules/system/service"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sAuth struct {
	base.BaseService
}

func (s *sAuth) Model(ctx context.Context) *gdb.Model {
	//TODO implement me
	panic("implement me")
}

func NewAuth() *sAuth {
	return &sAuth{}
}

func init() {
	userService.RegisterAuth(NewAuth())
}

func (s *sAuth) Login(ctx context.Context, in *req.LoginBody) (token string, expire int64, err error) {
	username := in.Username
	var userList []entity.AppUser
	err = userService.AppUser().Model(ctx).Where(dao.AppUser.Columns().Username, username).Scan(&userList)
	if utils.IsError(err) {
		return
	}
	if len(userList) != 1 {
		err = myerror.ValidationFailed(ctx, "用户记录异常")
		return
	}
	aesKey := config.GetConfigString(ctx, "settings.aesKey")
	if aesKey == "" {
		err = myerror.ValidationFailed(ctx, "系统加密配置异常")
		return
	}

	encryptedPass, err := secure.AESEncrypt(in.Password, aesKey)
	fmt.Println("--------------login-----------------")
	fmt.Println(encryptedPass)
	fmt.Println("==============login================")
	if utils.IsError(err) {
		g.Log().Errorf(ctx, "AES解密失败: %v", err)
		err = myerror.ValidationFailed(ctx, "密码解析失败")
		return
	}

	if encryptedPass != userList[0].PasswordHash {
		err = myerror.ValidationFailed(ctx, "用户或者密码错误")
		return
	}
	status := userList[0].Status
	if status != consts2.UserNormal {
		err = myerror.ValidationFailed(ctx, "用户状态异常")
	}
	appId := contexts.New().GetAppId(ctx)
	token, expire, err = service.Token().GenerateUserToken(ctx, consts2.ApiScene, appId, &model.Identity{
		Id:       userList[0].Id,
		AppId:    appId,
		RoleIds:  nil,
		DeptIds:  nil,
		Username: userList[0].Username,
	})

	if err != nil {
		return
	}
	return

}

func (s *sAuth) Register(ctx context.Context, in *req.RegisterBody) (result bool, err error) {
	var userList []entity.AppUser
	err = userService.AppUser().Model(ctx).Where(dao.AppUser.Columns().Username, in.Username).Scan(&userList)
	if utils.IsError(err) {
		return
	}
	if len(userList) > 0 {
		err = myerror.ValidationFailed(ctx, "用户名已存在")
		return
	}
	aesKey := config.GetConfigString(ctx, "settings.aesKey")
	if aesKey == "" {
		err = myerror.ValidationFailed(ctx, "系统加密配置异常")
		return
	}

	encryptedPass, err := secure.AESEncrypt(in.Password, aesKey)
	if err != nil {
		g.Log().Errorf(ctx, "AES解密失败: %v", err)
		err = myerror.ValidationFailed(ctx, "密码解析失败")
		return false, err
	}
	userService.AppUser().Save(ctx, &req.AppUserSave{
		Username:     in.Username,
		PasswordHash: encryptedPass,
		Nickname:     in.Nickname,
		Status:       consts2.UserNormal,
	})
	return true, nil
}
func (s *sAuth) GetUserInfo(ctx context.Context) (user *entity.AppUser, err error) {
	fromCtx := g.RequestFromCtx(ctx)
	appId := contexts.New().GetAppId(ctx)
	loginUser, err := service.Token().ParseLoginUser(fromCtx, appId)
	userInfo, err := userService.AppUser().GetById(ctx, loginUser.Id)
	if utils.IsError(err) {
		return nil, err
	}
	fmt.Println(loginUser.Username)
	return &entity.AppUser{
		Id:       loginUser.Id,
		Username: loginUser.Username,
		Nickname: userInfo.Nickname,
	}, nil
}
