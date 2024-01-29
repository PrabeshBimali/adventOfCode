package main

import (
  "os"
  "fmt"
)

func getBytes(filePath string) []byte {
  file, err := os.Open(filePath)

  defer file.Close()

  if err != nil {
    fmt.Printf("Error when opening file %s: %s\n", filePath, err)
    os.Exit(1)
  }

  fileInfo, err := file.Stat()

  if err != nil {
    fmt.Printf("Error loading file stat %s: %s\n", filePath, err)
    os.Exit(1)
  }

  size := fileInfo.Size()

  bytes := make([]byte, size)

  _, err = file.Read(bytes)

  if err != nil {
    fmt.Printf("Error when reading %s: %s\n", filePath, err)
    os.Exit(1)
  }

  return bytes
}

func findFloor(bytes []byte) {
  //fmt.Printf("Bytes: %s\n", string(bytes[:len(bytes)])) 
  sum := 0
  for i := 0; i < len(bytes); i++ {
    if string(bytes[i]) == "(" {
      sum += 1
    } else if string(bytes[i]) == ")" {
      sum -= 1
    }
  }

  fmt.Printf("%d\n", sum)
}

func firstCharToBasement(bytes []byte) {
  sum := 0
  for i := 0; i < len(bytes); i++ {
    if string(bytes[i]) == "(" {
      sum += 1
    } else if string(bytes[i]) == ")" {
      sum -= 1

      if sum == -1 {
        fmt.Printf("%d\n", i + 1)
        break
      }
    }
  } 
}

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("Usage: %s <filename>\n", os.Args[0])
  }else {
    bytes := getBytes(os.Args[1])
    firstCharToBasement(bytes)
  }
}
