package main

// ALIVE If a cell is alive it has this value
const ALIVE byte = 0xff

// DEAD If a cell is dead it has this value
const DEAD byte = 0x00

func countNeighbours(cell cell, world [][]byte) int {
	nebs := 0
	yworld := len(world)
	xworld := len(world[0])

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 || j == 1 {
				continue
			}
			if world[i%yworld][j%xworld] == ALIVE {
				nebs++
			}
		}
	}

	return nebs
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {

	aliveSlice := make([]cell, 0)

	for row, cellRow := range world {
		for col, current := range cellRow {
			nebs := countNeighbours(cell{col, row}, world)

			if current == ALIVE {
				if nebs == 2 || nebs == 3 {
					aliveSlice = append(aliveSlice, cell{col, row})
				}
			} else {
				if nebs == 3 {
					aliveSlice = append(aliveSlice, cell{col, row})
				}
			}
		}
	}

	return aliveSlice
}
