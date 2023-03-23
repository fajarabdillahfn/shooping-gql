package model

type Cart struct {
	Products   []*CheckoutProduct `json:"products"`
	TotalPrice float64            `json:"total_price"`
}

type CheckoutProduct struct {
	Sku        string  `json:"sku"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type Product struct {
	Sku      string  `json:"sku" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ProductInput struct {
	Sku string `json:"sku"`
}
