package transforms

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"reflect"
	"teste-apache-beam/entities"
)

func init() {
	beam.RegisterFunction(calculate)
	beam.RegisterType(reflect.TypeOf(entities.Pixel{}))
}

func CalculateTemporal(s beam.Scope, pixelList beam.PCollection) beam.PCollection {
	return beam.ParDo(s, calculate, pixelList)
}

func calculate(pixels []entities.Pixel, emit func(pixel entities.Pixel)) {
	var dataOne, dataTwo, dataThree string
	for _, pixel := range pixels {
		switch pixel.DataTypeID {
		case 1:
			dataOne += "1-" + pixel.DataInst
		case 2:
			dataTwo += "2-" + pixel.DataInst
		case 3:
			dataThree += "3-" + pixel.DataInst
		}
	}
	sourceId := pixels[0].SourceID
	dataList := []string{dataOne, dataTwo, dataThree}
	for i, v := range dataList {
		pixel := entities.Pixel{
			SourceID:     sourceId,
			DataTypeID:   i,
			SourceTypeID: 20,
			DataInst:     v,
		}
		emit(pixel)
	}
}
