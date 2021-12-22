package main

import "fmt"

//"pactice/routines"
/*return_data "practice/returnData"
struct_practice "practice/struct"*/

/*func YeniStructDeniyoz() *struct_practice.Sefa {
	var sefaa struct_practice.Sefa

	result := sefaa.Fetch(10)

	return result
}*/

func main() {

	/*cevap := return_data.SeeResponse(10)
	print(cevap)

	sonuc := YeniStructDeniyoz()

	print(sonuc)*/
	//routines.Routines()

	goroutine1 := make(chan string, 5)
	goroutine1 <- "Australia"
	data := <-goroutine1
	fmt.Printf("%s", data)
}
