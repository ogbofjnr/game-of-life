package main

import (
	"fmt"
	"time"
)

const gridSize = 25
const startX = 13
const startY = 13
const aliveSymbol = "+"
const deadSymbol = "-"
const dead = 0
const alive = 1

var neighborsCoordinates = [8][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {1, -1}, {-1, 1}, {1, 1}, {-1, -1}}

func main() {
	game := &GameOfLife{}
	game.placeGlider(startX, startY)

	for {
		game.print()
		game.tick()
	}
}

type GameOfLife struct {
	grid       [gridSize][gridSize]int8
	ticksCount int
}

// placeGlider place glider at specifies coordinates
func (g *GameOfLife) placeGlider(x, y int) {
	g.grid[x][y] = alive
	g.grid[x+1][y+1] = alive
	g.grid[x+2][y+1] = alive
	g.grid[x+2][y] = alive
	g.grid[x+2][y-1] = alive
}

// print clears previous grid and draw current one
func (g *GameOfLife) print() {
	fmt.Printf("\033[0;0H")
	fmt.Printf("Tick number %d\n", g.ticksCount)

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			cell := deadSymbol
			if g.grid[x][y] == alive {
				cell = aliveSymbol
			}
			fmt.Print(cell)
		}
		fmt.Println()
	}

	time.Sleep(time.Second)
}

// isAlive returns cell status according to the game rules
func (g *GameOfLife) isAlive(x, y int) int8 {
	var aliveNeighbors int8

	for i := 0; i < 8; i++ {
		//since we have a finite grid, stick the walls
		xx := (x + neighborsCoordinates[i][0] + gridSize) % gridSize
		yy := (y + neighborsCoordinates[i][1] + gridSize) % gridSize
		aliveNeighbors += g.grid[xx][yy]
	}

	if (g.grid[x][y] == alive && (aliveNeighbors == 2 || aliveNeighbors == 3)) || (g.grid[x][y] == dead && aliveNeighbors == 3) {
		return alive
	}

	return dead

}

// tick calculates next grid state
func (g *GameOfLife) tick() {
	var nextMatrix [gridSize][gridSize]int8
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			nextMatrix[i][j] = g.isAlive(i, j)
		}
	}
	g.grid = nextMatrix
	g.ticksCount++
}
