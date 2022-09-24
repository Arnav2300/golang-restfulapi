# golang-restfulapi

# Overview
A REST API (also known as RESTful API) is an application programming interface (API or web API) that conforms to the constraints of REST architectural style and allows for interaction with RESTful web services. REST stands for representational state transfer and was created by computer scientist Roy Fielding.

REST is a set of architectural constraints, not a protocol or a standard.

When a client request is made via a RESTful API, it transfers a representation of the state of the resource to the requester or endpoint. This information, or representation, is delivered in one of several formats via HTTP: JSON (Javascript Object Notation), HTML, XLT, Python, PHP, or plain text. JSON is the most generally popular file format to use because, despite its name, it’s language-agnostic, as well as readable by both humans and machines. 

This API allows users to create, delete, update and get the data from 'user data' from a mongo database. 

# Navigation
- [How to use the API](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#how-to-use-the-api) 
- [Functions and features](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#functions-and-features)
  - [Create request](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#create-request)
  - [Get request](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#get-request)
  - [Update request](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#update-request)
  - [Delete request](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#delete-request)
  - [Getall request](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#getall-request)
- [Developer Notes](https://github.com/Arnav2300/golang-restfulapi/edit/main/README.md#developer-notes)

# How to use the API
- Prerequisites: MongoDB, mongosh, golang compiler, postman/thunderclient (addon for vscode).
1. After installing the prerequisites, download this repository.
2. Open it in a code editor of your choice.
3. Folder structure should look like this. Ignore the readme and git files if any.

![image](https://user-images.githubusercontent.com/56104491/192092575-d532e9d7-ddbc-4788-aaf5-a09d704d2c4a.png)

4. Open cmd and type in ```mongosh```
5. In order to work correctly, mongoDB should be running on port 27017.
6. Now open ```main.go``` and open terminal.
7. Type the following command.
```
go run "{YOUR_PATH}\main.go"
```
8. Now open thunderclient or postman.
9. Create requests like the following.

![image](https://user-images.githubusercontent.com/56104491/192093037-2f7065f1-bc8b-4f86-87a5-9737b760b5cd.png)

10. The create request should look like the following.

![image](https://user-images.githubusercontent.com/56104491/192093253-052cec10-a89d-47e4-a17f-649f5968caaf.png)

JSON data to put in the body:
```
{
  "id": "1",
  "name": "arnav",
  "dob": {
    "day": 23,
    "month": 4,
    "year": 2000
  },
  "address":{
    "state": "delhi",
    "city": "delhi",
    "pincode":110096
  },
  "description":"student",
  "createdAt": {
    "day": 24,
    "month": 9,
    "year": 2022
  }
}
```
11. After sending the request, the status should return 200, like in the above screenshot.
12. Now navigate to the terminal with mongosh running.
13. Type ```show dbs```. This will return a list of databases.
14. Type ```use userdb```.
15. Type ```db.users.find();```. The following should be displayed.

![image](https://user-images.githubusercontent.com/56104491/192093428-61ea9de0-e49c-4d82-878a-ace5be644c22.png)

# Functions and features
## Create request

![image](https://user-images.githubusercontent.com/56104491/192093253-052cec10-a89d-47e4-a17f-649f5968caaf.png)

Request URI:
```
http://localhost:9090/v1/user/create
```
JSON to put in the body:
```
{
  "id": "1",
  "name": "arnav",
  "dob": {
    "day": 23,
    "month": 4,
    "year": 2000
  },
  "address":{
    "state": "delhi",
    "city": "delhi",
    "pincode":110096
  },
  "description":"student",
  "createdAt": {
    "day": 24,
    "month": 9,
    "year": 2022
  }
}
```
## Get request

![image](https://user-images.githubusercontent.com/56104491/192093725-5a3c98b0-a3d5-49fa-8927-4d5170b2f847.png)

Request URI:
```
http://localhost:9090/v1/user/get/1
```

## Update request

![image](https://user-images.githubusercontent.com/56104491/192093869-468444f0-20bf-4e95-a72b-59c9c598580b.png)

Request URI:
```
http://localhost:9090/v1/user/update
```
JSON data to put in body:
```
{
  "id": "1",
  "name": "abcdef",
  "dob": {
    "day": 23,
    "month": 4,
    "year": 2000
  },
  "address":{
    "state": "delhi",
    "city": "delhi",
    "pincode":110096
  },
  "description":"student",
  "createdAt": {
    "day": 24,
    "month": 9,
    "year": 2022
  }
}
```
Mongo shell should look like the following:

![image](https://user-images.githubusercontent.com/56104491/192093910-b491a0c5-1fa5-4d63-ae60-2382079d3454.png)


## Delete Request

![image](https://user-images.githubusercontent.com/56104491/192093943-f125c82d-2acb-4122-8010-5c347a8fd305.png)

Request UIR:
```
http://localhost:9090/v1/user/delete/1
```
## Getall request

![image](https://user-images.githubusercontent.com/56104491/192094019-86acf839-f5eb-49df-a760-af2ab51e771f.png)


Request URI:
```
http://localhost:9090/v1/user/getall
```
This request returns an array of all the data present in the database.
Example:
```
[
  {
    "id": "1",
    "name": "arnav",
    "dob": {
      "day": 23,
      "month": 4,
      "year": 2000
    },
    "address": {
      "state": "delhi",
      "city": "delhi",
      "pincode": 110096
    },
    "description": "student",
    "createdAt": {
      "day": 24,
      "month": 9,
      "year": 2022
    }
  },
  {
    "id": "2",
    "name": "abcd",
    "dob": {
      "day": 23,
      "month": 4,
      "year": 2000
    },
    "address": {
      "state": "delhi",
      "city": "delhi",
      "pincode": 110096
    },
    "description": "student",
    "createdAt": {
      "day": 24,
      "month": 9,
      "year": 2022
    }
  },
  {
    "id": "3",
    "name": "efgh",
    "dob": {
      "day": 23,
      "month": 4,
      "year": 2000
    },
    "address": {
      "state": "delhi",
      "city": "delhi",
      "pincode": 110096
    },
    "description": "student",
    "createdAt": {
      "day": 24,
      "month": 9,
      "year": 2022
    }
  }
]
```
# Developer Notes
An object oriented approach has been kept in mind while writing the code, following the MVC architecture. 

First databse models were defined in the [user.go](https://github.com/Arnav2300/golang-restfulapi/blob/main/models/user.go) script. 

[user.service.go](https://github.com/Arnav2300/golang-restfulapi/blob/main/services/user.service.go) defines the API contracts. 

[user.service.impl.go](https://github.com/Arnav2300/golang-restfulapi/blob/main/services/user.service.impl.go) implements the contracts. 

[user.controller.go](https://github.com/Arnav2300/golang-restfulapi/blob/main/controllers/user.controller.go) contains reference of user.service, user.service contains reference of mongo collection.
