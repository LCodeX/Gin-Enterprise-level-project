package models

type Packages struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	OriginalPrice float64 `json:"original_price"`
	DiscountPrice float64 `json:"discount_price"`
}

func (Packages) TableName() string {
	return "package"
}
