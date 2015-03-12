// Project Euler problem 1
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
    "fmt"
    "math"
  )

func multiples () int {
  sum := 3

  for i := 5; i < 1000; i++ {
    i := float64(i)

    if math.Mod(i,3) == 0 || math.Mod(i,5) == 0 {
      i := int(i)
      
      sum += i
    }
  }

  return sum
}

func main() {
  fmt.Println(multiples())
}
