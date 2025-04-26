package main
import "fmt"

func main(){
	var n, hasil int
	fmt.Scan(&n)
	hasil = genap(n)
	fmt.Print(hasil)
}
func genap(n int)int{
	if n < 2{
		return 0
	}
	if n%2 == 0{
		return n + genap(n-2)
	}else{
		return genap(n-1)
	}
}