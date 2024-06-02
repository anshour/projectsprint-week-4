package querybuilder

import (
	"fmt"
	"strconv"

	"github.com/lib/pq"
)

type Query struct {
	BaseQuery string
	Params    []interface{}
}

func (q *Query) nextParamIndex() string {
	return "$" + strconv.Itoa(len(q.Params)+1)
}

func (q *Query) AppendCondition(column string, operator string, value interface{}) {
	q.BaseQuery += " AND " + column + " " + operator + " " + q.nextParamIndex()
	q.Params = append(q.Params, value)
}

func (q *Query) AppendOrWhere(column1 string, operator1 string, value1 interface{}, column2 string, operator2 string, value2 interface{}) {
	q.BaseQuery += fmt.Sprintf(" AND (%s %s %s OR ", column1, operator1, q.nextParamIndex())
	q.Params = append(q.Params, value1)
	q.BaseQuery += fmt.Sprintf("%s %s %s)", column2, operator2, q.nextParamIndex())
	q.Params = append(q.Params, value2)
}

func (q *Query) AppendWhereAny(column string, value interface{}) {
	q.BaseQuery += " AND " + column + " = " + "ANY" + "(" + q.nextParamIndex() + ")"
	q.Params = append(q.Params, pq.Array(value))
}

func (q *Query) AppendOrder(column string, direction string) {
	q.BaseQuery += " ORDER BY " + column + " " + direction
}

func (q *Query) AppendLimit(limit int) {
	q.BaseQuery += " LIMIT " + q.nextParamIndex()
	q.Params = append(q.Params, limit)
}

func (q *Query) AppendOffset(offset int) {
	q.BaseQuery += " OFFSET " + q.nextParamIndex()
	q.Params = append(q.Params, offset)
}
