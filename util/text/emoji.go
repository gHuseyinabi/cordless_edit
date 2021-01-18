package text

// FindEmojiIndices will find all parts of a string (rune array) that
// could potentially be emoji sequences and will return them backwards.
// they'll be returned backwards in order to allow easily manipulating the
// data without invalidating the following indexes accidentally.
// Example:
//     Hello :world:, what a :nice: day.
// would result in
//     []int{22,27,6,12}
func FindEmojiIndices(runes []rune) []int {
	var sequencesBackwards []int
	return sequencesBackwards
}
