package model

type Article struct {
	Model
	Title     string `gorm:"type:varchar(50);size:50"`
	SubTitle  string `gorm:"type:varchar(100);size:100"`
	Content   string `gorm:"type:mediumtext"`
	Status    int    `gorm:"type:tinyint(1);default:1"`
	ViewCount uint64 `gorm:"default:0"`
}

const (
	ArticleDisabled = iota
	ArticleOk
)
