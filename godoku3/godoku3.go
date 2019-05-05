package godoku3

import (
    "fmt"
    "strconv"
)

const gridSize = 9

type Cell struct {
  Row int
  Col int
}

func Solve(s string) {
  grid := parseGrid(s)
  fmt.Printf("Original: %s\n", s)
  solveBruteForce(emptyCells(&grid), &grid)
  fmt.Printf("Solution: ")
  printString(grid)
}

func parseGrid(s string) [gridSize][gridSize]int {
  grid := [gridSize][gridSize]int{}
  for r := 0; r < gridSize; r++ {
    for c := 0; c < gridSize; c++ {
      i, _ := strconv.Atoi(string(s[c+r*gridSize]))
      grid[r][c] = i
    }
  }
  return grid
}

func printString(grid [gridSize][gridSize]int) {
  for r := 0; r < gridSize; r++ {
    for c := 0; c < gridSize; c++ {
      fmt.Printf("%d",grid[r][c])
    }
  }
  fmt.Println("")
}

func printGrid(grid [gridSize][gridSize]int) {
  for r := 0; r < gridSize; r++ {
    for c := 0; c < gridSize; c++ {
      g := grid[r][c]
      if g == 0 {
        fmt.Printf("-")
      } else {
        fmt.Printf("%d", g)
      }
    }
    fmt.Println("")
  }
  fmt.Println("---------")
}

func solveBruteForce(cells []Cell, grid *[gridSize][gridSize]int) bool {
  if !hasEmpty(grid) && isValid(grid) {
    return true
  }

  c := cells[0]
  for v := 1; v <= 9; v++ {
    grid[c.Row][c.Col] = v
    if isValid(grid) {
      if solveBruteForce(cells[1:], grid) {
        return true
      }
      grid[c.Row][c.Col] = 0
    } else {
      grid[c.Row][c.Col] = 0
    }
  }
  return false
}

func hasEmpty(grid *[gridSize][gridSize]int) bool {
  for r := 0; r < gridSize; r++ {
    for c := 0; c < gridSize; c++ {
      if grid[r][c] == 0 {
        return true
      }
    }
  }
  return false
}

func emptyCells(grid *[gridSize][gridSize]int) []Cell {
  var cells []Cell
  for r := 0; r < gridSize; r++ {
    for c := 0; c < gridSize; c++ {
      if grid[r][c] == 0 {
        cells = append(cells, Cell{r,c})
      }
    }
  }
  return cells
}

func isValid(grid *[gridSize][gridSize]int) bool {
  for i := 0; i < gridSize; i++ {
    switch {
    case !isRowValid(i, grid):
      return false
    case !isColValid(i, grid):
      return false
    case !isSubGridValid(i, grid):
      return false
    }
  }
  return true
}

func isRowValid(row int, grid *[gridSize][gridSize]int) bool {
  used := [10]bool{}
  var g = 0
  for c := 0; c < gridSize; c++ {
    g = grid[row][c]
    if g !=0 && used[g] {
      return false
    }
    used[g] = true
  }
  return true
}

func isColValid(col int, grid *[gridSize][gridSize]int) bool {
  used := [10]bool{}
  for r := 0; r < gridSize; r++ {
    g := grid[r][col]
    if g !=0 && used[g] {
      return false
    }
    used[g] = true
  }
  return true
}

// Subgrid numbering
// +-----+-----+-----+
// |     |     |     |
// |  0  |  1  |  2  |
// |     |     |     |
// +-----+-----+-----+
// |     |     |     |
// |  3  |  4  |  5  |
// |     |     |     |
// +-----+-----+-----+
// |     |     |     |
// |  6  |  7  |  8  |
// |     |     |     |
// +-----+-----+-----+

func isSubGridValid(sub int, grid *[gridSize][gridSize]int) bool {
  used := [10]bool{}
  r_offset := sub / 3
  c_offset := sub % 3
  for r := 0; r < 3; r++ {
    for c := 0; c < 3; c++ {
      g := grid[r+3*r_offset][c+3*c_offset]
      if g !=0 && used[g] {
        return false
      }
      used[g] = true
    }
  }
  return true
}
