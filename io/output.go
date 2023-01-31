package io

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

// Print2D prints a 2D matrix with each element separated by a separator
func Print2D[T any](matrix [][]T, sep string, converter func(T) string) string {
	outputSlice := make([]string, 0, len(matrix))
	for _, row := range matrix {
		outputSlice = append(outputSlice, asStringSep(row, sep, converter))
	}
	return strings.Join(outputSlice, "\n")
}

// Print2DInt prints a 2D int matrix with each element separated by a separator
func Print2DInt[T constraints.Integer](matrix [][]T, sep string) string {
	return Print2D(matrix, sep, func(t T) string {
		return strconv.Itoa(int(t))
	})
}

// Print2DString prints a 2D string matrix with each element separated by a separator
func Print2DString[T ~string](matrix [][]T, sep string) string {
	return Print2D(matrix, sep, func(t T) string {
		return string(t)
	})
}

// Print2DStringer prints a 2D string matrix with each element separated by a separator
func Print2DStringer(matrix [][]fmt.Stringer, sep string) string {
	return Print2D(matrix, sep, func(t fmt.Stringer) string {
		return t.String()
	})
}

func asStringSep[T any](row []T, sep string, converter func(T) string) string {
	stringSlice := make([]string, len(row))
	for i := range row {
		stringSlice[i] = converter(row[i])
	}
	return strings.Join(stringSlice, sep)
}
