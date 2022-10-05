package main 

import ( 
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"log"
)

var ( 

	h_share = []int{1 , 1 , 0 , 0}
	v_share = []int{1 , 0 , 0 , 1}  
	d_share = []int{1 , 0 , 1 , 0} 
	
	shares = [][]int{h_share , v_share , d_share}
)

func getComplement(x []int) []int { 
	
	m := make([]int , 0 , cap(x)) 
	for _ , v := range(x) { 
		m = append(m , 1 - v) 
	}
	return m 
}

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

func setShare(img *image.Gray , transparent []int , x, y int) { 

	for i:=0 ; i<=1 ; i++ { 
		for j:=0 ; j<=1 ; j++ { 
			c := color.Black 
			if transparent[i*2+j] == 0 { 
				c = color.White
			}	
			img.Set(2 * x + i , 2 * y + j , c)  
		}
	}

}
func setPixels(img1 , img2 *image.Gray, x , y int , color color.Color) { 

	id := rand.Intn(len(shares)) 
	if r ,_ ,_,_ := color.RGBA() ; r <= 200 { 
		setShare(img1 , shares[id] , x , y)	
		setShare(img2 , getComplement(shares[id]) , x , y)		
	} else { 
		setShare(img1 , shares[id] , x , y)	
		setShare(img2 , shares[id] , x , y)		
	}
		
	
}
func writeImage(img *image.Gray, filename string) { 
	 
	f , err := os.Create(filename) 

	if err != nil { 
		log.Fatal(err) 
	}	

	defer f.Close() 
	png.Encode(f , img) 	
}

func encrypt(filename string) { 

	img := readImage(filename)	
	startPoint , endPoint := img.Bounds().Min , img.Bounds().Max 
	rect := image.Rect(startPoint.X , startPoint.Y , 2 * endPoint.X , 2 * endPoint.Y) 
	img1 := image.NewGray(rect)
	img2 := image.NewGray(rect)

	for x := startPoint.X ; x < endPoint.X ; x++ { 
		for y := startPoint.Y ; y < endPoint.Y ; y++ { 
			setPixels(img1 , img2 , x , y , img.At(x , y)) 	
		}
	}

	writeImage(img1 , "img1.png") 
	writeImage(img2 , "img2.png") 
}

func main() { 
	
	encrypt(os.Args[1]) 
	fmt.Println()
}
	
