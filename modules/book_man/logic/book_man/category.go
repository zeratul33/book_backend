// Package book_man
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package book_man

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/book_man/model/req"
	"devinggo/modules/book_man/model/res"
	"devinggo/modules/book_man/service"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCategory struct {
	base.BaseService
}

func init() {
	service.RegisterCategory(NewCategory())
}

func NewCategory() *sCategory {
	return &sCategory{}
}

func (s *sCategory) Model(ctx context.Context) *gdb.Model {
	doObj := do.Category{}
	var params []string
	if utils.HasField(doObj, "CreatedBy") {
		params = append(params, "created_by")
	}
	if utils.HasField(doObj, "UpdatedBy") {
		params = append(params, "updated_by")
	}
	hookOptions := &hook.HookOptions{
		Params: params,
	}
	return dao.Category.Ctx(ctx).Hook(hook.Bind(hookOptions)).Cache(orm.SetCacheOption(ctx))
}

func (s *sCategory) handleSearch(ctx context.Context, in *req.CategorySearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.CategoryName) {
		m = m.Where("category_name", in.CategoryName)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	return
}

func (s *sCategory) GetList(ctx context.Context, inReq *model.ListReq, in *req.CategorySearch) (out []*res.Category, err error) {
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sCategory) GetPageList(ctx context.Context, req *model.PageListReq, in *req.CategorySearch) (rs []*res.Category, total int, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&rs, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sCategory) Save(ctx context.Context, in *req.CategorySave) (id int64, err error) {
	var saveData *do.Category
	if err = gconv.Struct(in, &saveData); err != nil {
		return
	}

	// 检查Level字段是否已存在且有值

	rs, err := s.Model(ctx).OmitNilData().Data(saveData).Insert()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = int64(tmpId)
	return
}

func (s *sCategory) GetById(ctx context.Context, id int64) (res *res.Category, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sCategory) Update(ctx context.Context, in *req.CategoryUpdate) (err error) {
	var updateData *do.Category
	if err = gconv.Struct(in, &updateData); err != nil {
		return
	}

	var categoryItem *entity.Category
	err = s.Model(ctx).Where("id", in.Id).Scan(&categoryItem)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(categoryItem) {
		err = fmt.Errorf("记录不存在")
		return
	}

	_, err = s.Model(ctx).OmitNilData().Data(updateData).Where("id", in.Id).Update()

	if utils.IsError(err) {
		return
	}

	return
}

func (s *sCategory) Delete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 删除指定的记录
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sCategory) RealDelete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 物理删除指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sCategory) Recovery(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 恢复指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sCategory) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}

	return
}

func (s *sCategory) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).OmitNilData().Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sCategory) GetExportList(ctx context.Context, req *model.ListReq, in *req.CategorySearch) (res []*res.CategoryExcel, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, req).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}
