# Purchase API

## Requirements

- Golang >= v1.19.3

### Endpoints

- Save Purchase:
  - HTTP VERB: `POST` - URL: `localhost:8080/v1/purchase`
  - BODY: `JSON`
  ```JSON
    {
      "description": "Example",
      "amount": 2.99,
      "transaction_date": "2001-09-25"
    }
  ```
  - RESPONSE: `SAVED PURCHASE`
  ```JSON
    {
      "id": 1,
      "description": "Example",
      "amount": 2.99,
      "transaction_date": "2001-09-25T00:00:00Z"
    }
  ```
- Get All Purchases:

  - HTTP VERB: `GET` - URL: `localhost:8080/v1/purchase`
  - RESPONSE: `LIST OF PURCHASES`

  ```JSON
    [
      {
        "id": 1,
        "description": "Example",
        "amount": 2.99,
        "transaction_date": "2001-09-25T00:00:00Z"
      }
    ]
  ```

- Convert Purchase amount:
  - HTTP VERB: `GET` - URL: `localhost:8080/v1/purchase/{PURCHASE_ID}/convert/{CURRENCY}`
  - RESPONSE: `CONVERTED PURCHASE`
  ```JSON
    {
      "converted_amount": 6.1,
      "description": "Example",
      "exchange_rate": 2.043,
      "id": 1,
      "purchase_amount": 2.99,
      "transaction_date": "2001-09-25T00:00:00Z"
    }
  ```
