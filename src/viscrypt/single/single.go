package single

import ( 
	"image"
	"utils"
)



func setPixels(transparents []*image.Gray, x , y , c int , black bool) { 

	n := len(transparents) 
	var shares []int 
	if black { 
		shares = GetBlackShares(n) 
	} else { 
		shares = GetWhiteShares(n) 
	}
	utils.SetTransparents(transparents , shares , x , y , c)   
	
}



func Encrypt(imgAddress string , n int) { 

	img := utils.ReadImage(imgAddress)	
	startPoint , endPoint := img.Bounds().Min , img.Bounds().Max 
	rect , c := utils.GetRectangle(startPoint , endPoint , n)
	transparents := utils.GetTransparents(n , rect)

	for x := startPoint.X ; x < endPoint.X ; x++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y++ { 
			setPixels(transparents , x , y , c , utils.IsBlack(img.At(x , y))) 
		}
	}

	utils.WriteImages(transparents)
}

