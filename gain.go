package main

type gainBase map[string]float32

func (g gainBase) Max() string {
	var maxValue float32
	var maxKey string
	for key, value := range g {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	return maxKey
}
