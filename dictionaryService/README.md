# Dictionary Service

## The DB
Run program locally and DB in Docker
```docker run --name jlpt-postgresql -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres```

Run everything together
```
docker-compose build
docker-compose up -d
```