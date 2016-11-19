package user

//SecurityQuestion model
type SecurityQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

//User model
type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email", omitempty`
	Collection []SecurityQuestion
}
