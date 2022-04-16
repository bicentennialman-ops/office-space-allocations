package interfaces

import "fmt"

type Position struct {
	Floor int
	Unit  int
}

func (position Position) ToString() string {
	return fmt.Sprintf("[%d-%d]", position.Floor+1, position.Unit+1)
}

type PositionList []Position

func (positionList PositionList) Len() int {
	return len(positionList)
}

func (positionList PositionList) Less(i, j int) bool {
	if positionList[i].Floor < positionList[j].Floor {
		return true
	}

	if positionList[i].Floor == positionList[j].Floor {
		return positionList[i].Unit < positionList[j].Unit
	}

	return false
}

func (positionList PositionList) Swap(i, j int) {
	positionList[i], positionList[j] = positionList[j], positionList[i]
}
