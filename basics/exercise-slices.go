package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	var res = make([][]uint8, dx)
	
	for i := 0; i < dx; i++ {
		
		res[i] = make([]uint8, dy)
		
		for j := 0; j < dy; j++ {
			res[i][j] = uint8((i * j))
		}
		
	}

	return res
}

func main() {
	pic.Show(Pic)
}
