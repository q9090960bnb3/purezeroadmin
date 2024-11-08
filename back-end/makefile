hello:
	goctl api go --api hello-api/api/hello.api --dir hello-api

user:
	goctl api  go --api user-api/api/user.api --dir user-api --home template

dsn="root:123456@tcp(127.0.0.1:55506)/purezeroadmin"
tables="tb_*"

mysql: 
	goctl model mysql datasource --url ${dsn}  -t=${tables} --dir user-api/models --home template

run:
	cd user-api && go run user.go