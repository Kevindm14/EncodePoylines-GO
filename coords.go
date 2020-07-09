package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

func main() {
	exampleEncodeCoords()
}

func exampleEncodeCoords() {
	coords := make([][]float64, 0)
	archivo, err := os.Open("example.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Location arrived") {
			result := strings.Fields(line)
			coords = append(coords, make([]float64, 0))

			rst := result[9]
			latitud, _ := strconv.ParseFloat(rst[2:13], 64)
			longitud, _ := strconv.ParseFloat(rst[14:26], 64)

			coords[count] = append(coords[count], latitud)
			coords[count] = append(coords[count], longitud)
			count++
		}

	}

	polilinea := polyline.EncodeCoords(coords)

	fmt.Printf("%s\n", polilinea)
}
