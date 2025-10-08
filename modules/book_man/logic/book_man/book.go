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

type sBook struct {
	base.BaseService
}

func init() {
	service.RegisterBook(NewBook())
}

func NewBook() *sBook {
	return &sBook{}
}

func (s *sBook) Model(ctx context.Context) *gdb.Model {
	doObj := do.Book{}
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
	return dao.Book.Ctx(ctx).Hook(hook.Bind(hookOptions)).Cache(orm.SetCacheOption(ctx))
}

func (s *sBook) handleSearch(ctx context.Context, in *req.BookSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.BookName) {
		m = m.Where("book_name", in.BookName)
	}

	if !g.IsEmpty(in.AuthorName) {
		m = m.Where("author_name", in.AuthorName)
	}

	if !g.IsEmpty(in.CategoryId) {
		m = m.Where("category_id", in.CategoryId)
	}

	return
}

func (s *sBook) GetList(ctx context.Context, inReq *model.ListReq, in *req.BookSearch) (out []*res.Book, err error) {
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sBook) GetPageList(ctx context.Context, req *model.PageListReq, in *req.BookSearch) (rs []*res.Book, total int, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&rs, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sBook) Save(ctx context.Context, in *req.BookSave) (id int64, err error) {
	var saveData *do.Book
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

func (s *sBook) GetById(ctx context.Context, id int64) (res *res.Book, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sBook) Update(ctx context.Context, in *req.BookUpdate) (err error) {
	var updateData *do.Book
	if err = gconv.Struct(in, &updateData); err != nil {
		return
	}

	var bookItem *entity.Book
	err = s.Model(ctx).Where("id", in.Id).Scan(&bookItem)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(bookItem) {
		err = fmt.Errorf("记录不存在")
		return
	}

	_, err = s.Model(ctx).OmitNilData().Data(updateData).Where("id", in.Id).Update()

	if utils.IsError(err) {
		return
	}

	return
}

func (s *sBook) Delete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 删除指定的记录
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sBook) RealDelete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 物理删除指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sBook) Recovery(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 恢复指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sBook) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}

	return
}

func (s *sBook) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).OmitNilData().Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sBook) GetExportList(ctx context.Context, req *model.ListReq, in *req.BookSearch) (res []*res.BookExcel, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, req).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}
