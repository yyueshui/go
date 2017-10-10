package main

import "fmt"

func main() {
	//修改字符串
	var HelloWorld string= "hello world"
	ByteHelloWorld := []byte(HelloWorld)
	ByteHelloWorld[1] = 'c'

	HelloWorld = string(ByteHelloWorld)

	fmt.Printf("%s\n", HelloWorld)

	var RawString = `
		hello
			world
	`

	fmt.Printf("%s\n", RawString)
}
