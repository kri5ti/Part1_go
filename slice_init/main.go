package main
import "fmt"

func main() {
	a:= []string{"Taras","Andrew","Roman"}
	names:= a[:]
	fmt.Println(names)
}
