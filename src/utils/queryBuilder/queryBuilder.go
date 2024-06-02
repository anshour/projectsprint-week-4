package querybuilder

import (
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
