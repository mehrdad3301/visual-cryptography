package main 

import ( 
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"log" 
	"flag"
	"utils"
	"strconv"
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

func isBlack(color color.Color) bool { 
	r , g , b ,a := color.RGBA() 
	if (r >= 200) && (g >= 200) && 
		(b >= 200) && (a>>8 >= 200) { 
		return true }
	return false 
} 

func setShare(transparent *image.Gray , share int , x , y , c int) { 
	
	for i := 0 ; i < c ; i++ { 
		for j := 0 ; j < c ; j++ { 

			bit := getKthBit(share , i*c + j)	
			clr := color.White 
			if bit == 1 { 
				clr = color.Black
			}
			transparent.Set(2*x + i , 2*y + j , clr)
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
	fmt.Println(black)
	if black { 
		shares = utils.GetBlackShares(n) 
	} else { 
		shares = utils.GetWhiteShares(n) 
	}
	setTransparents(transparents , shares , x , y , c)   
	
}
func writeImages(imgs []*image.Gray) { 
	 
	for i , img := range(imgs) { 

		f , err := os.Create("img_" + strconv.Itoa(i+1) + ".png") 
		if err != nil { 
			log.Fatal(err) 
		}	
	
		defer f.Close() 
		png.Encode(f , img) 	
	}	
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

func encrypt(imgAddress string , n int) { 

	img := readImage(imgAddress)	
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
	encrypt(os.Args[3] , *n) 
	fmt.Println(*n)
}
	
