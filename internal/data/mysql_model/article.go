package mysql_model

type Article struct {
}

func (Article) TableName() string {
	return ""
}
