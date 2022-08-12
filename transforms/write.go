package transforms

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"log"
	"teste-apache-beam/database"
	"teste-apache-beam/entities"
)

func init() {
	beam.RegisterFunction(writeFnPostgres)
}

func WriteOnDatabase(s beam.Scope, pixel beam.PCollection) {
	beam.ParDo0(s, writeFnPostgres, pixel)
}

func writeFnPostgres(pixel entities.Pixel) {
	db := database.DBConnection()
	defer db.Close()
	stmt, err := db.Prepare("insert into pixel_result values ($1, $2, $3, $4)")
	if err != nil {
		log.Panicln("prepare stmt:", err)
	}
	_, err = stmt.Exec(pixel.SourceID, pixel.DataTypeID, pixel.SourceTypeID, pixel.DataInst)
	if err != nil {
		log.Panicln("exec stmt:", err)
	}
}
