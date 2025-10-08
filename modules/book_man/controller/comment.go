// Package controller
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/modules/book_man/api"
	"devinggo/modules/book_man/model/res"
	"devinggo/modules/book_man/service"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/excel"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/request"
	sysService "devinggo/modules/system/service"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	CommentController = commentController{}
)

type commentController struct {
	base.BaseController
}

func (c *commentController) Index(ctx context.Context, in *api.IndexCommentReq) (out *api.IndexCommentRes, err error) {
	out = &api.IndexCommentRes{}
	items, totalCount, err := service.Comment().GetPageList(ctx, &in.PageListReq, &in.CommentSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Comment, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *commentController) Recycle(ctx context.Context, in *api.RecycleCommentReq) (out *api.RecycleCommentRes, err error) {
	out = &api.RecycleCommentRes{}
	pageListReq := &in.PageListReq
	pageListReq.Recycle = true
	items, totalCount, err := service.Comment().GetPageList(ctx, pageListReq, &in.CommentSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Comment, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *commentController) List(ctx context.Context, in *api.ListCommentReq) (out *api.ListCommentRes, err error) {
	out = &api.ListCommentRes{}
	rs, err := service.Comment().GetList(ctx, &in.ListReq, &in.CommentSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.Comment, 0)
	}
	return
}

func (c *commentController) Save(ctx context.Context, in *api.SaveCommentReq) (out *api.SaveCommentRes, err error) {
	out = &api.SaveCommentRes{}
	id, err := service.Comment().Save(ctx, &in.CommentSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *commentController) Read(ctx context.Context, in *api.ReadCommentReq) (out *api.ReadCommentRes, err error) {
	out = &api.ReadCommentRes{}
	info, err := service.Comment().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *commentController) Update(ctx context.Context, in *api.UpdateCommentReq) (out *api.UpdateCommentRes, err error) {
	out = &api.UpdateCommentRes{}
	err = service.Comment().Update(ctx, &in.CommentUpdate)
	if err != nil {
		return
	}
	return
}

func (c *commentController) Delete(ctx context.Context, in *api.DeleteCommentReq) (out *api.DeleteCommentRes, err error) {
	out = &api.DeleteCommentRes{}
	err = service.Comment().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
func (c *commentController) RealDelete(ctx context.Context, in *api.RealDeleteCommentReq) (out *api.RealDeleteCommentRes, err error) {
	out = &api.RealDeleteCommentRes{}
	err = service.Comment().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *commentController) Recovery(ctx context.Context, in *api.RecoveryCommentReq) (out *api.RecoveryCommentRes, err error) {
	out = &api.RecoveryCommentRes{}
	err = service.Comment().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *commentController) ChangeStatus(ctx context.Context, in *api.ChangeStatusCommentReq) (out *api.ChangeStatusCommentRes, err error) {
	out = &api.ChangeStatusCommentRes{}
	err = service.Comment().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *commentController) NumberOperation(ctx context.Context, in *api.NumberOperationCommentReq) (out *api.NumberOperationCommentRes, err error) {
	out = &api.NumberOperationCommentRes{}
	err = service.Comment().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *commentController) Export(ctx context.Context, in *api.ExportCommentReq) (out *api.ExportCommentRes, err error) {
	var (
		fileName  = "comment"
		sheetName = "Sheet1"
	)
	exports, err := service.Comment().GetExportList(ctx, &in.ListReq, &in.CommentSearch)
	if err != nil {
		return
	}
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.CommentExcel{})
	//销毁对象
	defer export.Close()
	newExports := []res.CommentExcel{}
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

func (c *commentController) Import(ctx context.Context, in *api.ImportCommentReq) (out *api.ImportCommentRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if err != nil {
		return nil, err
	}
	localPath := tmpPath + "/" + fileName
	var result []res.CommentExcel
	importFile := excel.NewExcelImportFile(localPath, res.CommentExcel{})
	defer importFile.Close()

	err = importFile.ImportDataToStruct(&result).Error()
	if err != nil {
		return nil, err
	} else {
		if !g.IsEmpty(result) {
			err = dao.Comment.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				for _, item := range result {
					var saveData *do.Comment
					if err = gconv.Struct(item, &saveData); err != nil {
						return
					}
					_, err = service.Comment().Model(ctx).OmitEmptyData().Data(saveData).Save()
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

func (c *commentController) DownloadTemplate(ctx context.Context, in *api.DownloadTemplateCommentReq) (out *api.DownloadTemplateCommentRes, err error) {
	var (
		fileName  = "comment_template"
		sheetName = "Sheet1"
		exports   = make([]res.CommentExcel, 0)
	)
	export := excel.NewExcelExport(sheetName, res.CommentExcel{})
	defer export.Close()
	err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *commentController) Remote(ctx context.Context, in *api.RemoteCommentReq) (out *api.RemoteCommentRes, err error) {
	out = &api.RemoteCommentRes{}
	r := request.GetHttpRequest(ctx)
	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.Comment().Model(ctx)
	var rs res.Comment
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
			out.Items = make([]res.Comment, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.Comment, 0)
		}
	}
	return
}

func (c *commentController) GetCommentByBook(ctx context.Context, in *api.GetCommentByBookReq) (out *api.GetCommentByBookRes, err error) {
	out = &api.GetCommentByBookRes{}
	r := g.RequestFromCtx(ctx)
	fmt.Println(r.Get("id"))
	comments, err := service.Comment().GetCommentByBook(ctx, r.Get("id").Int64())
	if !g.IsEmpty(comments) {
		for _, item := range comments {
			out.Comments = append(out.Comments, *item)
		}
	} else {
		out.Comments = make([]res.CommentApp, 0)
	}
	return
}

func (c *commentController) GetCommentByUser(ctx context.Context, in *api.GetCommentListByUserReq) (out *api.GetCommentListByUserRes, err error) {
	out = &api.GetCommentListByUserRes{}
	comments, err := service.Comment().GetCommentByUser(ctx)
	if utils.IsError(err) {
		return
	}
	if !g.IsEmpty(comments) {
		for _, item := range comments {
			out.Comments = append(out.Comments, *item)
		}
	} else {
		out.Comments = make([]res.CommentApp, 0)
	}
	return
}

func (c *commentController) PublishComment(ctx context.Context, in *api.PublishCommentReq) (out *api.PublishCommentRes, err error) {
	fmt.Println(in.Comment)
	fromCtx := g.RequestFromCtx(ctx)
	appId := contexts.New().GetAppId(ctx)
	loginUser, err := sysService.Token().ParseLoginUser(fromCtx, appId)
	result, err := service.Comment().Model(ctx).Data(&do.Comment{
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
		UserId:      loginUser.Id,
		UserComment: in.Comment,
		CommentTime: gtime.Now(),
		Status:      1,
		BookId:      in.BookId,
	}).FieldsEx("id", "created_by", "updated_by").Insert()
	if err != nil {
		return nil, err
	}
	gutil.Dump(result)
	return
}
