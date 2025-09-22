package bootstrap

import (
	"os"
	"log"
	"reflect"
	"strconv"
)

type Env struct {
	DBUri		string `env:"DB_URI" default:"mongodb://localhost:27017"`
	Port		int    `env:"PORT" default:"8080"`
	AccessTokenSecret string `env:"ACCESS_TOKEN_SECRET" default:"default_secret"`
	APPEnv		string `env:"APPEnv" default: "development"` 
}

func LoadEnv() *Env {
	env := Env{}
	if err := godotenv.LoadEnv(); err !=nil{
		log.Fatal("Error Loading .env file")
	}

	err := envMapToStruct(&env)
	if err !=nil{
		log.Fatal(err)
	}
	if env.APPEnv ==  "development" {
		log.println("Running in development mode")
	}
	return &env
}

func EnvMapToStruct(envStruct interface{}) error {
	val := reflect.ValueOf(envStruct).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		envVar := fieldType.Tag.Get("env")
		if envVar == "" {
			continue
		}

		defaultValue := fieldType.Tag.Get("default")
		if err := os.Setenv(envVar, defaultValue); err != nil {
			return err
		}
	}
	return nil
}