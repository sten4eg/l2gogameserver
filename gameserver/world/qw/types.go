package qw

import (
	"math"
	"unsafe"
)

// === Константы ===

// GeoHeightMask_1 = -16 (в бинарном виде: 1111111111110000)
// После сдвига вправо на 1: 0b11111111111110000 >> 1 = 0b1111111111111000 = 0xFFF8
// Используется для извлечения высоты из m_data
var GeoHeightMask1 int16 = -16

// GeoFieldMask_2 — маски для типов зон (битовые флаги)
var GeoFieldMask2 = [16]int16{
	-32768, 16384, 8192, 4096,
	2048, 1024, 512, 256,
	128, 64, 32, 16,
	8, 4, 2, 1,
}
var GeoDirMask = [4]int16{1, 2, 4, 8}

// === Enum'ы ===

type GeoDirEnum int32

const (
	GEO_DIR_EAST       GeoDirEnum = 0x0
	GEO_DIR_WEST       GeoDirEnum = 0x1
	GEO_DIR_SOUTH      GeoDirEnum = 0x2
	GEO_DIR_NORTH      GeoDirEnum = 0x3
	GEO_DIR_ABOVE      GeoDirEnum = 0x4
	GEO_DIR_BELOW      GeoDirEnum = 0x5
	GEO_DIR_HORIZONTAL GeoDirEnum = 0x6
	GEO_DIR_MAX        GeoDirEnum = 0x7
)

type JumpType int32

const (
	NoJump   JumpType = 0
	Jumpable JumpType = 1
)

type GeoFieldTypeEnum int32

const (
	GEO_FIELD_BATTLE_RESIDENCE GeoFieldTypeEnum = 0x0
	GEO_FIELD_BATTLE_PVP       GeoFieldTypeEnum = 0x1
	GEO_FIELD_FREEZE           GeoFieldTypeEnum = 0x2
	GEO_FIELD_PEACE            GeoFieldTypeEnum = 0x3
	GEO_FIELD_WATER            GeoFieldTypeEnum = 0x4
	GEO_FIELD_STATUS_CHANGE    GeoFieldTypeEnum = 0x5
	GEO_FIELD_INSTANT          GeoFieldTypeEnum = 0x6
	GEO_FIELD_SSQ              GeoFieldTypeEnum = 0x7
	GEO_FIELD_NO_RESTART       GeoFieldTypeEnum = 0x8
	GEO_FIELD_SHUTTLE          GeoFieldTypeEnum = 0x9
	GEO_FIELD_DRAGON_LAIR      GeoFieldTypeEnum = 0xA
	GEO_FIELD_NO_SKILL         GeoFieldTypeEnum = 0xB
	GEO_FIELD_TRAINING_ZONE    GeoFieldTypeEnum = 0xC
	GEO_FIELD_ONLY_PVE         GeoFieldTypeEnum = 0xD
	GEO_FIELD_MAX              GeoFieldTypeEnum = 0x10
)

// === Основные структуры ===

// FVector — 3D-вектор (позиция)
// В C++: long double (128 бит), но в Go используем float64 как приближение
type FVector struct {
	X, Y, Z float64
}

func (a *FVector) LenSquared2D() float64 {
	return a.X*a.X + a.Y*a.Y
}

func (a *FVector) Normalize2D() {
	length := math.Sqrt(a.LenSquared2D())
	if length > 1e-6 {
		a.X /= length
		a.Y /= length
	}
}
func (v *FVector) SetFrom(other *FVector) {
	v.X = other.X
	v.Y = other.Y
	v.Z = other.Z
}
func (v *FVector) Len2DTo(other *FVector) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return dx*dx + dy*dy
}

// CGeoCell — минимальная ячейка ландшафта (16x16 юнитов)
// Хранит высоту и флаги направлений в m_data
type CGeoCell struct {
	Data int16 // m_data: содержит высоту и флаги (GeoDirMask)
}

// CGeoSector — блок 128x128 (8x8 ячеек)
type CGeoSector struct {
	Field           int16    // m_Field — битовая маска зон (GEO_FIELD_*)
	DefaultCell     CGeoCell // m_DefaultCell — если сектор "плоский"
	TerritoryMinZ   int16    // m_TerritoryMinZ
	TerritoryMaxZ   int16    // m_TerritoryMaxZ
	BooleanFlag     int8     // m_BooleanFlag — флаги: 1 = flat, 2 = ?
	CellOffsetIndex int32    // cell_offset_index
	CellSegment     int32    // cell_segment
	NumCell         int16    // num_cell
}

// CSharedGeoZone — данные зоны (256x256 секторов = 32км x 32км)
type CSharedGeoZone struct {
	Sectors [256][256]CGeoSector // sectors[256][256]
	ZoneX   int8                 // zone_x
	ZoneY   int8                 // zone_y
	NumCell int32                // num_cell
}

// CTerritory — зона с особыми правилами (NoFly, Water и т.д.)
type CTerritory struct {
	Name            string      // sName
	TerritoryPoints []FVector2D // m_TerritoryPoints
	Bottom          int32       // m_Bottom
	Top             int32       // m_Top
	East            int32       // m_East
	West            int32       // m_West
	South           int32       // m_South
	North           int32       // m_North
}

// FVector2D — используется в CTerritory
type FVector2D struct {
	X, Y float64
}

// CGeoZone — логическая зона с полиморфным поведением
// Реализует интерфейс GeoZone для полиморфизма
type CGeoZone struct {
	Data                *CSharedGeoZone
	CastleID            int32
	NoFlyAreas          []*CTerritory
	NoCallPCAreas       []*CTerritory
	NoChangePCAreas     []*CTerritory
	NoDropItemAreas     []*CTerritory
	NoSaveBookmarkAreas []*CTerritory
	NoUseBookmarkAreas  []*CTerritory
	TransformableAreas  []*CTerritory
}

// Интерфейс для полиморфного поведения (вместо виртуальной таблицы)
type GeoZone interface {
	IsNoFlyArea(pos FVector) bool
	IsInNoCallPCArea(pos FVector) bool
	IsInNoChangePCArea(pos FVector) bool
	IsInNoDropItemArea(pos FVector) bool
	IsInNoSaveBookmarkArea(pos FVector) bool
	IsInNoUseBookmarkArea(pos FVector) bool
	IsInTransformableArea(pos FVector) bool
}

// Реализация методов для CGeoZone
func (z *CGeoZone) IsNoFlyArea(pos FVector) bool {
	return z.isInTerritory(pos, z.NoFlyAreas)
}

func (z *CGeoZone) IsInNoCallPCArea(pos FVector) bool {
	return z.isInTerritory(pos, z.NoCallPCAreas)
}

func (z *CGeoZone) IsInNoChangePCArea(pos FVector) bool {
	return z.isInTerritory(pos, z.NoChangePCAreas)
}

func (z *CGeoZone) IsInNoDropItemArea(pos FVector) bool {
	return z.isInTerritory(pos, z.NoDropItemAreas)
}

func (z *CGeoZone) IsInNoSaveBookmarkArea(pos FVector) bool {
	return z.isInTerritory(pos, z.NoSaveBookmarkAreas)
}

func (z *CGeoZone) IsInNoUseBookmarkArea(pos FVector) bool {
	return z.isInTerritory(pos, z.NoUseBookmarkAreas)
}

func (z *CGeoZone) IsInTransformableArea(pos FVector) bool {
	return z.isInTerritory(pos, z.TransformableAreas)
}

// Вспомогательная функция: проверка попадания в территорию
func (z *CGeoZone) isInTerritory(pos FVector, territories []*CTerritory) bool {
	for _, t := range territories {
		if pos.X >= float64(t.West) && pos.X <= float64(t.East) &&
			pos.Y >= float64(t.South) && pos.Y <= float64(t.North) &&
			int32(pos.Z) >= t.Bottom && int32(pos.Z) <= t.Top {
			return true
		}
	}
	return false
}

// === Основной класс геоданных ===

// CGeoDataIndexStruct — индексная структура
type CGeoDataIndexStruct struct {
	NumCell          [340]int32
	NumComplexSector [340]int32
	ZoneNum          int16
	Idx              [340]int8
	Idy              [340]int8
}

// CGeoData — главный класс геометрии
type CGeoData struct {
	HfmIndex      unsafe.Pointer // m_hfm_index
	ZoneIndex     *CGeoDataIndexStruct
	HfmZones      [20][17]unsafe.Pointer  // m_hfm_zones
	SharedZone    [20][17]*CSharedGeoZone // m_SharedZone
	Zones         [20][17]CGeoZone        // m_Zone
	HfmCells      [20][17]unsafe.Pointer  // m_hfm_cells
	Cells         [20][17][]CGeoCell      // m_cells
	HfmCellOffset [20][17]unsafe.Pointer  // m_hfm_cell_offset
	CellOffset    [20][17][]int16         // m_cell_offset
	BadMap        [27][27]int8            // m_bBadMap
	LstWaterPos   []FVector               // m_lstWaterPos
}

// === Навигация: PathNode ===

// CPathNode — нода для AI-движения
type CPathNode struct {
	X, Y, Z  int32
	Neighbor [8]uint32 // индексы соседей в m_Data
}

// PathNodeWorldInfo — метаданные для нод
type PathNodeWorldInfo struct {
	SectionOffsetCount [20][17][128][128]uint32 // смещение и количество нод
	MaxIndex           int32                    // максимальный индекс
}

// CPathNodeWorld — менеджер навигационных нод
type CPathNodeWorld struct {
	Data                    []*CPathNode       // m_Data
	PathNodeDataHandle      unsafe.Pointer     // m_PathNodeDataHandle
	PathNodeWorldInfoHandle unsafe.Pointer     // m_PathNodeWorldInfoHandle
	WorldInfo               *PathNodeWorldInfo // m_WorldInfo
}
