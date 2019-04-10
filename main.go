package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"log"
	"os"

	"github.com/imroc/req"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "novastore"
	app.Usage = "An easier way to get your stats and predictions"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "file",
			Value: "No File Given",
			Usage: "This speciciles a CSV file to our system to be either predicted on or analyzed",
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
			// This uses the following pattern novastoreCLI upload -file path-to-file
			Action: func(c *cli.Context) error {
				// a simple lookup function
				file, _ := os.Open(c.String("file"))

				_, filename := filepath.Split(c.String("file"))

				fmt.Println(filename)
				r, err := req.Post("http://dca-novanstoreapi.herokuapp.com/upload", req.FileUpload{
					File:      file,
					FieldName: "file",   // FieldName is form field name
					FileName:  filename, //Filename is the name of the file that you wish to upload. We use this to guess the mimetype as well as pass it onto the server
				})
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(r)
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
				r, err := req.Post("http://dca-novanstoreapi.herokuapp.com/login", header, req.BodyJSON(param))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(r)
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
				r, err := req.Post("http://dca-novanstoreapi.herokuapp.com/register", header, req.BodyJSON(param))
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
		{
			Name:  "stats",
			Usage: "returns your statistics on a given file in json format",
			Flags: myFlags,
			// This will get a user registerd into our system
			// we execute our `signup` command
			Action: func(c *cli.Context) error {

				bodyBuf := &bytes.Buffer{}
				bodyWriter := multipart.NewWriter(bodyBuf)

				// For the post request, it set up the form name attached to the file for posting. eg file=path-to-fie
				fileWriter, err := bodyWriter.CreateFormFile("file", c.String("file"))
				if err != nil {
					fmt.Println("error writing to buffer")
					return err
				}

				// Opens the file with that path from the -file flag
				fh, err := os.Open(c.String("file"))
				if err != nil {
					fmt.Println("error opening file")
					return err
				}
				defer fh.Close()

				// Allows copying from source to destination
				_, err = io.Copy(fileWriter, fh)
				if err != nil {
					return err
				}

				contentType := bodyWriter.FormDataContentType()
				bodyWriter.Close() // Closing bodyWriter

				response, err := http.Post("https://row2json.herokuapp.com/api", contentType, bodyBuf)
				if err != nil {
					log.Fatal(err)
				}
				// Defer closing untill output below is done
				defer response.Body.Close()
				body, err := ioutil.ReadAll(response.Body)
				// fmt.Println(body)
				responseJson := string(body)
				fmt.Println(responseJson)
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Welcoe to the Novastore CLI")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
