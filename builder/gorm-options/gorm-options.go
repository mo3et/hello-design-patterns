package main

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// 在需要使用opts (修改配置类)的函数中(GetUser or NewXXX)，传入配置好的，出参为Option函数的配置函数(WithXXX(type T) Option )

type User struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null;autoIncrementIncrement:2" json:"id"` // 主键ID
	Name       string    `gorm:"column:name;type:varchar(256);not null" json:"name"`                                         // 名称
	Age        int64     `gorm:"column:age;type:bigint(20) unsigned;not null" json:"age"`                                    // 名称
	CityID     int       `gorm:"column:city_id;type:int(11) unsigned;not null" json:"city_id"`                               // 城市ID
	Phone      string    `gorm:"column:phone;type:varchar(512);not null" json:"phone"`                                       // 手机号码
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null" json:"create_time"`                              // 创建时间
	ModifyTime time.Time `gorm:"column:modify_time;type:timestamp;not null" json:"modify_time"`                              // 修改时间
}

// User DAO 模块，负责代理和DB中表有关的交互操作
type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// 声明 Option 类型 是函数的类型，入参和出参都是 *gorm.DB
type Option func(db *gorm.DB) *gorm.DB

// 把预设的查询条件进行声明，每天通过闭包的方法接收使用方传入的属性值
// 在通过 gorm.DB链式调用的方式，将对应的筛选方式以 where 语句(或者我们需要的语法)组装到查询条件语句中

func WithId(id int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func WithName(name string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
}

func WithAge(age int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("age = ?", age)
	}
}

func WithCityID(cityID int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("city_id = ?", cityID)
	}
}

func WithPhone(phone string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("phone = ?", phone)
	}
}

func WithCreateTime(begin, end time.Time) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("create_time >= ? and create_time <= end", begin, end)
	}
}

// 只需声明好通用的查询用户方法 GetUser 然后在入参中追加 Option list
// 通过遍历 Option list，把使用方注入的查询条件拼接到条件语句中，最后一步到位查询操作
func (u *UserDAO) GetUser(ctx context.Context, opts ...Option) (*User, error) {
	db := u.db.WithContext(ctx).Model(&User{})
	for _, opt := range opts {
		db = opt(db)
	}
	var user User
	return &user, db.First(&user).Error
}

func (u *UserDAO) CountUser(ctx context.Context, opts ...Option) (int64, error) {
	db := u.db.WithContext(ctx).Model(&User{})
	for _, opt := range opts {
		db = opt(db)
	}
	var cnt int64
	return cnt, db.Count(&cnt).Error
}

func main() {
}
