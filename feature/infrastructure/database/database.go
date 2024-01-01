package database

import (
	"go.uber.org/fx"
)

// DatabaseTarget is parameter object for geting all Mongodb's dependency
type DatabaseTarget struct {
	fx.In
	MongoURL       string `name:"mongo_url" optional:"true"`
	GraphiQLEnable bool   `name:"graphiql_enable"`
	MONGO_URL      string `name:"MONGO_URL"`
	DBHost         string `name:"DB_HOST"`
	DBPort         string `name:"DB_PORT"`
	DBUser         string `name:"DB_USER"`
	DBPass         string `name:"DB_PASS"`
	DBName         string `name:"DB_NAME"`
	Lc             fx.Lifecycle
	// Logger         *zap.Logger
	// dbHandler      *DBHandler
}

// Connection is connection provider to access to global mongodb client
// type DBHandler struct {
// 	MongoClient mongo.Client
// 	database    *mongo.Database
// }

// func NewMongoDatabase2(target DatabaseTarget) (*DBHandler, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	dbHost := target.DBHost
// 	dbPort := target.DBPort
// 	dbUser := target.DBUser
// 	dbPass := target.DBPass
// 	dbName := target.DBName

// 	dbHost = config.DefaultIfEmpty(dbHost, string(constant.DBHost))
// 	dbPort = config.DefaultIfEmpty(dbPort, string(constant.DBPort))
// 	dbName = config.DefaultIfEmpty(dbName, string(constant.DB))

// 	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

// 	if dbUser == "" || dbPass == "" {
// 		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
// 	}

// 	client, err := mongo.NewClient(mongodbURI)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer CloseMongoDBConnection(client)

// 	target.dbHandler.MongoClient = client
// 	database := client.Database(dbName)
// 	target.dbHandler.database = &database

// 	return target.dbHandler, nil

// 	// return client, nil
// 	// return &Connection{
// 	// 	client: &client,
// 	// }, err
// }

// func CloseMongoDBConnection(client mongo.Client) {
// 	if client == nil {
// 		return
// 	}

// 	err := client.Disconnect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Connection to MongoDB closed.")
// }
