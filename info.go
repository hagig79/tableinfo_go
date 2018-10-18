/*
テーブル定義情報を解析する
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
列情報
*/
type Column struct {
	Field   string
	Type    string
	Null    string
	Key     sql.NullString
	Default sql.NullString
	Extra   sql.NullString
}

/*
指定したテーブルの列情報を列挙する
*/
func getTableInfo(db *sql.DB, tableName string) []Column {
	rows, err := db.Query("DESC " + tableName)
	if err != nil {
		panic(err.Error())
	}

	columns := []Column{}

	for rows.Next() {
		var r Column
		err = rows.Scan(&r.Field, &r.Type, &r.Null, &r.Key, &r.Default, &r.Extra)
		if err != nil {
			panic(err.Error())
		}
		columns = append(columns, r)
	}

	return columns
}

func main() {
	db, err := sql.Open("mysql", "user:@/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	columns := getTableInfo(db, "table")

	for _, column := range columns {
		fmt.Println(column)
	}
}
