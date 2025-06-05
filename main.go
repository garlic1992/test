package main

func main() {
	numTilings(10)

}

func numTilings(n int) int {
	var mod int = 1e9 + 7
	dp := [3]int{1, 0, 0}
	for i := 1; i <= n; i++ {
		temp := [3]int{0, 0, 0}
		temp[0] = (dp[0] + dp[1]*2 + dp[2]) % mod
		temp[1] = (dp[1] + dp[2]) % mod
		temp[2] = dp[0]
		dp = temp
	}
	return dp[0]
}

type H struct {
	nums []int
}

func (h *H) Len() int {
	return len(h.nums)
}

func (h *H) Less(i, j int) bool {
	return h.nums[i] < h.nums[j]
}

func (h *H) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *H) Push(_ any) {}

func (h *H) Pop() (_ any) {
	//x := h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	return
}
