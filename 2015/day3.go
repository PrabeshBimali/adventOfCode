package main

import (
  "fmt"
  "os"
)

var charMap map[string]point

func init() {
  charMap = make(map[string]point)
  charMap["^"] = point{x: 0, y: 1} 
  charMap[">"] = point{x: 1, y: 0}
  charMap["v"] = point{x: 0, y: -1}
  charMap["<"] = point{x: -1, y: 0}
}

type point struct {
  x int
  y int
}

type path []point

func createPath() path {
  p := path{}
  return append(p, point{x: 0, y: 0,})
}

func (pts path) exists(p point) bool {
  for _, val := range pts {
    if comparePoint(val, p) {
      return true
    }
  }

  return false
}

func (pts path) addUniquePoint(p point) path {
  if !pts.exists(p) {
     pts = append(pts, p)
  }

  return pts
}

func sumPoint(p1 point, p2 point) point {
  newPoint := point{x: 0, y: 0, }
  newPoint.x = p1.x + p2.x
  newPoint.y = p1.y + p2.y

  return newPoint
}

func comparePoint(p1 point, p2 point) bool {
  return p1.x == p2.x && p1.y == p2.y
}


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

func runPart1(filename string) {
  bytes := getBytes(filename)
  santaPath := createPath()
  // get first point from path
  lastPoint := santaPath[0]
  for _, val := range bytes {
    lastPoint = sumPoint(lastPoint, charMap[string(val)])
    santaPath = santaPath.addUniquePoint(lastPoint)
  }

  fmt.Printf("Houses received gifts: %d\n", len(santaPath)) 
}

func runPart2(filename string) {
  bytes := getBytes(filename)
  santaPath := createPath() 
  santaLastPoint := santaPath[0]
  roboLastPoint := santaPath[0]

  for idx, val := range bytes {
    if idx % 2 == 0 {
      santaLastPoint = sumPoint(santaLastPoint, charMap[string(val)])
      santaPath = santaPath.addUniquePoint(santaLastPoint)
    } else {
      roboLastPoint = sumPoint(roboLastPoint, charMap[string(val)])
      santaPath = santaPath.addUniquePoint(roboLastPoint)
    }
  }

  fmt.Printf("Houses received gifts: %d\n", len(santaPath)) // -1 because starting hous is same for both
  // fmt.Println(santaPath)
  // fmt.Println(roboPath)
}

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("Usage: %s <filename>\n", os.Args[0])
  }else {
    //runPart1(os.Args[1])
    runPart2(os.Args[1])
  }
}
