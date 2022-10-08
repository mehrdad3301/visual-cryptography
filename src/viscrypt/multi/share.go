package multi

/*
See single/share.go for definition of share.
*/ 

import ( 
	"time"
	"math/rand"
	"utils"
) 

var ( 
	
	threeShares = []int { 0b0111 , 0b1011 , 0b1101 , 0b1110 } 
	twoShares = []int { 0b1100 , 0b1010 , 0b1001 , 0b0110 , 0b0011 , 0b0101 } 
	subpixels = 4 
) 

func GetShares(targetIsBlack , img1IsBlack , img2IsBlack bool) []int { 

	rand.Seed(time.Now().UnixNano()) 
	
	var numRemWhitePixels int	
	if targetIsBlack { 
		numRemWhitePixels = 0 
	} else { 
		numRemWhitePixels = 1
	} 
		
	hole := 0  
	var shares []int
	for ; utils.CountOneBits(hole , subpixels) < subpixels - numRemWhitePixels ; { 
		shares = []int {getSingleShare(img1IsBlack) , 
						getSingleShare(img2IsBlack) }
		hole = utils.GetHole(shares , subpixels) 
	}

	return shares 
		
}

func getSingleShare(isImgBlack bool) int {
	if isImgBlack { 	
		return threeShares[rand.Intn(len(threeShares))] 	
	} else { 
		return twoShares[rand.Intn(len(twoShares))]	
	}
}

