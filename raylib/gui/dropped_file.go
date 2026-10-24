package gui

import (
	"encoding/json"
	"os"

	"github.com/Zenith-Kim-Light/WorldHexMap/internal/watabou_map"
)

func ProcessDroppedFile(path string) (*watabou_map.WatabouMap, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var WatabouMap watabou_map.WatabouMap
	err = json.Unmarshal(data, &WatabouMap)

	if err != nil {
		return nil, err
	}

	err = WatabouMap.Validate()
	if err != nil {
		return nil, err
	}

	return &WatabouMap, nil
}
