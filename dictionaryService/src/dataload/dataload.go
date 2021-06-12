package dataload

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
	_ "github.com/lib/pq"
)

var (
	jlpt1 = "./dataload/jlptData/jlpt_n1.csv"
	jlpt2 = "./dataload/jlptData/jlpt_n2.csv"
	jlpt3 = "./dataload/jlptData/jlpt_n3.csv"
	jlpt4 = "./dataload/jlptData/jlpt_n4.csv"
	jlpt5 = "./dataload/jlptData/jlpt_n5.csv"
)

var CSVs = []string{
	jlpt1,
	jlpt2,
	jlpt3,
	jlpt4,
	jlpt5,
}

func onFailure(DB *sql.DB, warn string, err error) {

	if err != nil {
		log.Fatal(warn+"\n", err)
	}

}

func Dataload(app models.App) {

	fmt.Println("CREATING TABLE")
	_, err := app.DB.Exec(
		"CREATE TABLE IF NOT EXISTS words (id SERIAL PRIMARY KEY,foreign1 VARCHAR(255) NOT NULL, foreign2 VARCHAR(255) NOT NULL,definitions TEXT NOT NULL,level  VARCHAR(255) NOT NULL)")

	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
		fmt.Println("skip the dataload, the table exists")
		return
	}

	for _, csvFile := range CSVs {
		openFile, err := os.Open(csvFile)
		onFailure(app.DB, "Error opening file", err)

		r := csv.NewReader(openFile)
		r.LazyQuotes = true

		fileData, err := r.ReadAll()
		if err != nil {
			fmt.Errorf("error in csvReadAll func: %w", err)
		}

		fmt.Printf("STARTING IMPORT FOR %s\n", csvFile)

		for i, record := range fileData {
			if i == 0 {
				fmt.Println("SKIPPING HEADER")
			} else {
				word := models.VocabWord{
					Foreign1:    record[0],
					Foreign2:    record[1],
					Definitions: record[2],
					Level:       record[3],
				}

				word.DbCreateWord(app.DB)
			}

		}

	}

	fmt.Println("SUCCESSFUL INITIAL DB MIGRATION")
}
