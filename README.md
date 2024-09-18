# National Weather Service api
a simple passthrough api that connects to the national weather service

# running the program
simply use the commands `go get` followed by `go run main.go` or `make run` (which just runs both for you)

# endpoints
as of now there are only 2 exposed endpoints:

`http://localhost:9001/swagger/`
`http://localhost:9001/forcast/get/{lat}/{long}`

## swagger
this provides documentation for all other endpoints available on the API as well as a simple way to test them

## /forcast/get/{lat}/{long}
this endpoint takes a latitude and a longitude then returns a forcast in that area. Any errors that occur in the National Weather Service API will be bubbled up and returned.