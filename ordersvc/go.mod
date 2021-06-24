module ordersvc

go 1.15

require (
	github.com/ascarter/requestid v0.0.0-20170313220838-5b76ab3d4aee
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2 // indirect
	github.com/go-resty/resty/v2 v2.6.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/iancoleman/strcase v0.1.3
	github.com/jmoiron/sqlx v1.3.4
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/unionj-cloud/go-doudou v0.2.9
	usersvc v0.0.0
)

replace github.com/unionj-cloud/go-doudou v0.2.9 => /Users/wubin1989/workspace/cloud/go-doudou

replace usersvc v0.0.0 => /Users/wubin1989/workspace/cloud/go-doudou-guide/usersvc
