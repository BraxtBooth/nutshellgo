package nutshellgo

// Product represents a product or service in Nutshell.
type Product struct {
	ID         string         `json:"id"`
	EntityType string         `json:"entityType"`
	Rev        string         `json:"rev"`
	Name       string         `json:"name"`
	Type       int            `json:"type"` // 0 = product, 1 = service
	SKU        string         `json:"sku,omitempty"`
	Unit       string         `json:"unit,omitempty"`
	Prices     []ProductPrice `json:"prices,omitempty"`
}

// ProductPrice represents the price details of a product.
type ProductPrice struct {
	ID       string `json:"id"`
	MarketID int    `json:"marketId"`
	Price    Price  `json:"price"`
}
