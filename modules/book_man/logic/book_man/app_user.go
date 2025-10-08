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

type sAppUser struct {
	base.BaseService
}

func init() {
	service.RegisterAppUser(NewAppUser())
}

func NewAppUser() *sAppUser {
	return &sAppUser{}
}

func (s *sAppUser) Model(ctx context.Context) *gdb.Model {
	doObj := do.AppUser{}
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
	return dao.AppUser.Ctx(ctx).Hook(hook.Bind(hookOptions)).Cache(orm.SetCacheOption(ctx))
}

func (s *sAppUser) handleSearch(ctx context.Context, in *req.AppUserSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.Id) {
		m = m.Where("id", in.Id)
	}

	if !g.IsEmpty(in.CreatedAt) {
		if len(in.CreatedAt) > 0 {
			m = m.WhereGTE("created_at", in.CreatedAt[0]+" 00:00:00")
		}
		if len(in.CreatedAt) > 1 {
			m = m.WhereLTE("created_at", in.CreatedAt[1]+"23:59:59")
		}
	}

	if !g.IsEmpty(in.Username) {
		m = m.Where("username", in.Username)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	return
}

func (s *sAppUser) GetList(ctx context.Context, inReq *model.ListReq, in *req.AppUserSearch) (out []*res.AppUser, err error) {
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sAppUser) GetPageList(ctx context.Context, req *model.PageListReq, in *req.AppUserSearch) (rs []*res.AppUser, total int, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&rs, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sAppUser) Save(ctx context.Context, in *req.AppUserSave) (id int64, err error) {
	var saveData *do.AppUser
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

func (s *sAppUser) GetById(ctx context.Context, id int64) (res *res.AppUser, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sAppUser) Update(ctx context.Context, in *req.AppUserUpdate) (err error) {
	var updateData *do.AppUser
	if err = gconv.Struct(in, &updateData); err != nil {
		return
	}

	var appUserItem *entity.AppUser
	err = s.Model(ctx).Where("id", in.Id).Scan(&appUserItem)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(appUserItem) {
		err = fmt.Errorf("记录不存在")
		return
	}

	_, err = s.Model(ctx).OmitNilData().Data(updateData).Where("id", in.Id).Update()

	if utils.IsError(err) {
		return
	}

	return
}

func (s *sAppUser) Delete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 删除指定的记录
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sAppUser) RealDelete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 物理删除指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sAppUser) Recovery(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 恢复指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sAppUser) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}

	return
}

func (s *sAppUser) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).OmitNilData().Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sAppUser) GetExportList(ctx context.Context, req *model.ListReq, in *req.AppUserSearch) (res []*res.AppUserExcel, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, req).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}
