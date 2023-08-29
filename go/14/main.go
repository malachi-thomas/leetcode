package main


func longestCommonPrefix(strs []string) string {
  output := ""
  for i := 0; i < len(strs); i++ {
    println(strs[i])
  }

  return output
}


func main() {
  // strs := []string{ "flower", "flow", "flight" }
  longestCommonPrefix([]string{ "flower", "flow", "flight" }) 
}
