package entities

type Pixel struct {
	SourceID     string
	DataTypeID   int
	SourceTypeID int
	DataInst     string
}

type PixelMap map[string][]Pixel
