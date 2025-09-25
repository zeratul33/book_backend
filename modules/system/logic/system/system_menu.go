// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemMenu struct {
	base.BaseService
}

func init() {
	service.RegisterSystemMenu(NewSystemMenu())
}

func NewSystemMenu() *sSystemMenu {
	return &sSystemMenu{}
}

func (s *sSystemMenu) Model(ctx context.Context) *gdb.Model {
	return dao.SystemMenu.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemMenu) GetRoutersByIds(ctx context.Context, menuIds []int64) (routes []*res.Router, err error) {
	systemMenuEntity := []entity.SystemMenu{}
	err = s.Model(ctx).WhereIn(dao.SystemMenu.Columns().Id, menuIds).Where(dao.SystemMenu.Columns().Status, 1).Order("parent_id, sort desc").Scan(&systemMenuEntity)
	if utils.IsError(err) {
		return
	}
	routes = s.treeList(systemMenuEntity)
	return
}

func (s *sSystemMenu) GetSuperAdminRouters(ctx context.Context) (routes []*res.Router, err error) {
	systemMenuEntity := []entity.SystemMenu{}
	err = s.Model(ctx).Where(dao.SystemMenu.Columns().Status, 1).Order("parent_id, sort desc").Scan(&systemMenuEntity)
	if utils.IsError(err) {
		return
	}
	routes = s.treeList(systemMenuEntity)
	return
}

func (s *sSystemMenu) treeList(nodes []entity.SystemMenu) (tree []*res.Router) {
	type itemTree map[int64]*res.Router
	itemList := make(itemTree)

	// 第一遍：创建所有节点并存储到map中
	for _, systemMenuEntity := range nodes {
		var item res.Router
		isHidden := false
		if systemMenuEntity.IsHidden == 1 {
			isHidden = true
		}
		route := "/" + systemMenuEntity.Route
		if systemMenuEntity.Type == "L" || systemMenuEntity.Type == "I" {
			route = systemMenuEntity.Route
		}
		item.Id = systemMenuEntity.Id
		item.ParentId = systemMenuEntity.ParentId
		item.Name = systemMenuEntity.Code
		item.Path = route
		item.Component = systemMenuEntity.Component
		item.Redirect = systemMenuEntity.Redirect
		item.Meta = res.Meta{Title: systemMenuEntity.Name, Icon: systemMenuEntity.Icon,
			Type: systemMenuEntity.Type, Hidden: isHidden, HiddenBreadcrumb: false}
		item.Children = make([]*res.Router, 0)
		itemList[systemMenuEntity.Id] = &item
	}

	// 第二遍：建立父子关系
	for _, systemMenuEntity := range nodes {
		item := itemList[systemMenuEntity.Id]
		if item == nil {
			continue
		}

		// 如果有父节点且父节点存在，则添加到父节点的children中
		if systemMenuEntity.ParentId != 0 && itemList[systemMenuEntity.ParentId] != nil {
			itemList[systemMenuEntity.ParentId].Children = append(itemList[systemMenuEntity.ParentId].Children, item)
		} else {
			// 否则作为根节点
			tree = append(tree, item)
		}
	}

	return
}

func (s *sSystemMenu) GetMenuCode(ctx context.Context, menuIds []int64) (menuCodes []string, err error) {
	result, err := s.Model(ctx).Fields(dao.SystemMenu.Columns().Code).WhereIn(dao.SystemMenu.Columns().Id, menuIds).Array()
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(result) {
		return
	}
	menuCodes = gconv.SliceStr(result)
	return
}

func (s *sSystemMenu) GetMenuByPermission(ctx context.Context, permission string, menuIds ...[]int64) (systemMenuEntity *entity.SystemMenu, err error) {
	m := s.Model(ctx).Where(dao.SystemMenu.Columns().Code, permission).Where(dao.SystemMenu.Columns().Status, 1)

	if len(menuIds) > 0 {
		m = m.WhereIn(dao.SystemMenu.Columns().Id, menuIds[0])
	}

	err = m.Order("parent_id, sort desc").Scan(&systemMenuEntity)

	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemMenu) GetTreeList(ctx context.Context, in *req.SystemMenuSearch) (tree []*res.SystemMenuTree, err error) {
	inReq := &model.ListReq{
		Recycle: in.Recycle,
	}
	params := g.Map{}
	if !g.IsEmpty(in.Status) {
		params["status"] = in.Status
	}

	m := s.Model(ctx)
	if !g.IsEmpty(in.Name) {
		m = m.Where("name like ? ", "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.NoButtons) {
		m = m.Where("type <> ? ", "B")
	}

	if !g.IsEmpty(in.Level) {
		m = m.Where("level like ? ", "%,"+in.Level+",%")
	}

	if !g.IsEmpty(in.CreatedAt) {
		if len(in.CreatedAt) > 0 {
			m = m.WhereGTE("created_at", in.CreatedAt[0]+" 00:00:00")
		}
		if len(in.CreatedAt) > 1 {
			m = m.WhereLTE("created_at", in.CreatedAt[1]+"23:59:59")
		}
	}
	m = orm.GetList(m, inReq, params)
	systemMenuEntity := []entity.SystemMenu{}
	err = m.Order("parent_id, sort desc").Scan(&systemMenuEntity)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(systemMenuEntity) {
		return
	}
	tree = s.treeItemList(ctx, systemMenuEntity)
	return
}

func (s *sSystemMenu) GetRecycleTreeList(ctx context.Context, in *req.SystemMenuSearch) (tree []*res.SystemMenuTree, err error) {
	inReq := &model.ListReq{
		Recycle: in.Recycle,
	}
	params := g.Map{}
	if !g.IsEmpty(in.Status) {
		params["status"] = in.Status
	}

	m := s.Model(ctx)
	if !g.IsEmpty(in.Name) {
		m = m.Where("name like ? ", "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.NoButtons) {
		m = m.Where("type <> ? ", "B")
	}

	if !g.IsEmpty(in.Level) {
		m = m.Where("level like ? ", "%,"+in.Level+",%")
	}

	if !g.IsEmpty(in.CreatedAt) {
		if len(in.CreatedAt) > 0 {
			m = m.WhereGTE("created_at", in.CreatedAt[0]+" 00:00:00")
		}
		if len(in.CreatedAt) > 1 {
			m = m.WhereLTE("created_at", in.CreatedAt[1]+"23:59:59")
		}
	}
	m = orm.GetList(m, inReq, params)
	systemMenuEntity := []entity.SystemMenu{}
	err = m.Order("parent_id, sort desc").Scan(&systemMenuEntity)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(systemMenuEntity) {
		return
	}
	tree = s.treeItemList(ctx, systemMenuEntity)
	return
}

func (s *sSystemMenu) treeItemList(ctx context.Context, nodes []entity.SystemMenu) (tree []*res.SystemMenuTree) {
	type itemTree map[int64]*res.SystemMenuTree
	itemList := make(itemTree)
	
	// 第一遍：创建所有节点并存储到map中
	for _, systemMenuEntity := range nodes {
		var item *res.SystemMenuTree
		if err := gconv.Struct(systemMenuEntity, &item); err != nil {
			g.Log().Error(ctx, "struct error:", err)
			continue
		}
		item.Children = make([]*res.SystemMenuTree, 0)
		itemList[systemMenuEntity.Id] = item
	}
	
	// 第二遍：建立父子关系
	for _, systemMenuEntity := range nodes {
		item := itemList[systemMenuEntity.Id]
		if item == nil {
			continue
		}
		
		// 如果有父节点且父节点存在，则添加到父节点的children中
		if systemMenuEntity.ParentId != 0 && itemList[systemMenuEntity.ParentId] != nil {
			itemList[systemMenuEntity.ParentId].Children = append(itemList[systemMenuEntity.ParentId].Children, item)
		} else {
			// 否则作为根节点
			tree = append(tree, item)
		}
	}
	return
}

func (s *sSystemMenu) treeSelectList(nodes []entity.SystemMenu) (tree []*res.SystemDeptSelectTree) {
	type itemTree map[int64]*res.SystemDeptSelectTree
	itemList := make(itemTree)
	
	// 第一遍：创建所有节点并存储到map中
	for _, systemMenuEntity := range nodes {
		var item res.SystemDeptSelectTree
		item.ParentId = systemMenuEntity.ParentId
		item.Id = systemMenuEntity.Id
		item.Label = systemMenuEntity.Name
		item.Value = systemMenuEntity.Id
		item.Children = make([]*res.SystemDeptSelectTree, 0)
		itemList[systemMenuEntity.Id] = &item
	}
	
	// 第二遍：建立父子关系
	for _, systemMenuEntity := range nodes {
		item := itemList[systemMenuEntity.Id]
		if item == nil {
			continue
		}
		
		// 如果有父节点且父节点存在，则添加到父节点的children中
		if systemMenuEntity.ParentId != 0 && itemList[systemMenuEntity.ParentId] != nil {
			itemList[systemMenuEntity.ParentId].Children = append(itemList[systemMenuEntity.ParentId].Children, item)
		} else {
			// 否则作为根节点
			tree = append(tree, item)
		}
	}
	return
}

func (s *sSystemMenu) GetSelectTree(ctx context.Context, userId int64, onlyMenu, scope bool) (routes []*res.SystemDeptSelectTree, err error) {
	m := s.Model(ctx).Where(dao.SystemMenu.Columns().Status, 1).Order("parent_id, sort desc")
	isSuperAdmin, err := service.SystemUser().IsSuperAdmin(ctx, userId)
	if err != nil {
		return
	}

	if scope && !isSuperAdmin {
		roleIds, err := service.SystemUser().GetRoles(ctx, userId)
		if err != nil {
			return nil, err
		}
		if !g.IsEmpty(roleIds) {
			menuIds, err := service.SystemRoleMenu().GetMenuIdsByRoleIds(ctx, roleIds)
			if err != nil {
				return nil, err
			}
			m = m.WhereIn(dao.SystemMenu.Columns().Id, menuIds)
		}
	}

	if onlyMenu {
		m = m.Where(dao.SystemMenu.Columns().Type, "M")
	}

	systemMenuEntity := []entity.SystemMenu{}
	err = m.Scan(&systemMenuEntity)
	if utils.IsError(err) {
		return
	}
	// 构建原有的菜单树
	originalTree := s.treeSelectList(systemMenuEntity)

	defaultNode := &res.SystemDeptSelectItem{
		Id:       0,
		ParentId: -1,
		Value:    0,
		Label:    "根节点",
	}
	// 创建默认根节点
	rootNode := &res.SystemDeptSelectTree{
		Children: originalTree,
	}

	rootNode.Id = defaultNode.Id
	rootNode.ParentId = defaultNode.ParentId
	rootNode.Value = defaultNode.Value
	rootNode.Label = defaultNode.Label

	// 返回包含根节点的树结构
	routes = []*res.SystemDeptSelectTree{rootNode}

	return
}

func (s *sSystemMenu) handleData(ctx context.Context, data *req.SystemMenuSave) (dataRes *req.SystemMenuSave, err error) {

	if g.IsEmpty(data.ParentId) {
		data.ParentId = 0
	}

	if !g.IsEmpty(data.Id) && (data.Id == data.ParentId) {
		return nil, gerror.New("id cannot be equal to parent_id")
	}

	var level string
	if data.ParentId == 0 {
		level = ",0,"
		if data.Type == "B" {
			data.Type = "M"
		}
	} else {
		var parentMenu *entity.SystemMenu
		err = s.Model(ctx).Where(dao.SystemMenu.Columns().Id, data.ParentId).Scan(&parentMenu)
		if utils.IsError(err) {
			return nil, err
		}
		if !g.IsEmpty(parentMenu) {
			level = fmt.Sprintf("%s%d,", parentMenu.Level, data.ParentId)
		} else {
			return nil, gerror.New("parent menu not found")
		}
	}
	data.Level = level
	dataRes = data
	return
}

func (s *sSystemMenu) Save(ctx context.Context, in *req.SystemMenuSave) (id int64, err error) {
	data, err := s.handleData(ctx, in)
	if err != nil {
		return
	}
	saveData := do.SystemMenu{
		Name:      data.Name,
		ParentId:  data.ParentId,
		Level:     data.Level,
		Sort:      data.Sort,
		Status:    data.Status,
		Code:      data.Code,
		Remark:    data.Remark,
		Type:      data.Type,
		Icon:      data.Icon,
		Route:     data.Route,
		Component: data.Component,
		Redirect:  data.Redirect,
		IsHidden:  data.IsHidden,
	}
	rs, err := s.Model(ctx).Data(saveData).Insert()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = gconv.Int64(tmpId)

	if data.Type == "M" && data.Restful == "1" {
		s.genButton(ctx, id, data.Name, data.Code)
	}

	return
}

func (s *sSystemMenu) genButton(ctx context.Context, id int64, name, code string) {
	m := make([]g.Map, 0)
	m = append(m, g.Map{
		"name": name + "列表",
		"code": code + ":index",
	})

	m = append(m, g.Map{
		"name": name + "回收站",
		"code": code + ":recycle",
	})

	m = append(m, g.Map{
		"name": name + "保存",
		"code": code + ":save",
	})

	m = append(m, g.Map{
		"name": name + "更新",
		"code": code + ":update",
	})

	m = append(m, g.Map{
		"name": name + "删除",
		"code": code + ":delete",
	})

	m = append(m, g.Map{
		"name": name + "读取",
		"code": code + ":read",
	})

	m = append(m, g.Map{
		"name": name + "恢复",
		"code": code + ":recovery",
	})

	m = append(m, g.Map{
		"name": name + "真实删除",
		"code": code + ":realDelete",
	})

	m = append(m, g.Map{
		"name": name + "导入",
		"code": code + ":import",
	})

	m = append(m, g.Map{
		"name": name + "导出",
		"code": code + ":export",
	})

	for _, v := range m {
		s.Save(ctx, &req.SystemMenuSave{
			Name:      gconv.String(v["name"]),
			ParentId:  id,
			Sort:      0,
			Status:    1,
			Code:      gconv.String(v["code"]),
			Remark:    "",
			Type:      "B",
			Icon:      "",
			Route:     "",
			Component: "",
			Redirect:  "",
			IsHidden:  2,
		})
	}
}

func (s *sSystemMenu) Update(ctx context.Context, in *req.SystemMenuSave) (err error) {
	data, err := s.handleData(ctx, in)
	if err != nil {
		return
	}
	var systemMenuItem *entity.SystemMenu
	err = s.Model(ctx).Where("id", in.Id).Scan(&systemMenuItem)
	if utils.IsError(err) {
		return
	}
	oldLevel := systemMenuItem.Level
	saveData := do.SystemMenu{
		Name:      data.Name,
		ParentId:  data.ParentId,
		Level:     data.Level,
		Sort:      data.Sort,
		Status:    data.Status,
		Code:      data.Code,
		Remark:    data.Remark,
		Type:      data.Type,
		Icon:      data.Icon,
		Route:     data.Route,
		Component: data.Component,
		Redirect:  data.Redirect,
		IsHidden:  data.IsHidden,
	}
	_, err = s.Model(ctx).Where(dao.SystemMenu.Columns().Id, data.Id).Data(saveData).Update()
	if utils.IsError(err) {
		return err
	}

	var menu []*entity.SystemMenu

	childLevelPrefix := fmt.Sprintf("%s%d,", oldLevel, data.Id)
	err = s.Model(ctx).Unscoped().WhereLike("level", childLevelPrefix+"%").Scan(&menu)
	if utils.IsError(err) {
		return err
	}
	if !g.IsEmpty(menu) {
		for _, item := range menu {
			newLevel := utils.ReplaceSubstr(item.Level, oldLevel, data.Level)
			_, err = s.Model(ctx).Unscoped().Where(dao.SystemMenu.Columns().Id, item.Id).Data(do.SystemMenu{
				Level: newLevel,
			}).Update()
			if utils.IsError(err) {
				return err
			}
		}
	}
	return
}

func (s *sSystemMenu) Delete(ctx context.Context, ids []int64) (err error) {
	for _, id := range ids {
		var targetRecord *entity.SystemMenu
		err = s.Model(ctx).Where("id", id).Scan(&targetRecord)
		if utils.IsError(err) {
			return err
		}

		if !g.IsEmpty(targetRecord) && !g.IsEmpty(targetRecord.Level) {
			childLevelPrefix := fmt.Sprintf("%s%d,", targetRecord.Level, id)
			// 批量删除所有子节点
			_, err = s.Model(ctx).
				Where("level LIKE ?", childLevelPrefix+"%").
				Delete()
			if utils.IsError(err) {
				return err
			}
		}
	}

	// 删除指定的记录
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemMenu) RealDelete(ctx context.Context, ids []int64) (err error) {
	// 为每个成功删除的ID处理其子菜单
	for _, id := range ids {
		var targetRecord *entity.SystemMenu
		err = s.Model(ctx).Unscoped().Where("id", id).Scan(&targetRecord)
		if utils.IsError(err) {
			return err
		}

		if !g.IsEmpty(targetRecord) && !g.IsEmpty(targetRecord.Level) {
			childLevelPrefix := fmt.Sprintf("%s%d,", targetRecord.Level, id)
			// 批量物理删除所有子节点
			_, err = s.Model(ctx).Unscoped().
				Where("level LIKE ?", childLevelPrefix+"%").
				Delete()
			if utils.IsError(err) {
				return err
			}
		}
	}

	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}

	return
}

func (s *sSystemMenu) Recovery(ctx context.Context, ids []int64) (err error) {
	// 为每个要恢复的ID处理其子菜单
	for _, id := range ids {
		var targetRecord *entity.SystemMenu
		err = s.Model(ctx).Unscoped().Where("id", id).Scan(&targetRecord)
		if utils.IsError(err) {
			return err
		}

		if !g.IsEmpty(targetRecord) && !g.IsEmpty(targetRecord.Level) {
			childLevelPrefix := fmt.Sprintf("%s%d,", targetRecord.Level, id)
			// 批量恢复所有子节点
			_, err = s.Model(ctx).Unscoped().
				Where("level LIKE ?", childLevelPrefix+"%").
				Update(g.Map{"deleted_at": nil})
			if utils.IsError(err) {
				return err
			}
		}
	}
	// 恢复指定的记录
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemMenu) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	doObj := do.SystemMenu{}
	needCalculateLevel := utils.HasField(doObj, "Level")
	if needCalculateLevel {
		// 2. 查询目标记录信息，获取level字段用于查找子节点
		var targetRecord *entity.SystemMenu
		err = s.Model(ctx).Where("id", id).Scan(&targetRecord)
		if utils.IsError(err) {
			return err
		}

		// 3. 如果目标记录存在且有level信息，查找并更新所有子节点
		if !g.IsEmpty(targetRecord) && !g.IsEmpty(targetRecord.Level) {
			childLevelPrefix := fmt.Sprintf("%s%d,", targetRecord.Level, id)
			// 批量更新所有子节点的状态
			_, err = s.Model(ctx).OmitNilData().Data(g.Map{"status": status}).
				Where("level LIKE ?", childLevelPrefix+"%").
				Where("id != ?", id). // 排除自身
				Update()
			if utils.IsError(err) {
				return err
			}
		}

	}
	return
}

func (s *sSystemMenu) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
