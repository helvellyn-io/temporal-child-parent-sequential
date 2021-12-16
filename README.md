# temporal-weather
Submits task to Temporal, retrieving the current temperature in Bolder CO 

#Run all tests 
go test ./...

#Create task / worker and sumbit. 

go rum worker/main.go
go run start/main.go 


