// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package user

import (
	"fmt"

	orm "github.com/gogf/gf-demos/global"
)

type User struct {
	Id        uint64 `gorm:"id,primary" json:"id"`                //
	Uuid      uint64 `gorm:"uuid"       json:"uuid"`              //
	Username  string `gorm:"username"   json:"username"`          //
	Name      string `gorm:"name"       json:"name"`              //
	Avatar    string `gorm:"avatar"     json:"avatar"`            //
	Password  string `gorm:"password"   json:"password"`          //
	Status    string `gorm:"status"     json:"status"`            //
	Sign      string `gorm:"sign"       json:"sign"`              //
	CreatedAt string `gorm:"column:created_at" json:"created_at"` //
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"` //
}

func (e *User) GetList() (User, error) {
	var user User
	if err := orm.Eloquent.Table("users").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (e *User) GetOne(username string) (User, error) {
	var user User
	if err := orm.Eloquent.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (e *User) GetProfile(id int8) (User, error) {
	fmt.Println(id)
	var user User
	if err := orm.Eloquent.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
