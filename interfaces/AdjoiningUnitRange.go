package interfaces

type AdjoiningUnitRange struct {
	Floor int
	Start int
	End   int
}

func (AdjoiningUnitRange AdjoiningUnitRange) GetRange() int {
	return AdjoiningUnitRange.End - AdjoiningUnitRange.Start + 1
}

func (AdjoiningUnitRange AdjoiningUnitRange) CalcScore() int {
	// Happier^2
	return (AdjoiningUnitRange.GetRange() - 1) * (AdjoiningUnitRange.GetRange() - 1)
}

type AdjoiningUnitRangeList []AdjoiningUnitRange

func (AdjoiningUnitRangeList AdjoiningUnitRangeList) Len() int {
	return len(AdjoiningUnitRangeList)
}

func (AdjoiningUnitRangeList AdjoiningUnitRangeList) Less(i, j int) bool {
	return AdjoiningUnitRangeList[i].GetRange() > AdjoiningUnitRangeList[j].GetRange()
}

func (AdjoiningUnitRangeList AdjoiningUnitRangeList) Swap(i, j int) {
	AdjoiningUnitRangeList[i], AdjoiningUnitRangeList[j] = AdjoiningUnitRangeList[j], AdjoiningUnitRangeList[i]
}
