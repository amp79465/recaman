package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
)

func main() {
  timeStart := time.Now()
  rMembers := []int{0}
  lastAdded, termLimit, sequenceCandidate := 0, 0, 0

  termLimitInput := os.Args[1]
  if s, err := strconv.Atoi(termLimitInput); err == nil {
		termLimit = s
	}

  for t := 1; t <= termLimit; t++ {
    sequenceCandidate = lastAdded - t
    if sequenceCandidate > 0 && !(inSequence(rMembers, sequenceCandidate)) {
      rMembers = append(rMembers, sequenceCandidate)
      lastAdded = sequenceCandidate
    } else {
      rMembers = append(rMembers, lastAdded + t)
      lastAdded = lastAdded + t
    }
  }
  timeElapsed := time.Since(timeStart)
  //fmt.Println(rMembers)
  fmt.Printf("Operation took %s", timeElapsed)
}

func inSequence(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
