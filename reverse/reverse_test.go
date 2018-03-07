package reverse

import "testing"

func TestReverseLine(t *testing.T) {
  input := "abcdefghijklmnopqrstuvwxyz"
  reversed := ReverseLine(input)
  if reversed != "zyxwvutsrqponmlkjihgfedcba" {
    t.Error("Expected zyxwvutsrqponmlkjihgfedcba, got ", reversed)
  }
}

func TestReverseSlice(t *testing.T) {
  s := make([]string, 3)
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"

  r := make([]string, 3)
  r[0] = "c"
  r[1] = "b"
  r[2] = "a"

  reversed := ReverseSlice(s)
  for i := range s {
    if r[i] != reversed[i] {
      t.Error("Expected %v, got ", r[i], reversed)
    }
  }
}
