package main 
import "fmt" 
func jianfa (n1 int, n2 int) int {
	var a int = 0
	a += n1
	a -= n2
	return a
}
func main() {
	a := jianfa(-45,31)
	fmt.Println(a)

}
