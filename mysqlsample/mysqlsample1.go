package mysqlsample

import (
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db1 *sqlx.DB

type Employee1 struct {
	Id     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Age    int    `db:"age" json:"age"`
	Gender string `db:"gender" json:"gender"`
}

func open1() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	var err error
	db1, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if err = db1.Ping(); err != nil {
			fmt.Println(err)
			return
		}
		db1.SetMaxOpenConns(10)
		db1.SetMaxIdleConns(2)
	}
}

func query1(id int) {
	var e Employee1
	if err := db1.Get(&e, "select id, name, age, gender from employee where id = ?", id); err != nil {
		fmt.Println(err)
		log.Fatalf("error occur %v\n", err)
		return
	} else {
		fmt.Println(e)
	}

	fmt.Println("==================query more====================")

	var empList []Employee1
	if err := db1.Select(&empList, "select id, name, age, gender from employee where id > ?", id); err != nil {
		log.Fatalf("error occur %v\n", err)
		return
	} else {
		fmt.Println(empList)
	}
}

/*
func stringFilter(src []interface{}) []interface{} {
	dst := make([]interface{}, 0, len(src))
	for _, d := range src {
		b, ok := d.([]byte)
		if ok {
			dst = append(dst, string(b))
		} else {
			dst = append(dst, d)
		}
	}
	return dst
}
*/

func queryWithStmt1(id int) {
	if rows, err := db1.NamedQuery("select id, name, age, gender from employee where id > :id", map[string]interface{}{"id": id}); err != nil {
		log.Fatalf("error occur %v\n", err)
		return
	} else {
		defer rows.Close()
		var uList []Employee1 = make([]Employee1, 0, 2)

		for rows.Next() {
			var e Employee1
			if err := rows.StructScan(&e); err != nil {
				log.Fatalf("error occur %v\n", err)
				return
			} else {
				uList = append(uList, e)
				//fmt.Println(len(uList), cap(uList))
			}
		}
		if data, err := json.Marshal(uList); err != nil {
			log.Fatalf("error occur %v\n", err)
			return
		} else {
			fmt.Println(string(data))
		}
	}
}

/*
func exec1(e Employee1, updated bool) {
	if updated {
		if result, err := db1.Exec("update employee set age=?,name=? where id=?", e.Age, e.Name, e.Id); err != nil {
			log.Fatalf("error occur %v\n", err)
			return
		} else {
			if row, err := result.RowsAffected(); err != nil {
				log.Fatalf("error occur %v\n", err)
				return
			} else {
				fmt.Println("row updated :", row)
			}
		}
	} else {
		if result, err := db1.Exec("insert into employee(name,age,gender) values(?,?,?)", e.Name, e.Age, e.Gender); err != nil {
			log.Fatalf("error occur %v\n", err)
			return
		} else {
			if id, err := result.LastInsertId(); err != nil {
				log.Fatalf("error occur %v\n", err)
				return
			} else {
				fmt.Println("new id :", id)
			}
		}
	}
}
*/

func Trans() {
	if tx, err := db1.Beginx(); err != nil {
		log.Fatalf("error occur %v\n", err)
		return
	} else {
		_, _ = tx.Exec("update employee set age=37 where id=3")
		_, err = tx.Exec("update employee set age=37 where id=4")
		if err != nil {
			log.Fatalf("error occur %v\n", err)
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}
}

func MySqlSample1() {
	open1()
	query1(1)
	//e := Employee1{
	//	Id:     4,
	//	Name:   "benson",
	//	Age:    36,
	//	Gender: "male",
	//}
	//exec1(e, true)
	queryWithStmt1(1)
	fmt.Println("=======================test tran========================")
	//Trans()
}
