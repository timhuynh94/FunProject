package models

type RespBody struct {
	Data Data `json:"data,omitempty"`
}

type Data struct {
	Product Product `json:"product,omitempty"`
}
type Product struct {
	Tcin string `json:"tcin,omitempty"`
	Item Item   `json:"item,omitempty"`
}

type Item struct {
	ProductDescription    ProductDescription    `json:"product_description,omitempty"`
	Enrichment            Enrichment            `json:"enrichment,omitempty"`
	ProductClassification ProductClassification `json:"product_classification,omitempty"`
	PrimaryBrand          PrimaryBrand          `json:"primary_brand,omitempty"`
	CurrentPrice          CurrentPrice          `json:"current_price,omitempty"`
}
type CurrentPrice struct {
	Value        string `json:"value,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}
type ProductDescription struct {
	Title                 string `json:"title,omitempty"`
	DownstreamDescription string `json:"downstream_description,omitempty"`
}
type Enrichment struct {
	Images Images `json:"images,omitempty"`
}

type Images struct {
	PrimaryImageURL string `json:"primary_image_url,omitempty"`
}

type ProductClassification struct {
	ProductTypeName     string `json:"product_type_name,omitempty"`
	MerchandiseTypeName string `json:"merchandise_type_name,omitempty"`
}

type PrimaryBrand struct {
	Name string `json:"name,omitempty"`
}
