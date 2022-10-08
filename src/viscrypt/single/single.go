package single

/*
Single package takes an image and argument n. It encrypts 
image into n transparencies. The hidden image is only revealed 
by stacking all the transparencies. A transparency is one of the
n resulting encrypted images which appears to be nothing more 
than randome noise.

See examples at assets/example_{2_2 , 3_3 , 4_4} 
*/

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

