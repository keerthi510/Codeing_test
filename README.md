# WeatherAPI

Before Running the API run the below command

### Environment Variable setting
```export env=local```

### cd into project directory and run
``` go run cmd/main.go```
### Send  a get request as
``` http://localhost:8080/weather/?city=london&country=uk&day=2```
day is the optional get parameter which gets the forecast for a specific  day

sqlite is used for in memory caching of previous 2 mins data 