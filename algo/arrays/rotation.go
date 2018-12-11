package arrays

import "github.com/joanlopez/golang-hacker-rank/math"

// Time complexity: O(n * d)
func OneByOneRotation(arr []int, d, n int) {
	for i := 0; i < d; i++ {
		leftRotateOne(arr, n)
	}
}

// Time complexity: O(n)
func TemporaryRotation(arr []int, d, n int) {
	tmp := make([]int, d)
	for i := 0; i < d; i++ {
		tmp[i] = arr[i]
	}

	for i := d; i < n; i++ {
		arr[i-d] = arr[i]
	}

	for i := 0; i < d; i++ {
		arr[i+n-d] = tmp[i]
	}
}

// Time complexity: O(n)
func ReverseRotation(arr []int, d, n int) {
	reverse(arr, 0, d-1);
	reverse(arr, d, n-1);
	reverse(arr, 0, n-1);
}

// Time complexity: O(n)
func JugglingRotation(arr []int, d, n int) {
	gcd := math.Gcd(d,n)
	for i := 0; i < gcd; i++ {
		temp := arr[i]
		j := i
		for true {
			k := j + d
			if k >= n {
				k = k - n
			}
			if k == i {
				break
			}
			arr[j] = arr[k]
			j = k
		}
		arr[j] = temp
	}
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start] // Swap
		start++
		end--
	}
}

func leftRotateOne(arr []int, n int) {
	temp := arr[0]
	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[n-1] = temp
}
