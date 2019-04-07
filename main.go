package main

import (
	"fmt"

	"github.com/imroc/req"
	"github.com/urfave/cli"

	// "github.com/imroc/req"
	"log"
	"os"
)

func main() {
	//authErr := godotenv.Load()
	//if authErr != nil {
	//	log.Fatal("Error loading .env file")
	//}
	//accessKey := os.Getenv("ACCESS")
	//secretKey := os.Getenv("SECRET")
	//format := "\nAccess: %s\nSecret: %s\n"
	//
	//_, authErr = fmt.Printf(format, accessKey, secretKey)
	//if authErr != nil {
	//	log.Fatal(authErr.Error())
	//}

	app := cli.NewApp()
	app.Name = "novastore"
	app.Usage = "An easier way to store your files on AWS"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "file",
			Value: "No File Given",
			Usage: "This uploads a target CSV file to our system to be either predicted on or analyzed",
		},
		cli.StringFlag{
			Name:  "email",
			Value: "No email given",
			Usage: "This sets the email for the account to be accessed",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "No password given",
			Usage: "This sets the password for the account to be accessed",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "upload",
			Usage: "Upload a given csv file to be  processed on our server.",
			Flags: myFlags,
			// the action, or code that will be executed when
			// we execute our `upload` command
			Action: func(c *cli.Context) error {
				// a simple lookup function
				fmt.Println(c.String("file"))
				//if err != nil {
				//	return err
				//}

				return nil
			},
		},
		{
			Name:  "login",
			Usage: "Logs a given user into our server to access their resources",
			Flags: myFlags,
			// the action, or code that will be executed when
			// we execute our `upload` command
			Action: func(c *cli.Context) error {
				// a simple lookup function
				fmt.Println(c.String("username"))
				fmt.Println(c.String("password"))
				header := req.Header{
					"Content-Type": "application/json",
				}
				param := req.Param{
					"email":    c.String("email"),
					"password": c.String("password"),
				}
				r, err := req.Post("http://localhost:3000/login", header, req.BodyJSON(param))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(r)
				return nil

				return nil
			},
		},
		{
			Name:  "register",
			Usage: "registers a system into our system to analyze their stats or ",
			Flags: myFlags,
			// This will get a user registerd into our system
			// we execute our `signup` command
			Action: func(c *cli.Context) error {
				// a simple lookup function
				fmt.Println(c.String("email"))
				fmt.Println(c.String("password"))
				header := req.Header{
					"Content-Type": "application/json",
				}
				param := req.Param{
					"email":    c.String("email"),
					"password": c.String("password"),
				}
				r, err := req.Post("http://localhost:3000/register", header, req.BodyJSON(param))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(r)
				return nil
			},
		},
		{
			Name:  "list_file",
			Usage: "This lists the file that the user can use to do stats, predictions",
			Flags: myFlags,
			// This will get a user registerd into our system
			// we execute our `signup` command
			Action: func(c *cli.Context) error {
				// a simple lookup function
				fmt.Println(c.String("file"))
				//if err != nil {
				//	return err
				//}

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("NovaStore CLI up and running")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
