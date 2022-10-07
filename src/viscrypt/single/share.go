package single 

import ( 
	"fmt"
	"math/rand"
	"time"
) 

var ( 

	blackShareTwo = [][]int{ []int{0b1100 , 0b0011} , []int{0b1010 , 0b0101} , []int{0b1001 , 0b0110} } 
	whiteShareTwo = []int{ 0b1100 , 0b1010 , 0b0011 , 0b0101 , 0b1001 , 0b0110 } 

	blackShareThree = []int { 0b1100 , 0b1010 , 0b1001 } 
	whiteShareThree = []int { 0b0011 , 0b0101 , 0b0110 } 

	blackShareFour = []int { 0b011011010 , 0b010111001 , 0b010110110 , 0b100111010 }
    whiteShareFour = []int { 0b011111000 , 0b010110011 , 0b001110101 , 0b000111110 } 
) 

func GetBlackShares(n int) []int { 

	rand.Seed(time.Now().UnixNano()) 
	if n == 4 { 
		permutate(blackShareFour , 9)
		return ShuffleShares(blackShareFour)
	} else if n == 2 { 
		x := rand.Intn(len(blackShareTwo))
		return ShuffleShares(blackShareTwo[x]) 
	} else if n == 3 { 
		permutate(blackShareThree , 4) 
		return ShuffleShares(blackShareThree) 
	} else { 
		return nil
	}
} 

func GetWhiteShares(n int) []int { 
		
	rand.Seed(time.Now().UnixNano()) 
	if n == 4 { 
		hole := getHole(whiteShareFour)
		permutate(whiteShareFour , 9)
		newHole := getHole(whiteShareFour)
		for ; hole == newHole ; { 
			newHole = getHole(permutate(whiteShareFour , 9))
		}
		return ShuffleShares(whiteShareFour)
	} else if n == 2 { 
		x := rand.Intn(len(whiteShareTwo))
		return []int{whiteShareTwo[x] , whiteShareTwo[x]}
	} else if n == 3 { 
		permutate(whiteShareThree , 4)
		return whiteShareThree 	
	} else { 
		return nil
	}
} 

func GetKthBit(number , k int) int { 
	return (number >> k) & 1 
} 

func setKthBit(number , k , bit int) int { 
	if bit == 1 { 
		return number | ( 1 << k ) 
	} else if bit == 0 { 
		return number &^ ( 1 << k ) 
	} else { 
		return number
	} 
} 

func permutate(share []int , x int) []int { 
	
	i := rand.Intn(x) 
	j := i 
	for ; j == i ; { 
		j = rand.Intn(x)
	}	
	
	for idx , v := range(share) { 
		bitI := GetKthBit(v , i) 
		bitJ := GetKthBit(v , j) 
		v = setKthBit(v , j , bitI) 
		v = setKthBit(v , i , bitJ)
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

func getHole(share []int) int { 
	
	hole := 0 
	for _ , n := range(share) { 
		for i:=0 ; i<9 ; i++ { 
			x := GetKthBit(n , i)
			if x == 1 { 	
				hole = setKthBit(hole , i , 1)
			}
		}
	}

	return hole 
} 

