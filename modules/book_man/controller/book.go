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
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model"
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
	BookController = bookController{}
)

type bookController struct {
	base.BaseController
}

func (c *bookController) Index(ctx context.Context, in *api.IndexBookReq) (out *api.IndexBookRes, err error) {
	out = &api.IndexBookRes{}
	items, totalCount, err := service.Book().GetPageList(ctx, &in.PageListReq, &in.BookSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Book, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *bookController) Recycle(ctx context.Context, in *api.RecycleBookReq) (out *api.RecycleBookRes, err error) {
	out = &api.RecycleBookRes{}
	pageListReq := &in.PageListReq
	pageListReq.Recycle = true
	items, totalCount, err := service.Book().GetPageList(ctx, pageListReq, &in.BookSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Book, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *bookController) List(ctx context.Context, in *api.ListBookReq) (out *api.ListBookRes, err error) {
	out = &api.ListBookRes{}
	rs, err := service.Book().GetList(ctx, &in.ListReq, &in.BookSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.Book, 0)
	}
	return
}

func (c *bookController) Save(ctx context.Context, in *api.SaveBookReq) (out *api.SaveBookRes, err error) {
	out = &api.SaveBookRes{}
	id, err := service.Book().Save(ctx, &in.BookSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *bookController) Read(ctx context.Context, in *api.ReadBookReq) (out *api.ReadBookRes, err error) {
	out = &api.ReadBookRes{}
	info, err := service.Book().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *bookController) Update(ctx context.Context, in *api.UpdateBookReq) (out *api.UpdateBookRes, err error) {
	out = &api.UpdateBookRes{}
	err = service.Book().Update(ctx, &in.BookUpdate)
	if err != nil {
		return
	}
	return
}

func (c *bookController) Delete(ctx context.Context, in *api.DeleteBookReq) (out *api.DeleteBookRes, err error) {
	out = &api.DeleteBookRes{}
	err = service.Book().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
func (c *bookController) RealDelete(ctx context.Context, in *api.RealDeleteBookReq) (out *api.RealDeleteBookRes, err error) {
	out = &api.RealDeleteBookRes{}
	err = service.Book().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *bookController) Recovery(ctx context.Context, in *api.RecoveryBookReq) (out *api.RecoveryBookRes, err error) {
	out = &api.RecoveryBookRes{}
	err = service.Book().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *bookController) ChangeStatus(ctx context.Context, in *api.ChangeStatusBookReq) (out *api.ChangeStatusBookRes, err error) {
	out = &api.ChangeStatusBookRes{}
	err = service.Book().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *bookController) NumberOperation(ctx context.Context, in *api.NumberOperationBookReq) (out *api.NumberOperationBookRes, err error) {
	out = &api.NumberOperationBookRes{}
	err = service.Book().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *bookController) Export(ctx context.Context, in *api.ExportBookReq) (out *api.ExportBookRes, err error) {
	var (
		fileName  = "book"
		sheetName = "Sheet1"
	)
	exports, err := service.Book().GetExportList(ctx, &in.ListReq, &in.BookSearch)
	if err != nil {
		return
	}
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.BookExcel{})
	//销毁对象
	defer export.Close()
	newExports := []res.BookExcel{}
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

func (c *bookController) Import(ctx context.Context, in *api.ImportBookReq) (out *api.ImportBookRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if err != nil {
		return nil, err
	}
	localPath := tmpPath + "/" + fileName
	var result []res.BookExcel
	importFile := excel.NewExcelImportFile(localPath, res.BookExcel{})
	defer importFile.Close()

	err = importFile.ImportDataToStruct(&result).Error()
	if err != nil {
		return nil, err
	} else {
		if !g.IsEmpty(result) {
			err = dao.Book.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				for _, item := range result {
					var saveData *do.Book
					if err = gconv.Struct(item, &saveData); err != nil {
						return
					}
					_, err = service.Book().Model(ctx).OmitEmptyData().Data(saveData).Save()
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

func (c *bookController) DownloadTemplate(ctx context.Context, in *api.DownloadTemplateBookReq) (out *api.DownloadTemplateBookRes, err error) {
	var (
		fileName  = "book_template"
		sheetName = "Sheet1"
		exports   = make([]res.BookExcel, 0)
	)
	export := excel.NewExcelExport(sheetName, res.BookExcel{})
	defer export.Close()
	err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *bookController) Remote(ctx context.Context, in *api.RemoteBookReq) (out *api.RemoteBookRes, err error) {
	out = &api.RemoteBookRes{}
	r := request.GetHttpRequest(ctx)
	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.Book().Model(ctx)
	var rs res.Book
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
			out.Items = make([]res.Book, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.Book, 0)
		}
	}
	return
}

func (c *bookController) GetBookList(ctx context.Context, in *api.GetBookListReq) (out *api.GetBookListRes, err error) {
	out = &api.GetBookListRes{}
	list, err := service.Book().GetList(ctx, &model.ListReq{}, &req.BookSearch{})
	if !g.IsEmpty(list) {
		for _, item := range list {
			out.Books = append(out.Books, *item)
		}

	} else {
		out.Books = make([]res.Book, 0)
	}
	return
}

func (c *bookController) GetBookById(ctx context.Context, in *api.GetBookByIdReq) (out *api.GetBookByIdRes, err error) {
	out = &api.GetBookByIdRes{}
	fromCtx := g.RequestFromCtx(ctx)
	id := fromCtx.Get("id")
	book, err := service.Book().GetById(ctx, id.Int64())
	if err != nil {
		return
	}
	out.Book = *book
	return
}
