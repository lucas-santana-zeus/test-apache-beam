package database

import (
	"fmt"
	"log"
	"strconv"
	"teste-apache-beam/entities"
)

func Populate() {
	db := DBConnection()
	defer db.Close()

	var rowCount int
	if err := db.QueryRow("select count(*) from pixel").Scan(&rowCount); err != nil {
		log.Fatal(err.Error())
		return
	}
	if rowCount > 0 {
		fmt.Println("Database already populated!")
		return
	}

	stmt, err := db.Prepare("insert into public.pixel values ( $1, $2, $3, $4 )")
	if err != nil {
		log.Fatal("prepare:", err)
		return
	}
	for i := 1; i <= 100; i++ {
		sourceId := fmt.Sprint("A" + strconv.Itoa(i))
		p := entities.Pixel{
			SourceID:     sourceId,
			SourceTypeID: 6,
			DataInst:     "10.0000",
		}
		for j := 1; j <= 3; j++ {
			p.DataTypeID = j
			_, err = stmt.Exec(p.SourceID, int64(p.DataTypeID), int64(p.SourceTypeID), p.DataInst)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}
