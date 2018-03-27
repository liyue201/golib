package main

import (
	"fmt"
	"github.com/liyue201/golib/path"
)

func main()  {
	fmt.Println("path:", path.GetExeDirPath())
	fmt.Println("AbsPath:", path.AbsPath("./fd/es"))
}


