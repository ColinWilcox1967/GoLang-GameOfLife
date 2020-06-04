
package main

import (
	"fmt"
	"os"
	"math/rand"
	"bufio"
	"strconv"
	"flag"
	"unicode"

	"github.com/eiannone/keyboard"
)

const (
	GAME_OF_LIFE_VERSION string = "1.0"
	MAX_GRID_WIDTH int = 10
	MAX_GRID_HEIGHT int = 10
	DEFAULT_CONFIG_FILE = "CONFIG.TXT"

	CELL_DEAD uint8 = 0
	CELL_ALIVE uint8 = 1
)

var (
	grid [MAX_GRID_WIDTH][MAX_GRID_HEIGHT]uint8

	grid_width int = MAX_GRID_WIDTH
	grid_height int = MAX_GRID_HEIGHT
	cycle int = 1
	configFile string = DEFAULT_CONFIG_FILE
	configLines []string
)

func main () {

	// get grid dimensions
	if len(os.Args) == 1 {
		setInitialRandomGridState ()
	}

	flag.StringVar(&configFile, "config", DEFAULT_CONFIG_FILE, "Board configuration file")
	flag.Parse ()
	
	if err, configLines := readConfigFileByLine (configFile); err == nil {
		grid_width,_ = strconv.Atoi(configLines[0])
		grid_height,_ = strconv.Atoi(configLines[1])
		clearGrid ()
		populateCells ()
	} else {
		setInitialRandomGridState ()
	}
    
	playGame (grid_width, grid_height)	
}

func readConfigFileByLine (filePath string) (error, []string) {
	
	 file, err := os.Open(filePath)
     if err != nil {
       return err, nil
     }
     defer file.Close()

	 scanner := bufio.NewScanner(file)
     for scanner.Scan() {
          configLines = append (configLines, scanner.Text ())
     }

     if err := scanner.Err(); err != nil {
     	return err, nil
     }

	 return nil, configLines
}

func clearGrid () {
	for y := 0; y < grid_height; y++ {
		for x := 0; x < grid_width; x++ {
			grid[x][y] = CELL_DEAD
		}
	}
}

func populateCells () {
	for line := 2; line < len(configLines); line++ {
		for offset := 0; offset < len(configLines[line]); offset++ {
			if configLines[line][offset] == 'o' {
				grid[offset][line-2] = CELL_ALIVE
			}
		}
	}	
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
	
	var exitGame bool

	for !allDead () && !exitGame {
		showGrid ()
		
		char, _, err := keyboard.GetSingleKey ()
		if err == nil {
			if byte(unicode.ToUpper(rune(char))) == 'X' {
				exitGame = true
			}	
		}

		if !exitGame {
			evolveNextStep ()
			cycle++
		}
	}
	showGrid ()
}

func showGrid (){
	for y := 0; y < grid_height; y++ {
		for x := 0; x < grid_width; x++ {
			switch grid[x][y] {
				case CELL_DEAD: fmt.Printf (".")
				case CELL_ALIVE: fmt.Printf ("o")
				default:
			}
		}
		if y == 0 {
			fmt.Printf ("   Cycle : %02d", cycle)
		} else if  y == grid_height - 1 {
			fmt.Printf ("   Press any key (X to Exit)")
		}
		fmt.Println ()
	}	
}

func countLiveNeighbours (x int, y int) int {
	var count int
	for deltaX := -1; deltaX <= 1; deltaX++ {
		for deltaY := -1; deltaY <= 1; deltaY++ {
				if onGrid (x+deltaX, y+deltaY) {
					if grid[x+deltaX][y+deltaY] == CELL_ALIVE {
						if !checkingSelf (deltaX, deltaY) {
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

func checkingSelf (dx int, dy int) bool {
	return (dx == 0) && (dy == 0)
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
			fmt.Println ()
		}
}