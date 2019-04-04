package main

import "fmt"
import "github.com/ehteshamz/greetings"

//ConcatSlice appends string
func ConcatSlice(sliceToConcat []byte) string {
	var strToReturn string
	for i := range sliceToConcat {
		if i != len(sliceToConcat)-1 {
			strToReturn += string(sliceToConcat[i]) + "-"
		} else {
			strToReturn += string(sliceToConcat[i])
		}
	}
	return strToReturn
}

//Encrypt encrypts string
func Encrypt(sliceToEncrypt []byte, ceaserCount int) []byte {
	for i := range sliceToEncrypt {
		sliceToEncrypt[i] += byte(ceaserCount)
	}
	return sliceToEncrypt
}

//EZGreetings fetches instructor package
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

func main() {

	sli := []byte{'1', '2', '3'}
	sli2 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(ConcatSlice(sli))
	fmt.Println(ConcatSlice(Encrypt(sli2, len(sli2))))
	fmt.Println(EZGreetings("faateh"))
}
