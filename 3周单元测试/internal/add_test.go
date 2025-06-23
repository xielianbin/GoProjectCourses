package internal
import "testing"
import "fmt"
func TestAdd(t *testing.T){
	re:=Add(1,3)
	if re!=3{
		t.Errorf("期待 %d, 实际 %d",3,re)
	}
}
// 如果是short模式就跳过，短时间的跳过，使用 go test -short
func TestAdd2(t *testing.T){
	if testing.Short(){
		t.Skip("short 模式下跳过")
	}
	fmt.Println("yes short模式")
	re:=Add(1,3)
	if re!=3{
		t.Errorf("期待 %d, 实际 %d",3,re)
	}
}
// 基于表格驱动的测试
func TestAdd3(t *testing.T){
	// 定义匿名结构体
	var dataset=[] struct{
		a int
		b int
		out int
	}{
		{1,2,3},
		{1,1,2},
		{0,0,0},}
	fmt.Println("yes table 模式")
	for _,value:=range dataset{

	re:=Add(value.a,value.b)
	
	if re!=value.out{
		t.Errorf("期待 %d, 实际 %d",value.out,re)
	}
	}
}

// 性能测试 Benchmark开头
// 测试核心函数
func BenchmarkADD()