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
