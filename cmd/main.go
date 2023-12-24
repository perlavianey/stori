package main

import (
	"encoding/csv"
	"log"
	"os"
	"stori/database"

	"github.com/oklog/ulid/v2"
)

func main() {
	var transactionList []database.Transaction

	// open file
	directory := "/app/cmd/directory"
	catFiles := ListarArchivo(directory)
	for _, archivo := range catFiles {
		fileIdentifier := ulid.Make()
		f, err := os.Open(directory + "/" + archivo.Name())
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		//convert file to B64 representation for attachment
		fileByte, _ := os.ReadFile(directory + "/" + archivo.Name())
		//fileB64 := base64.StdEncoding.EncodeToString(fileByte)

		// read csv values using csv.Reader
		csvReader := csv.NewReader(f)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		// convert records to array of structs
		transactionList = convertTransactions(data, fileIdentifier.String())

		//save the array data to the database
		for _, transaction := range transactionList {
			err := database.InsertRequest(pgclient, transaction)
			if err != nil {
				log.Print("Error saving records to database:", err)
				return
			}
		}

		//get customer's name and email
		name, email, err := database.GetAccountData(pgclient, transactionList[0].IdAccount)
		if err != nil {
			log.Print("Error getting customer's name and email: ", err)
			return
		}

		//get Summary
		summary, err := getSummary(transactionList)
		if err != nil {
			log.Fatal(err)
		}

		//send email
		err = sendEmail(name, email, summary, fileByte)
		if err != nil {
			log.Fatal(err)
		}
	}
}
