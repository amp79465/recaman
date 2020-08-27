package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
  )

var sequenceMembers [][]int = [][]int{[]int{0,0}}
var lastAdded int = 0
var workerCount int = 1

func searchSequence(seq [][]int, n int, result chan bool) {
  for i := 0; i < len(seq); i++ {
    if n >= seq[i][0] && n <= seq[i][1] {
      result <- true
    }
  }
  result <- false
}

func inSequence(n int) bool {
  divisor := 0
  if len(sequenceMembers) / workerCount >= 1 {
    divisor = workerCount
  } else {
    divisor = len(sequenceMembers)
  }
  resultChannel := make(chan divisor)
  for i := 0; i < divisor; i++ {
    go searchSequence(sequenceMembers[i:i], n, resultChannel)
  }

}

func addMember(n int) {
  for i := 0; i < len(sequenceMembers); i++ {
    if i == len(sequenceMembers) - 1 {
      // Handle the case of us looking at the last member of the list
      if n > sequenceMembers[i][1] + 1 {
        sequenceMembers = append(sequenceMembers, []int{n, n})
        return
       } else if n == sequenceMembers[i][1] + 1 {
        sequenceMembers[i] = []int{sequenceMembers[i][0], n}
        return
       }
      // n is to be added somewhere on the left of i
    } else if n < sequenceMembers[i][0] {
        // Add onto left side of current index, do not merge with i - 1
        if n == sequenceMembers[i][0] - 1 && n > sequenceMembers[i-1][1] + 1 {
          sequenceMembers[i] = []int{n, sequenceMembers[i][1]}
          return
        // Merge i and i - 1
        } else if n == sequenceMembers[i][0] - 1 && n == sequenceMembers[i-1][1] + 1 {
          sequenceMembers[i] = []int{sequenceMembers[i-1][0], sequenceMembers[i][1]}
          sequenceMembers = append(sequenceMembers[0:i-1], sequenceMembers[i:]...)
          return
        // Add onto right side of i - 1
        } else if n == sequenceMembers[i - 1][1] + 1 && n < sequenceMembers[i][0] - 1 {
          sequenceMembers[i-1] = []int{sequenceMembers[i-1][1], n}
          return
        // Insert [n, n] between i - 1 and i
        } else if n < sequenceMembers[i][0] - 1 && n > sequenceMembers[i-1][1] + 1 {
          sequenceMembers = append(sequenceMembers, []int{0,0})
          copy(sequenceMembers[i+1:], sequenceMembers[i:])
          sequenceMembers[i] = []int{n,n}
          return
        }
      // n is to be added somewhere on the right side of i
      } else if n > sequenceMembers[i][1] {
        // Add onto right side of current index, do not merge with i + 1
        if n == sequenceMembers[i][1] + 1 && n < sequenceMembers[i+1][0] - 1 {
          sequenceMembers[i] = []int{sequenceMembers[i][0], n}
          return
        // Merge i and i + 1
       } else if n == sequenceMembers[i][1] + 1 && n == sequenceMembers[i+1][0] - 1 {
         sequenceMembers[i] = []int{sequenceMembers[i][0], sequenceMembers[i+1][1]}
         if i == len(sequenceMembers) - 2 {
           sequenceMembers = sequenceMembers[0:i+1]
           return
         } else {
           sequenceMembers = append(sequenceMembers[0:i+1], sequenceMembers[i+2:]...)
           return
         }
      // Add onto left side of i + 1
      } else if n > sequenceMembers[i][1] + 1 && n == sequenceMembers[i+1][0] - 1 {
        sequenceMembers[i+1] = []int{n, sequenceMembers[i+1][1]}
        return
      // Insert between i and i + 1
      } else if n > sequenceMembers[i][1] + 1 && n < sequenceMembers[i+1][0] - 1 {
        sequenceMembers = append(sequenceMembers, []int{0,0})
        copy(sequenceMembers[i+2:], sequenceMembers[i+1:])
        sequenceMembers[i+1] = []int{n,n}
        return
      }
    }
  }
}

func recaman(termLimit int) {
  for t := 1; t <= termLimit; t++ {
    sequenceCandidate := lastAdded - t
    if sequenceCandidate > 0 && !(inSequence(sequenceCandidate)) {
      addMember(sequenceCandidate)
      lastAdded = sequenceCandidate
    } else {
      addMember(lastAdded + t)
      lastAdded = lastAdded + t
    }
  }
}

func main() {
  tL := 0
  timeStart := time.Now()
  termLimitInput := os.Args[1]
  if s, err := strconv.Atoi(termLimitInput); err == nil {
		tL = s
	}
  recaman(tL)
  timeElapsed := time.Since(timeStart)
  fmt.Println("Sequence object count: ", len(sequenceMembers))
  fmt.Println("Operation took", timeElapsed)
}
