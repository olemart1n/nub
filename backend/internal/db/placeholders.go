package db

import "fmt"

func makePlaceholders(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		out[i-1] = fmt.Sprintf("$%d", i)
	}
	return out
}

func stringSliceToInterfaces(slice []string) []any {
	out := make([]any, len(slice))
	for i, v := range slice {
		out[i] = v
	}
	return out
}
