# Fetch Receipt Processor

Webservice which processes the receipt and provides an id and calculates the reward points associated with the receipt

## Installation Steps

### Method 1 (Using Docker):

* Install Docker and make sure Docker daemon is running
* Build docker image using the command: `docker-compose build`
* Run the image in a container using the command: `docker-compose up -d`
* Check whether the container is running using the command: `docker ps`

### Method 2 (Using Go in local machine):

* Download Go from https://go.dev/doc/install and install the latest stable version
* In the fetch-receipt-processor directory, download dependencies using the following command: `go mod download`
* To build the webservice, run the command: `go build .`
* To start running the service, run: `./fetch-receipt-processor`

To test the service, run: `go test -v`

## Implementation

* Utilized the Factory and Strategy design patterns to enhance the scalability of the application.
  > <br> Idea behind Strategy pattern: Currently, our points calculator implements several rules. To make the application adaptable to
  > new calculation rules, we can dynamically provide the required rules at runtime, promoting better isolation and flexibility.
  > >
  > Idea behind Factory pattern: The Factory pattern allows us to create a points calculator with a custom strategy. Instead of tightly
  > coupling the creation logic in the main function, we can use the Factory to dynamically provide the appropriate calculator
  > based on the required strategy.
* Used UUID for ID Generation
* Utilized maps to retrieve points associated with an ID
* Followed DRY principle to promote reusability in the code
* Followed Test-Driven-Development for better code coverage and maintaing a clean code base

## Endpoints

Once the application/container is running, we can access the endpoints of the webservice

> [!NOTE]
> Application will run in port 8080 by default

> http://localhost:8080/receipts/process
> >
> http://localhost:8080/receipts/{id}/points

## API Endpoints Documentation

### Endpoint: Process Receipt 

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON
* Response: JSON containing an id for the receipt.


#### Sample Request: 
`http://localhost:8080/receipts/process`

#### Sample Request JSON:

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

#### Sample Response JSON:

```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

### Endpoint: Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`
* Response: A JSON object containing the number of points awarded.


#### Sample Request: 
`http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points`

#### Sample Response JSON:

```json
{ "points": 28 }
```




