package main

import (
  "strings"
  "fmt"
)
// str1 / str2
func gcdOfStrings(str1 string, str2 string) string {
  // arr := strings.Split(str1, str2)
  arr := strings.Replace(str2, "AB", "", -1)
  fmt.Println(arr)
  if len(arr) == 0 { return "" }
  // arr2 := strings.Split(str1, str2)
  // fmt.Println(len(arr1[0]), arr2[0])
  // println(strings.Split(str1, "y"))
  // arr := strings.Split(str1, " ")
  // fmt.Println(arr[0], arr[1])
  // fmt.Println(strings.Split(str1, "y")[0])

  // Println(strings.Split(str1, "y")[0])
  return "hi"
}

func main() {
    println(gcdOfStrings("ABABAB", "ABAB"))
  }


