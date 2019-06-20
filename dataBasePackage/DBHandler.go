package dataBasePackage

import (
	"database/sql"
	"fmt"
	"strconv"
)

type DataBaseHandler struct {
	COL_1 string
	COL_2 string
	COL_3 string
	COL_4 string
	COL_5 string
	TABLE_NAME string
	SQL_CREATE_ENTRIES string
}

func NewDB(colunms [7]string) *DataBaseHandler {
	database := &DataBaseHandler{
		COL_1 : colunms[0],
		COL_2 : colunms[1],
		COL_3 : colunms[2],
		COL_4 : colunms[3],
		COL_5 : colunms[4],
		TABLE_NAME : colunms[5],
		SQL_CREATE_ENTRIES : colunms[6],
	}
	return database
}

func (db DataBaseHandler) CreatDB() {
	db.SQL_CREATE_ENTRIES = "CREATE TABLE IF NOT EXISTS " + db.TABLE_NAME + " (" +
		db.COL_1 + " INTEGER PRIMARY KEY," +
		db.COL_2 + " TEXT," +
		db.COL_3 + " TEXT," +
		db.COL_4 + " TEXT," +
		db.COL_5 + " TEXT)"
	database, _ := sql.Open("sqlite3", "./Expenses.db")
	statement, _ := database.Prepare(db.SQL_CREATE_ENTRIES)
	_,err := statement.Exec()
	checkErr(err)
}

func (db DataBaseHandler) Insert(s string, s2 string, s3 string, s4 string) string {
	fmt.Println("Opening db ....")
	database, err := sql.Open("sqlite3", "./Expenses.db")
	checkErr(err)
	statement, err := database.Prepare("INSERT INTO " + db.TABLE_NAME + "(" + db.COL_2+ ","+ db.COL_3+ ","+db.COL_4+ ","+db.COL_5+")VALUES (?, ?, ?, ?)")
	checkErr(err)
	res,_ := statement.Exec(s,s2,s3,s4)
	id,_ := res.LastInsertId()
	return strconv.FormatInt(id,10)
}

func (db DataBaseHandler) Delete(id string) {
	fmt.Println("Opening db ....")
	database, err := sql.Open("sqlite3", "./Expenses.db")
	checkErr(err)

	statement, err := database.Prepare("DELETE FROM "+ db.TABLE_NAME+ " WHERE ID=?")
	checkErr(err)
	statement.Exec(id)
}

func (db DataBaseHandler) GetByCol(colType string,col string) *sql.Rows{
	database, err := sql.Open("sqlite3", "./Expenses.db")
	checkErr(err)
	// query
	ans, err := database.Query("SELECT * FROM " + db.TABLE_NAME + "where " + colType + "=" + col)
	checkErr(err)
	return ans
}
func (db DataBaseHandler) GetAll(s string) {
	database, err := sql.Open("sqlite3", "./Expenses.db")
	checkErr(err)
	// query
	rows, err := database.Query("SELECT * FROM " + db.TABLE_NAME )
	checkErr(err)
	var uid string
	var date string
	var category string
	var product string
	var value string

	for rows.Next() {
		err = rows.Scan(&uid, &date, &category, &product, &value)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(date)
		fmt.Println(category)
		fmt.Println(product)
		fmt.Println(value)
	}
}


func checkErr(err error) {
	if err != nil {
		fmt.Println("Error reading....")
		panic(err)
	}
}