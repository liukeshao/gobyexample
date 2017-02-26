package main

import "fmt"

// Persion struct
type Persion struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(Persion{"Bob", 20})
	fmt.Println(Persion{Name: "Alice", Age: 30})
	fmt.Println(Persion{Name: "Fred"})
	fmt.Println(&Persion{Name: "Ann", Age: 40})
	s := Persion{Name: "Sean", Age: 50}
	fmt.Println(s.Name)
	sp := &s
	fmt.Println(sp.Age)
	sp.Age = 51
	fmt.Println(sp.Age)
}
