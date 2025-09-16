package clause

import (
	"fmt"
	"strings"
)

type generateFunc func(...interface{}) (string, []interface{})

var clauseGenerators = map[ClauseType]generateFunc{
	SELECT:  _select,
	LIMIT:   _limit,
	ORDERBY: _orderBy,
	WHERE:   _where,
}

/*
	INSERT INTO $tableName ($fields)

arg[0] = tableName
arg[1] = column names
*/
func _insert() {

}

/*
	INSERT INTO $tableName ($fields)

arg[0] = tableName
arg[1] = select fields slices
*/
func _select(params ...interface{}) (string, []interface{}) {
	fmt.Println(len(params))
	tableName := params[0]
	fields := strings.Join(params[1].([]string), ",")

	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

/*
	INSERT INTO $tableName ($fields)

arg[0] = where desc
arg[1] = values
*/
func _where(params ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("WHERE %s", params[0]), params[1:]
}

func _orderBy(params ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", params[0]), []interface{}{}
}

func _limit(params ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("LIMIT ?"), params
}
