module github.com/TakaWakiyama/forcusing-backend/cmd/user

go 1.17

replace github.com/TakaWakiyama/forcusing-backend/cmd/user/pb => ./pb

require (
	github.com/TakaWakiyama/forcusing-backend/cmd/user/pb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/go-gorp/gorp v2.2.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
