package main

import (
	"fmt"
	"sort"

	"github.com/bicentenninal96/office-space-allocations/interfaces"
)

// BEGIN INPUT
const NUMBER_OF_POSITIONS_REQUIRED int = 8

const NUMBER_FLOORS int = 4
const NUMBER_UNITS int = 12

var SPACE = [][]int{
	{1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1},
	{1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1},
	{1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1},
	{1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0},
}

// END INPUT

var adjoiningUnitRanges interfaces.AdjoiningUnitRangeList

var countFreeUnits []int = make([]int, NUMBER_FLOORS)

var minFloorUsed = NUMBER_FLOORS
var maxFloorRelationScore = 0
var maxUnitsRelationScore = 0
var bestUnitsRangesSelected interfaces.AdjoiningUnitRangeList

// Count adjoining units and count free units for each floor
func Prepare() {
	for floor := 0; floor < NUMBER_FLOORS; floor++ {
		unit := 0
		freeUnits := 0
		for unit < NUMBER_UNITS {
			if SPACE[floor][unit] == 0 {
				adjoiningUnits := 0
				start := unit
				end := unit
				for unit < NUMBER_UNITS && SPACE[floor][unit] == 0 {
					adjoiningUnits++
					freeUnits++
					end = unit
					unit += 1

				}
				adjoiningUnitRanges = append(adjoiningUnitRanges, interfaces.AdjoiningUnitRange{
					Floor: floor,
					Start: start,
					End:   end,
				})

			}
			unit += 1
		}
		countFreeUnits[floor] = freeUnits
	}
	sort.Sort(adjoiningUnitRanges)

}

// Calc solution score by floor position
func CalcFloorsRelationScore(floorsSelected []int) int {
	floor := 0
	relationScore := 0
	for floor < len(floorsSelected) {
		if floorsSelected[floor] == 1 {
			continuous := -1
			for floor < len(floorsSelected) && floorsSelected[floor] == 1 {
				continuous += 1
				floor += 1
			}

			// Happier^2
			relationScore += continuous * continuous
		}
		floor += 1
	}
	return relationScore
}

// Calc solution score by floor position and unit position
func CalcUnitsRelationScore(floorsSelected []int) (int, interfaces.AdjoiningUnitRangeList) {
	countUnits := 0
	relationScore := 0
	var rangesSelected interfaces.AdjoiningUnitRangeList
	for _, item := range adjoiningUnitRanges {
		if floorsSelected[item.Floor] == 0 {
			continue
		}
		countUnits += item.GetRange()
		relationScore += item.CalcScore()
		rangesSelected = append(rangesSelected, item)
		if countUnits > NUMBER_OF_POSITIONS_REQUIRED {
			break
		}
	}
	return relationScore, rangesSelected
}

// Consider the right solution
func ConsiderSolution(floorsSelected []int) {
	floorsUsed := 0
	for floor := 0; floor < NUMBER_FLOORS; floor++ {
		if floorsSelected[floor] == 1 {
			floorsUsed += 1
		}
	}
	unitsRelationScore, unitsRangeSelected := CalcUnitsRelationScore(floorsSelected)
	if minFloorUsed > floorsUsed {
		minFloorUsed = floorsUsed
		maxFloorRelationScore = CalcFloorsRelationScore(floorsSelected)
		maxUnitsRelationScore = unitsRelationScore
		bestUnitsRangesSelected = unitsRangeSelected
		return
	}

	if minFloorUsed < floorsUsed {
		return
	}

	floorRelationScore := CalcFloorsRelationScore(floorsSelected)

	if maxFloorRelationScore < floorRelationScore {
		maxFloorRelationScore = floorRelationScore
		bestUnitsRangesSelected = unitsRangeSelected
		return
	}

	if maxFloorRelationScore == floorRelationScore {
		maxFloorRelationScore = floorRelationScore
		bestUnitsRangesSelected = unitsRangeSelected
	}
}

// Check all solutions
func CheckFloor(countUnits, numFloor int, floorsSelected []int) {
	if countUnits < NUMBER_OF_POSITIONS_REQUIRED && numFloor < len(floorsSelected) {
		floorsSelected[numFloor] = 1
		CheckFloor(countUnits+countFreeUnits[numFloor], numFloor+1, floorsSelected)
		floorsSelected[numFloor] = 0
		CheckFloor(countUnits, numFloor+1, floorsSelected)
		return
	}
	if countUnits < NUMBER_OF_POSITIONS_REQUIRED {
		return
	}
	ConsiderSolution(floorsSelected)
}

func PrintResults() {
	if bestUnitsRangesSelected.Len() == 0 {
		fmt.Println("Couldn't find any solution")
	}
	var positions interfaces.PositionList
	numberLocations := 0
	for _, unitsRange := range bestUnitsRangesSelected {
		for i := 0; i < unitsRange.GetRange(); i++ {
			positions = append(positions, interfaces.Position{
				Floor: unitsRange.Floor,
				Unit:  unitsRange.Start + i,
			})
			numberLocations++
			if NUMBER_OF_POSITIONS_REQUIRED == numberLocations {
				break
			}
		}
		if NUMBER_OF_POSITIONS_REQUIRED == numberLocations {
			break
		}
	}
	sort.Sort(positions)
	for _, position := range positions {
		fmt.Printf(position.ToString())
	}
}

func main() {
	Prepare()
	selectedFloors := make([]int, NUMBER_FLOORS)
	for floor := range selectedFloors {
		selectedFloors[floor] = 0
	}
	CheckFloor(0, 0, selectedFloors)
	PrintResults()
}
