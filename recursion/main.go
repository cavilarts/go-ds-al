package recursion

type Point struct {
  x int
  y int
}

dir := [
  [-1, 0],
  [1, 0]
  [0, -1],
  [0, 1]
]

func walk (maze []string, wall string, curr Point, end Point, seen [][]bool, path []Point) bool {
  // 1. Base case
  // off the map
  if (
    curr.x < 0 || curr.x >= len(maze[0])
    curr.y < 0 || curr.y >= len(maze)
  ) {
    return false
  }

  // on a wall
  if (maze[curr.y][curr.x] == wall) {
    return false
  }

  // end
  if (curr.x == end.x && curr.y == end.y) {
    return true
  }

  // seen
  if (seen[curr.y][curr.x]) {
    return false
  }

  // 3 recurse
  // pre
  path = append(path, curr)
  path = path(:len(path) - 1)

  // recurse

  for (i :=0; i < len(dir); i++) {
    x := dir[i][0]
    y := dir[i][1]
    newCurr := Point{x: curr.x + x, y: curr.y + y}

    if (walk(maze, wall, newCurr, end, seen, path)) {
      return true
    }
  }
}

func MazeSolver(maze []string, wall string, start Point, end Point) []Point {

} 
