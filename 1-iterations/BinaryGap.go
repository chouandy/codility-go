package iterations

// https://app.codility.com/programmers/lessons/1-iterations/binary_gap/

// 標準解概念如下：
// 1. 主要使用二元運算子 "右移 (>>)"。
// 2. 先宣告 "最大計算數量" 與 "目前計算數量"，均為 0。
// 3. 取得 unsigned int 轉換成二進位後，右邊連續為 0 的個數，採用的方法為
//    https://graphics.stanford.edu/~seander/bithacks.html#ZerosOnRightModLookup。
// 4. 如果有找到右邊連續為 0 的個數，使用二元運算子 "右移 (>>)"，將 N 直接往右移至非 0 的位址，再開始尋找。
// 5. 當 N 大於 0 則重複以下步驟
//   (1) N 除以 2 取餘數 (%)，如果為 1，則重置 "目前計算數量"，如果為 0，則累加 "目前計算數量"
//   (2) 如果 "目前計算數量" 大於 "最大計算數量"，則將 "最大計算數量" 更新為 "目前計算數量"
// 6. 返回 "最大計算數量"

var mod37BitPosition = []uint{
	32, 0, 1, 26, 2, 23, 27, 0, 3, 16, 24, 30, 28, 11, 0, 13, 4,
	7, 17, 0, 25, 22, 31, 15, 29, 10, 12, 6, 0, 21, 14, 9, 5,
	20, 8, 19, 18,
}

// BinaryGap solution for BinaryGap
func BinaryGap(N int) int {
	// New variables
	maxLength := 0
	currentLength := 0

	// Get trailing zeros
	trailing := mod37BitPosition[(N&-N)%37]

	// If trailing > 0, right shift
	if trailing > 0 {
		N = int(N >> uint(trailing))
	}

	for N > 0 {
		// Reset or increase the counter
		if N%2 == 1 {
			currentLength = 0
		} else {
			currentLength++
		}

		// Update max length
		if currentLength > maxLength {
			maxLength = currentLength
		}

		// N right shift 1
		N >>= 1
	}

	return maxLength
}
