package models

import (
	"errors"
	"time"
)

type Profile struct {
	Id        int64     ``
	Gender    string    `orm:"size(4)"`
	Age       int       ``
	Address   string    ``
	LastLogin time.Time `orm:"auto_now;type(datetime)"`
	Users     *Users    `orm:"null;rel(one);on_delete(cascade)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

func (u *Profile) TableName() string {
	return "users_profile"
}

type Users struct {
	Id       int64     ``
	Username string    `orm:"size(32)"`
	Email    string    `orm:"size(32)"`
	Uid      string    `orm:"size(32)"`
	Password string    ``
	Mobile   string    `orm:"size(18);index;unique"`
	IsSupper bool      ``
	IsStuff  bool      ``
	Activate bool      ``
	Profile  *Profile  `orm:"reverse(one)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

func (u *Users) TableName() string {
	return "users"
}

func AddUser(u Users) int64 {
	return u.Id
}

func GetUser(uid string) (u *Users, err error) {
	return nil, errors.New("User not exists")
}

func GetAllUsers() {
}

func UpdateUser(uid string, uu *Users) (a *Users, err error) {
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	return false
}

func DeleteUser(uid string) {
}
