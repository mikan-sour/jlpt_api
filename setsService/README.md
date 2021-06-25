# SET SERVICE
the service that holds the user's service

## The service


## The database
We're using Cassandra because I thought the set would best be represented by a json object

### Running the database in Docker
1. Run image
```docker run --name jlpt-mongodb -p 27017:27017 -d mongo```
2. Open in ternminal
```docker exec -it jlpt-mongodb bash```
3. Open cassandra
```mongo```


## OLD

## The database
We're using Cassandra because I thought the set would best be represented by a json object

### Running the database in Docker
1. Run image
```docker run --name jlpt-cassandra -p 9842:9842 -d cassandra```
2. Open in ternminal
```docker exec -it jlpt-cassandra bash```
3. Open cassandra
```cqlsh```