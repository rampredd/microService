package maria

import (
	"database/sql"
	"fmt"
	"microService/common"
	ds "microService/proto/datastore"
	query "microService/proto/query"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db    *sql.DB
	err   error
	dbCfg common.DbCfg
)

func Connect(dbName string) (*sql.DB, error) {
	// Create the database handle
	//	user:password@/dbname example:root:abc123@/db

	dbCfg, err = common.GetDbConfig("db")
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf("%s:%s@/%s", dbCfg.Username, dbCfg.Password, dbName)
	db, err = sql.Open("mysql", str)
	if err != nil {
		return nil, err
	}

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
	return db, nil
}

func getOperator(valType string, i interface{}) (res string) {
	res = ""
	switch valType {
	case "string":
		str := strings.Fields(i.(string))
		switch len(str) {
		case 0:
		case 1:
			res = " = " + "'" + str[0] + "'"
		case 2:
			switch str[0] {
			case "lte":
				res = " <= "
			case "lt":
				res = " < "
			case "gte":
				res = " >="
			case "gt":
				res = " > "
			}
			res += str[1]
		}
	case "int32":
		i32 := i.(int32)
		if i32 > 0 {
			res = fmt.Sprintf("%s%s", " = ", i32)
		}
	}
	return res
}

func getColVal(valType string, i interface{}) (res string, yes bool) {
	res = ""
	yes = false

	switch valType {
	case "string":
		res = i.(string)
		yes = true
	case "int32":
		i32 := i.(int32)
		if i32 > 0 {
			res = fmt.Sprintf("%s", i32)
			yes = true
		}
	}
	return res, yes
}

func SaveRecipe(req *ds.SaveRequest) (string, error) {
	i := 0
	colNames, res, s, colValues := "", "", "", "VALUES("
	cmd := fmt.Sprintf("INSERT INTO %s (", dbCfg.TableName)

	//structure elements into string
	v := reflect.ValueOf(req).Elem()
	typeOfRequest := v.Type()

	for i = 0; i < v.NumField()-1; i++ {
		f := v.Field(i)
		name := typeOfRequest.Field(i).Name
		valType := f.Type().String()
		val := f.Interface()

		//		fmt.Printf("%d: %s %s = %v %v\n", i, name, valType, val, f.IsValid())
		if f.IsValid() {
			if v, yes := getColVal(valType, val); yes == true {
				colNames += fmt.Sprintf("%s,", strings.ToLower(name))
				colValues += fmt.Sprintf("%s,", v)
			}
		}
	}
	//	replace last character by ')'
	//	colNames[len(colNames)-1] = ')'
	colNames = colNames[:len(colNames)-1]
	colNames += ")"
	colValues = colValues[:len(colValues)-1]
	colValues += ")"
	//	colValues[len(colValues)-1] = ')'
	//	cmd += colNames + colValues + ";"

	fmt.Println("cmd is ", cmd)

	rows, err := db.Query(cmd)
	if err != nil {
		return res, err
	}

	for i = 0; rows.Next(); i++ {
		err = rows.Scan(&s)
		if err != nil {
			return res, err
		}
		res += fmt.Sprintf("%s", s)
		s = ""
	}

	return res, nil
}

func GetRecipe(req *query.Request) (string, error) {
	i := 0
	str, res, s := "", "", ""
	cmd := fmt.Sprintf("SELECT recipe FROM %s WHERE ", dbCfg.TableName)

	//structure elements into string
	v := reflect.ValueOf(req).Elem()
	typeOfRequest := v.Type()

	for i = 0; i < v.NumField()-1; i++ {
		f := v.Field(i)
		name := typeOfRequest.Field(i).Name
		valType := f.Type().String()
		val := f.Interface()

		//		fmt.Printf("%d: %s %s = %v %v\n", i, name, valType, val, f.IsValid())
		if f.IsValid() {
			if opr := getOperator(valType, val); opr != "" {
				if cmd[len(cmd)-6:] == "WHERE " {
					str = fmt.Sprintf("%s%s", strings.ToLower(name), opr)
				} else {
					str = fmt.Sprintf("%s%s%s", " AND ", strings.ToLower(name), opr)
				}
				cmd += str
			}
		}
	}
	cmd += ";"

	fmt.Println("cmd is ", cmd)

	rows, err := db.Query(cmd)
	if err != nil {
		return res, err
	}

	for i = 0; rows.Next(); i++ {
		err = rows.Scan(&s)
		if err != nil {
			return res, err
		}
		res += fmt.Sprintf("%s", s)
		s = ""
	}

	return res, nil
}
