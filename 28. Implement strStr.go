// Implement strStr().

// Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.


func strStr(haystack string, needle string) int {
  if len(needle) > len(haystack) {
    return -1;
  }
  if needle == "" {
    return 0;
  }
  var j int;
  for i := 0; i < len(haystack) - len(needle) + 1; i++ {
    for j = 0; j < len(needle); j++ {
      if haystack[i + j] != needle[j] {
          break;
      }
    }
    if j == len(needle) {
      return i;
    }
  }
  return  -1;
}