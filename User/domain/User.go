package domain

type User struct {
	User_id int `json:"user_id,omitempty"`
	Name string `json:"name"`
	Surnames string `json:"surnames"`
	Email string `json:"email"`
	Password string `json:"password"`
	Premium bool `json:"premium"`
	Device_id int `json:"device_id"`
}

type UpdatePassword struct {
	User_id int `json:"user_id,omitempty"`
	Password string `json:"password"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}

type GetByID struct {
	User_id int `json:"user_id,omitempty"`
	Name string `json:"name"`
	Subnames string `json:"subnames"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type UpdateStatus struct {
	User_id int `json:"user_id,omitempty"`
	Premium bool `json:"premium"`
}