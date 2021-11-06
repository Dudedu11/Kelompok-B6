module main.go

go 1.17

require (
	github.com/go-kit/kit v0.12.0
	github.com/lib/pq v1.10.3
)

require (
	aph-go-service-master/datastruct v0.0.1 // indirect
	aph-go-service-master/logging v0.0.1 // indirect
	aph-go-service-master/service v0.0.1 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
)

require aph-go-service-master/transport v0.0.0

replace aph-go-service-master/transport => ./transport

replace aph-go-service-master/datastruct => ./datastruct

replace aph-go-service-master/logging => ./logging

replace aph-go-service-master/service => ./service
