# API

POST `localhost:8080/api/v0/signature-device` 
 - Creates a new signature device, receives a json object
```
{
    "id":"testing",
    "algorithm":"ECC",
    "label":""
}
```

GET `localhost:8080/api/v0/signature-device/list` 
 - Lists all signature devices, you can pass: id,label,algorithm as query parameter in order to filter
 - (`?id=&label=&algorithm=`)


GET `localhost:8080/api/v0/signature-device/{id}` 
 - Gets a specific signature device with the id passed on the path parameter


POST `localhost:8080/api/v0/sign-transaction` 
 - Creates a new transaction
```
{
    "device_id":"testing2",
    "data":"65d7"
}
```

GET `localhost:8080/api/v0/sign-transaction/list`
  - Lists all transactions, you can pass device_id as query parameter in order to filter
  - (`?device_id`)


GET `localhost:8080/api/v0/sign-transaction/{id}` 
 - Gets a specific transaction with the id passed on the path parameter
