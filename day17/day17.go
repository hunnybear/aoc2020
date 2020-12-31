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
	min NCell
	max NCell
}

type NCell struct{
	coordinates []int
}

type Cell struct {
	x int
	y int
	z int
}



type CoordinateSpace map[NCell]bool
type DimensionButNotLikeInGeometry CoordinateSpace

func (space DimensionButNotLikeInGeometry) Validate() bool {
	dimensions := -1
	cell_dimensions := -1
	for cell, _ := range space {
		cell_dimensions = cell.Dimensions()
		if cell_dimensions != dimensions {
			if dimensions == -1 {
				dimensions = cell_dimensions
			} else {
				return false
			}
		}
	}

	if dimensions < 1 {
		return false
	}
	return true
}

func (space DimensionButNotLikeInGeometry) Dimensions() int {
	space.Validate()
	for cell, _ := range space {
		return cell.Dimensions()
	}
}

func (cell NCell) Dimensions() int {return len(cell.coordinates)}

func (cell NCell) Equals(other []int) bool {
	if len(cell.coordinates) != len(other) { return false }
	for idx, this_val := range cell.coordinates {
		if this_val != other[idx] {
			return false
		}
	}
	return true
}

func (cell NCell) GetNeighbors() []NCell {
	var neighboars = []NCell{}

	for _, neighboar_coords := range get_permutations(-1, 1, cell.Dimensions()) {
		if cell.Equals(neighboar_coords){continue}
		neighboars = append(neighboars, NCell{coordinates:neighboar_coords})
	}

	return neighboars
}

func get_neighbors(cell Cell) []Cell {
	// Neighboars. Like pigs but kinda like horses.
	var neighboars = []Cell{}

	for _, n_c := range get_permutations(-1, 1, 3) {
		neighboars = append(neighboars, Cell{x:n_c[0], y:n_c[1], z:n_c[2]})
	}

	return neighboars
}

func get_permutations(min int, max int, n int) [][]int {
	permutations := [][]int{}
	for element := min; element <= max; element++ {
		if n > 1 {
			for _, remainder := range get_permutations(min, max, n-1) {
				permutations = append(permutations, append([]int{element}, remainder...))
			}
		} else {
			permutations = append(permutations, []int{element})
		}
		
	}

	return permutations
}

func get_cell_active(cell NCell, dimension DimensionButNotLikeInGeometry) bool {
	active_neighbors := 0
	is_active, present:= dimension[cell]
	if ! present {
		is_active = false
	}

	for _, neighbor := range cell.GetNeighbors() {
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
	max_int := int((^uint(0)) >> 1)
	min_int := -min_val -1

	// Don't validate here, already done inside .Dimensions()
	dimensions := dimension.Dimensions()

	mins := make([]int, dimensions)
	maxes := make([]int, dimensions)

	for i := 0; i < dimensions; i++ {
		mins[i] = max_int
		maxes[i] = min_int
	}

	for cell, is_active := range dimension {
		if ! is_active {
			continue
		}

		for i := 0; i < dimensions; i++ {
			if cell.coordinates[i] < mins[i] {
				mins[i] = cell.coordinates[i]
			} else if cell.coordinates[i] > maxes[i] {
				maxes[i] = cell.coordinates[i]
			}
		}
	}

	return Boundary {
		min: NCell{coordinates: mins},
		max: NCell{coordinates: maxes},
	}
}

func print(dimension DimensionButNotLikeInGeometry) {
	//bounds := get_bounds(dimension)
	//for z := bounds.min.z; z <= bounds.max.z; z++ {
	//	fmt.Printf("z: %d\n\n", z)
	//	for x := bounds.min.x; x <= bounds.max.x; x++ {
	//		for y := bounds.min.y; y <= bounds.max.y; y++ {
	//			is_active, _ := dimension[Cell{x:x, y:y, z:z}]
	//			char := INACTIVE
	//			if is_active{char=ACTIVE}
	//			fmt.Printf("%s", string(char))
	//		}
	//		fmt.Printf("\n")
	//	}
	//	fmt.Println("\n\n")
	//}
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
	if DEBUG {
		print(new_dimension)
		new_dimension.Validate()	
	}
	return new_dimension
}

func parse_input(input []byte, dimensions int) DimensionButNotLikeInGeometry {
	// from the problem we're projecting these two-d input states into 3|4 dim
	// space (writing as N-dim cos fun). so we assume that fixed 2 d and then
	// allow the n to be flexible, just create the suffix additional dimensions
	// and call it a day

	suffix := []int{}
	for i := 2; i < dimensions; i++{suffix = append(suffix, 0)}

	cell := make(NCell)

	dimension := make(DimensionButNotLikeInGeometry)
	input_rows := bytes.Split(input, []byte{'\n'})
	for x, row := range input_rows {
		for y, char := range row {
			cell = NCell{append([]int{x, y}, suffix...)}
			// in order to save loop length when iterating over all potentially
			// active points, don't store 'false' entries
			if char == ACTIVE {dimension[cell]=true}
		}
	}

	dimension.Validate()
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

	fmt.Println(NCell{coordinates:[]int{0,0,0,0}}.GetNeighbors())
}
