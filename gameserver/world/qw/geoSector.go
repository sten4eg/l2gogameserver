package qw

func (s *CGeoSector) GetFirstCellIndex(idx, idy int, offsetArray []int16) int {
	if s.BooleanFlag&1 != 0 {
		return 0
	}
	if s.BooleanFlag&2 != 0 {
		return s.CellSegment + 8*idx + idy
	}
	if idx != 0 || idy != 0 {
		i := (8*idx - 1) + s.CellOffsetIndex + idy
		return s.CellSegment + int(offsetArray[i])
	}
	return s.CellSegment
}

func (s *CGeoSector) GetLastCellIndex(idx, idy int, offsetArray []int16) int {
	if s.BooleanFlag&1 != 0 {
		return 0
	}
	if s.BooleanFlag&2 != 0 {
		return s.CellSegment + 8*idx + idy + 1
	}
	i := 8*idx + idy + s.CellOffsetIndex
	return s.CellSegment + int(offsetArray[i])
}
