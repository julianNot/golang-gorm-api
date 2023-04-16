# golang-gorm-api

##  Docker Instructions
Download Postgres image

```
docker pull postgres
```

Run postgres container with environment variables

```
docker run --name [nameContainer] -e POSTGRES_USER=[user] -e POSTGRES_PASSWORD=[password] -e POSTGRES_DB=[nameDB] -p 5432:5432 -d postgres
```
Access the container

```
docker exec -it [nameContainer] bash
```

Access the Postgres in Container

```
psql -U [user] --password --db [nameDB]
```

## Dependencies
* Mux
* GORM - ORM
* Driver Postgres - GORM
* AIR
* Godotenv