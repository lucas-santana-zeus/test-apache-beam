package main

import (
	"flag"
	"teste-apache-beam/database"
)

func main() {
	flag.Parse()

	database.Populate()

	//beam.Init()
	//
	//pipeline, scope := beam.NewPipelineWithRoot()
	//
	//if err := beamx.Run(context.Background(), pipeline); err != nil {
	//	fmt.Println(err.Error())
	//}
}
