## For User Test:

 - cd user 
 - Run `go test -v`

## For Client Test:

 - Uncomment the `restapi.RestUserService()` in the main function and comment the `GRPCUserService()` function in the main.go file
 - Run the REST Server first using `go run main.go`
 - run `go test -v` in another command prompt

## For GRPC Client Test:

 - Uncomment the `GRPCUserService()` in the main function and comment the `restapi.RestUserService()` function in the main.go file
 - Run the REST Server first using `go run main.go`
 - cd grpcClient
 - Run `go run grpcClient.go`