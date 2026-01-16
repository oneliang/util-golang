package common

// Queue 性能优化版泛型队列
type Queue[T any] struct {
	items []T // 存储元素的切片
	head  int // 队首索引（指向下一个要出队的元素）
}

// NewQueue 创建新队列（可选预分配容量）
func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, capacity),
		head:  0,
	}
}

// Enqueue 入队操作：添加元素到队尾
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue 出队操作：从队首移除并返回元素
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T // 泛型类型的零值

	// 队列为空
	if q.head >= len(q.items) {
		return zero, false
	}

	// 获取队首元素
	item := q.items[q.head]
	q.head++ // 移动队首指针

	// 定期清理已出队的元素，防止内存泄漏
	// 当已出队元素超过一半时，进行清理
	if q.head > len(q.items)/2 && len(q.items) > 0 {
		q.items = q.items[q.head:]
		q.head = 0

		// 可选：缩小底层数组容量（当实际使用不到容量一半时）
		if cap(q.items) > 2*len(q.items) && len(q.items) > 0 {
			newItems := make([]T, len(q.items))
			copy(newItems, q.items)
			q.items = newItems
		}
	}

	return item, true
}

// Front 查看队首元素（不移除）
func (q *Queue[T]) Front() (T, bool) {
	var zero T
	if q.head >= len(q.items) {
		return zero, false
	}
	return q.items[q.head], true
}

// Back 查看队尾元素
func (q *Queue[T]) Back() (T, bool) {
	var zero T
	if q.head >= len(q.items) {
		return zero, false
	}
	return q.items[len(q.items)-1], true
}

// IsEmpty 检查队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return q.head >= len(q.items)
}

// Size 获取队列长度（有效元素个数）
func (q *Queue[T]) Size() int {
	return len(q.items) - q.head
}

// Capacity 获取队列容量
func (q *Queue[T]) Capacity() int {
	return cap(q.items)
}

// Clear 清空队列（重置状态）
func (q *Queue[T]) Clear() {
	q.items = q.items[:0]
	q.head = 0
}

// EnqueueBatch 批量入队
func (q *Queue[T]) EnqueueBatch(items ...T) {
	q.items = append(q.items, items...)
}

// ForEach 遍历队列（不修改队列状态）
func (q *Queue[T]) ForEach(f func(T)) {
	for i := q.head; i < len(q.items); i++ {
		f(q.items[i])
	}
}

// ToSlice 转换为切片（复制一份数据）
func (q *Queue[T]) ToSlice() []T {
	if q.head >= len(q.items) {
		return []T{}
	}
	result := make([]T, q.Size())
	copy(result, q.items[q.head:])
	return result
}
