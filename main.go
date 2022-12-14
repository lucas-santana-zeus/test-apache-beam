package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"teste-apache-beam/database"
	"teste-apache-beam/transforms"
)

func main() {
	flag.Parse()

	database.Populate()

	beam.Init()

	pipeline, scope := beam.NewPipelineWithRoot()

	pixelMap := transforms.ReadDatabase(scope)

	temporalPixel := transforms.CalculateTemporal(scope, pixelMap)

	transforms.WriteOnDatabase(scope, temporalPixel)

	if err := beamx.Run(context.Background(), pipeline); err != nil {
		fmt.Println(err.Error())
	}
}
