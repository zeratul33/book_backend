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
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	sysService "devinggo/modules/system/service"
	"github.com/gogf/gf/v2/util/gutil"

	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sComment struct {
	base.BaseService
}

func init() {
	service.RegisterComment(NewComment())
}

func NewComment() *sComment {
	return &sComment{}
}

func (s *sComment) Model(ctx context.Context) *gdb.Model {
	doObj := do.Comment{}
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
	return dao.Comment.Ctx(ctx).Hook(hook.Bind(hookOptions)).Cache(orm.SetCacheOption(ctx))
}

func (s *sComment) handleSearch(ctx context.Context, in *req.CommentSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.UserId) {
		m = m.Where("user_id", in.UserId)
	}

	if !g.IsEmpty(in.UserComment) {
		m = m.Where("user_comment like ? ", "%"+in.UserComment+"%")
	}

	if !g.IsEmpty(in.CommentTime) {
		if len(in.CommentTime) > 0 {
			m = m.WhereGTE("comment_time", in.CommentTime[0]+" 00:00:00")
		}
		//if len(in.CreatedAt) > 1 {
		//    m = m.WhereLTE("comment_time", in.CommentTime[1]+"23:59:59")
		//}
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	if !g.IsEmpty(in.BookId) {
		m = m.Where("book_id", in.BookId)
	}

	return
}

func (s *sComment) GetList(ctx context.Context, inReq *model.ListReq, in *req.CommentSearch) (out []*res.Comment, err error) {
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sComment) GetPageList(ctx context.Context, req *model.PageListReq, in *req.CommentSearch) (rs []*res.Comment, total int, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&rs, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sComment) Save(ctx context.Context, in *req.CommentSave) (id int64, err error) {
	var saveData *do.Comment
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

func (s *sComment) GetById(ctx context.Context, id int64) (res *res.Comment, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sComment) Update(ctx context.Context, in *req.CommentUpdate) (err error) {
	var updateData *do.Comment
	if err = gconv.Struct(in, &updateData); err != nil {
		return
	}

	var commentItem *entity.Comment
	err = s.Model(ctx).Where("id", in.Id).Scan(&commentItem)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(commentItem) {
		err = fmt.Errorf("记录不存在")
		return
	}

	_, err = s.Model(ctx).OmitNilData().Data(updateData).Where("id", in.Id).Update()

	if utils.IsError(err) {
		return
	}

	return
}

func (s *sComment) Delete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 删除指定的记录
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sComment) RealDelete(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 物理删除指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sComment) Recovery(ctx context.Context, ids []int64) (err error) {
	// 检查是否需要处理Level字段

	// 恢复指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sComment) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}

	return
}

func (s *sComment) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).OmitNilData().Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sComment) GetExportList(ctx context.Context, req *model.ListReq, in *req.CommentSearch) (res []*res.CommentExcel, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, req).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sComment) GetCommentByBook(ctx context.Context, id int64) (out []*res.CommentApp, err error) {
	out = []*res.CommentApp{}
	g.Model("comment").Where("book_id", id).ScanList(&out, "Comment")
	valuesUnique := gdb.ListItemValuesUnique(out, "Comment", "UserId")
	gutil.Dump(valuesUnique)
	g.Model("book").Where("id", id).ScanList(&out, "Book", "Comment", "id:BookId")
	gutil.Dump(out)
	g.Model("app_user").Where("id", valuesUnique).ScanList(&out, "AppUser", "Comment", "id:UserId")
	return
}

func (s *sComment) GetCommentByUser(ctx context.Context) (out []*res.CommentApp, err error) {
	out = []*res.CommentApp{}
	fromCtx := g.RequestFromCtx(ctx)
	appId := contexts.New().GetAppId(ctx)
	loginUser, err := sysService.Token().ParseLoginUser(fromCtx, appId)
	if err != nil {
		return
	}
	//var result []do.Comment
	fmt.Println("----------loginUser----------")
	fmt.Println(loginUser.Id)
	//array, err := gdb.CatchSQL(ctx, func(ctx context.Context) error {
	//	err := service.Comment().Model(ctx).With(do.AppUser{}).Where("user_id", loginUser.Id).Scan(&out)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//g.Dump(array)
	g.Model("comment").Where("user_id", loginUser.Id).ScanList(&out, "Comment")
	valuesUnique := gdb.ListItemValuesUnique(out, "Comment", "UserId")
	unique := gdb.ListItemValuesUnique(out, "Comment", "BookId")
	gutil.Dump(valuesUnique)
	g.Model("app_user").Where("id", valuesUnique).ScanList(&out, "AppUser", "Comment", "id:UserId")
	g.Model("book").Where("id", unique).ScanList(&out, "Book", "Comment", "id:BookId")
	fmt.Printf("---------------getCommentByUser---------------")
	for _, comment := range out {
		gutil.Dump(comment)
		fmt.Printf("---------------getCommentByUser---------------")
	}
	fmt.Printf("---------------getCommentByUser---------------")
	//if !g.IsEmpty(comments) {
	//	for _, item := range comments {
	//		out = append(out, item)
	//	}
	//
	//} else {
	//	out = make([]*res.Comment, 0)
	//}
	return
}
