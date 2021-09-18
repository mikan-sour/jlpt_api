# Dictionary Service

## The DB
Run program locally and DB in Docker
```docker run --name jlpt-postgresql -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres```

Run everything together
```
docker-compose build
docker-compose up -d
```

## The Dataload
I am not doing DL with a bash script because the data in the csv needs to be parsed a specific way so I do it as part of the app startup script in GO (dataload package)