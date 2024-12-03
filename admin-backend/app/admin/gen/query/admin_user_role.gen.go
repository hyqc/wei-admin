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

func newAdminUserRole(db *gorm.DB, opts ...gen.DOOption) adminUserRole {
	_adminUserRole := adminUserRole{}

	_adminUserRole.adminUserRoleDo.UseDB(db, opts...)
	_adminUserRole.adminUserRoleDo.UseModel(&model.AdminUserRole{})

	tableName := _adminUserRole.adminUserRoleDo.TableName()
	_adminUserRole.ALL = field.NewAsterisk(tableName)
	_adminUserRole.AdminID = field.NewInt32(tableName, "admin_id")
	_adminUserRole.RoleID = field.NewInt32(tableName, "role_id")

	_adminUserRole.fillFieldMap()

	return _adminUserRole
}

type adminUserRole struct {
	adminUserRoleDo

	ALL     field.Asterisk
	AdminID field.Int32 // 管理员ID
	RoleID  field.Int32 // 角色ID

	fieldMap map[string]field.Expr
}

func (a adminUserRole) Table(newTableName string) *adminUserRole {
	a.adminUserRoleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminUserRole) As(alias string) *adminUserRole {
	a.adminUserRoleDo.DO = *(a.adminUserRoleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminUserRole) updateTableName(table string) *adminUserRole {
	a.ALL = field.NewAsterisk(table)
	a.AdminID = field.NewInt32(table, "admin_id")
	a.RoleID = field.NewInt32(table, "role_id")

	a.fillFieldMap()

	return a
}

func (a *adminUserRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminUserRole) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 2)
	a.fieldMap["admin_id"] = a.AdminID
	a.fieldMap["role_id"] = a.RoleID
}

func (a adminUserRole) clone(db *gorm.DB) adminUserRole {
	a.adminUserRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminUserRole) replaceDB(db *gorm.DB) adminUserRole {
	a.adminUserRoleDo.ReplaceDB(db)
	return a
}

type adminUserRoleDo struct{ gen.DO }

type IAdminUserRoleDo interface {
	gen.SubQuery
	Debug() IAdminUserRoleDo
	WithContext(ctx context.Context) IAdminUserRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminUserRoleDo
	WriteDB() IAdminUserRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminUserRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminUserRoleDo
	Not(conds ...gen.Condition) IAdminUserRoleDo
	Or(conds ...gen.Condition) IAdminUserRoleDo
	Select(conds ...field.Expr) IAdminUserRoleDo
	Where(conds ...gen.Condition) IAdminUserRoleDo
	Order(conds ...field.Expr) IAdminUserRoleDo
	Distinct(cols ...field.Expr) IAdminUserRoleDo
	Omit(cols ...field.Expr) IAdminUserRoleDo
	Join(table schema.Tabler, on ...field.Expr) IAdminUserRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminUserRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminUserRoleDo
	Group(cols ...field.Expr) IAdminUserRoleDo
	Having(conds ...gen.Condition) IAdminUserRoleDo
	Limit(limit int) IAdminUserRoleDo
	Offset(offset int) IAdminUserRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminUserRoleDo
	Unscoped() IAdminUserRoleDo
	Create(values ...*model.AdminUserRole) error
	CreateInBatches(values []*model.AdminUserRole, batchSize int) error
	Save(values ...*model.AdminUserRole) error
	First() (*model.AdminUserRole, error)
	Take() (*model.AdminUserRole, error)
	Last() (*model.AdminUserRole, error)
	Find() ([]*model.AdminUserRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminUserRole, err error)
	FindInBatches(result *[]*model.AdminUserRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminUserRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminUserRoleDo
	Assign(attrs ...field.AssignExpr) IAdminUserRoleDo
	Joins(fields ...field.RelationField) IAdminUserRoleDo
	Preload(fields ...field.RelationField) IAdminUserRoleDo
	FirstOrInit() (*model.AdminUserRole, error)
	FirstOrCreate() (*model.AdminUserRole, error)
	FindByPage(offset int, limit int) (result []*model.AdminUserRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminUserRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminUserRoleDo) Debug() IAdminUserRoleDo {
	return a.withDO(a.DO.Debug())
}

func (a adminUserRoleDo) WithContext(ctx context.Context) IAdminUserRoleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminUserRoleDo) ReadDB() IAdminUserRoleDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminUserRoleDo) WriteDB() IAdminUserRoleDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminUserRoleDo) Session(config *gorm.Session) IAdminUserRoleDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminUserRoleDo) Clauses(conds ...clause.Expression) IAdminUserRoleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminUserRoleDo) Returning(value interface{}, columns ...string) IAdminUserRoleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminUserRoleDo) Not(conds ...gen.Condition) IAdminUserRoleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminUserRoleDo) Or(conds ...gen.Condition) IAdminUserRoleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminUserRoleDo) Select(conds ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminUserRoleDo) Where(conds ...gen.Condition) IAdminUserRoleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminUserRoleDo) Order(conds ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminUserRoleDo) Distinct(cols ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminUserRoleDo) Omit(cols ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminUserRoleDo) Join(table schema.Tabler, on ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminUserRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminUserRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminUserRoleDo) Group(cols ...field.Expr) IAdminUserRoleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminUserRoleDo) Having(conds ...gen.Condition) IAdminUserRoleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminUserRoleDo) Limit(limit int) IAdminUserRoleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminUserRoleDo) Offset(offset int) IAdminUserRoleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminUserRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminUserRoleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminUserRoleDo) Unscoped() IAdminUserRoleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminUserRoleDo) Create(values ...*model.AdminUserRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminUserRoleDo) CreateInBatches(values []*model.AdminUserRole, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminUserRoleDo) Save(values ...*model.AdminUserRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminUserRoleDo) First() (*model.AdminUserRole, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserRole), nil
	}
}

func (a adminUserRoleDo) Take() (*model.AdminUserRole, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserRole), nil
	}
}

func (a adminUserRoleDo) Last() (*model.AdminUserRole, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserRole), nil
	}
}

func (a adminUserRoleDo) Find() ([]*model.AdminUserRole, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminUserRole), err
}

func (a adminUserRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminUserRole, err error) {
	buf := make([]*model.AdminUserRole, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminUserRoleDo) FindInBatches(result *[]*model.AdminUserRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminUserRoleDo) Attrs(attrs ...field.AssignExpr) IAdminUserRoleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminUserRoleDo) Assign(attrs ...field.AssignExpr) IAdminUserRoleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminUserRoleDo) Joins(fields ...field.RelationField) IAdminUserRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminUserRoleDo) Preload(fields ...field.RelationField) IAdminUserRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminUserRoleDo) FirstOrInit() (*model.AdminUserRole, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserRole), nil
	}
}

func (a adminUserRoleDo) FirstOrCreate() (*model.AdminUserRole, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserRole), nil
	}
}

func (a adminUserRoleDo) FindByPage(offset int, limit int) (result []*model.AdminUserRole, count int64, err error) {
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

func (a adminUserRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminUserRoleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminUserRoleDo) Delete(models ...*model.AdminUserRole) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminUserRoleDo) withDO(do gen.Dao) *adminUserRoleDo {
	a.DO = *do.(*gen.DO)
	return a
}
