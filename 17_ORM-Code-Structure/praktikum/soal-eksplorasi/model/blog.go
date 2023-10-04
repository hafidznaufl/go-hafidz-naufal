package model

type Blog struct {
	ID      uint   `json:"id" form:"id" gorm:"primaryKey"`
	Title   string `json:"title" form:"title"`
	Author  string `json:"author" form:"author"`
	Content string `json:"content" form:"content"`
	UserID  uint   `json:"user_id" form:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
}

type User struct {
	ID    uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name  string `json:"name" form:"name"`
	Blogs []Blog `gorm:"foreignKey:UserID"`
}
