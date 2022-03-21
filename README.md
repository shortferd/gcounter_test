# gcounter_test

simple golang test REST api application

Before the running of the application create the json directory in project folder:
```
cd $path_to_project/gcounter_test
mkdir json 
```

## Running the application
To run the application go to project directrory and execute:
```
cd $path_to_project/gcounter_test
go run cmd/main.go
```

## Testing the application
To provide simple testing execute the following in linux command line:
```
curl -X GET "localhost:8080"
```
