package models

type Users struct {
    ID int `gorm:"column:id" json:"id"`
}

func (d *Users) TableName() string {
    return "users"
}
