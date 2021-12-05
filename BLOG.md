# What Am I Learning Each Day?

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