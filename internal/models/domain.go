package models

type Domain struct {
	Origin string `gorm:"not null;unique,index;type:varchar(100);default:null"`
	Base
	IsActive bool `gorm:"not null;default:false"`
}
