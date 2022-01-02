package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	golog "log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/types"
)

func LoadAsLines(filename string) (lines []string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// I don't know what this line does
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}

func LoadInts(filename string) (nums []int) {
	vals := LoadAsLines(filename)

	for _, val := range vals {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	return
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

// https://golangbyexample.com/sort-string-golang/
type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func SortString(str string) string {
	arr := sortRuneString(str)

	sort.Sort(arr)

	return string(arr)
}

func SplitByEmptyNewline(str string) []string {
	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strings.TrimSpace(str), -1)
}

// TODO: would love to simply extend this
type ExtendedLogger struct {
	Logger golog.Logger
}

// used to make a new line after the Llongfile
func (log *ExtendedLogger) Println(v ...interface{}) {
	log.Logger.Output(2, "\n"+fmt.Sprintln(v...))
}

func (log *ExtendedLogger) SetOutput(w io.Writer) {
	log.Logger.SetOutput(w)
}

func (log *ExtendedLogger) Printf(format string, v ...interface{}) {
	log.Logger.Output(2, "\n"+fmt.Sprintf(format, v...))
}

// custom logger, purely to add a new line in Println
func Logger(args ...interface{}) *ExtendedLogger {
	// default empty prefix
	prefix := ""

	if len(args) > 0 {
		// first type switch
		switch v := args[0].(type) {
		case string:
			prefix = fmt.Sprintf("[%s] ", v)
		}
	}

	newLog := golog.New(os.Stdout, prefix, golog.Llongfile)

	return &ExtendedLogger{
		Logger: *newLog,
	}
}

func MinInt(nums ...int) int {
	min := nums[0]

	for _, val := range nums {
		if val < min {
			min = val
		}
	}

	return min
}

func MaxInt(nums ...int) int {
	max := nums[0]

	for _, val := range nums {
		if val > max {
			max = val
		}
	}

	return max
}

func BinaryToInt[T ~string](bin T) (int, error) {
	s := string(bin)

	val, err := strconv.ParseInt(s, 2, 64)

	return int(val), err
}

func AbsInt[T types.Numeric](num T) T {
	if num < 0 {
		return -num
	}
	return num
}
