package main

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_dao(t *testing.T) {
	// 连接 db
	db, err := gorm.Open(mysql.Open("dsn://127.0.0.1:3306"), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}

	// 构造 user dao
	dao := NewUserDAO(db)
	ctx := context.Background()

	// 根据id查询
	user1, err := dao.GetUser(ctx, WithId(1))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("user1: %+v", user1)

	// 根据名称+年龄查询
	user2, err := dao.GetUser(ctx, WithName("小洪"), WithAge(18))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("user2: %+v", user2)

	// 根据名称+城市+手机号查询
	user3, err := dao.GetUser(ctx, WithName("小红"), WithPhone("11822"), WithCityID(1))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("user3: %+v", user3)

	// 根据名称+创建时间查询
	user4, err := dao.GetUser(ctx, WithName("小张"), WithCreateTime(time.Now().AddDate(0, 0, -7), time.Now()))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("user4: %+v", user4)
}
