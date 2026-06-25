package stringutil

func reverse(s string) string {
 t := []rune(s)
 for i, e := 0, len(t)-1; i < len(t)/2; i, e = i+1, e-1 {
  t[i], t[e] = t[e], t[i]
 }
 return string(t)
}