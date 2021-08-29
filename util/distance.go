package util

type EditOperation int

type EditScript []EditOperation

type MatchFunction func(rune, rune) bool

// IdenticalRunes is the default MatchFunction: it checks whether two runes are
// identical.
func IdenticalRunes(a rune, b rune) bool {
	return a == b
}

type Options struct {
	InsCost int
	DelCost int
	SubCost int
	Matches MatchFunction
}

// DefaultOptions is the default options without substitution: insertion cost
// is 1, deletion cost is 1, substitution cost is 2 (meaning insert and delete
// will be used instead), and two runes match iff they are identical.
var DefaultOptions Options = Options{
	InsCost: 1,
	DelCost: 1,
	SubCost: 2,
	Matches: IdenticalRunes,
}

// DefaultOptionsWithSub is the default options with substitution: insertion
// cost is 1, deletion cost is 1, substitution cost is 1, and two runes match
// iff they are identical.
var DefaultOptionsWithSub Options = Options{
	InsCost: 1,
	DelCost: 1,
	SubCost: 1,
	Matches: IdenticalRunes,
}

// DistanceForStrings returns the edit distance between source and target.
//
// It has a runtime proportional to len(source) * len(target) and memory use
// proportional to len(target).
func DistanceForStrings(source []rune, target []rune) int {
	op := DefaultOptionsWithSub
	// Note: This algorithm is a specialization of MatrixForStrings.
	// MatrixForStrings returns the full edit matrix. However, we only need a
	// single value (see DistanceForMatrix) and the main loop of the algorithm
	// only uses the current and previous row. As such we create a 2D matrix,
	// but with height 2 (enough to store current and previous row).
	height := len(source) + 1
	width := len(target) + 1
	matrix := make([][]int, 2)

	// Initialize trivial distances (from/to empty string). That is, fill
	// the left column and the top row with row/column indices multiplied
	// by deletion/insertion cost.
	for i := 0; i < 2; i++ {
		matrix[i] = make([]int, width)
		matrix[i][0] = i * op.DelCost
	}
	for j := 1; j < width; j++ {
		matrix[0][j] = j * op.InsCost
	}

	// Fill in the remaining cells: for each prefix pair, choose the
	// (edit history, operation) pair with the lowest cost.
	for i := 1; i < height; i++ {
		cur := matrix[i%2]
		prev := matrix[(i-1)%2]
		cur[0] = i * op.DelCost
		for j := 1; j < width; j++ {
			delCost := prev[j] + op.DelCost
			matchSubCost := prev[j-1]
			if !op.Matches(source[i-1], target[j-1]) {
				matchSubCost += op.SubCost
			}
			insCost := cur[j-1] + op.InsCost
			cur[j] = min(delCost, min(matchSubCost, insCost))
		}
	}
	return matrix[(height-1)%2][width-1]
}

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}
