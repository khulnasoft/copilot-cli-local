// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package diff

type eqFunc func(inA, inB int) bool

// longestCommonSubsequence computes the longest common subsequence of two lists, and returns two lists that contain 
// the positions of the common items in the input lists, respectively.
// When multiple correct answers exists, the function picks one of them deterministically.
// Example:
//   input_a = ["a", "c", "b", "b", "d"], input_b = ["a", "B", "b", "c", "c", "d"]
//   One LCS is ["a","c","d"].
//   "a" is input_a[0] and input_b[0], "c" is in input_a[1] and input_b[3], "d" is input_a[4] and input_b[5].
//   Therefore, the output will be [0,1,4], [0,3,5]
func longestCommonSubsequence[T any](a []T, b []T, eq eqFunc) ([]int, []int) {
	if len(a) == 0 || len(b) == 0 {
		return nil, nil
	}
	// Initialize the matrix
	lcs := make([][]int, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		lcs[i] = make([]int, len(b)+1)
		lcs[i][len(b)] = 0
	}
	for j := 0; j < len(b)+1; j++ {
		lcs[len(a)][j] = 0
	}
	// Compute the lengths of the LCS for all sub lists.
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			switch {
			case eq(i, j):
				lcs[i][j] = 1 + lcs[i+1][j+1]
			case lcs[i+1][j] < lcs[i][j+1]:
				lcs[i][j] = lcs[i][j+1]
			default:
				lcs[i][j] = lcs[i+1][j]
			}
		}
	}
	// Backtrace to construct the LCS.
	var i, j int
	var indicesA, indicesB []int
	for {
		if i >= len(a) || j >= len(b) {
			break
		}
		switch {
		case eq(i, j):
			indicesA, indicesB = append(indicesA, i), append(indicesB, j)
			i++
			j++
		case lcs[i+1][j] < lcs[i][j+1]:
			j++
		default:
			i++
		}
	}
	return indicesA, indicesB
}