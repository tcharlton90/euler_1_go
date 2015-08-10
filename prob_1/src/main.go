package main

import "fmt"

func sum(x, y []int) int {
  // Sums two slices
  var ans int
  for _, n := range x {
    ans += n
    fmt.Println("%s ", ans)
  }
  fmt.Println("---------")
  for _, n := range y {
    ans += n
    fmt.Println("%s ", ans)
  }
  return ans
}

func getMultiples(multiplesOf, limit int) (ans []int){
  // Gets multiples of a number upto a lmit
  s := []int{}
  empty := []int{}
  for i := 0; i < limit; i++ {
    if i % multiplesOf == 0 {
      s = append(s, i)
    }
  }
  if len(s) > 0 {
    return s
  } else {
    return empty
  }
}

func runCalc(limit int) (ans int){
  // Calculates the sum of all multiples of 3 & 5 under limit
  threes := getMultiples(3, limit)
  fives  := getMultiples(5, limit)
  return sum(threes, fives)
}

func main(){
  ans := runCalc(1000)
  fmt.Println("Answer is %d\n", ans)
}
