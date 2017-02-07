package querylisttool

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

const QueryConfSQL = "select qc.key ,qc.sql,cc.name , cc.display,cc.displayname,cc.condition,cc.id from QUERY_CONF qc , COL_CONF cc where qc.key = cc.key and cc.key= ?"

type QueryConf struct {
	Key  string       `id:"key",type:"string`
	Sql  string       `id:"sql",type:"string"`
	Cols []ColumnConf ``
}

type ColumnConf struct {
	Key         string `id:"key",type:"string"`
	Name        string `id:"name",type:"string"`
	DisplayName string `id:"displayname",type"string"`
	Display     int64  `id:"display" ,type::"int"`
	Condition   int64  `id:"condition",type:"int"`
}

func (q *QueryConf) QueryByKey() error {

	if len(q.Key) <= 0 {
		return errors.New("key is nil or 0 length in queryconf")
	}

	colNames, valueMaps := Query(QueryConfSQL, q.Key)
	log.Println(colNames)

	if len(valueMaps) > 0 {
		r := valueMaps[0]
		qv := reflect.ValueOf(q)
		qt := reflect.TypeOf(*q)

		if num := qt.NumField(); num > 0 {
			for i := 0; i < num; i++ {
				elem := qv.Elem()
				fieldValue := elem.Field(i)
				field := qt.Field(i)
				log.Printf("this feild kind ", fieldValue.Kind())

				if len(field.Tag) > 0 {

					id := field.Tag.Get("id")

					v := reflect.ValueOf(r[id]).Elem().Convert(field.Type)

					fmt.Println(v)

					fieldValue.Set(v)
				}
			}
		}

	}

	for index, value := range valueMaps {
		log.Println(value)
		r := valueMaps[index]
		col := ColumnConf{}

		setConfValue(reflect.TypeOf(col), reflect.ValueOf(&col).Elem(), r)

		q.Cols = append(q.Cols, col)

	}

	return nil
}

func setConfValue(t reflect.Type, elem reflect.Value, r map[string]*sql.RawBytes) {

	log.Println("this type is ", t.Kind())
	if num := t.NumField(); num > 0 {
		for i := 0; i < num; i++ {
			fieldValue := elem.Field(i)
			field := t.Field(i)
			if fieldValue.Kind() != reflect.Struct {
				id := field.Tag.Get("id")

				log.Println("this colume is ", id)
				log.Println("this colume value is ", r[id])
				log.Println("will convert type is ", field.Type)

				vElem := reflect.ValueOf(r[id]).Elem()
				switch fieldValue.Kind() {
				case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:

					i, err := strconv.ParseInt(vElem.Convert(reflect.TypeOf("string")).Interface().(string), 10, 64)
					if err != nil {
						panic(err)
					}

					fieldValue.SetInt(i)

				case reflect.String:
					v := vElem.Convert(field.Type)
					fieldValue.Set(v)
				}

			}
		}
	}
}
