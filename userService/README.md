# User Service
stores users and handles authentication

## The DB
Run program locally and DB in Docker
```docker run --name jlpt-postgresql-users -p 5433:5432 -e POSTGRES_PASSWORD=password -d postgres```

## Auth Pattern

### Sign up
1. On successful create, return userId, username, isAdmin
2. Create session in table returning session ID (user ID is arg)
3. 

### Sign in
1. Login w/ Username and Password, return userId, username, isAdmin
2. On successful login (i.e. records match), create new record in "sessions" table, returning session ID. 
   1. Only required arg is user ID, which we got from successful login
3. If session ID returned successfully, create JWT w/ session ID as a claim
4. Return as HTTP-Only Secure cookie to the client
5. Create access token (short-lived) w/ userId and session ID to client