package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model         // Obrigatorio para usar a função .Create()
	Name       string  `gorm:"type:varchar(60); not null"`
	Price      float32 `gorm:"type:decimal(8,10); not null"`
	Stock      bool    `gorm:"type:tiny; not null"`
}
