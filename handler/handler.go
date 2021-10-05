package handler

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
)

var (

	// mysqlに接続　golang students データベースに接続している
	db, err = sql.Open("mysql", "admin:admin@tcp(localhost:3306)/golang")

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

)

type Student struct {
	ID int `db:"id"`
	Name string `db:"name"`
	Birthday string `db:"birthday"`
	Blood string `db:"blood_type"`
	Constellation string `db:"constellation"`
	Height string `db:"height"`
}

func SelectStudent() echo.HandlerFunc {
	return func(c echo.Context) error {

		// name パラメーターを受け取る
		name := c.Param("name")
		var students = []Student{}
		// select * from students where name = name; を実行し、membersに格納
		dbmap.Select(&students, "SELECT * FROM students WHERE name =" + "'" + name + "'" + ";")
		return c.JSON(http.StatusOK, students)
	}
}

func SelectStudents() echo.HandlerFunc {
	return func(c echo.Context) error {
		var students = []Student{}
		dbmap.Select(&students, "SLECT * FROM students;")
		return c.JSON(http.StatusOK, students)
	}
}