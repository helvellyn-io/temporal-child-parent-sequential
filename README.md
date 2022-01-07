# temporal-parent-child-activity
submits task to Temporal, retrieving the current temperature in Bolder CO 

#run all tests 
go test ./...

#create task / worker and sumbit. 

go run worker/main.go
go run start/main.go 


