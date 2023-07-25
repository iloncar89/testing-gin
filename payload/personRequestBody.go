package payload

type PersonRequestBody struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	YearOfBirth int    `json:"yearOfBirth"`
}
