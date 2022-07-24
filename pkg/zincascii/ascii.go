package zincascii

import (
	"fmt"
	"zincsearchstash/internal/setup"

	"github.com/common-nighthawk/go-figure"
)

func Apply() {
	zincBanner := figure.NewColorFigure("Zinc Search Stath", "chunky", "blue", true)
	zincBanner.Print()
	fmt.Printf("version %s\n\n", setup.Get().Version)
}
