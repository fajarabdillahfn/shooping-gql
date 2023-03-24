# shopping gql
This is a GraphQL Service to simulate a checkout products system with some promotions.
The promotions are:
* Each sale of a MacBook Pro comes with a free Raspberry Pi B
* Buy 3 Google Homes for the price of 2
* Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers

## Pre-requisite
### Create .env that contains
* POSTGRES_USER
* POSTGRES_PASSWORD
* POSTGRES_DB
* SQL_HOST
* SQL_PORT
* PORT (optional, default 8080)

## Run the app
go run server.go

# GQL
The GQL example is described below.

## Endpoints
GUI : `GET /`
URL : `POST /query`

## GQL Schema
```
mutation {
  checkout(input: $input) {
    products {
      sku
      name
      price
      quantity
      total_price
    }
    total_price
  }
}
```

## Input
```
input: [
    {
        "sku": <sku number>
    }
]
```

## Example Request
```
mutation {
  checkout(
    input: [
      {sku: "43N23P"},	 
      {sku: "43N23P"},	 
      {sku: "43N23P"},	 
    ]
  ) {
    products {
      sku
      name
      price
      quantity
      total_price
    }
    total_price
  }
}
```