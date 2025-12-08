package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction uint8

const (
	Left Direction = iota
	Right
)

type Rotation struct {
	direction Direction
	value     int
}

type Safe struct {
	position     int
	zero_stops   int
	zero_crosses int
}

const (
	SafeMax   = 99
	SafeRange = SafeMax + 1
)

func parseRotation(line string) Rotation {
	var dir Direction
	switch line[0] {
	case 'L':
		dir = Left
	case 'R':
		dir = Right
	default:
		log.Fatalf("Invalid direction in %s", line)
	}
	value, err := strconv.ParseInt(line[1:], 10, 0)
	if err != nil {
		log.Fatalf("Invalid value in %s", line)
	}
	return Rotation{dir, int(value)}
}

func rotate(safe *Safe, rotation Rotation) {
	if rotation.direction == Left {
		// Check if we can reach zero
		to_zero := safe.position

		if to_zero > 0 && rotation.value >= to_zero {
			safe.position = 0
			rotation.value -= to_zero
			safe.zero_crosses++
		}

		// Calculate how many full rotations are left
		full_rotations := rotation.value / SafeRange
		safe.zero_crosses += full_rotations
		rotation.value -= SafeRange * full_rotations

		safe.position -= rotation.value
		if safe.position < 0 {
			safe.position += SafeRange
		}
	} else {
		// Check if we can reach zero
		to_zero := SafeRange - safe.position
		if to_zero > 0 && rotation.value >= to_zero {
			safe.position = 0
			rotation.value -= to_zero
			safe.zero_crosses++

		}

		// Calculate how many full rotations are left
		full_rotations := rotation.value / SafeRange
		safe.zero_crosses += full_rotations
		rotation.value -= SafeRange * full_rotations

		safe.position += rotation.value
	}

	if safe.position == 0 {
		safe.zero_stops++
	}
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safe := Safe{position: 50}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rotation := parseRotation(line)
		crosses_before := safe.zero_crosses
		rotate(&safe, rotation)

		fmt.Printf("After rotation %s -> ", line)
		if safe.position == 0 {
			fmt.Printf("\033[31m%d\033[0m", safe.position)
		} else {
			fmt.Printf("%d", safe.position)
		}
		if crosses_before != safe.zero_crosses {
			fmt.Printf(", %d (%d) crosses", safe.zero_crosses-crosses_before, safe.zero_crosses)
		}
		fmt.Println("")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("============")
	fmt.Printf("Final position: %d\n", safe.position)
	fmt.Printf("Zero stops: %d\n", safe.zero_stops)
	fmt.Printf("Zero crosses: %d\n", safe.zero_crosses)
}
