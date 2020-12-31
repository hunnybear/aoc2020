//usr/bin/env go run $0 $@ ; exit
// part 1 == 273

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const TEST_RESULT = 112
const ACTIVE = '#'
const INACTIVE = '.'

var DEBUG = false

type Boundary struct {
	min Cell
	max Cell
}

type Cell struct {
	x int
	y int
	z int
}

type CoordinateSpace map[Cell]bool
type DimensionButNotLikeInGeometry CoordinateSpace

func get_neighbors(cell Cell) [26]Cell {
	// Neighboars. Like pigs but kinda like horses.
	var neighboars [26]Cell
	i := 0
	for x := cell.x - 1; x <= cell.x + 1; x++ {
		for y := cell.y - 1; y <= cell.y + 1; y++ {
			for z := cell.z - 1; z <= cell.z + 1; z++ {
				if x == cell.x && y == cell.y && z == cell.z {
					continue
				}
				neighboars[i] = Cell{x:x,y:y,z:z}
				i++
			}
		}
	}
	return neighboars
}

func get_cell_active(cell Cell, dimension DimensionButNotLikeInGeometry) bool {
	active_neighbors := 0
	is_active, present:= dimension[cell]
	if ! present {
		is_active = false
	}

	for _, neighbor := range get_neighbors(cell) {
		neighbor_active, _ := dimension[neighbor]
		if neighbor_active == true {active_neighbors ++}

	}
	if is_active && active_neighbors >=2 && active_neighbors <= 3 {
		return true
	} else if active_neighbors == 3 && ! is_active {
		return true
	}

	return false
}

func get_bounds(dimension DimensionButNotLikeInGeometry) Boundary {
	min_x := int((^uint(0)) >> 1)
	min_y := min_x
	min_z := min_x

	max_x := -min_x -1
	max_y := max_x
	max_z := max_x

	for cell, is_active := range dimension {
		if ! is_active {
			continue
		}

		if cell.x < min_x {
			min_x = cell.x
		} else if cell.x > max_x {
			max_x = cell.x
		}

		if cell.y < min_y {
			min_y = cell.y
		} else if cell.y > max_y {
			max_y = cell.y
		}

		if cell.z < min_z {
			min_z = cell.z
		} else if cell.z > max_z {
			max_z = cell.z
		}
	}

	return Boundary {
		min: Cell{x: min_x, y: min_y, z: min_z},
		max: Cell{x: max_x, y: max_y, z: max_z},
	}
}

func print(dimension DimensionButNotLikeInGeometry) {
	 bounds := get_bounds(dimension)
	 for z := bounds.min.z; z <= bounds.max.z; z++ {
	 	fmt.Printf("z: %d\n\n", z)
	 	for x := bounds.min.x; x <= bounds.max.x; x++ {
	 		for y := bounds.min.y; y <= bounds.max.y; y++ {
	 			is_active, _ := dimension[Cell{x:x, y:y, z:z}]
	 			char := INACTIVE
	 			if is_active{char=ACTIVE}
	 			fmt.Printf("%s", string(char))
	 		}
	 		fmt.Printf("\n")
	 	}
	 	fmt.Println("\n\n")
	 }
}

func generation(dimension DimensionButNotLikeInGeometry) DimensionButNotLikeInGeometry {
	
	to_check := CoordinateSpace{}
	new_dimension := DimensionButNotLikeInGeometry{}

	for current_cell, cell_active := range dimension{
		if cell_active {
			to_check[current_cell] = true
			for _, neighbor:= range get_neighbors(current_cell) {
				to_check[neighbor] = true
			}
		}
	}

	for cell, _ := range to_check {
		// if performance becomes an issue, delete rather than set to false
		new_dimension[cell] = get_cell_active(cell, dimension)
	}
	if DEBUG {print(new_dimension)}
	return new_dimension
}

func parse_input(input []byte) DimensionButNotLikeInGeometry {
	y := 0
	dimension := make(DimensionButNotLikeInGeometry)
	input_rows := bytes.Split(input, []byte{'\n'})
	for x, row := range input_rows {
		for z, char := range row {
			// in order to save loop length when iterating over all potentially
			// active points, don't store 'false' entries
			if char == ACTIVE {dimension[Cell{x:x,y:y,z:z}]=true}
		}
	}

	return dimension
}

func part_1(input []byte) int {
	dimension := parse_input(input)

	fmt.Println("begin")
	if DEBUG {print(dimension)}
	for i := 0; i < 6; i++ {
		if DEBUG {fmt.Printf("generation %d\n", i)}
		dimension = generation(dimension)
	}

	num_active := 0
	for _, is_active := range dimension {
		if is_active {num_active++}
	}
	return num_active
}

func main() {
	input_file := ""
	testing := false
	if len(os.Args) > 1  {
		testing = true
		input_file = "test_input"
	} else {
		testing = false
		input_file = "input"
	}
	input, err := ioutil.ReadFile(input_file)
	if err != nil {log.Fatal(err)}

	part_1_answer := part_1(input)

	if testing && part_1_answer != TEST_RESULT {
		log.Fatal(fmt.Sprintf("Test answer was wrong, got %d, expected %d", TEST_RESULT, part_1_answer))
	}

	fmt.Printf("Part 1: %d\n", part_1_answer)
}
