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

type sSubscribed struct {
	base.BaseService
}

func init() {
	service.RegisterSubscribed(NewSubscribed())
}

func NewSubscribed() *sSubscribed {
	return &sSubscribed{}
}

func (s *sSubscribed) Model(ctx context.Context) *gdb.Model {
	doObj := do.Subscribed{}
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
	return dao.Subscribed.Ctx(ctx).Hook(hook.Bind(hookOptions)).Cache(orm.SetCacheOption(ctx))
}

func (s *sSubscribed) handleSearch(ctx context.Context, in *req.SubscribedSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.Id) {
		m = m.Where("id", in.Id)
	}

	if !g.IsEmpty(in.SubscribedUser) {
		m = m.Where("subscribed_user", in.SubscribedUser)
	}

	if !g.IsEmpty(in.SubscriebedBook) {
		m = m.Where("subscriebed_book", in.SubscriebedBook)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	return
}

func (s *sSubscribed) GetList(ctx context.Context, inReq *model.ListReq, in *req.SubscribedSearch) (out []*res.Subscribed, err error) {
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSubscribed) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SubscribedSearch) (rs []*res.Subscribed, total int, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&rs, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSubscribed) Save(ctx context.Context, in *req.SubscribedSave) (id int64, err error) {
	var saveData *do.Subscribed
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

func (s *sSubscribed) GetById(ctx context.Context, id int64) (res *res.Subscribed, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSubscribed) Update(ctx context.Context, in *req.SubscribedUpdate) (err error) {
	var updateData *do.Subscribed
	if err = gconv.Struct(in, &updateData); err != nil {
		return
	}

	var subscribedItem *entity.Subscribed
	err = s.Model(ctx).Where("id", in.Id).Scan(&subscribedItem)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(subscribedItem) {
		err = fmt.Errorf("记录不存在")
		return
	}

	_, err = s.Model(ctx).OmitNilData().Data(updateData).Where("id", in.Id).Update()

	if utils.IsError(err) {
		return
	}

	return
}

func (s *sSubscribed) Delete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 删除指定的记录
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSubscribed) RealDelete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 物理删除指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSubscribed) Recovery(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 恢复指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSubscribed) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}

	return
}

func (s *sSubscribed) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).OmitNilData().Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSubscribed) GetExportList(ctx context.Context, req *model.ListReq, in *req.SubscribedSearch) (res []*res.SubscribedExcel, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, req).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}
