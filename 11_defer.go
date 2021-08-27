package main

import(
	"fmt"
	"error"
)



/**
func  main()  {
	defer func(){
		fmt.Println("first ...")
	}()

	defer func(){
		fmt.Println("second ...")
	}()

	defer func(){
		fmt.Println("third ...")
	}()

	fmt.Println("do first")
}
*/


func AwesomeFunction() (err error){
	defer func(){
		err = new errors.make("error...")
	}

	return nil
}



func main()  {
	AwesomeFunction()
	fmt.Println("hello ")
}