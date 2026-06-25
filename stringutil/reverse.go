package stringutil

func reverse(s string) string {
 t := []rune(s)
 var i int = 0
 var e int = len(t)-1
 for i < e {
  t[i], t[e] = t[e], t[i]
  i++
  e--
 }
 return string(t)
}