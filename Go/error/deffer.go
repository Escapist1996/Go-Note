package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic("Creat Failed")
	}
	defer file.Close() // 函数执行完毕前，关闭文件描述符

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 函数执行完毕前，将缓冲区中的内容刷新到文件中
	for i := 0; i < 100; i++ {
		fmt.Fprintln(writer, i) // 写入缓冲区
	}
}

func main() {
	writeFile("1.txt")
}
