# example-gorilla-rest-api
Example project with RESTful API on Gorilla mux router

### Create postgresDB in docker
```bash
docker run --name gorilla-api -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=gapi -p 5432:5432 -d postgres
```

### Use follow commands for migrations
##### For create migration, use this command in 'migration' folder
```bash
 goose create add_updated_at
```
