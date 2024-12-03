// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"admin/app/admin/gen/model"
)

func newAdminPermission(db *gorm.DB, opts ...gen.DOOption) adminPermission {
	_adminPermission := adminPermission{}

	_adminPermission.adminPermissionDo.UseDB(db, opts...)
	_adminPermission.adminPermissionDo.UseModel(&model.AdminPermission{})

	tableName := _adminPermission.adminPermissionDo.TableName()
	_adminPermission.ALL = field.NewAsterisk(tableName)
	_adminPermission.ID = field.NewInt32(tableName, "id")
	_adminPermission.MenuID = field.NewInt32(tableName, "menu_id")
	_adminPermission.Key = field.NewString(tableName, "key")
	_adminPermission.Name = field.NewString(tableName, "name")
	_adminPermission.Type = field.NewString(tableName, "type")
	_adminPermission.Describe = field.NewString(tableName, "describe")
	_adminPermission.IsEnabled = field.NewBool(tableName, "is_enabled")
	_adminPermission.CreatedAt = field.NewTime(tableName, "created_at")
	_adminPermission.UpdatedAt = field.NewTime(tableName, "updated_at")

	_adminPermission.fillFieldMap()

	return _adminPermission
}

type adminPermission struct {
	adminPermissionDo

	ALL    field.Asterisk
	ID     field.Int32
	MenuID field.Int32  // 所属菜单ID
	Key    field.String // 权限唯一标识名称
	Name   field.String // 权限显示名称
	/*
		权限的操作类型
		view：查看（只读）
		edit：编辑（读写）
		delete：删除（彻底删除）
	*/
	Type      field.String
	Describe  field.String // 权限描述
	IsEnabled field.Bool   // 是否启用：1启用，0禁用
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (a adminPermission) Table(newTableName string) *adminPermission {
	a.adminPermissionDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminPermission) As(alias string) *adminPermission {
	a.adminPermissionDo.DO = *(a.adminPermissionDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminPermission) updateTableName(table string) *adminPermission {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.MenuID = field.NewInt32(table, "menu_id")
	a.Key = field.NewString(table, "key")
	a.Name = field.NewString(table, "name")
	a.Type = field.NewString(table, "type")
	a.Describe = field.NewString(table, "describe")
	a.IsEnabled = field.NewBool(table, "is_enabled")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *adminPermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminPermission) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["menu_id"] = a.MenuID
	a.fieldMap["key"] = a.Key
	a.fieldMap["name"] = a.Name
	a.fieldMap["type"] = a.Type
	a.fieldMap["describe"] = a.Describe
	a.fieldMap["is_enabled"] = a.IsEnabled
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a adminPermission) clone(db *gorm.DB) adminPermission {
	a.adminPermissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminPermission) replaceDB(db *gorm.DB) adminPermission {
	a.adminPermissionDo.ReplaceDB(db)
	return a
}

type adminPermissionDo struct{ gen.DO }

type IAdminPermissionDo interface {
	gen.SubQuery
	Debug() IAdminPermissionDo
	WithContext(ctx context.Context) IAdminPermissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminPermissionDo
	WriteDB() IAdminPermissionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminPermissionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminPermissionDo
	Not(conds ...gen.Condition) IAdminPermissionDo
	Or(conds ...gen.Condition) IAdminPermissionDo
	Select(conds ...field.Expr) IAdminPermissionDo
	Where(conds ...gen.Condition) IAdminPermissionDo
	Order(conds ...field.Expr) IAdminPermissionDo
	Distinct(cols ...field.Expr) IAdminPermissionDo
	Omit(cols ...field.Expr) IAdminPermissionDo
	Join(table schema.Tabler, on ...field.Expr) IAdminPermissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminPermissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminPermissionDo
	Group(cols ...field.Expr) IAdminPermissionDo
	Having(conds ...gen.Condition) IAdminPermissionDo
	Limit(limit int) IAdminPermissionDo
	Offset(offset int) IAdminPermissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminPermissionDo
	Unscoped() IAdminPermissionDo
	Create(values ...*model.AdminPermission) error
	CreateInBatches(values []*model.AdminPermission, batchSize int) error
	Save(values ...*model.AdminPermission) error
	First() (*model.AdminPermission, error)
	Take() (*model.AdminPermission, error)
	Last() (*model.AdminPermission, error)
	Find() ([]*model.AdminPermission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminPermission, err error)
	FindInBatches(result *[]*model.AdminPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminPermission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminPermissionDo
	Assign(attrs ...field.AssignExpr) IAdminPermissionDo
	Joins(fields ...field.RelationField) IAdminPermissionDo
	Preload(fields ...field.RelationField) IAdminPermissionDo
	FirstOrInit() (*model.AdminPermission, error)
	FirstOrCreate() (*model.AdminPermission, error)
	FindByPage(offset int, limit int) (result []*model.AdminPermission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminPermissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminPermissionDo) Debug() IAdminPermissionDo {
	return a.withDO(a.DO.Debug())
}

func (a adminPermissionDo) WithContext(ctx context.Context) IAdminPermissionDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminPermissionDo) ReadDB() IAdminPermissionDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminPermissionDo) WriteDB() IAdminPermissionDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminPermissionDo) Session(config *gorm.Session) IAdminPermissionDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminPermissionDo) Clauses(conds ...clause.Expression) IAdminPermissionDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminPermissionDo) Returning(value interface{}, columns ...string) IAdminPermissionDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminPermissionDo) Not(conds ...gen.Condition) IAdminPermissionDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminPermissionDo) Or(conds ...gen.Condition) IAdminPermissionDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminPermissionDo) Select(conds ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminPermissionDo) Where(conds ...gen.Condition) IAdminPermissionDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminPermissionDo) Order(conds ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminPermissionDo) Distinct(cols ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminPermissionDo) Omit(cols ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminPermissionDo) Join(table schema.Tabler, on ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminPermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminPermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminPermissionDo) Group(cols ...field.Expr) IAdminPermissionDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminPermissionDo) Having(conds ...gen.Condition) IAdminPermissionDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminPermissionDo) Limit(limit int) IAdminPermissionDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminPermissionDo) Offset(offset int) IAdminPermissionDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminPermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminPermissionDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminPermissionDo) Unscoped() IAdminPermissionDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminPermissionDo) Create(values ...*model.AdminPermission) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminPermissionDo) CreateInBatches(values []*model.AdminPermission, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminPermissionDo) Save(values ...*model.AdminPermission) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminPermissionDo) First() (*model.AdminPermission, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminPermission), nil
	}
}

func (a adminPermissionDo) Take() (*model.AdminPermission, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminPermission), nil
	}
}

func (a adminPermissionDo) Last() (*model.AdminPermission, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminPermission), nil
	}
}

func (a adminPermissionDo) Find() ([]*model.AdminPermission, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminPermission), err
}

func (a adminPermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminPermission, err error) {
	buf := make([]*model.AdminPermission, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminPermissionDo) FindInBatches(result *[]*model.AdminPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminPermissionDo) Attrs(attrs ...field.AssignExpr) IAdminPermissionDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminPermissionDo) Assign(attrs ...field.AssignExpr) IAdminPermissionDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminPermissionDo) Joins(fields ...field.RelationField) IAdminPermissionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminPermissionDo) Preload(fields ...field.RelationField) IAdminPermissionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminPermissionDo) FirstOrInit() (*model.AdminPermission, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminPermission), nil
	}
}

func (a adminPermissionDo) FirstOrCreate() (*model.AdminPermission, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminPermission), nil
	}
}

func (a adminPermissionDo) FindByPage(offset int, limit int) (result []*model.AdminPermission, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminPermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminPermissionDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminPermissionDo) Delete(models ...*model.AdminPermission) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminPermissionDo) withDO(do gen.Dao) *adminPermissionDo {
	a.DO = *do.(*gen.DO)
	return a
}
