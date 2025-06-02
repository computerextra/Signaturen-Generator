package main

import (
	"Signaturen-Generator/db"
	"Signaturen-Generator/template"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
)

func main() {
	// TODO: Create Template for Signature
	// TODO: More Infos From .env!
	// TODO: Get "DW" from Employees

	// Delete all old Files
	_ = os.RemoveAll("Signatures")

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	database, err := db.NewDatabase(connString)
	if err != nil {
		panic(err)
	}
	employees, err := database.GetEmployees()
	if err != nil {
		panic(err)
	}
	err = os.Mkdir("Signatures", 0755)
	if err != nil {
		panic(err)
	}
	for _, employee := range employees {
		err := os.Mkdir(fmt.Sprintf("Signatures/%s", employee.Name), 0755)
		if err != nil {
			panic(err)
		}
		folder := path.Join("Signatures", employee.Name)
		file, err := os.Create(fmt.Sprintf("%s/%s (%s).html", folder, employee.Name, employee.Mail))
		if err != nil {
			panic(err)
		}

		err = template.GenerateTextFile(employee, folder)
		if err != nil {
			panic(err)
		}
		err = template.Html(employee).Render(context.Background(), file)
		if err != nil {
			panic(err)
		}
	}

}
