package bootstrap

import (
	"os"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	DBUri		string `env:"DB_URI" default:"mongodb://localhost:27017"`
	DBName		string `env:"DB_NAME" default:"production_mgmt"`
	DBUsername	string `env:"DB_USERNAME" default:""`
	DBPassword	string `env:"DB_PASSWORD" default:""`
	Port		int    `env:"PORT" default:"8080"`
	AccessTokenSecret string `env:"ACCESS_TOKEN_SECRET" default:"default_secret"`
	APPEnv		string `env:"APPEnv" default:"development"` 
}

func LoadEnv() *Env {
	env := Env{}
	if err := godotenv.Load(); err !=nil{
		log.Fatal("Error Loading .env file")
	}

	err := EnvMapToStruct(&env)
	if err !=nil{
		log.Fatal(err)
	}
	if env.APPEnv ==  "development" {
		log.Println("Running in development mode")
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

		envValue := os.Getenv(envVar)
		
		if envValue != "" {
			defaultValue := fieldType.Tag.Get("default")
			envValue = defaultValue
		}
		
		switch field.Kind() {
		case reflect.String:
			field.SetString(envValue)
		case reflect.Int:
			intValue, err := strconv.Atoi(envValue)
			if err != nil {
				intValue = 0
			}
			field.SetInt(int64(intValue))
		case reflect.Bool:
			boolval := strings.ToLower(envValue) == "true"
			field.SetBool(boolval)
		case reflect.Float64:
			floatValue, err := strconv.ParseFloat(envValue, 64)
			if err != nil {
				floatValue = 0.0
			}
			field.SetFloat(floatValue)
		default:
			continue
		}
	}	
	return nil
}