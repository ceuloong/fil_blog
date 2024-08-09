package monitor

import (
	"bufio"
	"fmt"
	"os"
)

func WriteToFile(filename string, content string) int {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	data := []byte(content)
	rel, err := writer.Write(data)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	writer.Flush() // 刷新缓冲区，将数据写入文件
	fmt.Println("Data written successfully.")
	return rel
}

// ReadLines 使用 bufio.Scanner 读取文件的每一行
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
