package main

import (
	"strconv"
	"strings"
)

func getLucky(s string, k int) int {
	var t strings.Builder
	for _, c := range s {
		t.WriteString(strconv.Itoa(int(c - 'a' + 1)))
	}
	s = t.String()
	println(s)
	res, _ := strconv.Atoi(s)
	println(res)
	s = strconv.Itoa(res)
	tt := 0
	for _, c := range s {
		tt += int(c - '0')
	}
	// for ; k > 0; k--{
	//     cnt := 0
	//     for res != 0{
	//         cnt += res % 10
	//         res /= 10
	//     }
	//     res = cnt
	// }
	return tt
}

func main() {
	tt := getLucky("fleyctuuajsr", 1)
	println(tt)
}

// func getLucky(s string, k int) int {
// 	var t strings.Builder
// 	for _, c := range s {
// 		t.WriteString(strconv.Itoa(int(c - 'a' + 1)))
// 	}
// 	s = t.String()
// 	for k > 0 {
// 		k--
// 		t := 0
// 		for _, c := range s {
// 			t += int(c - '0')
// 		}
// 		s = strconv.Itoa(t)
// 	}
// 	ans, _ := strconv.Atoi(s)
// 	return ans
// }
