package calculator

// CalculatorServer 提供基础的计算功能
type CalculatorServer struct {}

// Add 执行加法运算
func (s *CalculatorServer) Add(a, b int) int {
    return a + b
}

// Subtract 执行减法运算
func (s *CalculatorServer) Subtract(a, b int) int {
    return a - b
}

// NewCalculatorServer 创建一个新的计算器服务
func NewCalculatorServer() *CalculatorServer {
    return &CalculatorServer{}
}
