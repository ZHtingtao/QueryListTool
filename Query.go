package querylisttool

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Query(sqlStr string, args ...interface{}) ([]string, map[int]map[string]*sql.RawBytes) {

	rows, err := db.Query(sqlStr, args...)
	checkErr(err)
	var valueMap = make(map[int](map[string](*sql.RawBytes)))
	colNames, err := rows.Columns()
	checkErr(err)
	for i := 0; rows.Next(); i++ {
		valMap := make(map[string]*sql.RawBytes)

		scanArgs := make([]interface{}, len(colNames))
		for i := range colNames {
			valMap[colNames[i]] = new(sql.RawBytes)

			scanArgs[i] = valMap[colNames[i]]
		}

		rows.Scan(scanArgs...)

		valueMap[i] = valMap
		log.Println(valMap)
	}
	log.Println(len(valueMap))
	return colNames, valueMap

}

func checkErr(err error) {
	if err != nil {
		log.Println("database error: ")
		panic(err)
	}
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(192.168.0.26:3306)/gowebpractice?charset=utf8")
	checkErr(err)
}
