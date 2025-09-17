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
	INSERT:  _insert,
	VALUES:  _values,
}

/*
	INSERT INTO $tableName ($fields)

arg[0] = tableName
arg[1] = column names
*/
func _insert(params ...interface{}) (string, []interface{}) {
	tableName := params[0]
	columns := strings.Join(params[1].([]string), ",")

	return fmt.Sprintf("INSERT INTO %s (%s)", tableName, columns), []interface{}{}
}

/*
	INSERT INTO $tableName ($fields)

arg[0] = tableName
arg[1] = select fields slices
*/
func _select(params ...interface{}) (string, []interface{}) {
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

/*
VALUES ($v1), ($v2), ...
*/
func _values(params ...interface{}) (string, []interface{}) {
	var sql strings.Builder

	v0 := params[0].([]interface{})
	placeholder := []string{}
	for i := 0; i < len(v0); i++ {
		placeholder = append(placeholder, "?")
	}
	bindStr := strings.Join(placeholder, ", ")

	sql.WriteString("VALUES ")
	vars := []interface{}{}
	for i, vi := range params {
		v := vi.([]interface{})
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))

		if i != len(params)-1 {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}

	return sql.String(), vars
}
