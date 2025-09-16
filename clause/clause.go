package clause

import "strings"

type ClauseType int

const (
	SELECT ClauseType = iota
	WHERE
	ORDERBY
	LIMIT
)

type ClauseBuilder struct {
	clauseSql    map[ClauseType]string
	clauseParams map[ClauseType][]interface{}
}

func (c *ClauseBuilder) Set(cType ClauseType, params ...interface{}) {
	// lazy initial
	if c.clauseSql == nil {
		c.clauseSql = make(map[ClauseType]string)
		c.clauseParams = make(map[ClauseType][]interface{})
	}

	genFunc := clauseGenerators[cType]
	clauseSql, clauseParams := genFunc(params...)
	c.clauseSql[cType] = clauseSql
	c.clauseParams[cType] = clauseParams
}

func (c *ClauseBuilder) Build(order ...ClauseType) (string, []interface{}) {
	var clauseSqls []string
	var sqlParams []interface{}

	for _, clause := range order {
		clauseSql := c.clauseSql[clause]
		clauseParams := c.clauseParams[clause]

		clauseSqls = append(clauseSqls, clauseSql)
		sqlParams = append(sqlParams, clauseParams...)
	}
	return strings.Join(clauseSqls, " "), sqlParams
}
