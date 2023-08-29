package main

func mergeAlternately(word1 string, word2 string) string {
  word3 := ""
  for i := 0; len(word3) != len(word1) + len(word2); i++ {
    if len(word1) > i { word3 += string(word1[i]) }
    if len(word2) > i { word3 += string(word2[i]) }
  }
  return word3
}
func main() {
  println(mergeAlternately("hatsx", "cowxzz"))
}
