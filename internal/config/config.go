package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	MongoDB string
	ServerPort string
}

func Load()(Config, error){
	// godotenv.Load() - Reads the .env and sets them into the process environment.
	if err:=godotenv.Load(); err!=nil{
		return Config{},fmt.Errorf("⚠️ FAILED to load .env! - %v",err)
	} 
	
		mongoURI,err:= extractEnv("MONGODB_URI")
		if err != nil {
			return Config{},err
		}
		mongoDB,err:= extractEnv("MONGODB_NAME")
		if err != nil {
			return Config{},err
		}
		port,err:= extractEnv("PORT")
		if err != nil {
			return Config{},err
		}

		return Config{
			MongoURI: mongoURI,
			MongoDB: mongoDB,
			ServerPort: port,
		},nil

}

// Helper f(x)
func extractEnv(key string)(string,error){
	val:=os.Getenv(key)

	if val==""{
		return "",fmt.Errorf("⚠️ Missing required .env!")
	}
	return val,nil
}