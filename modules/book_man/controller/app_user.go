// Package controller
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
	"context"
	"devinggo/modules/book_man/api"
	"devinggo/modules/book_man/model/res"
	"devinggo/modules/book_man/service"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils/request"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"

	"devinggo/modules/system/pkg/excel"

	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	AppUserController = appUserController{}
)

type appUserController struct {
	base.BaseController
}

func (c *appUserController) Index(ctx context.Context, in *api.IndexAppUserReq) (out *api.IndexAppUserRes, err error) {
	out = &api.IndexAppUserRes{}
	items, totalCount, err := service.AppUser().GetPageList(ctx, &in.PageListReq, &in.AppUserSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.AppUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *appUserController) Recycle(ctx context.Context, in *api.RecycleAppUserReq) (out *api.RecycleAppUserRes, err error) {
	out = &api.RecycleAppUserRes{}
	pageListReq := &in.PageListReq
	pageListReq.Recycle = true
	items, totalCount, err := service.AppUser().GetPageList(ctx, pageListReq, &in.AppUserSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.AppUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *appUserController) List(ctx context.Context, in *api.ListAppUserReq) (out *api.ListAppUserRes, err error) {
	out = &api.ListAppUserRes{}
	rs, err := service.AppUser().GetList(ctx, &in.ListReq, &in.AppUserSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.AppUser, 0)
	}
	return
}

func (c *appUserController) Save(ctx context.Context, in *api.SaveAppUserReq) (out *api.SaveAppUserRes, err error) {
	out = &api.SaveAppUserRes{}
	id, err := service.AppUser().Save(ctx, &in.AppUserSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *appUserController) Read(ctx context.Context, in *api.ReadAppUserReq) (out *api.ReadAppUserRes, err error) {
	out = &api.ReadAppUserRes{}
	info, err := service.AppUser().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *appUserController) Update(ctx context.Context, in *api.UpdateAppUserReq) (out *api.UpdateAppUserRes, err error) {
	out = &api.UpdateAppUserRes{}
	err = service.AppUser().Update(ctx, &in.AppUserUpdate)
	if err != nil {
		return
	}
	return
}

func (c *appUserController) Delete(ctx context.Context, in *api.DeleteAppUserReq) (out *api.DeleteAppUserRes, err error) {
	out = &api.DeleteAppUserRes{}
	err = service.AppUser().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
func (c *appUserController) RealDelete(ctx context.Context, in *api.RealDeleteAppUserReq) (out *api.RealDeleteAppUserRes, err error) {
	out = &api.RealDeleteAppUserRes{}
	err = service.AppUser().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *appUserController) Recovery(ctx context.Context, in *api.RecoveryAppUserReq) (out *api.RecoveryAppUserRes, err error) {
	out = &api.RecoveryAppUserRes{}
	err = service.AppUser().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *appUserController) ChangeStatus(ctx context.Context, in *api.ChangeStatusAppUserReq) (out *api.ChangeStatusAppUserRes, err error) {
	out = &api.ChangeStatusAppUserRes{}
	err = service.AppUser().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *appUserController) NumberOperation(ctx context.Context, in *api.NumberOperationAppUserReq) (out *api.NumberOperationAppUserRes, err error) {
	out = &api.NumberOperationAppUserRes{}
	err = service.AppUser().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *appUserController) Export(ctx context.Context, in *api.ExportAppUserReq) (out *api.ExportAppUserRes, err error) {
	var (
		fileName  = "appUser"
		sheetName = "Sheet1"
	)
	exports, err := service.AppUser().GetExportList(ctx, &in.ListReq, &in.AppUserSearch)
	if err != nil {
		return
	}
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.AppUserExcel{})
	//销毁对象
	defer export.Close()
	newExports := []res.AppUserExcel{}
	if !g.IsEmpty(exports) {
		for _, item := range exports {
			newExports = append(newExports, *item)
		}
	}
	err = export.ExportSmallExcelByStruct(newExports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *appUserController) Import(ctx context.Context, in *api.ImportAppUserReq) (out *api.ImportAppUserRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if err != nil {
		return nil, err
	}
	localPath := tmpPath + "/" + fileName
	var result []res.AppUserExcel
	importFile := excel.NewExcelImportFile(localPath, res.AppUserExcel{})
	defer importFile.Close()

	err = importFile.ImportDataToStruct(&result).Error()
	if err != nil {
		return nil, err
	} else {
		if !g.IsEmpty(result) {
			err = dao.AppUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				for _, item := range result {
					var saveData *do.AppUser
					if err = gconv.Struct(item, &saveData); err != nil {
						return
					}
					_, err = service.AppUser().Model(ctx).OmitEmptyData().Data(saveData).Save()
					if err != nil {
						return err
					}
				}
				return
			})
			if err != nil {
				return
			}
		} else {
			err = myerror.ValidationFailed(ctx, "没有数据!")
		}
	}
	return
}

func (c *appUserController) DownloadTemplate(ctx context.Context, in *api.DownloadTemplateAppUserReq) (out *api.DownloadTemplateAppUserRes, err error) {
	var (
		fileName  = "appUser_template"
		sheetName = "Sheet1"
		exports   = make([]res.AppUserExcel, 0)
	)
	export := excel.NewExcelExport(sheetName, res.AppUserExcel{})
	defer export.Close()
	err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *appUserController) Remote(ctx context.Context, in *api.RemoteAppUserReq) (out *api.RemoteAppUserRes, err error) {
	out = &api.RemoteAppUserRes{}
	r := request.GetHttpRequest(ctx)
	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.AppUser().Model(ctx)
	var rs res.AppUser
	remote := orm.NewRemote(m, rs)
	openPage := params.GetVar("openPage")
	items, totalCount, err := remote.GetRemote(ctx, params)
	if err != nil {
		return
	}
	if !g.IsEmpty(openPage) && openPage.Bool() {
		if !g.IsEmpty(items) {
			for _, item := range items {
				out.Items = append(out.Items, item)
			}
		} else {
			out.Items = make([]res.AppUser, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.AppUser, 0)
		}
	}
	return
}
