package stringutil

import (
   "testing"
   "fmt"
)

func TestReverse(t *testing.T) {
   cases := []struct {
      in, want string
   }{
      {"Hello, world", "dlrow ,olleH"},
      {"Hello, 世界", "界世 ,olleH"},
      {"", ""},
   }
   for i, c := range cases {
      got := Reverse(c.in)
      fmt.Printf("[%03d]: Reverse(%q) == %q, want %q", i, c.in, got, c.want)
      if got != c.want {
         t.Errorf("\tFailed!\n")
      }else{
         fmt.Printf("\tSuccess!\n")
      }
   }
}
