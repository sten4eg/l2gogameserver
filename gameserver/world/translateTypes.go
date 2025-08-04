package world

import (
	"math"
)

// FVector структура для координат
type FVector struct {
	X, Y, Z float64
}

// FVector2D структура для 2D координат
type FVector2D struct {
	X, Y float64
}

// CGeoCell структура ячейки
type CGeoCell struct {
	MData int16 // __int16 m_data
}

// CGeoSector структура сектора
type CGeoSector struct {
	MField          int16    // __int16 m_Field
	MDefaultCell    CGeoCell // CGeoCell m_DefaultCell
	MTerritoryMinZ  int16    // __int16 m_TerritoryMinZ
	MTerritoryMaxZ  int16    // __int16 m_TerritoryMaxZ
	MBooleanFlag    byte     // char m_BooleanFlag
	CellOffsetIndex int      // int cell_offset_index
	CellSegment     int      // int cell_segment
	NumCell         int16    // __int16 num_cell
}

// CSharedGeoZone структура общего гео зоны
type CSharedGeoZone struct {
	Sectors [256][256]CGeoSector // CGeoSector sectors[256][256]
	ZoneX   byte                 // char zone_x
	ZoneY   byte                 // char zone_y
	NumCell int                  // int num_cell
}

// CGeoDataIndexStruct структура индекса данных
type CGeoDataIndexStruct struct {
	NumCell          [340]int  // int num_cell[340]
	NumComplexSector [340]int  // int num_complex_sector[340]
	ZoneNum          int16     // __int16 zone_num
	Idx              [340]byte // char idx[340]
	Idy              [340]byte // char idy[340]
}

// CTerritory структура территории
type CTerritory struct {
	SName            string      // std::basic_string<wchar_t,std::char_traits<wchar_t>,std::xallocator<wchar_t> >
	MTerritoryPoints []FVector2D // std::vector<FVector2D,std::xallocator<FVector2D> >
	MBottom          int         // int m_Bottom
	MTop             int         // int m_Top
	MEast            int         // int m_East
	MWest            int         // int m_West
	MSouth           int         // int m_South
	MNorth           int         // int m_North
}

// CGeoZone структура зоны
type CGeoZone struct {
	MData                *CSharedGeoZone // CSharedGeoZone *m_data
	CastleId             int             // int castleId
	MNoFlyAreas          []*CTerritory   // std::vector<CTerritory *,std::xallocator<CTerritory *> >
	MNoCallPCAreas       []*CTerritory   // std::vector<CTerritory *,std::xallocator<CTerritory *> >
	MNoChangePCAreas     []*CTerritory   // std::vector<CTerritory *,std::xallocator<CTerritory *> >
	MNoDropItemAreas     []*CTerritory   // std::vector<CTerritory *,std::xallocator<CTerritory *> >
	MNoSaveBookmarkAreas []*CTerritory   // std::vector<CTerritory *,std::xallocator<CTerritory *> >
	MNoUseBookmarkAreas  []*CTerritory   // std::vector<CTerritory *,std::xallocator<CTerritory *> >
	MTransformableAreas  []*CTerritory   // std::vector<CTerritory *,std::xallocator<CTerritory *> >
}

// CGeoData структура данных геометрии
type CGeoData struct {
	MHfmIndex      interface{}             // void *m_hfm_index
	MzoneIndex     *CGeoDataIndexStruct    // CGeoDataIndexStruct *m_zone_index
	MHfmZones      [20][17]interface{}     // void *m_hfm_zones[20][17]
	MSharedZone    [20][17]*CSharedGeoZone // CSharedGeoZone *m_SharedZone[20][17]
	MZone          [20][17]CGeoZone        // CGeoZone m_Zone[20][17]
	MHfmCells      [20][17]interface{}     // void *m_hfm_cells[20][17]
	MCells         [20][17][]*CGeoCell     // CGeoCell *m_cells[20][17]
	MHfmCellOffset [20][17]interface{}     // void *m_hfm_cell_offset[20][17]
	MCellOffset    [20][17][]int16         // __int16 *m_cell_offset[20][17]
	MBadMap        [27][27]byte            // char m_bBadMap[27][27]
	MlstWaterPos   []FVector               // std::vector<FVector,std::xallocator<FVector> > m_lstWaterPos
}

// Структура для хранения массива указателей на массивы (вместо двойных указателей)
type CGeoCellArray struct {
	Cells []*CGeoCell
}
