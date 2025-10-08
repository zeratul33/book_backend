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
	CategoryController = categoryController{}
)

type categoryController struct {
	base.BaseController
}

func (c *categoryController) Index(ctx context.Context, in *api.IndexCategoryReq) (out *api.IndexCategoryRes, err error) {
	out = &api.IndexCategoryRes{}
	items, totalCount, err := service.Category().GetPageList(ctx, &in.PageListReq, &in.CategorySearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Category, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *categoryController) Recycle(ctx context.Context, in *api.RecycleCategoryReq) (out *api.RecycleCategoryRes, err error) {
	out = &api.RecycleCategoryRes{}
	pageListReq := &in.PageListReq
	pageListReq.Recycle = true
	items, totalCount, err := service.Category().GetPageList(ctx, pageListReq, &in.CategorySearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.Category, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *categoryController) List(ctx context.Context, in *api.ListCategoryReq) (out *api.ListCategoryRes, err error) {
	out = &api.ListCategoryRes{}
	rs, err := service.Category().GetList(ctx, &in.ListReq, &in.CategorySearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.Category, 0)
	}
	return
}

func (c *categoryController) Save(ctx context.Context, in *api.SaveCategoryReq) (out *api.SaveCategoryRes, err error) {
	out = &api.SaveCategoryRes{}
	id, err := service.Category().Save(ctx, &in.CategorySave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *categoryController) Read(ctx context.Context, in *api.ReadCategoryReq) (out *api.ReadCategoryRes, err error) {
	out = &api.ReadCategoryRes{}
	info, err := service.Category().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *categoryController) Update(ctx context.Context, in *api.UpdateCategoryReq) (out *api.UpdateCategoryRes, err error) {
	out = &api.UpdateCategoryRes{}
	err = service.Category().Update(ctx, &in.CategoryUpdate)
	if err != nil {
		return
	}
	return
}

func (c *categoryController) Delete(ctx context.Context, in *api.DeleteCategoryReq) (out *api.DeleteCategoryRes, err error) {
	out = &api.DeleteCategoryRes{}
	err = service.Category().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
func (c *categoryController) RealDelete(ctx context.Context, in *api.RealDeleteCategoryReq) (out *api.RealDeleteCategoryRes, err error) {
	out = &api.RealDeleteCategoryRes{}
	err = service.Category().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *categoryController) Recovery(ctx context.Context, in *api.RecoveryCategoryReq) (out *api.RecoveryCategoryRes, err error) {
	out = &api.RecoveryCategoryRes{}
	err = service.Category().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *categoryController) ChangeStatus(ctx context.Context, in *api.ChangeStatusCategoryReq) (out *api.ChangeStatusCategoryRes, err error) {
	out = &api.ChangeStatusCategoryRes{}
	err = service.Category().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *categoryController) NumberOperation(ctx context.Context, in *api.NumberOperationCategoryReq) (out *api.NumberOperationCategoryRes, err error) {
	out = &api.NumberOperationCategoryRes{}
	err = service.Category().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *categoryController) Export(ctx context.Context, in *api.ExportCategoryReq) (out *api.ExportCategoryRes, err error) {
	var (
		fileName  = "category"
		sheetName = "Sheet1"
	)
	exports, err := service.Category().GetExportList(ctx, &in.ListReq, &in.CategorySearch)
	if err != nil {
		return
	}
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.CategoryExcel{})
	//销毁对象
	defer export.Close()
	newExports := []res.CategoryExcel{}
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

func (c *categoryController) Import(ctx context.Context, in *api.ImportCategoryReq) (out *api.ImportCategoryRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if err != nil {
		return nil, err
	}
	localPath := tmpPath + "/" + fileName
	var result []res.CategoryExcel
	importFile := excel.NewExcelImportFile(localPath, res.CategoryExcel{})
	defer importFile.Close()

	err = importFile.ImportDataToStruct(&result).Error()
	if err != nil {
		return nil, err
	} else {
		if !g.IsEmpty(result) {
			err = dao.Category.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				for _, item := range result {
					var saveData *do.Category
					if err = gconv.Struct(item, &saveData); err != nil {
						return
					}
					_, err = service.Category().Model(ctx).OmitEmptyData().Data(saveData).Save()
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

func (c *categoryController) DownloadTemplate(ctx context.Context, in *api.DownloadTemplateCategoryReq) (out *api.DownloadTemplateCategoryRes, err error) {
	var (
		fileName  = "category_template"
		sheetName = "Sheet1"
		exports   = make([]res.CategoryExcel, 0)
	)
	export := excel.NewExcelExport(sheetName, res.CategoryExcel{})
	defer export.Close()
	err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *categoryController) Remote(ctx context.Context, in *api.RemoteCategoryReq) (out *api.RemoteCategoryRes, err error) {
	out = &api.RemoteCategoryRes{}
	r := request.GetHttpRequest(ctx)
	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.Category().Model(ctx)
	var rs res.Category
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
			out.Items = make([]res.Category, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.Category, 0)
		}
	}
	return
}
func (c *categoryController) GetCategoryList(ctx context.Context, in *api.GetCategoryListReq) (out *api.GetCategoryListRes, err error) {
	out = &api.GetCategoryListRes{}
	categorys, err := service.Category().GetList(ctx, &model.ListReq{}, &req.CategorySearch{})
	if err != nil {
		return
	}
	if !g.IsEmpty(categorys) {
		for _, category := range categorys {
			out.Categorys = append(out.Categorys, *category)
		}
	} else {
		out.Categorys = make([]res.Category, 0)
	}
	return

}
