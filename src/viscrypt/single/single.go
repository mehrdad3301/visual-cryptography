package single

import ( 
	"image"
	"image/color"
	"utils"
)


func setShare(transparent *image.Gray , share int , x , y , c int) { 
	
	for i := 0 ; i < c ; i++ { 
		for j := 0 ; j < c ; j++ { 

			bit := GetKthBit(share , i*c + j)	
			clr := color.White 
			if bit == 1 { 
				clr = color.Black
			}
			transparent.Set(c*x + i , c*y + j , clr)
		} 
	}
} 

func setTransparents(transparents []*image.Gray , shares []int , x , y , c int) { 

	for i := range(transparents) { 
		setShare(transparents[i] , shares[i] , x , y , c) 
	}
} 

func setPixels(transparents []*image.Gray, x , y , c int , black bool) { 

	n := len(transparents) 
	var shares []int 
	if black { 
		shares = GetBlackShares(n) 
	} else { 
		shares = GetWhiteShares(n) 
	}
	setTransparents(transparents , shares , x , y , c)   
	
}

func getTransparents(n int , rect image.Rectangle) []*image.Gray { 
	
	transparents := make([]*image.Gray , 0 , n)  
	for i:=0 ; i<n ; i++ { 
		transparents = append(transparents , image.NewGray(rect)) 
	}
	return transparents

}

func getRectangle(a , b image.Point , n int) (image.Rectangle , int) { 
	
	var multiplier int 
	
	if n == 4 { 
		multiplier = 3 
	} else { 
		multiplier = 2 
	}
	return image.Rect(a.X , a.Y , multiplier * b.X , multiplier * b.Y) , multiplier
} 

func Encrypt(imgAddress string , n int) { 

	img := utils.ReadImage(imgAddress)	
	startPoint , endPoint := img.Bounds().Min , img.Bounds().Max 
	rect , c := getRectangle(startPoint , endPoint , n)
	transparents := getTransparents(n , rect)

	for x := startPoint.X ; x < endPoint.X ; x++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y++ { 
			setPixels(transparents , x , y , c , utils.IsBlack(img.At(x , y))) 
		}
	}

	utils.WriteImages(transparents)
}

