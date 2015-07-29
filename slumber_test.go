package slumber

import "fmt"

func ExampleNew() {
	slumber := New("http://example.org/")
	fmt.Println(slumber.Base)
	// Output: http://example.org/
}

func ExampleChild() {
	slumber := New("http://example.org/")
	fmt.Println(slumber.Child("test").Base)
	// Output: http://example.org/test
}

func ExampleChild_deeper() {
	slumber := New("http://example.org/")
	fmt.Println(slumber.Child("test").Child("42").Child("things").Base)
	// Output: http://example.org/test/42/things
}
