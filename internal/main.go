package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Zenith-Kim-Light/WorldHexMap/internal/watabou_map"
)

func main() {

	path := os.Args[1]
	data, err := os.ReadFile(path)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		os.Exit(1)
	}

	var WatabouMap watabou_map.WatabouMap
	err = json.Unmarshal(data, &WatabouMap)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		os.Exit(1)
	}

	hexes := extractNorthCoast(WatabouMap.Hexes)
	for name, _ := range hexes {
		fmt.Println(name)
	}
}

func extractNorthCoast(hexes map[string]watabou_map.WatabouHex) map[string]watabou_map.WatabouHex {
	hexesOfNorthCoast := make(map[string]watabou_map.WatabouHex)
	var stringBuilder strings.Builder
	stringBuilder.Grow(6)
	for currentHexName, currentHex := range hexes {

		/*
			currentHex.Terrain == nil ||
			*currentHex.Terrain == "mountain" ||
						*currentHex.Terrain == "swamp" ||
						*currentHex.Terrain == "forest-dark" ||
						*currentHex.Terrain == "forest-light" ||
						*currentHex.Terrain == "forest-dead" ||
						*currentHex.Terrain == "desert"
		*/
		if currentHex.Terrain != nil && *currentHex.Terrain == "water" {
			continue
		}

		stringBuilder.WriteRune('q')
		stringBuilder.WriteString(strconv.Itoa(currentHex.Q))
		stringBuilder.WriteRune('_')
		stringBuilder.WriteRune('r')
		stringBuilder.WriteString(strconv.Itoa(currentHex.R - 1))

		nameOfNorthNeighbourHex := stringBuilder.String()
		stringBuilder.Reset()
		northNeighbourHex, ok := hexes[nameOfNorthNeighbourHex]
		if ok {
			if northNeighbourHex.Terrain != nil && *northNeighbourHex.Terrain == "water" {
				hexesOfNorthCoast[currentHexName] = currentHex
			}
		}
	}

	return hexesOfNorthCoast
}
