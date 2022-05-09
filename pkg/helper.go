package helper

type Numeric interface {
	int | int32 | int64 | float32 | float64
}

func ChMin[T Numeric](a *T, b T) {
	if *a > b {
		*a = b
	}
}

func ChMax[T Numeric](a *T, b T) {
	if *a < b {
		*a = b
	}
}

func LowerBound(key int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	left, right := 0, len(nums)-1
	if nums[right] < key {
		return right + 1
	}
	for 1 < right-left {
		mid := (left + (right - 1)) / 2
		if key <= nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	if key <= nums[left] {
		return left
	}
	return right
}

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	v := make(Queue[T], 0)
	return v
}

func (q *Queue[T]) Enqueue(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Dequeue() T {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *Queue[T]) Front() T {
	return (*q)[0]
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}
