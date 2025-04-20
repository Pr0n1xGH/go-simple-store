package model

type CartItem struct {
	ID       uint `gorm:"primarykey"`
	CartID   uint
	Product  string `validate:"required"`
	Quantity int    `validate:"gte=1"`

	Cart Cart `gorm:"foreingKey:CartID"`
}
