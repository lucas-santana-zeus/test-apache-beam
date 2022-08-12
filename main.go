package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

func main() {
	flag.Parse()

	beam.Init()

	pipeline, scope := beam.NewPipelineWithRoot()

	if err := beamx.Run(context.Background(), pipeline); err != nil {
		fmt.Println(err.Error())
	}
}
