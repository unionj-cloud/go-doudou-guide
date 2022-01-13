module usersvc

go 1.15

require (
	github.com/ascarter/requestid v0.0.0-20170313220838-5b76ab3d4aee
	github.com/brianvoe/gofakeit/v6 v6.9.0
	github.com/go-resty/resty/v2 v2.6.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/handlers v1.5.1
	github.com/iancoleman/strcase v0.1.3
	github.com/jmoiron/sqlx v1.3.4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/slok/goresilience v0.2.0
	github.com/unionj-cloud/cast v1.3.2
	github.com/unionj-cloud/go-doudou v0.9.2
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

//replace github.com/unionj-cloud/go-doudou v0.9.2 => /Users/wubin1989/workspace/cloud/go-doudou
