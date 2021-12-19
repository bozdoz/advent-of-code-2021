package main

import (
	"os"
	"testing"
)

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func TestParsing(t *testing.T) {
	input := "[[[[[9,8],1],2],3],4]"
	pair := parsePairs(input)

	if pair.String() != input {
		t.Logf("Answer should be %v, but got %v", input, pair.String())
		t.Fail()
	}
}

func TestExploding(t *testing.T) {
	tests := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
	}

	for input, expected := range tests {
		pair := parsePairs(input)

		// manually run once
		explode := pair.getNestedPairAtDepth(4)

		explode.explode()

		reduced := pair.String()

		if reduced != expected {
			t.Logf("Answer should be %v, but got %v", expected, reduced)
			t.Fail()
		}
	}
}

func TestGettingValueGTE(t *testing.T) {
	pair := parsePairs("[15,[0,13]]")
	splitPair := pair.getPairWithValueGTE(10)

	if pair != splitPair {
		t.Logf("Answer should be %v, but got %v", pair, splitPair)
		t.Fail()
	}

	pair = parsePairs("[[0,7],[15,[0,13]]]")
	splitPair = pair.getPairWithValueGTE(10)

	if splitPair != pair.right.pair {
		t.Logf("Answer should be %v, but got %v", pair.right.pair, splitPair)
		t.Fail()
	}

	pair = parsePairs("[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[11,9],[11,0]]]]")
	splitPair = pair.getPairWithValueGTE(10)
	expected := "[15,5]"

	if splitPair.String() != expected {
		t.Logf("Answer should be %v, but got %v", expected, splitPair)
		t.Fail()
	}
}

func TestSplitting(t *testing.T) {
	tests := map[string]string{
		"[[[[0,7],4],[15,[0,13]]],[1,1]]": "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
	}

	for input, expected := range tests {
		pair := parsePairs(input)

		// manually run split once
		splitting := pair.getPairWithValueGTE(splitVal)

		splitting.split(10)

		reduced := pair.String()

		if reduced != expected {
			t.Logf("Answer should be %v, but got %v", expected, reduced)
			t.Fail()
		}
	}
}

func TestReducing(t *testing.T) {
	pair := parsePairs("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
	expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"

	pair.reduce()

	reduced := pair.String()

	if reduced != expected {
		t.Logf("Answer should be %v, but got %v", expected, reduced)
		t.Fail()
	}
}

func TestSumming(t *testing.T) {
	tests := map[string][2]string{
		"[[1,2],[[3,4],5]]": {
			"[1,2]", "[[3,4], 5]",
		},
		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]": {
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		},
	}

	for expected, data := range tests {
		first := parsePairs(data[0])
		second := parsePairs(data[1])
		result := first.add(second)

		if result.String() != expected {
			t.Logf("Answer should be:\n %v, but got:\n %v", expected, result.String())
			t.Fail()
		}
	}

}

func TestMagnitude(t *testing.T) {
	tests := map[string]int{
		"[[1,2],[[3,4],5]]":                                     143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                     1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":                         445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":                         791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":                         1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}

	for str, expected := range tests {
		result := parsePairs(str)

		magnitude := result.getMagnitude()

		if magnitude != expected {
			t.Logf("Answer should be %v, but got %v", expected, result.String())
			t.Fail()
		}
	}
}

func TestPartOne(t *testing.T) {
	content := FileLoader("example.txt")
	expected := 4140

	result, err := PartOne(content)

	if err != nil {
		t.Log("got err, but shouldn't have: ", err)
		t.Fail()
	}

	if result != expected {
		t.Logf("Answer should be %v, but got %v", expected, result)
		t.Fail()
	}
}

func TestFailingSplit(t *testing.T) {
	pair := parsePairs("[[[[7,7],[7,8]],[[9,5],[8,0]]],[[[9,10],20],[8,[9,0]]]]")

	splitPair := pair.getPairWithValueGTE(10)
	expected := "[9,10]"

	if splitPair.String() != expected {
		t.Logf("expected %v got: %v", expected, splitPair)
		t.Fail()
	}
}

func TestStepByStep(t *testing.T) {
	pairs := parsePairs("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")
	cur := pairs.add(parsePairs("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"))
	expected := "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"

	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"))
	expected = "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"

	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]"))
	expected = "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"

	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[7,[5,[[3,8],[1,4]]]]"))
	expected = "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]"

	if cur.String() != expected {
		t.Logf("Expected: \n%v got: \n%v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[[2,[2,2]],[8,[8,1]]]"))
	expected = "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]"

	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[2,9]"))
	expected = "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]"

	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[1,[[[9,3],9],[[9,0],[0,7]]]]"))
	expected = "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]"

	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[[[5,[7,4]],7],1]"))
	expected = "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]"
	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}

	cur = cur.add(parsePairs("[[[[4,2],2],6],[8,7]]"))
	expected = "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"

	if cur.String() != expected {
		t.Logf("Expected: %v got: %v", expected, cur)
		t.Fail()
	}
}
