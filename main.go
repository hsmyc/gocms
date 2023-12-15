package main

func Osso(a func(int) *int) {
	println(a(4))
}
func Messo(b int) *int {
	c := b
	return &c
}
func main() {
	a := 4
	Osso(Messo)
	println("x: ", &a)
}
