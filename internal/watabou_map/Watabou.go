package watabou_map

const (
	LayoutOddR  = "odd-r"
	LayoutEvenQ = "even-q"
)

type WatabouMap struct {
	Name      string                `json:"name"`
	Origin    string                `json:"origin"`
	Bp        WatabouBp             `json:"bp"`
	Layout    string                `json:"layout"`
	Hexes     map[string]WatabouHex `json:"hexes"`
	Rivers    map[string]River      `json:"rivers"`
	Roads     map[string][]string   `json:"roads"`
	Searoutes map[string][]string   `json:"searoutes"`
	Features  []Feature             `json:"features"`
}

// WatabouBp Parameters(blueprint) for the Watabou map: dimensions, tags, and seed. Actually, I only need the dimensions.
type WatabouBp struct {
	Width  int32    `json:"width"`
	Height int32    `json:"height"`
	Tags   []string `json:"tags"`
	Seed   int64    `json:"seed"`
}

type WatabouHex struct {
	Q       int     `json:"q"`
	R       int     `json:"r"`
	Town    *Town   `json:"town"`
	Terrain *string `json:"terrain"`
}

type Town struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Walled bool   `json:"walled"`
	Info   string `json:"info"`
	Link   string `json:"link"`
	Seed   int    `json:"seed"`
}

type River struct {
	Parent  *River   `json:"parent"`
	Channel []string `json:"channel"`
}

type Feature struct {
	Name  string   `json:"name"`
	Hexes []string `json:"hexes"`
}

func (wm WatabouMap) Validate() error {
	var err error
	err = wm.validateLayout()
	if err != nil {
		return err
	}

	return nil
}

func (wm WatabouMap) validateLayout() error {
	switch wm.Layout {
	case LayoutOddR, LayoutEvenQ:
		return nil
	}
	return &Error{message: "Not supported layout: " + wm.Layout}
}
