package multi 

/*
Multi is a encryption scheme in visual-cryptography. 
It takes three images as input and returns two transparents. 
First image contains the secret to be encrypted.The rest two 
images are used to conceive. The transparents correspond to the 
misleading images fed to the program as input. 

For more information on definition of transparent see single 
package doc. To see examples of this scheme see assets/example_multi. 
*/

import ( 

	"utils"
	"image"
)


func setPixels(transparencies[]*image.Gray , x , y , c int , 
				targetIsBlack , img1IsBlack , img2IsBlack bool) { 
	
	shares := GetShares(targetIsBlack , img1IsBlack , img2IsBlack) 
	utils.SetTransparency(transparencies , shares , x , y , c)

}
	


func Encrypt(filenames []string) { 
	
	images := utils.ReadImages(filenames) 
	startPoint , endPoint := images[0].Bounds().Min , images[0].Bounds().Max 
	rect , c := utils.GetRectangle(startPoint , endPoint , 2) 
	
	transparencies := utils.GetTransparency(2 , rect) 
	
	for x := startPoint.X ; x < endPoint.X ; x ++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y ++ { 

			setPixels(transparencies, x , y , c , utils.IsBlack(images[0].At(x , y)),
												 utils.IsBlack(images[1].At(x , y)),
												 utils.IsBlack(images[2].At(x , y)))	
		}
	}  

	utils.WriteImages(transparencies) 	
} 
