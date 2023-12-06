// Package test
// @Description:
// @Author AN 2023-12-06 23:17:26
package test

import (
	"fmt"
	"testing"
)

func TestApp(t *testing.T) {
	a := 1
	b := "hello, world"
	ap := &a
	bp := &b
	fmt.Println(ap)
	fmt.Println(bp)
	fmt.Printf("ap: %T, bp: %T\n", a, b)
	fmt.Printf("ap: %T, bp: %T\n", ap, bp)
	*ap = 3
	*bp = "4"
	fmt.Println(*ap)
	fmt.Println(*bp)
	fmt.Printf("ap: %T, bp: %T\n", *ap, *bp)

	var ta int
	tb := 2
	sum(&ta, &tb)
	fmt.Println(ta)

}

func sum(a *int, b *int) {
	*a = *b + 1
}
