package data


type Diary struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       string `json:"id"`
}

func (name *Diary) SetUsername(username string) {
	name.Username = username
}

func (username *Diary) GetUsername() string {
	return username.Username
}

func (password *Diary) SetPassword(key string) {
	password.Password = key
}

func (password *Diary) GetPassword() string {
	return password.Password
}

func (id *Diary) SetId(identity string) {
	id.Id = identity
}

func (id *Diary) GetId() string {
	return id.Id
}

