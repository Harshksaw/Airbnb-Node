package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load(){
	err := godotenv.Load()
	if err!= nil {
		fmt.Println("Error Loading .env file")
	}

}

func GetString(key string, fallback string)string{


	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value

}

func GetInt(key string , fallback int)int {


	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intValue, err  :=  strconv.Atoi(value)
	if(err != nil){
		return 2
	}

		return intValue

}


// func getKey(key string , fallback any)any{
// 	value , ok  := os.LookupEnv(key)

// 	if !ok {
// 		return fallback
// 	}
// 	return value
// }