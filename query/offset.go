package query

type Offset int

func (offset Offset) Build(query *Query) {
	query.OffsetResult = offset
}