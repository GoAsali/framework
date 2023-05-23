package repositories

type Operator string

var (
	OrOperator  Operator = "or"
	AndOperator Operator = "and"
)

type ListQueryOption struct {
	operator   Operator
	conditions []interface{}
	limit      int
	offset     int
}

type ListQueryExecuteFn func(option *ListQueryOption)

func defaultListQuery() *ListQueryOption {
	return &ListQueryOption{
		operator: AndOperator,
		limit:    10,
		offset:   0,
	}
}

// WithOperator List query with a dedicated operator [Or, and]
func WithOperator(op Operator) ListQueryExecuteFn {
	return func(option *ListQueryOption) {
		option.operator = op
	}
}

// WithLimit Set limit for query
// Ex: select * from users limit 10
func WithLimit(limit int) ListQueryExecuteFn {
	return func(option *ListQueryOption) {
		option.limit = limit
	}
}

// WithOffset Set offset sql query
// Ex: select * from users offset 10
func WithOffset(offset int) ListQueryExecuteFn {
	return func(option *ListQueryOption) {
		option.offset = offset
	}
}
