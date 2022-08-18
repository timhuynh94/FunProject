package models

type RespBody struct {
	Data Data `json:"data"`
}

type Data struct {
	Product Product `json:"product"`
}
type Product struct {
	Tcin string `json:"tcin"`
	Item Item   `json:"item"`
}

type Item struct {
	ProductDescription    ProductDescription    `json:"product_description"`
	Enrichment            Enrichment            `json:"enrichment"`
	ProductClassification ProductClassification `json:"product_classification"`
	PrimaryBrand          PrimaryBrand          `json:"primary_brand"`
	CurrentPrice          CurrentPrice          `json:"current_price"`
}
type CurrentPrice struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency_code"`
}
type ProductDescription struct {
	Title                 string `json:"title"`
	DownstreamDescription string `json:"downstream_description"`
}
type Enrichment struct {
	Images Images `json:"images"`
}

type Images struct {
	PrimaryImageURL string `json:"primary_image_url"`
}

type ProductClassification struct {
	ProductTypeName     string `json:"product_type_name"`
	MerchandiseTypeName string `json:"merchandise_type_name"`
}

type PrimaryBrand struct {
	Name string `json:"name"`
}
