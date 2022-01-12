package entity

type User struct {
	UserName string `json:"userName"`
	Name     string `json:"name"`
	SurName  string `json:"surName"`
	Password string `json:"password"`
	Age      string `json:"age"`
}

func (user User) ControlUserInfo() bool {
	if user.UserName == "" {
		return false
	}

	if user.Name == "" {
		return false
	}

	if user.SurName == "" {
		return false
	}

	if user.Password == "" {
		return false
	}

	if user.Age == "" {
		return false
	}

	return true
}
