package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

type IdRange struct {
	min ProductId
	max ProductId
}

type ProductId = []byte

func idToNumber(id ProductId) int {
	value, err := strconv.ParseInt(string(id), 10, 0)
	if err != nil {
		log.Fatalf("Invalid id %s", id)
	}
	return int(value)
}

func increment(id *ProductId) {
	digit_idx := len(*id) - 1
	for {
		(*id)[digit_idx]++
		if (*id)[digit_idx] <= '9' {
			break
		}
		(*id)[digit_idx] = '0'
		if digit_idx == 0 {
			*id = append([]byte{'1'}, *id...)
			break
		}
		digit_idx--
	}
}

func forEachId(id_range IdRange, f func(*ProductId)) {
	id := slices.Clone(id_range.min)
	for !bytes.Equal(id, id_range.max) {
		f(&id)
		increment(&id)
	}
	f(&id)
}

func validatePartOne(id ProductId) bool {
	half_size := len(id) / 2
	first := id[:half_size]
	second := id[half_size:]
	return !bytes.Equal(first, second)
}

func validatePartTwo(id ProductId) bool {
	size := len(id)
	half_size := size / 2
	seq_len := 1
	for seq_len <= half_size {
		if size%seq_len != 0 {
			seq_len++
			continue
		}
		all_equal := true
		first_view := id[0:seq_len]
		seq_start := 0
		for seq_start+seq_len <= size {
			view := id[seq_start : seq_start+seq_len]
			if !bytes.Equal(first_view, view) {
				all_equal = false
				break
			}
			seq_start += seq_len
		}
		if all_equal {
			return false
		}

		seq_len++
	}
	return true
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	invalid_sum_one := 0
	invalid_sum_two := 0
	clb := func(id *ProductId) {
		if !validatePartOne(*id) {
			invalid_sum_one += idToNumber(*id)
		}
		if !validatePartTwo(*id) {
			invalid_sum_two += idToNumber(*id)
		}
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadBytes(',')
		if len(str) > 0 {
			parts := bytes.Split(str[:len(str)-1], []byte("-"))
			if len(parts) != 2 {
				log.Fatal("Invalid range")
			}
			id_range := IdRange{min: parts[0], max: parts[1]}
			fmt.Printf("=== Range %s-%s ===\n", id_range.min, id_range.max)
			forEachId(id_range, clb)
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read file")
		}
	}

	fmt.Println("1. Invalid ID sum", invalid_sum_one)
	fmt.Println("2. Invalid ID sum", invalid_sum_two)
}
