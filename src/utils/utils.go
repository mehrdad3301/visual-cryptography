package utils 

import ( 
	
	"image"
	"image/color" 
	"image/png"
	"log" 
	"os"
	"strconv"
) 


func IsBlack(c color.Color) bool {

	r , g , b , _ := c.RGBA() 
	y := 0.299 * float32(r) + 
		 0.587 * float32(g) +  
		 0.114 * float32(b) 
	
	if y <= 255/2 { 
		return true 
	} else { 
		return false 
	}
} 

func ReadImage(filename string) image.Image { 

	f , err := os.Open(filename) 
	if err != nil { 
		log.Fatal(err) 
	} 

	defer f.Close()
	img , _ , err := image.Decode(f) 
	if err != nil {
		log.Fatal(err) 
	}
	
	return img 	
}

func WriteImage(filename string , img *image.Gray) { 


	f , err := os.Create(filename) 
	if err != nil { 
		log.Fatal(err)
	}
	
	defer f.Close() 
	png.Encode(f , img)
}

func WriteImages(imgs []*image.Gray) { 

	for i , img := range(imgs) { 
		filename := "img_" + strconv.Itoa(i) + ".png" 
		WriteImage(filename , img) 
	} 
} 


func ReadImages(names []string) []image.Image { 
	
	images := make([]image.Image , 0) 
		for _ , name := range(names) { 
			images = append(images , ReadImage(name))
		}
	return images 
}

func GetRectangle(a , b image.Point , n int) (image.Rectangle , int) { 
	
	var multiplier int 
	
	if n == 4 { 
		multiplier = 3 
	} else { 
		multiplier = 2 
	}
	return image.Rect(a.X , a.Y , multiplier * b.X , multiplier * b.Y) , multiplier
}
 
func GetTransparents(n int , rect image.Rectangle) []*image.Gray { 
	
	transparents := make([]*image.Gray , 0 , n)  
	for i:=0 ; i<n ; i++ { 
		transparents = append(transparents , image.NewGray(rect)) 
	}
	return transparents

}

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

func SetTransparents(transparents []*image.Gray , shares []int , x , y , c int) { 

	for i := range(transparents) { 
		setShare(transparents[i] , shares[i] , x , y , c) 
	}
} 

func GetKthBit(number , k int) int { 

	return (number >> k) & 1 
}

func SetKthBit(number , k , bit int) int { 
	if bit == 1{ 
		return number | (1 << k) 
	} else {
		return number &^ (1 << k)
	} 
}


func GetHole(share []int , subpixels int) int { 
	
	hole := 0 
	for _ , n := range(share) { 
		for i:=0 ; i<subpixels ; i++ { 
			x := GetKthBit(n , i)
			if x == 1 { 	
				hole = SetKthBit(hole , i , 1)
			}
		}
	}
	return hole 
} 

func CountOneBits(num , k int) int { 

	sum := 0 
	for i := 0 ; i < k ; i++ { 
		bit := utils.GetKthBit(num , i)
		if bit == 1{ 
			sum++
		}
	}
	return sum
} 

