package pop

// Having will append a HAVING clause to the query
func (q *Query) Having(condition string, args ...interface{}) *Query {
	if q.RawSQL.Fragment != "" {
		Logger.Warning("Query is set up to use raw SQL, not adding HAVING clause")
		return q
	}
	q.havingClauses = append(q.havingClauses, HavingClause{condition, args})

	return q
}
