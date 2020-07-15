module gomicro/server

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/grpc-gateway v1.14.5 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.8.0
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gotest.tools v2.2.0+incompatible
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
