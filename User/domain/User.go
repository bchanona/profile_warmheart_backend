package domain

type User struct {
	User_id int `json:"id_usuario,omitempty"`
	Name string `json:"name"`
	Subnames string `json:"subnames"`
	Email string `json:"email"`
	Password string `json:"password"`
	Premium bool `json:"premium"`
	Device_id int `json:"device_id"`
}

type UpdateUser struct {
	User_id int `json:"id_usuario,omitempty"`
	Name string `json:"name"`
	Subnames string `json:"subnames"`
	Email string `json:"email"`
	Device_id int `json:"device_id"`
}

type Login struct {
	User_id int `json:"id_usuario,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
	Premium bool `json:"premium"`
}

type UpdateStatus struct {
	User_id int `json:"id_usuario,omitempty"`
	Email string `json:"email"`
	Premium bool `json:"premium"`
	Device_id int `json:"device_id"`
}