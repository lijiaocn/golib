//create: 2015/05/14 15:55:20 Change: 2019/05/05 18:58:54 author:lijiao
package v2

import (
	"fmt"
)

var (
	VERSION string
	COMPILE string
)

func Show() {
	fmt.Printf("version: %s   compile at: %s( just for test v1)\n", VERSION, COMPILE)
}
