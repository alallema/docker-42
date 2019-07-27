package Users

type User struct {
	Lastname		string `json:"lastName"`
	Firstname		string `json:"firstName"`
	Email			string `json:"email"`
	Phone			string `json:"phoneNumber"`
}

type GetUserResponse struct {
	User	User `json:"user"`
}

type AllUserResponse struct {
	Total		int `json:"total"`
	User		[]User `json:"user"`
}

//** ERROR **//

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
