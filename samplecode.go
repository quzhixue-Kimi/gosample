package main

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

//func sendData(ch1 chan int) {
//	for i := 0; i < 10; i++ {
//		ch1 <- i
//		time.Sleep(500 * time.Millisecond)
//		fmt.Println("写入数据 ", i)
//	}
//	close(ch1)
//}
//
//func readData(ch1 chan int, exit chan bool) {
//	for {
//		v, ok := <-ch1
//		if !ok {
//			fmt.Println("已经读取了所有的数据，", ok)
//			break
//		}
//		time.Sleep(700 * time.Millisecond)
//		fmt.Println("取出数据：", v)
//	}
//	exit <- true
//	close(exit)
//}

func main() {
	open()
	query()
	e1 := Employee{
		name:   "kitty",
		age:    22,
		gender: "female",
	}
	exec(e1, true)
	query()
	fmt.Println("==========================")
	queryWithStmt(2)
	//jobs := make(chan int, 5)
	//results := make(chan int, 5)

	//for i := 1; i <= 2; i++ {
	//	go workpool.WorkerPool(i, jobs, results)
	//}

	//for i := 1; i <= 5; i++ {
	//	jobs <- i
	//}
	//close(jobs)

	////for v := range results {
	////	fmt.Println(v)
	////}

	////for {
	////	if v, ok := <-results; !ok {
	////		break
	////	} else {
	////		fmt.Println(v)
	////	}
	////}

	//for i := 1; i <= 5; i++ {
	//	x := <-results
	//	fmt.Println(x)
	//}

	//	selectsample.SelectChannel()

	//msg := make(chan *io.LogInfo, 5)
	//exit := make(chan bool, 1)

	//l := io.NewLogInfo("hello", "Info")

	//for {
	//	go l.FileReadOperator(msg)
	//	go l.FileCreateOperator(msg, exit)
	//	if _, ok := <-exit; !ok {
	//		break
	//	}
	//}

	//ch1 := make(chan int, 5)
	//exit := make(chan bool, 1)
	//for {
	//	go sendData(ch1)
	//	go readData(ch1, exit)
	//	if _, ok := <-exit; !ok {
	//		break
	//	}
	//}
}
