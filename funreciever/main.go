package main

type myfun func()

func (f myfun) set() {

}
func main() {
	var a myfun
	a.set()
}
