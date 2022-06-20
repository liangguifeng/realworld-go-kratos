package mysql_model

type Comment struct {
}

func (Comment) TableName() string {
	return ""
}
