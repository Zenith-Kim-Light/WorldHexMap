package util

import "testing"

type testDataForCalcNearestPowersOfTwo struct {
	caseMessage          string
	width                int32
	height               int32
	expectedResultWidth  int32
	expectedResultHeight int32
}

func TestCalcNearestPowersOfTwo(t *testing.T) {
	var testData = []testDataForCalcNearestPowersOfTwo{
		{caseMessage: "case 1", width: 0, height: 0, expectedResultWidth: 1, expectedResultHeight: 1},
		{caseMessage: "case 2", width: 3, height: 3, expectedResultWidth: 4, expectedResultHeight: 4},
		{caseMessage: "case 3", width: 4, height: 4, expectedResultWidth: 4, expectedResultHeight: 4},
		{caseMessage: "case 4", width: 800, height: 0, expectedResultWidth: 1024, expectedResultHeight: 1},
		{caseMessage: "case 5", width: 0, height: 800, expectedResultWidth: 1, expectedResultHeight: 1024},
	}

	for _, td := range testData {
		resultWidth, resultHeight := CalcNearestPowersOfTwo(td.width, td.height)
		if resultWidth != td.expectedResultWidth {
			t.Error(td.caseMessage, "| Width", resultWidth, "not equals expected", td.expectedResultWidth)
		}
		if resultHeight != td.expectedResultHeight {
			t.Error(td.caseMessage, "| Height", resultHeight, "not equals expected", td.expectedResultHeight)
		}
	}
}
