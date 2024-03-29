module ordersvc

go 1.15

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/iancoleman/strcase v0.1.3
	github.com/jmoiron/sqlx v1.3.4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/unionj-cloud/go-doudou v1.0.2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	usersvc v0.0.0
)

replace usersvc v0.0.0 => ../usersvc

replace github.com/unionj-cloud/go-doudou v1.0.2 => /Users/wubin1989/workspace/cloud/go-doudou

//replace github.com/unionj-cloud/go-doudou v1.0.2 => D:\project\cloud\go-doudou
