package main

import "fmt"


func asciiReverseString(s []byte)  {

	l := len(s)
	for i := range s {
		if l/2 <= i {
			break
		}
		s[i],s[l-i-1] = s[l-i-1],s[i]
	}

	fmt.Println(string(s))
}

func testReverseInt(s []int)  {

	l := len(s)
	for i := range s {
		if l/2 <= i {
			break
		}
		fmt.Println(i,l-i-1)
		s[i],s[l-i-1] = s[l-i-1],s[i]
	}

	fmt.Println((s))

}

func main() {
	asciiReverseString([]byte("hello"))
	testReverseInt([]int{1,2,3,4})
}