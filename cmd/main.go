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

	// get files stores in mounted directory /app/cmd/directory
	directory := "/app/cmd/directory"
	catFiles, err := listFiles(directory)
	if err != nil {
		log.Fatal("Error listing files in " + directory + ":" + err.Error())
	}

	//iterate over files
	for _, archivo := range catFiles {
		fileIdentifier := ulid.Make()
		fileName := directory + "/" + archivo.Name()
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal("Error opening file "+fileName+":", err)
		}

		defer f.Close()

		//read csv file for attachment
		fileByte, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal("Error reading file "+fileName+":", err)
		}

		// read csv values using csv.Reader
		csvReader := csv.NewReader(f)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal("Error getting data from csv:", err)
		}

		//upload file to S3
		err = uploadFileToS3(fileIdentifier.String(), fileByte)
		if err != nil {
			log.Print("Error to upload file", fileIdentifier, "to S3:", err)
		}

		// convert records to array of structs
		transactionList = convertTransactions(data, fileIdentifier.String())

		//save the data to the database
		counterError := 0
		for _, transaction := range transactionList {
			err := database.InsertRequest(pgclient, transaction)
			if err != nil {
				log.Print("Error saving record of ", fileIdentifier, " to database:", err)
				counterError++
			}
		}

		//if any error occurs, print the error with the file identifier so the system administrator can make the corresponding troubleshooting
		if counterError > 0 {
			log.Print("There were errors saving data. Please check process of file ", fileIdentifier)
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
			log.Fatal("Error getting summary for customer: ", err)
		}

		//send email
		err = sendEmail(name, email, summary, fileByte)
		if err != nil {
			log.Fatal("Error to send email: ", err)
		}
	}
}
