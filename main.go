
package main

import (
	"fmt"
	"os"
	"math/rand"

)

const (
	GAME_OF_LIFE_VERSION string = "1.0"
	MAX_GRID_WIDTH int = 10
	MAX_GRID_HEIGHT int = 10

	CELL_DEAD uint8 = 0
	CELL_ALIVE uint8 = 1
)

var (
	grid [MAX_GRID_WIDTH][MAX_GRID_HEIGHT]uint8

	grid_width int = MAX_GRID_WIDTH
	grid_height int = MAX_GRID_HEIGHT
	cycle int = 1
)

func main () {

	// get grid dimensions
	if len(os.Args) == 1 {
		setInitialRandomGridState ()
	}
	
	playGame (grid_width, grid_height)	
}

func allDead () bool {
	var count int

	for y := 0; y < grid_height; y++ {
		for x := 0; x < grid_width; x++ {
			if grid[x][y] == CELL_DEAD {
				count++
			}
		}
	}

	if count == grid_width * grid_height {
		return true
	}

	return false
}

func setInitialRandomGridState () {
	count := 10+rand.Intn (10)

	for cell := 0; cell < count; cell++ {
		xRand := rand.Intn (grid_width)
		yRand := rand.Intn (grid_height)
		grid[xRand][yRand] = CELL_ALIVE
	}
}

func playGame (x int, y int) {
	for !allDead () {
		showGrid ()
		evolveNextStep ()
		cycle++
	}
	showGrid ()
}

func showGrid (){
	fmt.Printf ("Cycle : %02d\n", cycle)
	for y := 0; y < grid_height; y++ {
		for x := 0; x < grid_width; x++ {
			switch grid[x][y] {
				case CELL_DEAD: fmt.Printf (".")
				case CELL_ALIVE: fmt.Printf ("O")
				default:
			}
		}
		fmt.Println ()
	}	
	fmt.Println ()
}

func countLiveNeighbours (x int, y int) int {
	var count int
	for deltaX := -1; deltaX <= 1; deltaX++ {
		for deltaY := -1; deltaY <= 1; deltaY++ {
			if deltaX !=0 && deltaY != 0 {
				if onGrid (x+deltaX, y+deltaY) {
					if grid[x+deltaX][y+deltaY] == CELL_ALIVE {
						count++
					}
				}
			}
		}
	}

	return count
}

func onGrid (x int, y int) bool {
	return (x >= 0) && (x < grid_width) && (y >= 0) && (y < grid_height)
}

func evolveNextStep () {
	 for y := 0; y < grid_height; y++ {
		for x := 0; x < grid_width; x++ {
				count := countLiveNeighbours (x,y)
				if grid[x][y] == CELL_DEAD 	{
					if count == 3 {
						grid[x][y] = CELL_ALIVE
					}
				} else 	{
					if count == 2 || count == 3 {
						grid[x][y] = CELL_ALIVE
					} else {
						grid[x][y] = CELL_DEAD
					}
				}
			}
		}
}