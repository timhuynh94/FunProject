# myRetail RESTful service

## Description
A RESTful service that can retrieve product and
price details by ID, which will aggregate product data from multiple
sources and return it as JSON to the calle

## Tech Stack
- Golang
- Gin Framework
- Redis
- Docker

## Tech used for testing
- Insomina `https://insomnia.rest/`

## Future improvements
- Unit tests using mock for redis db and redsky api call
- Separate GET and PUT functions from main.go into a handler file
- Setup docker-compose to build/run app and run redis instance
- Setup DockerFile to have multi stages such as `development, builder, production`

## Recommendations for Prod
- Separate environments for dev, stage, prod
- A separate instance/cluster of REDIS 
- CI/CD pipeline
- Security scanning for possible vulnerabilities in code

## How to run

1. Initiate REDIS database

- Clone the repo `git clone git@github.com:timhuynh94/TargetChallenge.git`

- cd into folder, run
 ```shell script 
docker compose up -d
```

- or 

```shell script
docker run --name redis -p 6379:6379 redis
```

2. Startup app
```shell script
go run github.com/timhuynh94/TargetChallenge
```


## How to test
```shell script
go test -v server.go main.go handlers_test.go  -covermode=count  -coverprofile=./bin/coverage.out
```

## Initial Structure
```
    .
    ├── bin
    │   └── main
    ├── models
    │   └── product.go
    ├── rdb_service.go
    ├── .gitignore
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    ├── main_test.go
    ├── main.go
    └── README.md
```

### Rest api
**Object: Product**
```go
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
```

#### Endpoints

**Get product**
```http request
GET http://localhost:8080/products/13860427
Accept: application/json
###
```

**Update a product pricing details**
```http request
PUT http://localhost:8080/products/13860427
Content-Type: application/json

{
	"product": {
		"tcin": "13860427",
		"item": {
			"product_description": {
				"title": "Conan the Barbarian (DVD)",
				"downstream_description": "The most legendary Barbarian of all time is back. Having thrived and evolved for eight consecutive decades in the public imagination -- in prose and graphics, on the big screen and small, in games and properties of all kinds -- Conan's exploits in the Hyborian Age now come alive like never before in a colossal action-adventure film.    A quest that begins as a personal vendetta for the fierce Cimmerian warrior soon turns into an epic battle against hulking rivals, horrific monsters, and impossible odds, as Conan realizes he is the only hope of saving the great nations of Hyboria from an encroaching reign of supernatural evil."
			},
			"enrichment": {
				"images": {
					"primary_image_url": "https://target.scene7.com/is/image/Target/GUEST_45ee9424-04da-4eb7-923d-5150304cf28d"
				}
			},
			"product_classification": {
				"product_type_name": "ELECTRONICS",
				"merchandise_type_name": "Movies"
			},
			"primary_brand": {
				"name": "Lionsgate"
			},
			"current_price": {
				"value": "12",
				"currency_code": "USD"
			}
		}
	}
}
###
```


