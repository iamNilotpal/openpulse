package user

import "github.com/iamNilotpal/openpulse/business/data/order"

var DefaultOrderBy = order.NewBy(OrderByID, order.ASC)

const (
	OrderByID        = "id"
	OrderByFirstName = "first_name"
	OrderByLastName  = "last_name"
	OrderByEmail     = "email"
)
