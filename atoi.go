package main

import (
  "fmt"
  "math"
)

func myAtoi(s string) int {
    if len(s) == 0 {
      return 0
    }

    var sign int = 1
    var i int = 0
    var res int = 0

    for i < len(s) && s[i] == ' ' {
      i++
    }
    
    if i == len(s) {
      return 0
    }

    if s[i] == '-' {
      sign = -1
      i++
    } else if s[i] == '+' {
      i++
    }
    
    for i < len(s) && isDigit(s[i]) {
        res = res * 10 + int(s[i] - '0')

        if res * sign > math.MaxInt32 {
          return math.MaxInt32
        } else if res * sign < math.MinInt32 {
          return math.MinInt32
        }
        i++
    }

    return sign * res 
  }

    
func isDigit(ch byte) bool {
    if ch  >= '0' && ch <= '9' {
      return true
    }
    return false
}

func main() {
    test := "  025Bla0000"
    fmt.Println(myAtoi(test))

}
