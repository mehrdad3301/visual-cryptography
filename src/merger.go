
package main 

import ( 
	"os"
	"image"
	"image/png"
	"image/color"
	"log"
) 
	


func openImages(names []string) []image.Image { 

	images := make([]image.Image , 0) 
	for _ , name := range(names) { 

		file , err := os.Open(name)
		if err != nil { 
			log.Fatal(err) 
		}
		
		image , _  , err := image.Decode(file)
		if err != nil { 
			log.Fatal(err) 
		}

		images = append(images , image) 	
		file.Close() 
	}
	return images 
}

func isBlack(c color.Color) bool { 
	r , g , b, a := c.RGBA() 
	if (r == 0) && (b == 0) &&
	   (g == 0) && (a >> 8 == 255) { 
		return true 
	}
	return false
}

func merge(images []image.Image) { 
	
	startPoint , endPoint := images[0].Bounds().Min , images[0].Bounds().Max
	mergedImage := image.NewGray(image.Rect(startPoint.X , startPoint.Y , 
										   endPoint.X ,endPoint.Y))
	for x := startPoint.X ; x < endPoint.X ; x++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y++ { 
			c := color.White 
			for _ , image := range(images) { 
				if isBlack(image.At(x , y)) { 
					c = color.Black 
				}
			}
			mergedImage.Set(x , y , c) 	
		}
	}
	writeImage("merged.png" , mergedImage) 
} 

func writeImage(filename string , img *image.Gray) { 

	file , err := os.Create(filename) 
	if err != nil { 
		log.Fatal(err) 
	} 
	defer file.Close()

	png.Encode(file , img) 
} 

func main() { 

	imgNames := os.Args[1:] 
	images := openImages(imgNames) 
	merge(images) 
} 
