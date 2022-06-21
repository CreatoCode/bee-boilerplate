package json

// very fast json lib

import (
	"fmt"

	"github.com/bytedance/sonic"
)

type stu struct {
	Name string
}

func test() {
	var data stu
	// Marshal
	output, err := sonic.Marshal(&data)
	// Unmarshal
	// err := sonic.Unmarshal(output, &data)
	if nil != err {
		fmt.Printf("error:%v", err)
	} else {
		fmt.Print(output)
	}
}
