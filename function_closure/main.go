package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
  defer fmt.Println("main func final-finish")
  defer fmt.Println("main func semi-finish")
  fmt.Println("Hello World")
}

func trimExtension(files ...string) []string {
  out := make([]string, 0, len(files))
  for _, v := range files {
    out = append(out, strings.TrimSuffix(v, ".csv"))
  }
  return out
}

func fileChecker(name string) (string, error) {
  f, err := os.Open(name)
  if err != nil {
    return "", errors.New("file not fount")
  }
  defer f.Close()
  return name, nil
}

func addExt(f func(file string) string, name string) {
  fmt.Println(f(name))
}

func multiply() func(int) int {
  return func(n int) int {
    return n * 1000
  }
}

func countUP() func(int) int {
  count := 0
  return func(n int) int {
    count += n
    return count
  }
}

func main() {
  funcDefer()
  files := []string{"file.csv", "file2.csv", "file.csv"}
  fmt.Println(trimExtension(files...))
  name, err := fileChecker("file.txt")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(name)

  i := 1
  func(i int) {
    fmt.Println(i)
  }(i)

  f1 := func(i int) int {
    return i + 1
  }
  fmt.Println(f1(3))

  f2 := func(file string) string {
    return file + ".csv"
  }
  addExt(f2, "file1")

  f3 := multiply()
  fmt.Println(f3(2))

  f4 := countUP()
  for i := 1; i <= 5; i++ {
    v := f4(2)
    fmt.Printf("%v\n", v)
  }
}