package main

import (
    "fmt"
    "log"

    "github.com/rainemng/mcp-server-test/pkg/calculator"
)

func main() {
    // 创建服务器和客户端实例
    server := calculator.NewCalculatorServer()
    client := calculator.NewCalculatorClient(server)

    // 测试加法
    result, err := client.ProcessRequest("add", 5, 3)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("5 + 3 = %d\n", result)

    // 测试减法
    result, err = client.ProcessRequest("subtract", 10, 4)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("10 - 4 = %d\n", result)
}
