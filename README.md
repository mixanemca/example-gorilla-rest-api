# example-gorilla-rest-api
Example project with RESTful API on Gorilla mux router

## Help
You can pass the `--help` flag to see all the flags and description.

```
Usage of ./example-gorilla-rest-api:
  -c, --config string            path to config file
  -H, --database.host string     database host (default "localhost")
  -N, --database.name string     database name (default "gapi")
  -P, --database.port int        database port (default 5432)
  -U, --database.user string     database user (default "postgres")
  -a, --http.address string      http listening address (default "127.0.0.1:8080")
  -t, --http.timeout.read int    http read timeout (default 5)
  -w, --http.timeout.write int   http write timeout (default 5)
  -f, --log.format log.format    log format (default json)
  -l, --log.level log.level      log level (default info)
pflag: help requested
```

## Swagger
To generate documentation we will use [swag](https://github.com/swaggo/swag). How to install:
1. Add comments to your API source code. At least fill follow:
```go
// @version 
// @title 
// @description 
```
2. Download swag by using:
`go install github.com/swaggo/swag/cmd/swag@latest`
3. Run the Swag in your Go project root folder which contains main.go file, Swag will parse comments and generate required files(docs folder and docs/doc.go).
` swag init -g cmd/example-gorilla-rest-api/main.go`
4. Download http-swagger by using:
`go get -u github.com/swaggo/http-swagger`
And import following in your code:
`import "github.com/swaggo/http-swagger"`
5. Add import to `docs` directory
`_ "github.com/mixanemca/example-gorilla-rest-api/docs"`
6. Add to your router - depends on the library you are using (see examples in documentation).
