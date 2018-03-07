package reverse

// ReverseLine reverses a string line
func ReverseLine(input string) string {
  n := 0
  rune := make([]rune, len(input))
  for _, r := range input {
    rune[n] = r
    n++
  }
  rune = rune[0:n]

  for i := 0; i < n/2; i++ {
    rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
  }
  reversedText := string(rune)
  return reversedText
}

// ReverseSlice reverses a slice of strings (bottom to top)
func ReverseSlice(input []string) []string {
  last := len(input) - 1
  for i := 0; i < len(input)/2; i++ {
    input[i], input[last-i] = input[last-i], input[i]
  }

  return input
}
