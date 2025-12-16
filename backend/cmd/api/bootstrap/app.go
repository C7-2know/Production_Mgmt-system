package bootstrap

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"soap_factory/infrastructure/database"
)

type Application struct {
	Env   *Env
	DB 	  *mongo.Database

}

func App() *Application {
	app := &Application{}
	app.Env = LoadEnv()
	
	cfg := database.Config{
		URI:      app.Env.DBUri,
		Username: app.Env.DBUsername,
		Password: app.Env.DBPassword,
		Database: app.Env.DBName,
	}
	db, err := database.ConnectMongoDB(cfg)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	app.DB = db.Database
	return app
}

func (app *Application) Close() {
	if app.DB != nil && app.DB.Client() != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := app.DB.Client().Disconnect(ctx); err != nil {
			panic("Failed to disconnect from database: " + err.Error())
		}
	}
}