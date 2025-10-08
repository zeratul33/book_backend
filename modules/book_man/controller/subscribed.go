// Package controller
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
	"context"
	"devinggo/modules/book_man/api"
	"devinggo/modules/book_man/model/req"
	"devinggo/modules/book_man/model/res"
	"devinggo/modules/book_man/service"
	consts2 "devinggo/modules/system/consts"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/contexts"
	service2 "devinggo/modules/system/service"
	"github.com/gogf/gf/v2/util/gutil"

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
	SubscribedController = subscribedController{}
)

type subscribedController struct {
	base.BaseController
}

func (c *subscribedController) Index(ctx context.Context, in *api.IndexSubscribedReq) (out *api.IndexSubscribedRes, err error) {
	out = &api.IndexSubscribedRes{}
	items, totalCount, err := service.Subscribed().GetPageList(ctx, &in.PageListReq, &in.SubscribedSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Subscribed, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *subscribedController) Recycle(ctx context.Context, in *api.RecycleSubscribedReq) (out *api.RecycleSubscribedRes, err error) {
	out = &api.RecycleSubscribedRes{}
	pageListReq := &in.PageListReq
	pageListReq.Recycle = true
	items, totalCount, err := service.Subscribed().GetPageList(ctx, pageListReq, &in.SubscribedSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Subscribed, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *subscribedController) List(ctx context.Context, in *api.ListSubscribedReq) (out *api.ListSubscribedRes, err error) {
	out = &api.ListSubscribedRes{}
	rs, err := service.Subscribed().GetList(ctx, &in.ListReq, &in.SubscribedSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.Subscribed, 0)
	}
	return
}

func (c *subscribedController) Save(ctx context.Context, in *api.SaveSubscribedReq) (out *api.SaveSubscribedRes, err error) {
	out = &api.SaveSubscribedRes{}
	id, err := service.Subscribed().Save(ctx, &in.SubscribedSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *subscribedController) Read(ctx context.Context, in *api.ReadSubscribedReq) (out *api.ReadSubscribedRes, err error) {
	out = &api.ReadSubscribedRes{}
	info, err := service.Subscribed().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *subscribedController) Update(ctx context.Context, in *api.UpdateSubscribedReq) (out *api.UpdateSubscribedRes, err error) {
	out = &api.UpdateSubscribedRes{}
	err = service.Subscribed().Update(ctx, &in.SubscribedUpdate)
	if err != nil {
		return
	}
	return
}

func (c *subscribedController) Delete(ctx context.Context, in *api.DeleteSubscribedReq) (out *api.DeleteSubscribedRes, err error) {
	out = &api.DeleteSubscribedRes{}
	err = service.Subscribed().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
func (c *subscribedController) RealDelete(ctx context.Context, in *api.RealDeleteSubscribedReq) (out *api.RealDeleteSubscribedRes, err error) {
	out = &api.RealDeleteSubscribedRes{}
	err = service.Subscribed().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *subscribedController) Recovery(ctx context.Context, in *api.RecoverySubscribedReq) (out *api.RecoverySubscribedRes, err error) {
	out = &api.RecoverySubscribedRes{}
	err = service.Subscribed().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *subscribedController) ChangeStatus(ctx context.Context, in *api.ChangeStatusSubscribedReq) (out *api.ChangeStatusSubscribedRes, err error) {
	out = &api.ChangeStatusSubscribedRes{}
	err = service.Subscribed().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *subscribedController) NumberOperation(ctx context.Context, in *api.NumberOperationSubscribedReq) (out *api.NumberOperationSubscribedRes, err error) {
	out = &api.NumberOperationSubscribedRes{}
	err = service.Subscribed().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *subscribedController) Export(ctx context.Context, in *api.ExportSubscribedReq) (out *api.ExportSubscribedRes, err error) {
	var (
		fileName  = "subscribed"
		sheetName = "Sheet1"
	)
	exports, err := service.Subscribed().GetExportList(ctx, &in.ListReq, &in.SubscribedSearch)
	if err != nil {
		return
	}
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.SubscribedExcel{})
	//销毁对象
	defer export.Close()
	newExports := []res.SubscribedExcel{}
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

func (c *subscribedController) Import(ctx context.Context, in *api.ImportSubscribedReq) (out *api.ImportSubscribedRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if err != nil {
		return nil, err
	}
	localPath := tmpPath + "/" + fileName
	var result []res.SubscribedExcel
	importFile := excel.NewExcelImportFile(localPath, res.SubscribedExcel{})
	defer importFile.Close()

	err = importFile.ImportDataToStruct(&result).Error()
	if err != nil {
		return nil, err
	} else {
		if !g.IsEmpty(result) {
			err = dao.Subscribed.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				for _, item := range result {
					var saveData *do.Subscribed
					if err = gconv.Struct(item, &saveData); err != nil {
						return
					}
					_, err = service.Subscribed().Model(ctx).OmitEmptyData().Data(saveData).Save()
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

func (c *subscribedController) DownloadTemplate(ctx context.Context, in *api.DownloadTemplateSubscribedReq) (out *api.DownloadTemplateSubscribedRes, err error) {
	var (
		fileName  = "subscribed_template"
		sheetName = "Sheet1"
		exports   = make([]res.SubscribedExcel, 0)
	)
	export := excel.NewExcelExport(sheetName, res.SubscribedExcel{})
	defer export.Close()
	err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *subscribedController) Remote(ctx context.Context, in *api.RemoteSubscribedReq) (out *api.RemoteSubscribedRes, err error) {
	out = &api.RemoteSubscribedRes{}
	r := request.GetHttpRequest(ctx)
	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.Subscribed().Model(ctx)
	var rs res.Subscribed
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
			out.Items = make([]res.Subscribed, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.Subscribed, 0)
		}
	}
	return
}

func (c *subscribedController) SubscribeBook(ctx context.Context, in *api.SubscribeBookReq) (out *api.SubscribeBookRes, err error) {
	out = &api.SubscribeBookRes{Result: false}
	requestFromCtx := g.RequestFromCtx(ctx)
	appId := contexts.New().GetAppId(ctx)
	user, err := service2.Token().ParseLoginUser(requestFromCtx, appId)
	_, err = service.Subscribed().Save(ctx, &req.SubscribedSave{
		SubscribedUser:  user.Id,
		SubscriebedBook: in.Id,
		Status:          consts2.UserNormal,
	})
	if err != nil {
		return
	}
	out.Result = true
	return

}

func (c *subscribedController) GetSubscribedList(ctx context.Context, in *api.GetSubscribedListReq) (out *api.GetSubscribedListRes, err error) {
	out = &api.GetSubscribedListRes{}
	requestFromCtx := g.RequestFromCtx(ctx)
	appId := contexts.New().GetAppId(ctx)
	user, err := service2.Token().ParseLoginUser(requestFromCtx, appId)
	if err != nil {
		return
	}
	list, err := service.Subscribed().GetList(ctx, &model.ListReq{}, &req.SubscribedSearch{SubscribedUser: user.Id})
	if err != nil {
		return
	}
	if !gutil.IsEmpty(list) {
		for _, item := range list {
			out.SubscribedList = append(out.SubscribedList, *item)
		}
	}
	return
}
