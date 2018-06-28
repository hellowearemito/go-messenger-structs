package template

const TemplateTypeReceipt TemplateType = "receipt"

type ReceiptTemplate struct {
	TemplateBase
	RecipientName string            `json:"recipient_name"`
	Id            string            `json:"order_number"`
	Currency      string            `json:"currency"`
	PaymentMethod string            `json:"payment_method"`
	Timestamp     int64             `json:"timestamp,omitempty"`
	Url           string            `json:"order_url,omitempty"`
	Items         []OrderItem       `json:"elements"`
	Address       *OrderAddress     `json:"address,omitempty"`
	Summary       OrderSummary      `json:"summary"`
	Adjustments   []OrderAdjustment `json:"adjustments,omitempty"`
}

type OrderItem struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
	Price    int64  `json:"price,omitempty"`
	Currency string `json:"currency,omitempty"`
	ImageURL string `json:"image_url,omiempty"`
}

type OrderAddress struct {
	Street1    string `json:"street_1"`
	Street2    string `json:"street_2,omitempty"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Country    string `json:"country"`
}

type OrderSummary struct {
	TotalCost    int `json:"total_cost,omitempty"`
	Subtotal     int `json:"subtotal,omitempty"`
	ShippingCost int `json:"shipping_cost,omitempty"`
	TotalTax     int `json:"total_tax,omitempty"`
}

type OrderAdjustment struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

func (ReceiptTemplate) Type() TemplateType {
	return TemplateTypeReceipt
}

func (ReceiptTemplate) SupportsButtons() bool {
	return false
}
