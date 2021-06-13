package main

import "fmt"

type Tile int

const (
	TileStart Tile = iota
	TileBlank
	TileDoubles
	TileRosette
	TileEyes
	TileFinish
)

var BlueGameBoard = [16]Tile{TileStart, TileEyes, TileBlank, TileEyes, TileRosette, TileBlank, TileBlank, TileDoubles, TileRosette, TileBlank, TileDoubles, TileEyes, TileBlank, TileBlank, TileRosette, TileFinish}
var RedGameBoard = [16]Tile{TileStart, TileEyes, TileBlank, TileEyes, TileRosette, TileBlank, TileBlank, TileDoubles, TileRosette, TileBlank, TileDoubles, TileEyes, TileBlank, TileBlank, TileRosette, TileFinish}

func isSharedTile(tileIndex int) bool {
	return tileIndex >= 5 && tileIndex <= 12
}

func main() {
	tileIndex := 0
	if isSharedTile(tileIndex) {
		fmt.Printf("%d is shared!!!\n", tileIndex)
	} else {
		fmt.Printf("%d is NOT shared!!!\n", tileIndex)
	}
}
