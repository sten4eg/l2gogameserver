package qw

func (s *CGeoSector) GetFirstCellIndex(idx, idy int, offsetArray []int16) int {
	if s.BooleanFlag&1 != 0 {
		return 0
	}
	if s.BooleanFlag&2 != 0 {
		return int(s.CellSegment) + 8*idx + idy
	}

	if idx != 0 || idy != 0 {
		i := (8*idx - 1) + int(s.CellOffsetIndex) + idy
		return int(s.CellSegment) + int(offsetArray[i])
	}
	return int(s.CellSegment)
}

func (s *CGeoSector) GetLastCellIndex(idx, idy int, offsetArray []int16) int {
	if s.BooleanFlag&1 != 0 {
		return 0
	}
	if s.BooleanFlag&2 != 0 {
		return int(s.CellSegment) + 8*idx + idy + 1
	}
	i := 8*idx + idy + int(s.CellOffsetIndex)
	return int(s.CellSegment) + int(offsetArray[i])
}
