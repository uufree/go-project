package user_manager

import (
	"database/sql"
	"log"
	"time"
)

const (
	OperatorFromUid        = 1
	OperatorFromName       = 2
	OperatorFromAge        = 3
	OperatorFromUpdateTime = 4
	OperatorFromAll        = 5
)

type User struct {
	Uid        int
	Name       string
	Age        int
	UpdateTime time.Time
}

func (u *User) String() {
	log.Printf("Uid: %d, Name: %s, Age: %d, UpdateTime: %v", u.Uid, u.Name, u.Age, u.UpdateTime)
}

type Mysql struct {
	db      *sql.DB
	LastUid int
}

func (*Mysql) QueryUser(operatorType int, key interface{}) ([]*User, error) {

}

func (*Mysql) DeleteUser(operatorType int, key interface{}) error {

}

func (*Mysql) InsertUser(operatorType int, key interface{}) error {

}

func (*Mysql) UpdateUser(operatorType int, key interface{}) error {

}
