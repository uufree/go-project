package user_manager

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	OperatorFromUid        = 1
	OperatorFromName       = 2
	OperatorFromAge        = 3
	OperatorFromUpdateTime = 4
	OperatorFromAll        = 5
)

const (
	MysqlTableName = "userinfo"
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
	mutex   sync.Mutex
}

func (m *Mysql) Init() error {
	var err error
	m.db, err = sql.Open("mysql", "uuchen:342100@/details?charset=utf8")
	if err != nil {
		log.Println("Open MySql Faild.")
		return errors.New("Open MySQL Driver Failed.")
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.LastUid = 0

	return nil
}

func (m *Mysql) QueryUser(operatorType int, key interface{}) ([]*User, error) {
	var users []*User
	var queryCommand string
	switch operatorType {
	case OperatorFromAge:
		queryKey, ok := key.(int)
		if !ok {
			log.Println("operator and real type not match")
			return nil, errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("select * from %s where age=%d", MysqlTableName, queryKey)
	case OperatorFromName:
		queryKey, ok := key.(string)
		if !ok {
			log.Println("operator and real type not match")
			return nil, errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("select * from %s where name=%s", MysqlTableName, queryKey)
	case OperatorFromUid:
		queryKey, ok := key.(int)
		if !ok {
			log.Println("operator and real type not match")
			return nil, errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("select * from %s where uid=%d", MysqlTableName, queryKey)
	case OperatorFromUpdateTime:
		queryKey, ok := key.(time.Time)
		if !ok {
			log.Println("operator and real type not match")
			return nil, errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("select * from %s where update_time<=%d", MysqlTableName, queryKey.Unix())
	case OperatorFromAll:
		queryKey, ok := key.(User)
		if !ok {
			log.Println("operator and real type not match")
			return nil, errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("select * from %s where age=%d and name=%s and uid=%s and update_time<=%d", MysqlTableName, queryKey.Age, queryKey.Name, queryKey.Uid, queryKey.UpdateTime)
	default:
		log.Println("Invalid Operator Type")
	}

	if len(queryCommand) == 0 {
		log.Println("query command is empty")
		return nil, errors.New("query command is empty")
	}
	log.Println(queryCommand)

	rows, err := m.db.Query(queryCommand)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		if err = rows.Scan(&u.Uid, &u.Name, &u.Age, &u.UpdateTime); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}

func (m *Mysql) DeleteUser(operatorType int, key interface{}) error {
	var queryCommand string
	switch operatorType {
	case OperatorFromAge:
		queryKey, ok := key.(int)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("delete * from %s where age=%d", MysqlTableName, queryKey)
	case OperatorFromName:
		queryKey, ok := key.(string)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("delete * from %s where name=%s", MysqlTableName, queryKey)
	case OperatorFromUid:
		queryKey, ok := key.(int)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("delete * from %s where uid=%d", MysqlTableName, queryKey)
	case OperatorFromUpdateTime:
		queryKey, ok := key.(time.Time)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("delete * from %s where update_time<=%d", MysqlTableName, queryKey.Unix())
	case OperatorFromAll:
		queryKey, ok := key.(User)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("delete * from %s where age=%d and name=%s and uid=%s and update_time<=%d", MysqlTableName, queryKey.Age, queryKey.Name, queryKey.Uid, queryKey.UpdateTime)
	default:
		log.Println("Invalid Operator Type")
	}

	if len(queryCommand) == 0 {
		log.Println("query command is empty")
		return errors.New("query command is empty")
	}
	log.Println(queryCommand)

	result, err := m.db.Exec(queryCommand)
	if err != nil {
		log.Println(err)
		return err
	}
	update_number, _ := result.RowsAffected()
	log.Println("delete rows: ", update_number)
	return nil
}

func (m *Mysql) InsertUser(operatorType int, key interface{}) error {
	u, ok := key.(User)
	if !ok {
		log.Println("convert interface{} failed.")
		return errors.New("convet interface{} failed.")
	}
	queryCommand := fmt.Sprintf("insert into table(uid, name, age, update_time) values(%d, %s, %d, %d)", u.Uid, u.Name, u.Age, u.UpdateTime)
	log.Println(queryCommand)

	result, err := m.db.Exec(queryCommand)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m *Mysql) UpdateUser(operatorType int, uid int, key interface{}) error {
	var queryCommand string
	switch operatorType {
	case OperatorFromAge:
		queryKey, ok := key.(int)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("update %s set age=%d where uid=%d", MysqlTableName, queryKey, uid)
	case OperatorFromName:
		queryKey, ok := key.(string)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("update %s set name=%s where uid=%d", MysqlTableName, queryKey, uid)
	case OperatorFromUid:
		queryKey, ok := key.(int)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("update %s set uid=%d where uid=%d", MysqlTableName, queryKey, uid)
	case OperatorFromUpdateTime:
		queryKey, ok := key.(time.Time)
		if !ok {
			log.Println("operator and real type not match")
			return errors.New("operator and real type not match")
		}
		queryCommand = fmt.Sprintf("update %s set update_time=%d where uid=%d", MysqlTableName, queryKey.Unix(), uid)
	default:
		log.Println("Invalid Operator Type")
	}

	if len(queryCommand) == 0 {
		log.Println("query command is empty")
		return errors.New("query command is empty")
	}
	log.Println(queryCommand)

	result, err := m.db.Exec(queryCommand)
	update_number, _ := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("update rows: ", update_number)
	return nil
}

func (m *Mysql) Destory() error {
	if err := m.db.Close(); err != nil {
		return errors.New("Close MySql Driver Failed.")
	}
	return nil
}
