package gameserver

import (
	_ "embed"
	"l2gogameserver/packets"
	"log"
	"math"
)

// 1 файл геодаты - это 1 регион. Регион разделен на 256*256 блоков
// блоков 3 вида - FloatBlock, ComplexBlock, MultiBlock
// REGION - это 256*256 любых из 3 блоков, 1 регион это то что лежит в файле

//go:embed geodata/18_21.l2j
var b []byte

// Cell минимальная структура в геодате, может иметь ограничения прохода по NSWE
type Cell int16

// FlatBlock единый блок
type FlatBlock int16

// ComplexBlock состоит из 8*8 ячеек (Cell), разных по Z
type ComplexBlock struct {
	Cell []Cell
}

const (
	REGION_BLOCKS_X = 256
	REGION_BLOCKS_Y = 256
	REGION_BLOCKS   = REGION_BLOCKS_X * REGION_BLOCKS_Y
	REGION_CELLS_X  = REGION_BLOCKS_X * BlockCellsX
	REGION_CELLS_Y  = REGION_BLOCKS_Y * BlockCellsY
	REGION_CELLS    = REGION_CELLS_X * REGION_CELLS_Y

	GEO_REGIONS_Y = 32

	BlockCellsX = 8
	BlockCellsY = 8
	BLOCK_CELLS = BlockCellsX * BlockCellsY
)

func fillComplexBlock(r *packets.Reader) ComplexBlock {
	var cb ComplexBlock
	for i := 0; i < 64; i++ {
		cb.Cell = append(cb.Cell, Cell(r.ReadInt16()))
	}
	return cb
}

func fillFlatBlock(r *packets.Reader) FlatBlock {
	var k FlatBlock
	k = FlatBlock(r.ReadInt16())
	return k
}

//
func fillMultilayerBlock(r *packets.Reader) []int8 {
	start := r.CurrentIndex()
	for i := 0; i < 64; i++ {
		var kn int
		nLayer := r.ReadSingleByte() // if layer <=0 || >125 - panic
		if nLayer > 125 {
			panic("nLayer > 125 ! Invalid layers count")
		}
		kn = int(nLayer * 2)

		_ = r.ReadBytes(kn)
	}
	cur := r.CurrentIndex()
	r.UnreadBytes(int(cur - start))
	var res []int8
	for i := 0; i < int(cur-start); i++ {
		res = append(res, r.ReadInt8())
	}

	return res
}

var Blocks []interface{}

func Load() {

	r := packets.NewReader(b)
	res := make([]interface{}, 0, 65536)
	for i := 0; i < 65536; i++ {
		blockType := r.ReadSingleByte()
		switch blockType {
		case 0:
			res = append(res, fillFlatBlock(r))
		case 1:
			res = append(res, fillComplexBlock(r))
		case 2:
			res = append(res, fillMultilayerBlock(r))
			break
		}
		_ = blockType
	}
	Blocks = res
	var Co int
	qq := res[2835]
	qqw := qq.(ComplexBlock)

	var zZ []int16
	for _, v := range qqw.Cell {
		zZ = append(zZ, v.getHeight())
	}
	for _, v := range res {
		switch va := v.(type) {
		case ComplexBlock:
			Co += 64
		case FlatBlock:
			Co += 1
		case []int8:
			Co += len(va)
		}
	}

	//for _, v := range res {
	//	switch va := v.(type) {
	//	case ComplexBlock:
	//		for _, vv := range va.Cell {
	//			log.Fatal(vv.getHeight())
	//		}
	//	}
	//}

}

//func х(val int16) int16 {
//	x := (-7405<< 8) | (-7405 &0xff)
//	r := x & 0x0fff0
//		rr := r >> 1
//
//}

// Z для ячейки
func (cb Cell) getHeight() int16 {
	h := int16(int32(cb) & int32(0x0000FFF0))
	h = h >> 1
	return h
}

func getGeoX(worldX int32) int32 {
	if worldX >= -655360 && worldX <= 393215 {
		return (worldX - -655360) / 16
	}
	log.Fatal("Illegal world X in getGeoX")
	return 0
}

func getGeoY(worldY int32) int32 {
	if worldY >= -589824 && worldY <= 458751 {
		return (worldY - -589824) / 16
	}
	log.Fatal("Illegal world X in getGeoX")
	return 0
}

func FindPath(x, y, z, tx, ty, tz int) {
	gx := getGeoX(int32(x))
	gy := getGeoY(int32(y))

	rg := getRegion(gx, gy)
	bg := getBlock(gx, gy)

	t := Blocks[bg]
	tt := t.(ComplexBlock)

	gz := tt.getHeight(gx, gy, int32(z))
	_ = gz

	gtx := getGeoX(int32(tx))
	gty := getGeoX(int32(ty))

	gtz := tt.getHeight(gtx, gty, int32(tz))

	//alloc
	mapSize = int(64 + (2 * math.Max(math.Abs(float64(gx-gtx)), math.Abs(float64(gy-gty)))))

	buffer = make([][]CellNode, mapSize)
	for i, _ := range buffer {

		buffer[i] = make([]CellNode, mapSize)
	}

	ke := buffer
	_ = ke

	cellNodeFindPath(gx, gy, int32(gz), gtx, gty, int32(gtz))

	_, _ = gx, gy
	_ = rg
	_ = bg
	var i int
	i = 32
	_ = i
}

var (
	baseX int
	baseY int

	targetX int32
	targetY int32
	targetZ int32

	buffer [][]CellNode

	current CellNode
	mapSize int
)

type CellNode struct {
	next    *CellNode
	I       int
	isInUse bool
}

func cellNodeFindPath(x, y, z, tx, ty, tz int32) {

	baseX = int(x) + ((int(tx-x) - mapSize) / 2) // middle of the line (x,y) - (tx,ty)
	baseY = int(y) + ((int(ty-y) - mapSize) / 2) // will be in the center of the buffer

	targetX = tx
	targetY = ty
	targetZ = tz

	current = cellNodeGetNode(x, y, z)
}
func newC() CellNode {

	return CellNode{}
}
func cellNodeGetNode(x, y, z int32) CellNode {
	aX := int(x) - baseX

	if (aX < 0) || (aX >= mapSize) {
		return CellNode{}
	}

	aY := int(y) - baseY

	if (aY < 0) || (aY >= mapSize) {
		return CellNode{}
	}

	result := buffer[aX][aY]
	_ = result

	//if result == CellNode{}; {
	//
	//}
	return CellNode{}
}

func (b *ComplexBlock) getHeight(x, y, z int32) int16 {
	return b.getNearestZ(x, y, z)
}

func (b *ComplexBlock) getNearestZ(geoX, geoY, worldZ int32) int16 {
	return b.getCellHeight(geoX, geoY)
}

func (b *ComplexBlock) getCellHeight(geoX, geoY int32) int16 {
	height := int16(b.getCellData(geoX, geoY) & 0x0FFF0)
	return height >> 1
}

func (b *ComplexBlock) getCellData(geoX, geoY int32) int32 {
	n := ((geoX % BlockCellsX) * BlockCellsY) + (geoY % BlockCellsY)
	return int32(b.Cell[n])

}

//getRegion получение региона из глобальной мапы со всеми регионами карты
func getRegion(geoX, geoY int32) int32 {
	return ((geoX / REGION_CELLS_X) * GEO_REGIONS_Y) + (geoY / REGION_CELLS_Y)
}

func getBlock(geoX, geoY int32) int32 {
	return (((geoX / BlockCellsX) % REGION_BLOCKS_X) * REGION_BLOCKS_Y) + ((geoY / BlockCellsY) % REGION_BLOCKS_Y)
}

/**
 * проверка проходимости по прямой
 */
//jts GeoEngine 506 Line
//func canMove( __x,  __y,  _z,  __tx,  __ty,  _tz int,  withCollision bool,  geoIndex int)  {
//
//	_x := __x - models.MapMinX >> 4
//	_y := __y - models.MapMinY >> 4
//	_tx := __tx - models.MapMinX >> 4
//	_ty := __ty - models.MapMinY >> 4
//
//	diff_x := _tx - _x
//	diff_y := _ty - _y
//
//	incx := sign(diff_x)
//	incy := sign(diff_y)
//
//	overRegionEdge := (_x >> 11) != (_tx >> 11) || (_y >> 11) != (_ty >> 11)
//
//	if diff_x < 0 {
//		diff_x = - diff_x
//	}
//
//	if diff_y < 0 {
//		diff_y = -diff_y
//	}
//
//	var pdx, pdy, es, el int
//
//	if diff_x > diff_y {
//		pdx = incx
//		pdy = 0
//		es = diff_y
//		el = diff_x
//	} else {
//		pdx = 0
//		pdy = incy
//		es = diff_x
//		el = diff_y
//	}
//
//	err := el / 2
//
//	curr_x := _x
//	curr_y := _y
//	curr_z := _z
//
//	next_x := curr_x
//	next_y := curr_y
//	next_z := curr_z
//
//	 next_layer := []int16{2}
//	 temp_layer := []int16{2}
//	 curr_layer := []int16{2}
//
//	//NGetLayers(curr_x, curr_y, curr_layers, geoIndex);
//	//if (curr_layers[0] == 0) {
//	//	return true;
//	//}
//
//}
