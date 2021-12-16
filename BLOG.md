# What Am I Learning Each Day?

### Day 15

Tripped up by copying a [][]Cell over to a map[int]map[int]Cell implementation.

Caught in recursion trying to add neighbours to string representation

Learned Dijkstra's algorithm.

Wrote an implementation of a priority queue.

First time using type switch.

Really just running long processes today.

Initially tried my algorithms from day 12, but it was awfully long.

Finally got it:

```bash
> time ./run.sh 
Part One: 707 
Part Two: 2942 

real    0m37.778s
user    0m47.089s
sys     0m4.114s
```

### Day 14

I thought I figured this puzzle out before I began.  I thought to save a map of pairs, and never again think of the original template string.  This made it difficult to count the occurences of letters afterwards.  I had to refactor to also continually keep track of chars, which wasn't hard at all:

```go
for steps > 0 {
	steps--
	nextElements := newElements("")

	for ref, replacement := range polymer.insertionRules {
		count, ok := elements.pairs[ref]

		if !ok {
			continue
		}

		// remove pair from pairs so that we can merge
		// any unmatched pairs (they should persist)
		delete(elements.pairs, ref)

		// increment char count for replacement
		nextElements.charCount[replacement] += count

		// the replacement creates two new pairs
		// NN -> C creates NC & CN
		newPairs := []string{
			ref[0:1] + replacement,
			replacement + ref[1:2],
		}

		for _, newPair := range newPairs {
			nextElements.pairs[newPair] += count
		}
	}

	// merge any unmatched pairs
	elements.merge(nextElements)
}
```

This is my first time using `delete`.  Also it was nice to implement a merge function:

```go
type Elements struct {
	pairs, charCount map[string]int
}

func (elements *Elements) merge(otherElements Elements) {
	for key, val := range otherElements.pairs {
		elements.pairs[key] += val
	}

	for key, val := range otherElements.charCount {
		elements.charCount[key] += val
	}
}
```

I reused the `splitByEmptyNewline` function from Day 13, but ran into a problem with `fmt.Sscanf` because I hit an EOF!  So I refactored the function, adding `strings.TrimSpace`) and moved it to utils:

```go
func SplitByEmptyNewline(str string) []string {
	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strings.TrimSpace(str), -1)
}
```

I also hit a snag with `createDay.sh`, because I've started adding more than one go file in my daily puzzle directories.  I'm not sure how to adjust for that: maybe I should just have a template directory or something to copy from... (update: fixed)

First day implementing a proper min/max function too:

```go
func (elements *Elements) getMinMax() (int, int) {
	max := math.Inf(-1)
	min := math.Inf(1)
	...
}
```

Even though I hated comparing `float64` to `int`

### Day 13

Successfully altered today's script to use the `log` package instead of `fmt` for printing debug information to the console.  I made it so that the debug prints did not output to stdout while running, but only while testing.  I did this by introducing [init functions](https://go.dev/doc/effective_go#init) (for the first time).

thirteen.go
```go
// discard logs when script is run (overwritten in test file)
func init() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(ioutil.Discard)
}
```

thirteen_test.go
```go
// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}
```

Adding `log.SetFlags(log.Llongfile)` showed me the absolute path and line number, and vscode took me directly to the log statement!

I'm also glad that I've been getting into the habit of creating `String()` methods for my types, so that I can mimic the outputs on the site.  There was no way I was going to program a way to transform ascii art into a string.

### Day 12

I should really figure out how to use Logging instead of printing so much.

I used the debugger a lot for this puzzle.  For whatever reason, I was getting this output for the first example:

```
start,A,b,A,c,A,end
start,A,b,A,end
start,A,b,end
start,A,c,A,b,end
start,A,c,A,b,end,end
start,A,c,A,end
start,A,end
start,b,A,c,A,end
start,b,A,end
start,b,end
```

Note the single line with **"end,end"**. 😩

So my puzzle became figuring that out; which took hours, and I still don't know *why* it happened.

My failing code was:

```go
func (paths *Paths) addCave(path Path) {
	lastCave := path[len(path)-1]
```

And the *fix* was:

```go
func (paths *Paths) addCave(prevPath Path) {
	path := make(Path, len(prevPath))

	// make a copy!
	copy(path, prevPath)

	lastCave := path[len(path)-1]
```

So, unfortunately my day became a lesson in debugging and custom string representations of pointers.

**Update**

I took some time to split `twelve.go` into separate files.  I find Go's import/hoisting system bizarre, but thank goodness my IDE can follow the references...  I still don't know best practices around when a type should be in it's own file, for example. AFAIK, the package names in the files must be identical.

Also added this repo to Github Actions for continuously testing (make sure I don't break things when updating subsequent days).

**Second Update**

Slices are headers describing contiguous sections of backing arrays.  They just point to the head of an array (which holds the data).  Also `append`:

> The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated.

So the capacity of the path which was passed recursively to the traverse function was probably changing.  And the path itself was just a header for an array which held the underlying data.  So it was the same header being manipulated.  Explicitly adding a `copy()` I think was the right answer.

### Day 11

Felt better about today's puzzle.  Seemed complex enough to earn my attention and respect.

Tried to add a deferred function today, but it didn't actually work:

```go
	// any octopus with an energy level greater than 9 flashes
	for _, cell := range cells {
		// check for cell.flashed because
		// it may have been flashed by a neighbour
		// ...scandalous!
		if cell.energy > maxEnergy && !cell.flashed {
			cell.flash()
		}
	}

	// wanted to make this a deferred statement, but no
	for _, cell := range cells {
		if cell.flashed {
			flashes++
			cell.resolveFlash()
		}
	}
```

That was going to be:

```go
	for _, cell := range cells {
		if cell.energy > maxEnergy && !cell.flashed {
			cell.flash()
			// thought I could resolve the flash,
			// but forgot that neighbours flash themselves in another function...
			defer cell.resolveFlash()
		}
	}
```

Today was also my first time writing a custom string representation for printing (since that's typically how I debug)
:

```go
// custom string representation
func (cell *Cell) String() string {
	if cell.energy > maxEnergy {
		// flashing
		return "x"
	}

	return fmt.Sprintf("%d", cell.energy)
}
```

I did this because the `Grid` has references to `[][]*Cell`, and `Cell` has references to `[]*Cell`; so when I printed, I kept getting their addresses.  I believe I set up the pointers correctly this time, and successfully linked all of the cells to their references, otherwise I wouldn't have been able to have the `flash` method on the `Cell` type (would have needed access to the grid).  Also I found the cell's neighbours when setting up the grid, so that I didn't have to traverse the grid on each step/update.

I'm still unsure if I need to be returning a pointer in a generic constructor-like function or not:

```go
func newGrid(data []string) *Grid {
	grid := &Grid{}
	// ... etc	
	return grid
}
```

In previous days, I had just returned the struct, like this:

```go
func newGrid(data []string) (grid Grid) {
	return
}
```

But I ran into issues with a named result parameter for a pointer, trying this:

```go
func newGrid(data []string) (grid *Grid) {
	// grid is not initialized!! It's nil!
	return
}
```

**Update**

I found [this link](https://stackoverflow.com/a/44827234/488784) which suggests returning a value duplicates the memory allocation; though [another answer](https://medium.com/@philpearl/bad-go-pointer-returns-340f2da8289) suggests that returning pointers have more overhead if the data is small and/or short-lived

> In the pointer case the memory has to be allocated on the heap, which will take about 25ns, then the data is set up (which should take about the same time in both cases), then the pointer is written to the stack to return the struct to the caller. In the value case there’s no allocation, but the whole struct has to be copied onto the stack to return it to the caller.

> If the lifetime of the returned data was longer the results could be very different. But perhaps this is an indication that returning pointers to structs that have a short lifetime is bad Go.

I don't think the data structure is very big or long-lived, so maybe returning a pointer for `newGrid` is a mistake.  And maybe not!? 🤷‍♀️

So, I thought let's benchmark it.  I copy-pasted the `newGrid` function and made one return a value, and the other return a pointer.   Created my first two benchmark tests:

```go
func BenchmarkGridValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo(vals)
	}
}

func BenchmarkGridPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwoPointer(vals)
	}
}
```

```bash
> go test -bench BenchmarkGrid

BenchmarkGridValue-2                2239            519795 ns/op
BenchmarkGridPointer-2              2103            518339 ns/op
```

So I guess they are about equal.

### Day 10

Definitely puzzle fatigue.

I genuinely spent more time reading the puzzle than answering it.  I had issues understanding what I was being asked to do.  VSCode also had a bracket-pair colorizer which kind of threw me off.  And I think there was less hand-holding for this puzzle than the last several.

Had a rather large type structure for brackets and lines (which grew to be larger after part two):

```go
type Bracket struct {
	isOpen                          bool
	pair                            rune
	corruptScore, autocompleteScore int
}

type Stack []rune

type Line struct {
	isCorrupted, isIncomplete bool
	corruptedBy               rune
	incompleteBrackets        Stack
}
```

I defined each bracket type immediately as to whether it was **open** or not; I also wanted to explicitly pair them for reference.  So my logic was: 1. I generally ignore open brackets (add them to the stack), and compared closing brackets with the tail of the stack. Corrupted lines need to know which bracket it's corrupted by, and incomplete lines need to know which open brackets are unpaired.

I learned that I ought to be using pointers as opposed to copying data structures in every method.  I wasn't aware that was happening yesterday, and I should revisit Day 9.

### Day 9

Think I'm getting puzzle fatigue.  Kind of fun today as I've never done any path-tracing, or whatever that's called algorithms before.  I think what I did with the search function worked fine enough.  Only did one pass, so I haven't critiqued my own work yet.  I think the separation of concerns are mostly there.  Much better than yesterday's first pass anyway.

```go
func (b *basin) search(r, c int) {
	// we only search cells we know are within the basin
	if b.included[r] == nil {
		b.included[r] = make(map[int]bool, b.cols)
	}

	b.included[r][c] = true

	neighbours, err := b.heights.findNeighbours(r, c)

	if err != nil {
		panic(fmt.Sprintf("failed to find neighbours for %d, %d", r, c))
	}

	for _, coords := range neighbours {
		r, c = coords[0], coords[1]
		included := b.isIncluded(r, c)

		if included || b.heights[r][c] == 9 {
			continue
		}

		// search again!
		b.search(r, c)
	}
}
```

Did some panicking.  Maybe that's appropriate for these advent challenges.  Why bother setting up full error catching for a non-production app?

IDE helped me simplify a type, from:

```go
// go helped fix this previously redundant type
indices := [][]int{
	[]int{col, row - 1},
	[]int{col, row + 1},
	[]int{col - 1, row},
	[]int{col + 1, row},
}
```

To: 

```go
// go helped fix this previously redundant type
indices := [][]int{
	{col, row - 1},
	{col, row + 1},
	{col - 1, row},
	{col + 1, row},
}
```

### Day 8

Added string sorting.  Tried to do some kind of string intersection to figure out the values, but the functions are probably lacking.  I saw a bunch of errors and warnings, but kind of let them be for today.  

Part one was awfully boring. 

Part two could have been fun.  I wish I had more time to clean it up.

I did object-oriented because I wanted to keep track of two datasets: the segments and the digits

I ran into this issue, where it was just a pain to assign to a nested struct:

```go
// kind of a mess working around UnaddressableFieldAssign
//
// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source=gopls#UnaddressableFieldAssign
func (entry *Entry) setPatternValue(pattern string, value int) {
	pat := entry.patterns[pattern]
	pat.value = value
	entry.patterns[pattern] = pat
}
```

I ignored all errors, had to use the test debugger, and am generally unhappy with the naming/logic of some functions.  Brute force!

**Update**

After refactor I had an easier time creating custom types and methods.  I think it made it a little more readable and consistent:

```go
type Pattern string

func newPattern(p string) Pattern {
	p = utils.SortString(p)

	return Pattern(p)
}

// all parts of b are inside of a
func (a *Pattern) contains(b Pattern) bool {
outer:
	for _, x := range b {
		for _, y := range *a {
			if x == y {
				continue outer
			}
		}
		return false
	}

	return true
}
```

After that I had an easier time calling `pattern.contains(candidate)`. And then I made an alias for []Pattern:

```go
// created purely to make the remove method
type PatternArr []Pattern

func (arr PatternArr) remove(pattern Pattern) (out PatternArr) {
	for _, val := range arr {
		if val != pattern {
			out = append(out, val)
		}
	}

	return
}
```

Now I can call `remaining = remaining.remove(pattern)`


### Day 7

Wrote some stats functions which I thought would help solve the problem, if not outright solve it.  Saved them in utils: `Mean`, `Median`, `Mode`, `Stdev`.

Decided to just iterate all possible solutions and pick the best.  I did this by iterating the values as unique sorted sets.

Realized I could probably pick the last descending value (i.e. the values should continue to go down until the minimum value is reached and the values go up again).  I couldn't imagine a scenario where the values would again decrease (in a sorted set).  So that's the logic I went with.

In part one I only iterated the data provided.  But seeing as the solution to part two was a number/position that was not present in the puzzle input, I decided I needed to iterate from the min value to max value by 1 step.

So far the most boring puzzle solution...

### Day 6

You can run a go module with `go run .` and it will just find the file you want to run (maybe it will run all of them)

Ran into issues casting `[9]int` as `[]int`, and not sure how to get around it to use `utils.Sum`.   I thought about changing the data structure to just use []int and initialize with `make([]int, 9)`, but it seemed annoying to do that in every place that needed it initialized. (**edit** someone on r/adventofcode correctly pointed out that I could pass a slice of `state[:]` to `utils.Sum`)

```go
return utils.Sum(state[:]...), err
```

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
