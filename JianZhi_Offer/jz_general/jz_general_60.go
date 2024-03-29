package jz_general

// n 个骰子 求数字的 可能性

func dicesProbability(n int) []float64 {
	res := make([]float64, 6)
	for i := 0; i < len(res); i++ {
		res[i] = 1.0 / 6
	}
	for k := 2; k <= n; k++ {
		temp := make([]float64, k*5+1)
		for i := 0; i < len(res); i++ {
			for j := 0; j < 6; j++ {
				temp[i+j] += res[i] / 6
			}
		}
		res = make([]float64, k*5+1) // 记住 这里不是 res：=  make([]float64, k*5+1)  多了一个冒号，最后结果是res 永远都是只有6个。
		copy(res, temp)
	}
	return res
}
