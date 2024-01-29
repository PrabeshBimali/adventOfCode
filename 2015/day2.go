package main

import (
  "slices"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "os"
)

func areaOfSmallSide(arr ...int) int {
  slices.Sort(arr)
  return arr[0] * arr[1]
}

func surfaceArea(l, w, h int) int {
  surfaceArea := (2*l*w) + (2*l*h) + (2*w*h) 
  return surfaceArea
}

func wrapperArea(l, w, h int) int {
  return surfaceArea(l, w, h) + areaOfSmallSide(l, w, h)
}

func smallPerimeter(arr ...int) int {
  slices.Sort(arr)
  return 2 * (arr[0] + arr[1])
}

func volume(l, w, h int) int {
  return l * w * h
}

func ribbonLength(l, w, h int ) int {
  return smallPerimeter(l, w, h) + volume(l, w, h)
}

func readLines(filename string) []string {
  file, err := os.Open(filename)

  if err != nil {
    fmt.Printf("Error when opening file: %s\n", err)
    os.Exit(1)
  }

  defer file.Close()
  lineSlice := []string{}

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lineSlice = append(lineSlice, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }

  return lineSlice
}

func run(filename string) { 
  lineSlice := readLines(filename)
  cuboid := make([]int, 3)
  area := 0
  ribbon := 0

  for _, element := range lineSlice {
    arr := strings.Split(element, "x")
    
    for index, val := range arr {
      i, err := strconv.Atoi(val)

      if err != nil {
        fmt.Printf("Could not convert %s to int: %s", val, err)
      }

      cuboid[index] = i
    }

    area += wrapperArea(cuboid[0], cuboid[1], cuboid[2])
    ribbon += ribbonLength(cuboid[0], cuboid[1], cuboid[2])
  }

  fmt.Printf("Area: %d Perimeter: %d\n", area, ribbon)
}

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("Usage: %s <filename>\n", os.Args[0])
  }else {
    run(os.Args[1])
  }
}
