package mysql_model

type User struct {
}

func (User) TableName() string {
	return ""
}
