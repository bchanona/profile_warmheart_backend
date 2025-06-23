package domain

type Supervisor struct {
	Supervisor_id int    `json:"supervisor_id"`
	Name          string `json:"supervisor_name"`
	Surnames      string `json:"supervisor_surname"`
	Email         string `json:"supervisor_email"`
	Password      string `json:"supervisor_password"`
	User_id       int    `json:"user_id"`
}

type UpdateSupervisor struct {
	Name     string `json:"supervisor_name"`
	Surnames string `json:"supervisor_surname"`
	Email    string `json:"supervisor_email"`
}

type GetSupervisorByUserID struct {
	User_id       int    `json:"user_id"`
	Supervisor_id int    `json:"supervisor_id"`
	Name          string `json:"supervisor_name"`
	Surnames      string `json:"supervisor_surname"`
	Email         string `json:"supervisor_email"`
}

type GetUserBySupervisorID struct {
	Supervisor_id int    `json:"supervisor_id"`
	User_id       int    `json:"user_id"`
	Device_id     int    `json:"user_device"`
	Name          string `json:"user_name"`
	Surnames      string `json:"user_surname"`
	Email         string `json:"user_email"`
	Premium       bool   `json:"user_premium"`
}
type UpdatePassword struct {
	Password string `json:"supervisor_password"`
}
