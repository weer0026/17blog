package model

import "time"

type Model struct {
	ID      uint64 `gorm:"primary_key"`
	Created time.Time
	Updated time.Time
	Deleted *time.Time
}

var Models = []interface{}{
&Article{}, &Category{},
}
