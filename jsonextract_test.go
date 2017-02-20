package jsonextract

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	jsondata := `{"kind":"List","apiVersion":"v1","metadata":{"selfLink":"/api/v1","resourceVersion":"3811716"},"items":[{"name":"test1","image":"meitu"},{"name":"abcd","image":"meitu2"}]}`

	t1 := MyJson(jsondata)
	fmt.Println(t1) //&{map[metadata:map[resourceVersion:3811716 selfLink:/api/v1] items:[map[name:test1 image:meitu] map[name:abcd image:meitu2]] kind:List apiVersion:v1]}

	t2 := MyJson(jsondata).Getkey("kind")
	fmt.Println(t2) //&{List}

	t3 := MyJson(jsondata).Getkey("kind").ToString()
	fmt.Println(t3) //List

	t4 := MyJson(jsondata).Getkey("items").GetArrylen()
	fmt.Println(t4) //2

	t5 := MyJson(jsondata).Getkey("metadata")
	fmt.Println(t5) //&{map[selfLink:/api/v1 resourceVersion:3811716]}

	t6 := MyJson(jsondata).Getkey("metadata").GetData()
	fmt.Println(t6) //map[selfLink:/api/v1 resourceVersion:3811716]

	t7 := MyJson(jsondata).Getkey("metadata").Getkey("selfLink")
	fmt.Println(t7) //&{/api/v1}

	t8 := MyJson(jsondata).Getkey("metadata").Getkey("selfLink").ToString()
	fmt.Println(t8) ///api/v1
}
