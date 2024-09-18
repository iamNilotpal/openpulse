package order

const (
	ASC  = "ASC"
	DESC = "DESC"
)

var Directions = map[string]string{
	ASC:  "ASC",
	DESC: "DESC",
}

type by struct {
	Field     string
	Direction string
}

func NewBy(field string, direction string) by {
	return by{
		Field:     field,
		Direction: direction,
	}
}
