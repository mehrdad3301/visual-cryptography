package main 

import ( 
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"log" "flag"
	"share"
)

func readImage(filename string) (image.Image) { 

	file , err := os.Open(filename)	
	if err != nil { 
		log.Fatal(err)
	}

	defer file.Close() 
	img ,_ ,  err := image.Decode(file)	
	if err != nil { 
		log.Fatal(err)
	}

	return img 
}

func getKthBit(number , k int) int {
	
	return (number >> k) & 1
	
}
func setShare(transparent *image.Gray , share int , x , y , c int) { 
	
	for i := 0 ; i < c ; i++ { 

		for j := 0 ; j < c ; j++ { 
			bit := getKthBit(share , i*c +j)	
			color := color.White 
			if bit == 1 { 
				color = color.Black
			}
			transparent.Set(i , j , color)
		} 

	}


} 

func isBlack(color color.Color) bool { 
	r , g , b ,a := color.RGBA() 
	if (r == 255) && (g == 255) && 
		(b == 255) && (a>>8 == 255) { 
		return true }
	return false 
} 

func setTransparents(transparents []*image.Gray , shares []int , x , y , c int) { 

	for i := range(transparents) { 
		setShare(transparents[i] , shares[i] , x , y ,c) 
	}
} 

func setPixels(transparents []*image.Gray, x , y , c int , black bool) { 

	n := len(transparents) 
	var shares []int if black { shares = share.GetBlackShares(n) } else { 
		shares = share.GetWhiteShares(n) 
	}
	setTransparents(transparents , shares , x , y , c)   
	
}
func writeImages(imgs []*image.Gray, filename string) { 
	 
	for i , img := range(imgs) { 

		f , err := os.Create(filename + i) 
	
		if err != nil { 
			log.Fatal(err) 
		}	
	
		defer f.Close() 
		png.Encode(f , img) 	
	}	
}

func getTransparents(n int , rect image.Rectangle) []*image.Gray { 
	
	transparents := make([]*imageGray , 0 , n)  
	for i=0 ; i<n ; i++ { 
		transparents.append(image.NewGray(rect)) 
	}
	return transparents

}

func getRectangle(a , b Point , n int) (image.Rectangle , int) { 
	
	var multiplier int 
	
	if n == 4 { 
		multiplier = 3 
	} else { 
		multiplier = 2 
	}
	return image.Rect(a.X , a.Y , multiplier * b.X , multiplier * b.Y) , multiplier
} 

func encrypt(filename string , n int) { 

	img := readImage(filename)	
	startPoint , endPoint := img.Bounds().Min , img.Bounds().Max 
	rect , c := getRectangle(startPoint , endPoint , n)
	transparents := getTransparents(n , rect)

	for x := startPoint.X ; x < endPoint.X ; x++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y++ { 
			setPixels(transparents , x , y , c , isBlack(img.At(x , y))) 	
		}
	}

	writeImages(transparents)
}

func main() { 
	
	n := flag.Int("n" , 2 , "number of transparents") 
	flag.Parse() 
	encrypt(os.Args[1] , n) 
	fmt.Println(*n)
}
	
