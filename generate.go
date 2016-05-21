package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/husl-colors/husl-go"
)

func randRange(r1, r2 int) int {
	r2 = r2 + 1 // make it inclusive
	return rand.Intn(r2-r1) + r1
}

func main() {
	rand.Seed(time.Now().UnixNano())

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
		data = append(data, husl.HuslToHex(float64(H[0]), float64(backS), float64(darkL)+float64(rangeL)*math.Pow(float64(i)/7.0, 1.5)))
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
		data = append(data, husl.HuslToHex(float64(h), float64(s), float64(l)))
	}

	fp, err := os.Create("./styles/colorsgen.less")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	for k := 0; k <= 15; k++ {
		fmt.Fprintf(fp, "@base0%X: %s;\n", k, data[k])
	}
}
