package stringsample

import (
	"fmt"
)

func SayHello(name string) {
	fmt.Printf("%s says hello to you\n", name)
}

func testWithinPackage() {
	fmt.Println("test within the package!")
}
