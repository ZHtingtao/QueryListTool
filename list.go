package querylisttool

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"
)

func QueryDataList(key string, params map[string]interface{}) (string, error) {
	qc := QueryConf{
		Key: key,
	}

	err := qc.QueryByKey()
	if err != nil {
		return "", err
	}

	paramValues := []interface{}{}
	paramWheres := []string{}

	queryDataSql := "select * from (" + qc.Sql + " ) t where 1=1 "
	if len(params) > 0 {

		for _, v := range qc.Cols {

			if params[v.Name] != nil {
				paramValues = append(paramValues, params[v.Name])
				paramWheres = append(paramWheres, v.Name)
			}

		}
	}

	var valueMap map[int]map[string]*sql.RawBytes
	if len(paramWheres) > 0 {
		log.Println("query params : ", paramValues)
		queryDataSql += " and " + strings.Join(paramWheres, "=? and ") + " =?"

		log.Println("query sql :", queryDataSql)
		_, valueMap = Query(queryDataSql, paramValues...)
	} else {
		_, valueMap = Query(queryDataSql)
	}
	log.Println("query data sql is :", queryDataSql)

	list_p := parseData(&valueMap, &qc)

	j, err := json.Marshal(struct {
		List    []map[string]string
		Columes []ColumnConf
	}{
		List:    *list_p,
		Columes: qc.Cols,
	})

	if err != nil {
		panic(err)
	}
	return string(j), nil

}

func parseData(data *map[int]map[string]*sql.RawBytes, qc *QueryConf) *[]map[string]string {

	colMap := make(map[string]*ColumnConf)
	list := make([]map[string]string, len(*data))
	for _, v := range qc.Cols {
		colMap[v.Name] = &v
	}

	for i, m := range *data {
		list[i] = make(map[string]string)
		for key, v := range m {

			list[i][key] = string(*v)
		}
	}
	return &list
}
