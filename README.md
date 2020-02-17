SmartMEI dev test
=======

This is a dev test for Backend Software Engineer test at SmartMEI.

This project was made using go1.13 as programming language.

Pre-requisites
------

For run this project, you must have `docker` and `docker-compose` installed.

Run
------

* Clone this repository.
* Execute `docker-composer up`.
* Open `http://localhost:18000/` in your browser.

Unit Tests
------  
* Clone this repository.
* Execute `docker-composer up` (if not executing yet).
* Execute `docker exec -it smartmei-dev-test-2k20_smartmei-app_1 go test`.

Stop
------
* Execute `docker-compose down --rmi all`

Routes
------

##### GET /

Open GraphiQL client

##### POST /query

Query into app

Request example
```bash
curl --request POST \
  --url http://127.0.0.1:18000/query \
  --header 'content-type: application/json' \
  --data '{"query":"query ($url: String!) {\n  getProfessionalTaxTransfer(Url: $url) {\n    consultDate\n    description\n    amount {\n      usd {\n        integer\n        decimal\n        __typename\n      }\n      brl {\n        integer\n        decimal\n        __typename\n      }\n      eur {\n        integer\n        decimal\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n","variables":{"url":"https://www.smartmei.com.br/#planos-e-tarifas"}}'
```

Success response example
```json
{
  "data": {
    "getProfessionalTaxTransfer": {
      "consultDate": "2020-02-17 04:50:48",
      "description": "Transferência",
      "amount": {
        "usd": {
          "integer": 162,
          "decimal": 1.62,
          "__typename": "Amount"
        },
        "brl": {
          "integer": 700,
          "decimal": 7,
          "__typename": "Amount"
        },
        "eur": {
          "integer": 149,
          "decimal": 1.49,
          "__typename": "Amount"
        },
        "__typename": "ProfessionalTaxTransferAmount"
      },
      "__typename": "ProfessionalTaxTransfer"
    },
    "__typename": "Query"
  }
}
```

Reference
------
* https://graphql.org/learn/schema/
* https://levelup.gitconnected.com/setup-simple-go-development-environment-with-docker-b8b9c0d4e0a8
* https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/
* https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
* https://golang.org/pkg/fmt/
* https://github.com/graph-gophers/graphql-go
* https://github.com/deltaskelta/graphql-go-pets-example
* https://github.com/ZAYDEK/graphql-go-walkthrough
* https://www.callicoder.com/docker-golang-image-container-example/