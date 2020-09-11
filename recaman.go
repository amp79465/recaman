package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
  "sync"
  "math"
  )

var sequenceMembers [][]int = [][]int{[]int{0,0}}
var lastAdded int = 0
var workerCount int = 2

// Takes a slice of the sequence and unpacks it, searching for the int n
func searchSequence(seq [][]int, n int, result chan bool, wg *sync.WaitGroup) {
  defer wg.Done()
  for i := 0; i < len(seq); i++ {
    if n >= seq[i][0] && n <= seq[i][1] {
      result <- true
      return
    }
  }
  result <- false
  return
}

func monitorWorker(wg *sync.WaitGroup, resultChannel chan bool) {
  wg.Wait()
  close(resultChannel)
}

func inSequence(n int) bool {
  //Set up the divisor to split up the sequence appropriately
  divisor := 0
  wg := &sync.WaitGroup{}
  sliceLength := int(math.Floor(float64(len(sequenceMembers)) / float64(workerCount))) + 1
  if sliceLength >= 1 {
    divisor = workerCount
  } else {
    divisor = len(sequenceMembers)
  }
  wg.Add(divisor)
  //Deploy the goroutines to search the function
  resultChannel := make(chan bool)

  fmt.Println("Sequence: ", sequenceMembers)

  for i := 0; i < divisor; i++ {
    //Handle if we are on the last member of the sequence object and capture the remainder of the slice
    if i == divisor - 1 {
      go searchSequence(sequenceMembers[i*divisor:], n, resultChannel, wg)
      fmt.Println(sequenceMembers[i*divisor:])
    } else {
      go searchSequence(sequenceMembers[i*divisor:i*divisor+sliceLength], n, resultChannel, wg)
      fmt.Println(sequenceMembers[i*divisor:i*divisor+sliceLength])
    }
  }

  go monitorWorker(wg, resultChannel)

  // Get the results of the goroutines and return true if any are true
  for result := range resultChannel {
    if result == true {
      return true
    }
  }
  return false
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

// The sequence generating function
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
    fmt.Println(t, lastAdded)
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
