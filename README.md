# ADS API

- [GitHub](https://github.com/jafarsirojov/ads) - GitHub repository.

## Frameworks
- [UberFX micro framework](https://godoc.org/go.uber.org/fx) - DI framework.
- [gorilla/mux](https://github.com/gorilla/mux) - Gorilla MUX.
- [UberZAP logging](https://godoc.org/go.uber.org/zap) - Blazing fast, structured, leveled logging in Go.


## API


### Add ad
{api_address}:7777/api/shop/v1/ad

Method: POST

Request:

```
{
  "title": "IPHONE 16",
  "description": "prodayu7898056789",
  "price": 39090,
  "linksToPhotos": ["https://www.svyaznoy.ru/catalog/phone/225/apple", "https://www.apple.com/shop/buy-iphone/iphone-12"]
}
```

Responses:

```
{
    "code": 200,
    "message": "Success",
    "payload": {
        "id": 2,
        "title": "IPHONE 16",
        "description": "prodayu7898056789",
        "price": 39090,
        "linksToPhotos": [
            "https://www.svyaznoy.ru/catalog/phone/225/apple",
            "https://www.apple.com/shop/buy-iphone/iphone-12"
        ],
        "createdAt": "2022-09-11T13:34:22.01175+05:00"
    }
}
```

```
{
    "code": 400,
    "message": "BadRequest",
    "payload": null
}
```

```
{
    "code": 500,
    "message": "InternalErr",
    "payload": null
}
```




### Get list
{api_address}:1111/api/shop/v1/ads

Method: GET

#### Headers:

offset: 0/10/20...

sortByDate: desc/asc

sortByPrice: desc/asc



### Responses:

```
{
    "code": 200,
    "message": "Success",
    "payload": [
        {
            "id": 2,
            "title": "IPHONE 16",
            "price": 39090,
            "linkToPhoto": "https://www.svyaznoy.ru/catalog/phone/225/apple",
            "createdAt": "2022-09-11T13:34:22.01175+05:00"
        },
        {
            "id": 1,
            "title": "IPHONE 14",
            "price": 10000,
            "linkToPhoto": "https://www.svyaznoy.ru/catalog/phone/225/apple",
            "createdAt": "2022-09-11T12:48:58.394415+05:00"
        }
    ]
}
```

```
{
    "code": 400,
    "message": "BadRequest",
    "payload": null
}
```

```
{
    "code": 404,
    "message": "NotFound",
    "payload": null
}
```

```
{
    "code": 500,
    "message": "InternalErr",
    "payload": null
}
```




### Get by id
{api_address}:1111/api/shop/v1/ad/{id}

Method: GET

Responses:

```
{
    "code": 200,
    "message": "Success",
    "payload": {
        "id": 2,
        "title": "IPHONE 16",
        "description": "prodayu7898056789",
        "price": 39090,
        "linksToPhotos": [
            "https://www.svyaznoy.ru/catalog/phone/225/apple",
            "https://www.apple.com/shop/buy-iphone/iphone-12"
        ],
        "createdAt": "2022-09-11T13:34:22.01175+05:00"
    }
}
```

```
{
    "code": 400,
    "message": "BadRequest",
    "payload": null
}
```

```
{
    "code": 404,
    "message": "NotFound",
    "payload": null
}
```

```
{
    "code": 500,
    "message": "InternalErr",
    "payload": null
}
```

