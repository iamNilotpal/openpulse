package modified_by

type ModifiedBy struct {
	Id        int
	Email     string
	FirstName string
	LastName  string
}

func New(id int, email, firstName, lastName string) ModifiedBy {
	return ModifiedBy{
		Id:        id,
		Email:     email,
		LastName:  lastName,
		FirstName: firstName,
	}
}
