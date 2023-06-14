package demo

import (
	"fmt"
	"mygorm/database"
	"mygorm/model"
)

func FindOne() {
	var user model.User
	db := database.DB

	// 根据整型主键查找
	result := db.First(&user, 1)

	// 获取第一条记录（主键升序）
	// result := db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	// result := db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	// result := db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	if result.Error != nil {
		// 检查 ErrRecordNotFound 错误
		// errors.Is(result.Error, gorm.ErrRecordNotFound)
		fmt.Println("Not found")
	} else {
		fmt.Println("Found user:", user.Name)
		fmt.Println("RowsAffected", result.RowsAffected)
	}

	// 避免ErrRecordNotFound
	// result := db.Limit(1).Find(&user)

}

func FindAll() {
	var users []model.User
	db := database.DB

	// 根据整型主键查找
	result := db.Find(&users)
	// result := db.Find(&users, []int{1,2,3})

	if result.Error != nil {
		// 检查 ErrRecordNotFound 错误
		// errors.Is(result.Error, gorm.ErrRecordNotFound)
		fmt.Println("error", result.Error)
	} else {
		fmt.Println("Found user num:", result.RowsAffected)
	}
}

func FindWhere() {
	var user model.User
	var users []model.User
	db := database.DB

	// String 条件
	db.Where("name = ?", "jinzhu").First(&user)
	db.Where("age BETWEEN ? AND ?", 18, 30).Find(&users)

	// Struct 条件
	db.Where(&model.User{Name: "jinzhu", Age: 20}).First(&user)

	// Map条件
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)

	// Slice of primary keys
	db.Where([]int64{20, 21, 22}).Find(&users)
}
