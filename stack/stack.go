package stack

type Stack struct {
	Data []int
}

// 后进先出（LIFO）的栈
// 將元素x推入棧頂
func (s *Stack) Push(x int) {
	s.Data = append(s.Data, x)
}

// 移除棧頂元素並返回他
func (s *Stack) Pop() int {
	if len(s.Data) == 0 {
		return -1
	}
	val := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1] // 堆疊 MyStack 中的最後一個元素彈出（移除)
	return val
}

// 返回棧頂元素
func (s *Stack) Top() int {
	if len(s.Data) == 0 {
		return -1
	}

	return s.Data[len(s.Data)-1]
}

// 檢查棧是否為空
func (s *Stack) Empty() bool {
	return len(s.Data) == 0
}
