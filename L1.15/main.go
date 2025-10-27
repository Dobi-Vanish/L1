package main

import (
	"strings"
)

func byteCopyFunc() string {
	v := createHugeString(1 << 10)

	justString := string([]byte(v[:100]))

	return justString
}

func builderFunc() string {
	v := createHugeString(1 << 10)

	var builder strings.Builder
	builder.Grow(100)
	builder.WriteString(v[:100])

	return builder.String()
}

func main() {
	result := byteCopyFunc()
	//	result := builderFunc() // Второй вариант через string.Builder. Не совсем понятно как хотят в задании - исправить или свою реализацию сделать.
	_ = result // Используем хоть как-то, ибо иначе компилятор не пропустит.
}

// createHugeString создаём чтобы пропустил компилятор.
func createHugeString(size int) string {
	return strings.Repeat("1", size)
}
