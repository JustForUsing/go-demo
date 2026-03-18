package repos

import (
	"gorm.io/gorm"
)

type QueryBuilder struct {
	model interface{}
	db    *gorm.DB
}

// Value 根据字段名查询单个字段的值
// field 数据库字段名
// value 接收查询结果的变量指针
func (qb *QueryBuilder) Value(filed string, value interface{}) error {
	return qb.db.Model(qb.model).
		Select(filed).
		Scan(value).Error
}

// Count 获取记录数
func (qb *QueryBuilder) Count() (int64, error) {
	var count int64
	if err := qb.db.Model(qb.model).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Exist 检查是否存在
func (qb *QueryBuilder) Exist() (bool, error) {
	var count int64
	if err := qb.db.Model(qb.model).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// First 获取一条记录
// model 接收查询结果的结构体指针
func (qb *QueryBuilder) First(model interface{}) error {
	return qb.db.Model(qb.model).
		Limit(1).
		First(model).Error
}

// Select 选择字段
func (qb *QueryBuilder) Select(query interface{}, args ...interface{}) *QueryBuilder {
	return &QueryBuilder{model: qb.model, db: qb.db.Select(query, args...)}
}

//// FindInt64 根据多个and条件查询用户ID
//func (qb *QueryBuilder) FindInt64(int64Filed string) (int64, error) {
//	var int64Value int64
//	if err := qb.db.Model(qb.model).
//		Select(int64Filed).
//		Scan(&int64Value).Error; err != nil {
//		return 0, err
//	}
//	return int64Value, nil
//}
//
//// Exist 检查是否存在，实测比count更耗时
//func (qb *QueryBuilder) Exist() (bool, error) {
//	var exists bool
//	//转换查询sql
//	sql := qb.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
//		return tx.Model(qb.model).Select("1").Scan(nil)
//	})
//	if err := qb.db.Model(qb.model).
//		Raw("SELECT EXISTS(" + sql + ")").
//		Scan(&exists).Error; err != nil {
//		return false, err
//	}
//	return exists, nil
//}
