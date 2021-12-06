# What Am I Learning Each Day?

### Day 6

You can run a go module with `go run .` and it will just find the file you want to run (maybe it will run all of them)

Ran into issues casting `[9]int` as `[]int`, and not sure how to get around it to use `utils.Sum`.   I thought about changing the data structure to just use []int and initialize with `make([]int, 9)`, but it seemed annoying to do that in every place that needed it initialized.

On the initial run, I used an awful data structure to run the code (exponentially growing array).

Assuming I should give up, I reached out to thatguygriff who said he had to redo his.  So, fine, I changed the logic to just keep track of state of counters (0 - 9) and just a count of counters `[9]int`.  Worked great.  Didn't crash the compiler.

I thought against using custom object methods, and I initially tried passing state by reference, but found it difficult to pass it through two functions; also the current logic might be difficult to maintain, since the day incrementer manages two states:

```go
// incrementDay decrements fish counters
func incrementDay(state State) (newState State) {
	for i, val := range state {
		if i == 0 {
			// 0-day counters are both moved to 8 and copied to 6
			newState[8] = val
			newState[6] += val
		} else {
			// move counter value down a day
			newState[i-1] += val
		}
	}

	return
}
```

if I did this with one state, I could try a different pattern:

```go
// incrementDay decrements fish counters
func incrementDay(state *State) {
	temp := state[0]

	// moves all counters to previous index
	for i, val := range state[1:] {
		state[i] = val
	}

	// 0-day counters are both moved to 8 and copied to 6
	state[8] = temp
	state[6] += temp
}
```

I have no idea which is better 🤷‍♀️

### Day 5

First time doing OOP with types and type methods.  I think I wanted the load function to calculate the space at the same time as loading the lines, so maybe it made sense as a method instead of a pure function. First time using fmt.Fscanf.

Today seemed straight forward, but maybe because I've done so much work with 2d space.

It wasn't apparent that I needed to calculate the 2d space, given the example was a 10x10 grid, but I figured I should just to be safe, and I'm glad I accounted for it.

I used a 1d grid to hold the space, because I wasn't sure how to fill an empty 2d grid with the proper size.  Maybe a for loop?

```go
g.space = make([][]int, maxY+1)

for i := range g.space {
  g.space[i] = make([]int, maxX + 1)
}
```

Could have also done a `map`, I think:

```go
g.space = map[int]map[int]int{}

if g.space[0] == nil {
  g.space[0] = map[int]int{}
}
```

Discovered a better way to incrementally test the app: use a `map[int][int]` for each part's answers, and check for `ok` in the tests:

```go
var answers = map[int]int{
	1: 5,
	2: 12,
}

func TestExampleOne(t *testing.T) {
	expected, ok := answers[1]

  // Especially helpful in PartTwo, which always fails when I'm only 
  // testing PartOne!  Now it will just not run the full test if I
  // don't yet give it an expected answer.
	if !ok {
		return
	}
  // ...
}
```

### Day 4

Finally created run.sh and test.sh which accepts env vars to test individual directories

Giving up on the idea of generic array filter functions.

This was my first problem using references.  I found the typings difficult.

I also typed the boards `type Bingo [5][5]int`.  But still have yet to add any type methods. 

I read the entire input as a string, which actually seemed pretty easy.

My idea was to maintain two boards: one of values, and one of whether the values were drawn (`[5][5]bool`) (I could have done a single board which matched both, whatever). 

For each draw I iterated each board, each row, each value. If the value was there I immediately checked the row and column for a line of truthy values; if it was there I skipped to the next board (first time using labeled for loops). As soon as the board is a winner, I remove it from the list of boards.

### Day 3

Learned how difficult it is to do this without vscode's remote container. Had to flesh out the Dockerfile to continually run `go test` in the day folder.

Ran into painpoints like creating generic array filter functions (like JavaScript).

Worked with the binaries as strings, which likely caused a lot of logic in place of using binary operators which I'm far less familiar with.

My logic was to transpose the input data to work on the columns first.  Found it annoying to convert from the string binary to int, and also to convert `int64` to `int`

This was my first problem using named return results.

### Day 2

Learned about `go test ./...` which will test all directories.

Failed to recognize that single quotes was `rune` type, so found out that `strings.Fields(str)` could work just as well as `strings.Split(str, " ")`

### Day 1

Learned a lot of setup, Docker; learned that I'm not using a version of golang that has Generics 🤦‍♀️

Decided I wanted to keep using one.go, two.go, etc. instead of main.

Learned about spreading
