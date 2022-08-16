package transforms

import "github.com/apache/beam/sdks/v2/go/pkg/beam"

func GroupPixel(s beam.Scope, kvPixel beam.PCollection) beam.PCollection {
	return beam.GroupByKey(s, kvPixel)
}
