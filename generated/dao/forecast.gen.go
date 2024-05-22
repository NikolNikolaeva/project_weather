// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"project_weather/generated/dao/model"
)

func newForecast(db *gorm.DB, opts ...gen.DOOption) forecast {
	_forecast := forecast{}

	_forecast.forecastDo.UseDB(db, opts...)
	_forecast.forecastDo.UseModel(&model.Forecast{})

	tableName := _forecast.forecastDo.TableName()
	_forecast.ALL = field.NewAsterisk(tableName)
	_forecast.ID = field.NewString(tableName, "id")
	_forecast.CityID = field.NewString(tableName, "city_id")
	_forecast.ForecastDate = field.NewTime(tableName, "forecast_date")
	_forecast.Temperature = field.NewField(tableName, "temperature")
	_forecast.Condition = field.NewString(tableName, "condition")
	_forecast.CreatedAt = field.NewTime(tableName, "created_at")
	_forecast.UpdatedAt = field.NewTime(tableName, "updated_at")

	_forecast.fillFieldMap()

	return _forecast
}

type forecast struct {
	forecastDo

	ALL          field.Asterisk
	ID           field.String
	CityID       field.String
	ForecastDate field.Time
	Temperature  field.Field
	Condition    field.String
	CreatedAt    field.Time
	UpdatedAt    field.Time

	fieldMap map[string]field.Expr
}

func (f forecast) Table(newTableName string) *forecast {
	f.forecastDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f forecast) As(alias string) *forecast {
	f.forecastDo.DO = *(f.forecastDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *forecast) updateTableName(table string) *forecast {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewString(table, "id")
	f.CityID = field.NewString(table, "city_id")
	f.ForecastDate = field.NewTime(table, "forecast_date")
	f.Temperature = field.NewField(table, "temperature")
	f.Condition = field.NewString(table, "condition")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *forecast) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *forecast) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["id"] = f.ID
	f.fieldMap["city_id"] = f.CityID
	f.fieldMap["forecast_date"] = f.ForecastDate
	f.fieldMap["temperature"] = f.Temperature
	f.fieldMap["condition"] = f.Condition
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f forecast) clone(db *gorm.DB) forecast {
	f.forecastDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f forecast) replaceDB(db *gorm.DB) forecast {
	f.forecastDo.ReplaceDB(db)
	return f
}

type forecastDo struct{ gen.DO }

type IForecastDo interface {
	gen.SubQuery
	Debug() IForecastDo
	WithContext(ctx context.Context) IForecastDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IForecastDo
	WriteDB() IForecastDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IForecastDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IForecastDo
	Not(conds ...gen.Condition) IForecastDo
	Or(conds ...gen.Condition) IForecastDo
	Select(conds ...field.Expr) IForecastDo
	Where(conds ...gen.Condition) IForecastDo
	Order(conds ...field.Expr) IForecastDo
	Distinct(cols ...field.Expr) IForecastDo
	Omit(cols ...field.Expr) IForecastDo
	Join(table schema.Tabler, on ...field.Expr) IForecastDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IForecastDo
	RightJoin(table schema.Tabler, on ...field.Expr) IForecastDo
	Group(cols ...field.Expr) IForecastDo
	Having(conds ...gen.Condition) IForecastDo
	Limit(limit int) IForecastDo
	Offset(offset int) IForecastDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IForecastDo
	Unscoped() IForecastDo
	Create(values ...*model.Forecast) error
	CreateInBatches(values []*model.Forecast, batchSize int) error
	Save(values ...*model.Forecast) error
	First() (*model.Forecast, error)
	Take() (*model.Forecast, error)
	Last() (*model.Forecast, error)
	Find() ([]*model.Forecast, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Forecast, err error)
	FindInBatches(result *[]*model.Forecast, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Forecast) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IForecastDo
	Assign(attrs ...field.AssignExpr) IForecastDo
	Joins(fields ...field.RelationField) IForecastDo
	Preload(fields ...field.RelationField) IForecastDo
	FirstOrInit() (*model.Forecast, error)
	FirstOrCreate() (*model.Forecast, error)
	FindByPage(offset int, limit int) (result []*model.Forecast, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IForecastDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f forecastDo) Debug() IForecastDo {
	return f.withDO(f.DO.Debug())
}

func (f forecastDo) WithContext(ctx context.Context) IForecastDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f forecastDo) ReadDB() IForecastDo {
	return f.Clauses(dbresolver.Read)
}

func (f forecastDo) WriteDB() IForecastDo {
	return f.Clauses(dbresolver.Write)
}

func (f forecastDo) Session(config *gorm.Session) IForecastDo {
	return f.withDO(f.DO.Session(config))
}

func (f forecastDo) Clauses(conds ...clause.Expression) IForecastDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f forecastDo) Returning(value interface{}, columns ...string) IForecastDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f forecastDo) Not(conds ...gen.Condition) IForecastDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f forecastDo) Or(conds ...gen.Condition) IForecastDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f forecastDo) Select(conds ...field.Expr) IForecastDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f forecastDo) Where(conds ...gen.Condition) IForecastDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f forecastDo) Order(conds ...field.Expr) IForecastDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f forecastDo) Distinct(cols ...field.Expr) IForecastDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f forecastDo) Omit(cols ...field.Expr) IForecastDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f forecastDo) Join(table schema.Tabler, on ...field.Expr) IForecastDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f forecastDo) LeftJoin(table schema.Tabler, on ...field.Expr) IForecastDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f forecastDo) RightJoin(table schema.Tabler, on ...field.Expr) IForecastDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f forecastDo) Group(cols ...field.Expr) IForecastDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f forecastDo) Having(conds ...gen.Condition) IForecastDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f forecastDo) Limit(limit int) IForecastDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f forecastDo) Offset(offset int) IForecastDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f forecastDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IForecastDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f forecastDo) Unscoped() IForecastDo {
	return f.withDO(f.DO.Unscoped())
}

func (f forecastDo) Create(values ...*model.Forecast) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f forecastDo) CreateInBatches(values []*model.Forecast, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f forecastDo) Save(values ...*model.Forecast) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f forecastDo) First() (*model.Forecast, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Forecast), nil
	}
}

func (f forecastDo) Take() (*model.Forecast, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Forecast), nil
	}
}

func (f forecastDo) Last() (*model.Forecast, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Forecast), nil
	}
}

func (f forecastDo) Find() ([]*model.Forecast, error) {
	result, err := f.DO.Find()
	return result.([]*model.Forecast), err
}

func (f forecastDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Forecast, err error) {
	buf := make([]*model.Forecast, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f forecastDo) FindInBatches(result *[]*model.Forecast, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f forecastDo) Attrs(attrs ...field.AssignExpr) IForecastDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f forecastDo) Assign(attrs ...field.AssignExpr) IForecastDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f forecastDo) Joins(fields ...field.RelationField) IForecastDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f forecastDo) Preload(fields ...field.RelationField) IForecastDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f forecastDo) FirstOrInit() (*model.Forecast, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Forecast), nil
	}
}

func (f forecastDo) FirstOrCreate() (*model.Forecast, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Forecast), nil
	}
}

func (f forecastDo) FindByPage(offset int, limit int) (result []*model.Forecast, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f forecastDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f forecastDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f forecastDo) Delete(models ...*model.Forecast) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *forecastDo) withDO(do gen.Dao) *forecastDo {
	f.DO = *do.(*gen.DO)
	return f
}