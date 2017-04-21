package stringutil

import "testing"

func Test(t *testing.T){
    var tests = []struct{
      input, expected string
    }{
    {"Backward", "drawkcaB"},
    {"Hello, 世界", "界世 ,olleH"},
    {"", ""},
  }
  for _, s := range tests {
    got := Reverse(s.input)
    if got != s.expected {
      t.Errorf("Reverse(%q) == %q, expected %q", s.input, got, s.expected)
    }
  }
}
