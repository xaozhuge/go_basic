package main

import (
	"fmt"

	"rsc.io/quote"

	"example/calc"
)

func main() {
	fmt.Println(quote.Hello())  // Ahoy, world!
	fmt.Println(calc.Add(10, 3))
}

/*
docker exec go_c sh -c "cd modules;go run ."
*/


/*
1. 在当前文件夹下生成了go.mod，这个文件记录当前模块的模块名以及所有依赖包的版本。


*/