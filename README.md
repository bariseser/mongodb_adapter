# MongoDB Adapter for Go

This module help to you some mongo command which add, get and del

### Connect to Redis
````go
mongoClient = ConnectMongo("mongodb://localhost:23176", , "dbname", "CollectionName")
````

###  Find
````go
err := mongoClient.GetData(bson.D{{filters}}, interface{})

if err != nil {
    log.Fatal(err.Error())
}
````

###  All Rows
````go
err := mongoClient.GetAll(bson.D{{filters}}, interface{})

if err != nil {
    log.Fatal(err.Error())
}
````

###  Insert
````go
err := mongoClient.GetAll(bson.D{{filters}}, interface{})

if err != nil {
    log.Fatal(err.Error())
}
````

###  Update
````go
err := mongoClient.GetAll(bson.D{{filters}}, bson.D{{filters}})

if err != nil {
    log.Fatal(err.Error())
}
````

###  Delete
````go
err := mongoClient.DeleteData(bson.D{{filters}})

if err != nil {
    log.Fatal(err.Error())
}
````