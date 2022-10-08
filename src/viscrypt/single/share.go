package single 

import ( 
	"fmt"
	"math/rand"
	"time"
	"utils"
) 

var ( 

	blackShareTwo = [][]int{ []int{0b1100 , 0b0011} , []int{0b1010 , 0b0101} , []int{0b1001 , 0b0110} } 
	whiteShareTwo = []int{ 0b1100 , 0b1010 , 0b0011 , 0b0101 , 0b1001 , 0b0110 } 
	shareTwoSubPixels = 4 
	blackShareThree = []int { 0b1100 , 0b1010 , 0b1001 } 
	whiteShareThree = []int { 0b0011 , 0b0101 , 0b0110 } 

	blackShareFour = []int { 0b011011010 , 0b010111001 , 0b010110110 , 0b100111010 }
    whiteShareFour = []int { 0b011111000 , 0b010110011 , 0b001110101 , 0b000111110 } 
	shareFourSubPixels = 9
) 

func GetBlackShares(n int) []int { 

	rand.Seed(time.Now().UnixNano()) 
	if n == 4 { 
		permutate(blackShareFour , shareFourSubPixels)
		return ShuffleShares(blackShareFour)
	} else if n == 2 { 
		x := rand.Intn(len(blackShareTwo))
		return ShuffleShares(blackShareTwo[x]) 
	} else if n == 3 { 
		permutate(blackShareThree , shareTwoSubPixels) 
		return ShuffleShares(blackShareThree) 
	} else { 
		return nil
	}
} 

func GetWhiteShares(n int) []int { 
		
	rand.Seed(time.Now().UnixNano()) 
	if n == 4 { 
		hole := utils.GetHole(whiteShareFour , shareFourSubPixels)
		permutate(whiteShareFour , shareFourSubPixels)
		newHole := utils.GetHole(whiteShareFour , shareFourSubPixels)
		for ; hole == newHole ; { 
			newHole = utils.GetHole(permutate(whiteShareFour , shareFourSubPixels),
														 shareFourSubPixels)
		}
		return ShuffleShares(whiteShareFour)
	} else if n == 2 { 
		x := rand.Intn(len(whiteShareTwo))
		return []int{whiteShareTwo[x] , whiteShareTwo[x]}
	} else if n == 3 { 
		permutate(whiteShareThree , shareTwoSubPixels)
		return whiteShareThree 	
	} else { 
		return nil
	}
} 



func permutate(share []int , x int) []int { 
	
	i := rand.Intn(x) 
	j := i 
	for ; j == i ; { 
		j = rand.Intn(x)
	}	
	
	for idx , v := range(share) { 
		bitI := utils.GetKthBit(v , i) 
		bitJ := utils.GetKthBit(v , j) 
		v = utils.SetKthBit(v , j , bitI) 
		v = utils.SetKthBit(v , i , bitJ)
		share[idx] = v 
	}

	return share
} 

func ShuffleShares(share []int) []int { 
	
	for i := range share { 
		j := rand.Intn(i + 1) 
		share[i] , share[j] = share[j] , share[i]
	}

	return share 
}

func printTransparency(num int) { 

	fmt.Printf("%03b\n" , num >> 6) 
	fmt.Printf("%03b\n" , (num >> 3) - (num >> 6 << 3)) 
	fmt.Printf("%03b\n" , num - (num >> 3) << 3 ) 
	fmt.Println()
}

func printShare(share []int) { 

	for _ , v := range(share) { 
		printTransparency(v) 
	}
} 

