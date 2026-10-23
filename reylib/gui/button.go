package gui

type ButtonAction uint8

const (
	ActionField ButtonAction = iota
	ActionForest
	ActionMount
	ActionWater
	ActionExit
)

type Button struct {
	ID    ButtonAction
	Label string
}
