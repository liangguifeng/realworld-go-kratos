package mysql_model

type Favorite struct {
}

func (Favorite) TableName() string {
	return ""
}
