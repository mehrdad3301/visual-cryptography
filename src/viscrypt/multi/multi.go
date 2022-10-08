package multi 


import ( 

	"utils"
	"image"
)


func setPixels(transparents []*image.Gray , x , y , c int , 
				targetIsBlack , img1IsBlack , img2IsBlack bool) { 
	
	shares := GetShares(targetIsBlack , img1IsBlack , img2IsBlack) 
	utils.SetTransparents(transparents , shares , x , y , c)

}
	


func Encrypt(filenames []string) { 
	
	images := utils.ReadImages(filenames) 
	startPoint , endPoint := images[0].Bounds().Min , images[0].Bounds().Max 
	rect , c := utils.GetRectangle(startPoint , endPoint , 2) 
	fmt.Println(images[0].Bounds())
	
	transparents := utils.GetTransparents(2 , rect) 
	
	for x := startPoint.X ; x < endPoint.X ; x ++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y ++ { 

			setPixels(transparents , x , y , c , utils.IsBlack(images[0].At(x , y)),
												 utils.IsBlack(images[1].At(x , y)),
												 utils.IsBlack(images[2].At(x , y)))	
		}
	}  

	utils.WriteImages(transparents) 	

} 
