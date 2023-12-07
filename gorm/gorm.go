package gorm

import (
	"blog/common"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var db = common.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Find() {
	var p Product
	db.First(&p, 1)
	fmt.Printf("p: %v\n", p)
	db.First(&p, "code = ?", "1001")
	fmt.Printf("p: %v\n", p)
}

func update() {
	var p Product
	db.First(&p, 1)
	//Update - 将product的price更新为200
	//db.Model(&p).Update("price", 200)
	//Updates - 更新多个字段
	//db.Model(&p).Updates(Product{Price: 201, Code: "1002"})//仅更新非零值字段
	db.Model(&p).Updates(map[string]interface{}{"Price": 100, "Code": "1001"})

}

func delete() {
	var p Product
	db.First(&p, 1)
	db.Delete(&p, 1)
}

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
	Active   bool
}

func CreateUser() {
	user := User{
		Name:     "kite",
		Age:      22,
		Birthday: time.Now(),
		Active:   false,
	}
	db.Create(&user)
}

//会话测试
func test1() {
	db.Session(&gorm.Session{DryRun: true})
}

func txTest() {
	user := User{
		Name:     "lily",
		Age:      22,
		Birthday: time.Now(),
		Active:   false,
	}

	// user2 := User{
	// 	Name:     "rose",
	// 	Age:      22,
	// 	Birthday: time.Now(),
	// 	Active:   false,
	// }

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if err := tx.Create(nil).Error; err != nil {
			return err
		}
		return nil
	})
}

func main() {
	//create()

	//find()

	//update()

	//delete()

	//db.AutoMigrate(&User{})

	//CreateUser()

	txTest()
}
