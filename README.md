# WOOKER API
Wooker API is the core backend for open source word book app.

## REST RESOURCES
### USER

|OPERATION|METHOD|PATH|
|-------|------|----|
|Get All|GET|/v1/users/|
|Get Single|GET|/v1/users/:username/|
|Update|PUT|/v1/users/:username|

**GET ALL RESPONSE *Example***
```json
{
  "code": 200,
  "error": "",
  "users": [
    {
      "id": 2,
      "username": "mehmet",
      "is_admin": false,
      "email": "ahmetberke9009@gmail.com",
      "email_verified": true,
      "picture": "https://lh3.googleusercontent.com/a-/AOh14Ghm1L8Gr0OZ54yl6QPEWUSQUCfMTA4D_KOQo_xt1Q=s96-c"
    },
    {
      "id": 1,
      "username": "mandalina",
      "is_admin": false,
      "email": "tuncel.ahmet.berke@gmail.com",
      "email_verified": true,
      "picture": "https://lh3.googleusercontent.com/a-/AOh14Gir6NmnNJfmxRgKiIfrCWqyQvRxcH9kcVKJEOZe=s96-c"
    }
  ]
}
```
**GET SINGLE *Example***
```json
{
    "code": 200,
    "error": "",
    "user": {
        "id": 1,
        "username": "mandalina",
        "is_admin": false,
        "email": "tuncel.ahmet.berke@gmail.com",
        "email_verified": true,
        "picture": "https://lh3.googleusercontent.com/a-/AOh14Gir6NmnNJfmxRgKiIfrCWqyQvRxcH9kcVKJEOZe=s96-c"
    },
    "logged_in": true
}
```
**UPDATE *Example***
```json
{
    "code": 200,
    "error": "",
    "user": {
        "id": 0,
        "username": "ahmet",
        "is_admin": false,
        "email": "",
        "email_verified": false,
        "picture": ""
    },
    "logged_in": true
}
```