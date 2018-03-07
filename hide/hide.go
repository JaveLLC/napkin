package hide

import (
  "bufio"
  "fmt"
  "os"
)

// HideFile hides your files, sort of, not really, shhhh
func HideFile(inputFile string, outputFile string) error {
  reversedLines, err := readLinesReverse(inputFile)
  if err != nil {
    return err
  }

  err = writeOutputFile(reversedLines, outputFile)
  if err != nil {
    return err
  }

  return nil
}

func readLinesReverse(inputFile string) ([]string, error) {
  file, err := os.Open(inputFile)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    n := 0
    rune := make([]rune, len(scanner.Text()))
    for _, r := range scanner.Text() {
      rune[n] = r
      n++
    }
    rune = rune[0:n]

    for i := 0; i < n/2; i++ {
      rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
    }
    reversedText := string(rune)

    lines = append(lines, reversedText)
  }

  last := len(lines) - 1
  for i := 0; i < len(lines)/2; i++ {
    lines[i], lines[last-i] = lines[last-i], lines[i]
  }

  return lines, scanner.Err()
}

func writeOutputFile(lines []string, outputFile string) error {
  file, err := os.Create(outputFile)
  if err != nil {
    return err
  }

  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}
