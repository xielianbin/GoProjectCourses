package main
import "fmt"

type Shoe struct {
	brand  string	
	size   int
	color  string
	price  float64
}
// 定义一个函数类型，使用函数返回函数类型，并把返回的函数类型作为新的函数的参数
type Option func(*Shoe)
func WithBrand(brand string) Option {
	return func(s *Shoe) {
		s.brand = brand
	}
}
func WithSize(size int) Option{
	return func(s *Shoe){
		s.size = size
	}
}
func WithColor(color string) Option {
	return func(s *Shoe) {
		s.color = color
	}
}
// 返回的是一个闭包函数
func WithPrice(price float64) Option {
	return func(s *Shoe) {
		s.price = price
	}
}	
func NewShoe(opts ...Option) *Shoe{
	s:=&Shoe{
		brand: "Nike",		
		size:  42,
		color: "Black",	
		price: 99.99,
		}
	// _,opt中，第一个是索引，第二个才是值
	for _,opt:=range opts{
		opt(s)
	}
	return s
}
func main(){
	s1:=NewShoe()
	fmt.Println("Shoe 1:", s1)
	s2:=NewShoe(WithBrand("Anta"),
	)
	fmt.Println("Shoe 2:", s2)
	s3:=NewShoe(WithBrand("Adidas"),
	 WithSize(44), 
	 WithColor("White"),
	  WithPrice(89.99),
	)
	fmt.Println("Shoe 3:", s3)
	s4:=NewShoe(WithBrand("Puma"),)
	WithColor("Red")(s4) // 使用闭包函数
	fmt.Println("Shoe 4:", s4)
}