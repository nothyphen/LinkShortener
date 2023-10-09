package entity

type Short struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Link 	  string `gorm:"type:varchar(250)" json:"-"`
	Short     string `gorm:"type:varchar(7)" json:"-"`
}