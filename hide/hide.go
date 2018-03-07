package hide

import (
  "bufio"
  "fmt"
  "github.com/JaveLLC/napkin/reverse"
  "log"
  "os"
)

// HideFile hides your files, sort of, not really, shhhh
func HideFile(inputFile string, outputFile string) error {
  log.Printf("Reading the origin file...")
  reversedLines, err := readLinesReverse(inputFile)
  if err != nil {
    return err
  }

  log.Printf("Writing the destination file...")
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

  log.Printf("Hiding...")
  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    reversedText := reverse.ReverseLine(scanner.Text())

    lines = append(lines, reversedText)
  }

  log.Printf("Further hiding...")
  lines = reverse.ReverseSlice(lines)

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
