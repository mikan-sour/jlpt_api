# SET SERVICE
the service that holds the user's service

## The service
This service is for getting all the available sets that a user has access to. Each set contains an array of integers, which represent the unique IDs of a word in the dictionary service. 

Once BFF is implemented, the BFF will get all sets from this service, then get the words from the dictionary based on those unique id's, then return to the client

### Routes

#### Healthcheck (/healthcheck): 
returns API status

#### Sets (/sets):

##### Accepts
- Get 
- Post

returns an array of available sets 

##### QueryParams
- isPublic (bool)
- setName (string) (optional)
- isMine (bool)
- limit (int) (should have a default)
- offset (int) (should have a default)

##### Response Object
the IDs of the cards will be requested via a different route

- id (mongodb _id)
- owner (int) < this needs to be processed on the BFF to get the user's actual name! >
- name (string)
- isPublic (bool)

#### Words in Set (/sets/<setID>)
returns the array of ints that represent the words belonging to a set

##### Response Object
Array of ints (this could change if we implement an object for the word object)

####


# OLD
Not using Cassandra because it doesn't make sense for the service (Use the right tools for the job!)

## The database
We're using Cassandra because I thought it might be fun to try

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