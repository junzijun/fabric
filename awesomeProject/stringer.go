package main
import "fmt"

type Person struct{
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years),p.Name,p.age")
}
func main(){
	a := Person{"aaaa",21}
	z := Person{"zzzz",25}
	fmt.Println(a,z)
}