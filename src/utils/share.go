package utils

import ( 
	"math/rand"
	"time"
) 

var ( 

	blackShareTwo = [][]int{ []int{0b1100 , 0b0011} , []int{0b1010 , 0b0101} , []int{0b1001 , 0b0110}} 
	whiteShareTwo = []int{ 0b1100 , 0b1010 , 0b0011 , 0b0101 , 0b1001 , 0b0110 } 

	blackShareFour = []int { 0b011011010 , 0b010111001 , 0b010110110 , 0b100111010 }
    whiteShareFour = []int { 0b011111000 , 0b010110011 , 0b001110101 , 0b000111110 } 
) 

func GetBlackShares(n int) []int { 

	rand.Seed(time.Now().UnixNano()) 
	if n == 4 { 
		return ShuffleShares(blackShareFour) 
	} else if n == 2 { 
		x := rand.Intn(len(blackShareTwo))
		return ShuffleShares(blackShareTwo[x]) 
	} else { 
		return nil
	}
} 

func GetWhiteShares(n int) []int { 
		
	rand.Seed(time.Now().UnixNano()) 
	if n == 4 { 
		return ShuffleShares(whiteShareFour) 		

	} else if n == 2 { 
		x := rand.Intn(len(whiteShareTwo))
		return []int{whiteShareTwo[x] , whiteShareTwo[x]}
	} else { 
		return nil
	}

} 

func ShuffleShares(share []int) []int { 
	
	for i := range share { 
		j := rand.Intn(i + 1) 
		share[i] , share[j] = share[j] , share[i]
	}

	return share 
}

