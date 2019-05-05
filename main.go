package main

import (
    "github.com/bryanbrady/godoku/godoku1"
    "github.com/bryanbrady/godoku/godoku2"
    "github.com/bryanbrady/godoku/godoku3"
    "bufio"
    "flag"
    "log"
    "os"
    "runtime/pprof"
)

const gridSize = 9

type Cell struct {
  Row int
  Col int
}

var filePtr    = flag.String("file", "", "puzzle file, 1 puzzle per line")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var solver     = flag.Int("solver", 1, "1, 2, 3")
var solve = godoku1.Solve
func main() {
  flag.Parse()
  switch *solver {
  case 1:
    solve = godoku1.Solve
  case 2:
    solve = godoku2.Solve
  case 3:
    solve = godoku3.Solve
  }
  if *cpuprofile != "" {
    f, err := os.Create(*cpuprofile)
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
  }
  if *filePtr != "" {
    file, err := os.Open(*filePtr)
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
      solve(scanner.Text())
    }
  } else {
    for _, e := range flag.Args() {
      solve(e)
    }
  }
}
