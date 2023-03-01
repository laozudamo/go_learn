package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 于正常的结构体字段，你也可以通过标签 embedded 将其嵌入，例如

// 模型
type Product struct {
	gorm.Model // 继承关系
	Code       string
	Price      uint
}

// var db *mysql

func main() {
	dsn := "root:root123456@/gp_learn?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 创建表
	db.AutoMigrate(&Product{})

}

func create_data(db *gorm.DB) {
	db.Create(&Product{Code: "21142", Price: 1000})
}

/*
	1.根据主键查询
	2.根据条件查询
*/
func search_data() {

}

func delete_data() {

}

/*
	1.更新多个字段
	2.更新某个字段
*/
func update_data() {

}

/*
	增删改查
*/

/*
	对象关系映射 (Object Relational Mapping，简称ORM）模式是一种为了解决面向对象与关系数据库(如mysql数据库）存在的互不匹配的现象的技术。简单的说，ORM是通过使用
	描述对象和数据库之间映射的元数据，将程序中的对象自动持久化到关系数据库中。
*/
