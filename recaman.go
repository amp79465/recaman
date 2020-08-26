package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
  )

var sequenceMembers [][]int = [][]int{[]int{0,0}}

func inSequence(n int) bool {
  for i := 0; i < len(sequenceMembers); i++ {
    if n >= sequenceMembers[i][0] && n <= sequenceMembers[i][1] {
      return true
    }
  }
  return false
}

func addMember(n int) {
  for i := 0; i < len(sequenceMembers); i++ {
    if i == len(sequenceMembers) - 1 {
      if n > sequenceMembers[i][1] + 1 {
        sequenceMembers = append(sequenceMembers, []int{n, n})
        return
      } else if n == sequenceMembers[i][1] + 1 {
        sequenceMembers[i] = sequenceMembers[i][0][n]
        return
      }
      } else if n < sequenceMembers[i-1][0] || n > sequenceMembers[i+1][1] {
        continue
      } else if n < sequenceMembers[i][0] {
        //Add onto left side of current index, do not merge with i - 1
        if n == sequenceMembers[i][0] - 1 && n > sequenceMembers[i-1][1] + 1 {
          sequenceMembers[i] = []int{n, sequenceMembers[i][1]}
          return
        //Merge i and i - 1
        } else if n == sequenceMembers[i][0] - 1 && n == sequenceMembers[i-1][1] + 1 {
          sequenceMembers[i] = []int{sequenceMembers[i-1][0], sequenceMembers[i][1]}
          sequenceMembers = append(sequenceMembers[0:i-1], sequenceMembers[i:])
          return
        //Add onto right side of i - 1
        } else if n == sequenceMembers[i - 1][1] + 1 && n < sequenceMembers[i][0] - 1 {
          sequenceMembers[i-1] = []int{sequenceMembers[i-1][1], n}
          return
        //Insert [n, n] between i - 1 and i
        } else if n < sequenceMembers[i][0] - 1 && n > sequenceMembers[i-1][1] + 1 {
          sequenceMembers = append(append(sequenceMembers[0:i], []int{n, n}), sequenceMembers[i:])
          return
        }
      } else if n > sequenceMembers[i][1] {

      }
    }
  }
}

func recaman(termLimit int) {
  for t := 1; t <= termLimit; t++ {
    sequenceCandidate = lastAdded - t
    if sequenceCandidate > 0 && !(inSequence(sequenceMembers, sequenceCandidate)) {
      addMember(sequenceCandidate)
      lastAdded = sequenceCandidate
    } else {
      addMember(lastAdded + t)
      lastAdded = lastAdded + t
    }
  }
}

func main() {
  timeStart := time.Now()
  termLimitInput := os.Args[1]
  if s, err := strconv.Atoi(termLimitInput); err == nil {
		termLimit = s
	}

  timeElapsed := time.Since(timeStart)
  fmt.Printf("Operation took %s", timeElapsed)
}
