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

func newAdminRolePermission(db *gorm.DB, opts ...gen.DOOption) adminRolePermission {
	_adminRolePermission := adminRolePermission{}

	_adminRolePermission.adminRolePermissionDo.UseDB(db, opts...)
	_adminRolePermission.adminRolePermissionDo.UseModel(&model.AdminRolePermission{})

	tableName := _adminRolePermission.adminRolePermissionDo.TableName()
	_adminRolePermission.ALL = field.NewAsterisk(tableName)
	_adminRolePermission.RoleID = field.NewInt32(tableName, "role_id")
	_adminRolePermission.PermissionID = field.NewInt32(tableName, "permission_id")

	_adminRolePermission.fillFieldMap()

	return _adminRolePermission
}

type adminRolePermission struct {
	adminRolePermissionDo

	ALL          field.Asterisk
	RoleID       field.Int32 // 角色ID
	PermissionID field.Int32 // 权限ID

	fieldMap map[string]field.Expr
}

func (a adminRolePermission) Table(newTableName string) *adminRolePermission {
	a.adminRolePermissionDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminRolePermission) As(alias string) *adminRolePermission {
	a.adminRolePermissionDo.DO = *(a.adminRolePermissionDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminRolePermission) updateTableName(table string) *adminRolePermission {
	a.ALL = field.NewAsterisk(table)
	a.RoleID = field.NewInt32(table, "role_id")
	a.PermissionID = field.NewInt32(table, "permission_id")

	a.fillFieldMap()

	return a
}

func (a *adminRolePermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminRolePermission) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 2)
	a.fieldMap["role_id"] = a.RoleID
	a.fieldMap["permission_id"] = a.PermissionID
}

func (a adminRolePermission) clone(db *gorm.DB) adminRolePermission {
	a.adminRolePermissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminRolePermission) replaceDB(db *gorm.DB) adminRolePermission {
	a.adminRolePermissionDo.ReplaceDB(db)
	return a
}

type adminRolePermissionDo struct{ gen.DO }

type IAdminRolePermissionDo interface {
	gen.SubQuery
	Debug() IAdminRolePermissionDo
	WithContext(ctx context.Context) IAdminRolePermissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminRolePermissionDo
	WriteDB() IAdminRolePermissionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminRolePermissionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminRolePermissionDo
	Not(conds ...gen.Condition) IAdminRolePermissionDo
	Or(conds ...gen.Condition) IAdminRolePermissionDo
	Select(conds ...field.Expr) IAdminRolePermissionDo
	Where(conds ...gen.Condition) IAdminRolePermissionDo
	Order(conds ...field.Expr) IAdminRolePermissionDo
	Distinct(cols ...field.Expr) IAdminRolePermissionDo
	Omit(cols ...field.Expr) IAdminRolePermissionDo
	Join(table schema.Tabler, on ...field.Expr) IAdminRolePermissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRolePermissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminRolePermissionDo
	Group(cols ...field.Expr) IAdminRolePermissionDo
	Having(conds ...gen.Condition) IAdminRolePermissionDo
	Limit(limit int) IAdminRolePermissionDo
	Offset(offset int) IAdminRolePermissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRolePermissionDo
	Unscoped() IAdminRolePermissionDo
	Create(values ...*model.AdminRolePermission) error
	CreateInBatches(values []*model.AdminRolePermission, batchSize int) error
	Save(values ...*model.AdminRolePermission) error
	First() (*model.AdminRolePermission, error)
	Take() (*model.AdminRolePermission, error)
	Last() (*model.AdminRolePermission, error)
	Find() ([]*model.AdminRolePermission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRolePermission, err error)
	FindInBatches(result *[]*model.AdminRolePermission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminRolePermission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminRolePermissionDo
	Assign(attrs ...field.AssignExpr) IAdminRolePermissionDo
	Joins(fields ...field.RelationField) IAdminRolePermissionDo
	Preload(fields ...field.RelationField) IAdminRolePermissionDo
	FirstOrInit() (*model.AdminRolePermission, error)
	FirstOrCreate() (*model.AdminRolePermission, error)
	FindByPage(offset int, limit int) (result []*model.AdminRolePermission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminRolePermissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminRolePermissionDo) Debug() IAdminRolePermissionDo {
	return a.withDO(a.DO.Debug())
}

func (a adminRolePermissionDo) WithContext(ctx context.Context) IAdminRolePermissionDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminRolePermissionDo) ReadDB() IAdminRolePermissionDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminRolePermissionDo) WriteDB() IAdminRolePermissionDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminRolePermissionDo) Session(config *gorm.Session) IAdminRolePermissionDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminRolePermissionDo) Clauses(conds ...clause.Expression) IAdminRolePermissionDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminRolePermissionDo) Returning(value interface{}, columns ...string) IAdminRolePermissionDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminRolePermissionDo) Not(conds ...gen.Condition) IAdminRolePermissionDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminRolePermissionDo) Or(conds ...gen.Condition) IAdminRolePermissionDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminRolePermissionDo) Select(conds ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminRolePermissionDo) Where(conds ...gen.Condition) IAdminRolePermissionDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminRolePermissionDo) Order(conds ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminRolePermissionDo) Distinct(cols ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminRolePermissionDo) Omit(cols ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminRolePermissionDo) Join(table schema.Tabler, on ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminRolePermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminRolePermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminRolePermissionDo) Group(cols ...field.Expr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminRolePermissionDo) Having(conds ...gen.Condition) IAdminRolePermissionDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminRolePermissionDo) Limit(limit int) IAdminRolePermissionDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminRolePermissionDo) Offset(offset int) IAdminRolePermissionDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminRolePermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRolePermissionDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminRolePermissionDo) Unscoped() IAdminRolePermissionDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminRolePermissionDo) Create(values ...*model.AdminRolePermission) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminRolePermissionDo) CreateInBatches(values []*model.AdminRolePermission, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminRolePermissionDo) Save(values ...*model.AdminRolePermission) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminRolePermissionDo) First() (*model.AdminRolePermission, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRolePermission), nil
	}
}

func (a adminRolePermissionDo) Take() (*model.AdminRolePermission, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRolePermission), nil
	}
}

func (a adminRolePermissionDo) Last() (*model.AdminRolePermission, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRolePermission), nil
	}
}

func (a adminRolePermissionDo) Find() ([]*model.AdminRolePermission, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminRolePermission), err
}

func (a adminRolePermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRolePermission, err error) {
	buf := make([]*model.AdminRolePermission, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminRolePermissionDo) FindInBatches(result *[]*model.AdminRolePermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminRolePermissionDo) Attrs(attrs ...field.AssignExpr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminRolePermissionDo) Assign(attrs ...field.AssignExpr) IAdminRolePermissionDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminRolePermissionDo) Joins(fields ...field.RelationField) IAdminRolePermissionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminRolePermissionDo) Preload(fields ...field.RelationField) IAdminRolePermissionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminRolePermissionDo) FirstOrInit() (*model.AdminRolePermission, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRolePermission), nil
	}
}

func (a adminRolePermissionDo) FirstOrCreate() (*model.AdminRolePermission, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRolePermission), nil
	}
}

func (a adminRolePermissionDo) FindByPage(offset int, limit int) (result []*model.AdminRolePermission, count int64, err error) {
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

func (a adminRolePermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminRolePermissionDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminRolePermissionDo) Delete(models ...*model.AdminRolePermission) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminRolePermissionDo) withDO(do gen.Dao) *adminRolePermissionDo {
	a.DO = *do.(*gen.DO)
	return a
}
