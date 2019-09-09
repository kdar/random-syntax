package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
	"github.com/hsluv/hsluv-go"
)

func randRange(r1, r2 int) int {
	r2 = r2 + 1 // make it inclusive
	return rand.Intn(r2-r1) + r1
}

func hsluvPalette() []string {
	var data []string

	// 6 hues to pick from
	h := randRange(0, 360)
	H := []int{0, 60, 120, 180, 240, 300}
	for i := range H {
		H[i] = (h + H[i]) % 360
	}

	// 8 shades of low-saturated color
	backS := randRange(4, 50)
	darkL := randRange(0, 10)
	rangeL := 90 - darkL
	for i := 0; i <= 7; i++ {
		data = append(data, hsluv.HsluvToHex(float64(H[0]), float64(backS), float64(darkL)+float64(rangeL)*math.Pow(float64(i)/7.0, 1.5)))
	}

	// 8 Random shades
	minS := randRange(30, 70)
	maxS := minS + 30
	minL := randRange(50, 70)
	maxL := minL + 20
	for j := 0; j <= 7; j++ {
		h := H[randRange(0, 5)]
		s := randRange(minS, maxS)
		l := randRange(minL, maxL)
		data = append(data, hsluv.HsluvToHex(float64(h), float64(s), float64(l)))
	}

	return data
}

// func testPalette() []string {
// 	palette := hsluvPalette()[:8]
// 	// var palette []string
// 	// base, err := colorful.SoftPaletteEx(8, colorful.SoftPaletteSettings{
// 	// 	CheckColor: func(l, a, b float64) bool {
// 	// 		return l < 0.1
// 	// 	},
// 	// 	Iterations:  100,
// 	// 	ManySamples: true,
// 	// })
// 	//
// 	// for _, p := range base {
// 	// 	palette = append(palette, p.Hex())
// 	// }
//
// 	pal, err := colorful.HappyPalette(8)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// pal2 := colorful.FastWarmPalette(10)
// 	// pal3, err3 := colorful.HappyPalette(10)
// 	// pal4 := colorful.FastHappyPalette(10)
// 	// pal5, err5 := colorful.SoftPalette(10)
//
// 	for _, p := range pal {
// 		palette = append(palette, p.Hex())
// 	}
//
// 	return palette
// }

func main() {
	rand.Seed(time.Now().UnixNano())

	data := hsluvPalette()
	buf := &bytes.Buffer{}
	for k := 0; k <= 15; k++ {
		fmt.Fprintf(buf, "@base0%X: %s;\n", k, data[k])
	}

	fp, err := os.Create("./styles/colorsgen.less")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	buf.WriteTo(fp)
}
