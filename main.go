
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
)


func main(){
	authErr := godotenv.Load()
	if authErr != nil {
		log.Fatal("Error loading .env file")
	}
	accessKey := os.Getenv("ACCESS")
	secretKey := os.Getenv("SECRET")
	format := "\nAccess: %s\nSecret: %s\n"

	_, authErr = fmt.Printf(format, accessKey, secretKey)
	if authErr != nil {
		log.Fatal(authErr.Error())
	}

	app := cli.NewApp()
	app.Name = "novastore"
	app.Usage = "An easier way to store your files on AWS"

	app.Action = func(c *cli.Context) error {
		fmt.Println("NovaStore CLI up and running")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}