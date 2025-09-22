package bootstrap

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *Env
	mongo *mongo.Client

}

func App() *Application {
	app := &Application{}
	app.Env = LoadEnv()
	app.mongo = NewMongoDatabase(app.Env)
	return app
}
