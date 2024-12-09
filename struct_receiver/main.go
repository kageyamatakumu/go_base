package main

import (
	"fmt"
	"unsafe"
)

type Task struct {
  Title string
  Estimate int
}

func main() {
  task1 := Task {
    Title: "Learn Golang",
    Estimate: 3,
  }
  task1.Title = "Learning Go"
  fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

  var task2 Task = task1
  task2.Title = "new"
  fmt.Printf("%[1]T %+[1]v %v\n", task2, task2.Title)

  task1p := &Task {
    Title: "Learn concurrency",
    Estimate: 2,
  }
  fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
  task1p.Title = "Changed"
  fmt.Printf("task1p: %+v\n", *task1p)
  var task2p *Task = task1p
  task2p.Title = "Changed by Task2"
  fmt.Printf("task1: %+v\n", *task1p)
  fmt.Printf("task2: %+v\n", *task2p)

}