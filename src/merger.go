
package main 

import ( 
	"os"
	"image"
	"image/color"
	"utils"
) 
	
func merge(images []image.Image) { 
	
	startPoint , endPoint := images[0].Bounds().Min , images[0].Bounds().Max
	mergedImage := image.NewGray(image.Rect(startPoint.X , startPoint.Y , 
										   endPoint.X ,endPoint.Y))
	for x := startPoint.X ; x < endPoint.X ; x++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y++ { 
			c := color.White 
			for _ , image := range(images) { 
				if utils.IsBlack(image.At(x , y)) { 
					c = color.Black 
				}
			}
			mergedImage.Set(x , y , c) 	
		}
	}
	utils.WriteImage("merged.png" , mergedImage) 
} 

func main() { 

	imgNames := os.Args[1:] 
	images := utils.ReadImages(imgNames) 
	merge(images) 
} 
