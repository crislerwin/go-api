# My #1  CRUD MVC API with Golang

## To run you need to have installed docker and docker-compose

1. Clone this repository
2. Run `docker-compose up --build`
3. Run `go run main.go`
4. Open `localhost:5000/api/v1/books`

## Request Body

- GET /api/v1/books
- POST /api/v1/books
- PUT /api/v1/books/:id
- DELETE /api/v1/books/:id

**Example**

application/json

```json
{
  "name": "My book",
  "description": "This is my fist book",
  "medium_price": 10,
  "author": "Jay Jones",
  "img_url": "/src/images"
}
```
