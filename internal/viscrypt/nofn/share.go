package nofn

/*
A share is a slice of integer. Each integer in binary representation
contain information on how to set subpixels in transparencies. For
example 0b011111000 corresponds to a 3x3 matrix of subpixels in tra-
nsparency. Below matrix illustrates how an integer is interpreted
and used to set pixels in transparency.

		0 1 1
		1 1 1
		0 0 0
*/

import (
	"image"
	"math/rand"
)

var (

	// shares that are used for black pixels when number of transparencies are two
	nTwoBlackShares = [][]int{{0b1100, 0b0011}, {0b1010, 0b0101}, {0b1001, 0b0110}}

	// shares that are used for white pixels when number of transparencies are two
	nTwoWhiteShares = []int{0b1100, 0b1010, 0b0011, 0b0101, 0b1001, 0b0110}

	// number of subpixels per pixel
	nTwoSubpixels = 4

	// shares that are used for black pixels when number of transparencies are three
	nThreeBlackShares = []int{0b1100, 0b1010, 0b1001}

	// shares that are used for black pixels when number of transparencies are three
	nThreeWhiteShares = []int{0b0011, 0b0101, 0b0110}

	// number of subpixels per pixel when number of transparencies are three
	nThreeSubpixels = 4

	// shares that are used for black pixels when number of transparencies are four
	nFourBlackShares = []int{0b011011010, 0b010111001, 0b010110110, 0b100111010}

	// shares that are used for white pixels when number of transparencies are four
	nFourWhiteShares = []int{0b011111000, 0b010110011, 0b001110101, 0b000111110}

	// number of subpixels per pixel
	nFourSubpixels = 9
)

// permutate columns in a share
func permutate(share []int, subpixels int) []int {

	i := rand.Intn(subpixels)
	j := i
	for j == i {
		j = rand.Intn(subpixels)
	}

	for idx, v := range share {
		bitI := getKthBit(v, i)
		bitJ := getKthBit(v, j)
		v = setKthBit(v, j, bitI)
		v = setKthBit(v, i, bitJ)
		share[idx] = v
	}

	return share
}

func shuffle(shares []int) []int {

	for i := range shares {
		j := rand.Intn(i + 1)
		shares[i], shares[j] = shares[j], shares[i]
	}

	return shares
}

func getKthBit(number, k int) int {
	return (number >> k) & 1
}

func setKthBit(number, k, bit int) int {
	if bit == 1 {
		return number | (1 << k)
	} else {
		return number &^ (1 << k)
	}
}

func getBlackShares(numTransparencies int) []int {

	if numTransparencies == 4 {
		permutate(nFourBlackShares, nFourSubpixels)
		return shuffle(nFourBlackShares)
	} else if numTransparencies == 2 {
		x := rand.Intn(len(nTwoBlackShares))
		return shuffle(nTwoBlackShares[x])
	} else if numTransparencies == 3 {
		permutate(nThreeBlackShares, nThreeSubpixels)
		return shuffle(nThreeBlackShares)
	} else {
		return nil
	}
}

func getWhiteShares(n int) []int {

	if n == 4 {
		hole := getHole(nFourWhiteShares, nFourSubpixels)
		permutate(nFourWhiteShares, nFourSubpixels)
		newHole := getHole(nFourWhiteShares, nFourSubpixels)
		for hole == newHole {
			newHole = getHole(permutate(nFourWhiteShares, nFourSubpixels), nFourSubpixels)
		}
		return nFourWhiteShares
	} else if n == 2 {
		x := rand.Intn(len(nTwoWhiteShares))
		return shuffle([]int{nTwoWhiteShares[x], nTwoWhiteShares[x]})
	} else if n == 3 {
		permutate(nThreeWhiteShares, nThreeSubpixels)
		return shuffle(nThreeWhiteShares)
	} else {
		return nil
	}
}

func getRectangle(a, b image.Point, n int) (image.Rectangle, int) {

	var multiplier int

	if n == 4 {
		multiplier = 3  
	} else {
		multiplier = 2 
	}
	return image.Rect(a.X, a.Y, multiplier*b.X, multiplier*b.Y), multiplier
}

// TODO remove this function
func getHole(share []int, subpixels int) int {

	hole := 0
	for _, n := range share {
		for i := 0; i < subpixels; i++ {
			x := getKthBit(n, i)
			if x == 1 {
				hole = setKthBit(hole, i, 1)
			}
		}
	}
	return hole
}
