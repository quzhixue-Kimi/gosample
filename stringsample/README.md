go test -v -cover -coverprofile=result.out
go tool cover -html=result.out
go test -v -run=TestSuiteSplit/case3
go test -bench=Split
go test -bench=Split -benchmem
go test -bench=. --benchmem
go test -bench=. --benchmem -benchtime=10s
go test -bench=. --benchmem -benchtime=5s -cpu=1
