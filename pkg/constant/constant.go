package constant

type key string

const (
	Environment key = "local"
	ServerHost  key = "127.0.0.1"
	ServerPort  key = "8778"
	DBName      key = "GoServer"
	DB_USER_COL key = "UsersInfo"
	DBHost      key = "127.0.0.1"
	DBPort      key = "27017"
)
