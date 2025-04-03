package calculator

import (
    "fmt"
)

// CalculatorClient 处理计算请求
type CalculatorClient struct {
    server *CalculatorServer
}

// NewCalculatorClient 创建一个新的计算器客户端
func NewCalculatorClient(server *CalculatorServer) *CalculatorClient {
    return &CalculatorClient{server: server}
}

// ProcessRequest 处理计算请求
func (c *CalculatorClient) ProcessRequest(operation string, a, b int) (int, error) {
    switch operation {
    case "add":
        return c.server.Add(a, b), nil
    case "subtract":
        return c.server.Subtract(a, b), nil
    default:
        return 0, fmt.Errorf("unsupported operation: %s", operation)
    }
}
