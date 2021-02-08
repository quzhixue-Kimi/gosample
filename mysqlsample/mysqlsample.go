package mysqlsample

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Employee struct {
	id     int
	name   string
	age    int
	gender string
}

func open() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if err = db.Ping(); err != nil {
			fmt.Println(err)
			return
		}
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(2)
	}
}

func query() {
	var e1 Employee
	if err := db.QueryRow("select id, name, age, gender from employee where id=?", 1).Scan(&e1.id, &e1.name, &e1.age, &e1.gender); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e1)
	fmt.Println("===========query more===========")
	if rows, err := db.Query("select id, name, age, gender from employee"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer rows.Close()

		for rows.Next() {
			var e Employee
			if err := rows.Scan(&e.id, &e.name, &e.age, &e.gender); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(e)
			}
		}
	}
}

func queryWithStmt(id int) {
	if stmt, err := db.Prepare("select id, name, age, gender from employee where id > ?"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer stmt.Close()
		if rows, err := stmt.Query(id); err != nil {
			fmt.Println(err)
			return
		} else {
			defer rows.Close()

			for rows.Next() {
				var e Employee
				if err := rows.Scan(&e.id, &e.name, &e.age, &e.gender); err != nil {
					fmt.Println(err)
					return
				} else {
					fmt.Println(e)
				}
			}
		}
	}
}

func exec(e Employee, updated bool) {
	if !updated {
		if result, err := db.Exec("insert into employee(name, age, gender) values(?,?,?)", e.name, e.age, e.gender); err != nil {
			fmt.Println(err)
			return
		} else {
			if id, err := result.LastInsertId(); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(id)
			}
		}
	} else {
		if result, err := db.Exec("update employee set age = ? where id = ?", e.age, e.id); err != nil {
			fmt.Println(err)
			return
		} else {
			if rowCount, err := result.RowsAffected(); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println("updated rows", rowCount)
			}
		}
	}
}

func MySqlSample() {
	open()
	query()
	e1 := Employee{
		name:   "kimi",
		age:    36,
		gender: "male",
	}
	exec(e1, false)
	query()
	fmt.Println("==========================")
	queryWithStmt(2)
}
