package transforms

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"log"
	"teste-apache-beam/database"
	"teste-apache-beam/entities"
)

func init() {
	beam.RegisterDoFn(getPixels)
}

func ReadDatabase(s beam.Scope) beam.PCollection {
	return beam.ParDo(s, getPixels, beam.Impulse(s))
}

func getPixels(_ []byte, emit func(key string, pixel entities.Pixel)) {
	db := database.DBConnection()
	defer db.Close()

	//pixelMap := make(entities.PixelMap)
	rows, err := db.Query("select source_id, datatype_id, sourcetype_id, data_inst from pixel")
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		var p entities.Pixel
		if err := rows.Scan(&p.SourceID, &p.DataTypeID, &p.SourceTypeID, &p.DataInst); err != nil {
			log.Panic(err)
		}
		emit(p.SourceID, p)
	}
}
