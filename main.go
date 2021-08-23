package main

import "fmt"

func main() {
	fmt.Println("hello world")
goland()
hahaha()
tmich("tobi", 26 )
b, t :=tmich("tobi" ,26)

fmt.Println("result:", b , t)


T :=mich(30)
fmt.Println(T)
}
func goland() {
	fmt.Println("my name is tobi" )
}
func hahaha() {
	fmt.Println("47")


}


func tmich(name string, age int) (string, int) {
	fmt. Println("name is", name)
	fmt. Println("age is" , age)

	age = age + 10

	return name,age

}
func mich(mikes int) (int) {

	mikes = mikes * 2

	return mikes
}