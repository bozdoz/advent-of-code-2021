package utils

import (
	"bufio"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func LoadFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// I don't know what this line does
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func LoadInts(filename string) []int {
	vals := LoadFile(filename)
	nums := []int{}

	for _, val := range vals {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	return nums
}

func LoadAsString(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

// TODO: how can this output an []int with a size: [9]int
func LoadCSVInt(filename string) (out []int) {
	data := LoadAsString(filename)
	lines := strings.Split(data, "\n")
	vals := strings.Split(lines[0], ",")

	for _, val := range vals {
		i, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		out = append(out, i)
	}

	return
}

func Sum(nums ...int) (s int) {
	for _, val := range nums {
		s += val
	}

	return
}

func Mean(arr []int) float64 {
	n := len(arr)

	sum := Sum(arr...)

	return float64(sum) / float64(n)
}

func Median(arr []int) float64 {
	min := arr[0]
	max := min

	for _, val := range arr {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}

	return float64(max-min)/float64(2) + float64(min)
}

func Mode(arr []int) int {
	count := map[int]int{}

	for _, val := range arr {
		count[val]++
	}

	maxCount := 0
	maxVal := arr[0]

	for key, val := range count {
		if val > maxCount {
			maxCount = val
			maxVal = key
		}
	}

	return maxVal
}

func Stdev(vals []int) float64 {
	mean := Mean(vals)
	n := len(vals)

	var sum float64

	for _, val := range vals {
		x := float64(val) - mean
		sum += x * x
	}

	return math.Sqrt(sum / float64(n))
}
