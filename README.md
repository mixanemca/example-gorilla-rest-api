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
