package query

// Join defines join information in query.
type Join struct {
	Mode       string
	Collection string
	From       string
	To         string
	Arguments  []interface{}
}

func (join Join) Build(query *Query) {
	// TODO: infer from and to when not specified
	query.JoinClause = append(query.JoinClause, join)
}

func JoinWith(mode string, collection string, from string, to string) Join {
	return Join{
		Mode:       mode,
		Collection: collection,
		From:       from,
		To:         to,
	}
}

func JoinFragment(expr string, args ...interface{}) Join {
	return Join{
		Mode:      expr,
		Arguments: args,
	}
}

// func Join(collection string) Join {
// 	return JoinWith("JOIN", collection, "", "")
// }

func JoinOn(collection string, from string, to string) Join {
	return JoinWith("JOIN", collection, from, to)
}

func JoinInner(collection string) Join {
	return JoinInnerOn(collection, "", "")
}

func JoinInnerOn(collection string, from string, to string) Join {
	return JoinWith("INNER JOIN", collection, from, to)
}

func JoinLeft(collection string) Join {
	return JoinLeftOn(collection, "", "")
}

func JoinLeftOn(collection string, from string, to string) Join {
	return JoinWith("LEFT JOIN", collection, from, to)
}

func JoinRight(collection string) Join {
	return JoinRightOn(collection, "", "")
}

func JoinRightOn(collection string, from string, to string) Join {
	return JoinWith("RIGHT JOIN", collection, from, to)
}

func JoinFull(collection string) Join {
	return JoinFullOn(collection, "", "")
}

func JoinFullOn(collection string, from string, to string) Join {
	return JoinWith("FULL JOIN", collection, from, to)
}