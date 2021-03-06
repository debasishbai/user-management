
# User Management Overview
 
This application provides APIs for simple CRUD operation to manage users.

## API overview

Below is the list of enpoints with their respective input and output

### Get User
#### Endpoint

```
GET
/user_management/user
```
Get all users
#### Output
```json
[
    {
        "id": 1,
        "first_name": "Bruce",
        "last_name": "Wayne",
        "age": 30,
        "occupation": "Entrepreneur"
    },
    {
        "id": 2,
        "first_name": "Robert",
        "last_name": "Downey Jr",
        "age": 30,
        "occupation": "Entrepreneur"
    }
]
```
```
GET
/user_management/user?id=2
```
`id`: primary id of the user
#### Output
```json 
[
    {
        "id": 2,
        "first_name": "Robert",
        "last_name": "Downey Jr",
        "age": 30,
        "occupation": "Entrepreneur"
    }
]
```

### Create new user
### Endpoint
```
POST
/user_management/user
```
### Input
```json
    {
        "first_name": <string type>,
        "last_name": <string type>,
        "age": <int type>,
        "occupation": <string type>
    }
```

### Update user
### Endpoint
```
PATCH
/user_management/user/:id
```
`id`: primary id of the user
### Input
```json
    {
        "id": 2,
        "first_name": "Steve",
        "last_name": "Rogers",
        "age": 30,
        "occupation": "Soldier"
    }
```  

### Delete user
### Endpoint
```
DELETE
/user_management/user/:id
```
`id`: primary id of the user

## Setup
```
Install Go 1.18
Run Commands
`go mod download`
`go build`
`go run main.go`
```
