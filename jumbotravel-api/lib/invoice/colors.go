package invoice

import "github.com/johnfercher/maroto/pkg/color"

func getBlue() color.Color {
	return color.Color{
		Red:   0,
		Green: 182,
		Blue:  215,
	}
}

func getGray() color.Color {
	return color.Color{
		Red:   209,
		Green: 213,
		Blue:  219,
	}
}

var (
	blue  = getBlue()
	gray  = getGray()
	white = color.NewWhite()
)
