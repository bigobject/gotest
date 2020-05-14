package hello

//#include <hello.h>
import "C"

func SayHello(words string) int {
	return int(C.SayHello(C.CString(words)))
}
