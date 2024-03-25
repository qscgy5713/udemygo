package queue

type Queue struct {
	Data []int
}

// 将元素添加到队列的末尾
func (q *Queue) Push(x int) {
	q.Data = append(q.Data, x)
}

// 从队列的头部移除并返回元素。
func (q *Queue) Pop() int {
	if len(q.Data) == 0 {
		return -1
	}
	val := q.Data[0]
	q.Data = q.Data[1:]
	return val

}

// 返回队列的头部元素，但不将其从队列中移除。
func (q *Queue) Peek() int {
	if len(q.Data) == 0 {
		return -1
	}
	return q.Data[0]
}

func (q *Queue) Size() int {
	return len(q.Data)
}

// 检查队列中是否没有元素
func (q *Queue) Empty() bool {
	return len(q.Data) == 0
}
