package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init")
}

var name string = "app"

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}
type intTest func(int) bool

func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
func filter(nums []int, f intTest) []int {
	var result []int
	for _, num := range nums {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}
func printPerson(p person) {
	fmt.Println(p.Name, p.Age, p.Sex)
}
func main() {
	//	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//	var oddResult = filter(nums, isOdd)
	//	var evenResult = filter(nums, isEven)
	//	fmt.Printf("oddNums:%T", oddResult)
	//	fmt.Println(oddResult)
	//	fmt.Printf("evenNums:%T", evenResult)
	//	fmt.Println(evenResult)
	//	var p = person{Name: "张三", Age: 18, Sex: "男"}
	//	printPerson(p)
	//	var p1 = new(person)
	//	p1.Name = "张三"
	//	p1.Age = 18
	//	p1.Sex = "男"
	//	printPerson(*p1)
	//	fmt.Println(_const.Monday)
	//
	//	fmt.Printf("%v\n", "你好")            // 可以作为任何值的占位符输出
	//	fmt.Printf("%v %T\n", "枫枫", "枫枫") // 打印类型
	//	fmt.Printf("%d\n", 3)                 // 整数
	//	fmt.Printf("%.1f\n", 1.25)            // 小数
	//	fmt.Printf("%s\n", "哈哈哈")          // 字符串
	//	fmt.Printf("%#v\n", "")               // 用go的语法格式输出，很适合打印空字符串
	//
	//	fmt.Println(math.MaxInt)
	//
	//	var s = `nihao
	//saxas
	//`
	//	fmt.Println(s)
	//	var array1 = []int{1, 2, 3}
	//	array1[2] = 10 // 根据索引找到对应的元素位置，然后替换
	//	array1 = append(array1, 4)
	//	fmt.Println(array1)
	//
	//	var list = [...]string{"a", "b", "c"}
	//	slices := list[:2] // 左一刀，右一刀  变成了切片
	//	fmt.Printf("%#v", slices)
	//testMap()
	//testFor()
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(testClosure(2)(1, 2, 3, 4, 5))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(_const.Monday)
	//var s Student = Student{Name: "张三", Age: 0, Password: "123456"}
	//bytes, _ := json.Marshal(s)
	//fmt.Println(string(bytes))
	//var code MyCode = 1
	//fmt.Printf("%T", code)
	testInterface()
}
func testInterface() {
	var a Animal
	a = Chicken{Name: "chick"}
	a.Eat()
}

type MyCode int

type Student struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"`
}

func testMap() {
	var m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	for index, value := range m {
		fmt.Println(index, value)
	}
	fmt.Printf("%#v", m)
}
func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func testWhile(numList ...int) (int, error) {
	l := len(numList)
	for i := 0; i < l; i++ {
		fmt.Println(numList[i])
	}
	return l, nil
}

func testClosure(t int) func(numList ...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	return func(numList ...int) int {
		sum := 0
		for i := 0; i < len(numList); i++ {
			sum += numList[i]
		}
		return sum
	}
}
