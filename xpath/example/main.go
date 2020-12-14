package main

import (
	"fmt"
	"github.com/liyue201/golib/xpath"
)

func main()  {
	fmt.Println("xpath:", xpath.GetExeDirPath())
	fmt.Println("AbsPath:", xpath.AbsPath("./fd/es"))
}


