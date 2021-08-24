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

H :=f(15, 12, 4)
fmt.Println("result", H)
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
func strings(){
a := "haba"
b := "blackbonez"
c := "fendi"

d := a + b + c
fmt. Println(d)
}

func f(a int, b int, c int) (int) {
	a = 15
	b = 12
	c = 4


	d := a + b + c
	return d
}