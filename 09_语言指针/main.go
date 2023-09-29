package main

import "fmt"

func fun1() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)
}

func fun2() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

func main() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)
}
