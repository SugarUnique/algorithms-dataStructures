package main

import(
	"fmt"
	"sort"
	"reflect"
)
type People struct{
	Name string
	Age int
}
type PeopleArr []People 
func (p PeopleArr)Less(i,j int)bool{
	return p[i].Age < p[j].Age
}

func (p PeopleArr)Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

func (p PeopleArr)Len()int{
	return len(p)
}

//冒泡排序
func BubbleSort(arr []int){
	n := len(arr)
	for i:=1;i<n;i++{
		for j:=0;j<n-i;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}
//直接插入排序

func InsertSort(arr []int){
	n := len(arr)
	for i:=1;i<n;i++{
		temp := arr[i]
		j:=i-1
		for ;j>=0;j--{
			if arr[j] > temp{
				arr[j+1] = arr[j]
			}else{
				break
			}
		}
		arr[j+1] = temp
	}
}
//选择排序
func SelectSort(arr []int){
	n := len(arr)
	for i:=0;i<n-1;i++{
		for j:=i+1;j<n;j++{
			if arr[i] > arr[j]{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
}
//快速排序
func Partition(start,end int,arr[]int)int{
	temp := arr[start]

	for start<end{
		for start<end&&arr[end]>=temp{
			end--
		}
		arr[start] = arr[end]
		for start<end&&arr[start]<=temp{
			start++
		}
		arr[end] = arr[start]
	}
	arr[start] = temp
	return start
}

func QuickSort(start,end int,arr[]int){
	if start<end{
		mid := Partition(start,end,arr)
		QuickSort(start,mid-1,arr)
		QuickSort(mid+1,end,arr)
	}
}
//二路归并排序
func MergeSort(arr[]int){
	n := len(arr)
	buf := make([]int, n)
	BranchSort(0,n-1,arr,buf)
}
func BranchSort(start,end int,arr[]int,buf[]int){
	if start<end{
		mid := (start+end)/2
		BranchSort(start,mid,arr,buf)
		BranchSort(mid+1,end,arr,buf)
		Merge(start,mid,end,arr,buf)
	}
}
func Merge(start,mid,end int, arr[]int,buf[]int){
	i := start
	j := mid+1
	k := start
	for i<=mid&&j<=end{
		if arr[i]<arr[j]{
			buf[k] = arr[i]
			i++
			k++
		}else{
			buf[k] = arr[j]
			j++
			k++
		}
	}
	
	for i<=mid{
		buf[k] = arr[i]
		i++
		k++
	}

	for j<=end{
		buf[k] = arr[j]
		j++
		k++
	}

	for m:=0;m<=end;m++{
		arr[m] = buf[m]
	}
}
//堆排序
func HeapSort(arr []int){
	n := len(arr)

	for i:=n/2-1;i>=0;i--{
		Sift(i,n-1,arr)
	}

	for j:=n-1;j>=0;j--{
		arr[0],arr[j] = arr[j],arr[0]
		Sift(0,j-1,arr)
	}
}

func Sift(start,end int,arr[]int){
	i := start
	j := 2*i+1
	for j<=end{
		if j<end&&arr[j]<arr[j+1]{
			j++
		}
		if arr[j]<arr[i]{
			break
		}
		arr[i],arr[j] = arr[j],arr[i]
		i = j
		j = 2*i+1
	}
}

type User struct{
	Id int
	Name string
	Age int
}

func (user User)Hello(s string){
	fmt.Println("Hello World ",s)
}

func Info(o interface {}){
	t := reflect.TypeOf(o)
	fmt.Println("Type:",t.Name())
	if kind := t.Kind();kind!=reflect.Struct{
		fmt.Println("Type error")
		return 
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i:=0;i<t.NumField();i++{
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Type,f.Name,val)
	}
	for j:=0;j<t.NumMethod();j++{
		m := t.Method(j)
		fmt.Println(m.Name,m.Type)
	}
}

//闭包的应用
func test(x int)(func(),func()){
	return func(){
		fmt.Println(x)
		x+=10
	},func(){
		fmt.Println(x)
	}
}
func main(){
	testMap := map[string]int{"zhan1":12,"si":34,"cheng":6,"li":6}
	testStructArr := make(PeopleArr, len(testMap))
	i := 0
	for k,v := range testMap{
		testStructArr[i] = People{k,v}
		i++
	}
	fmt.Println(testStructArr)
	sort.Sort(testStructArr)
	fmt.Println(testStructArr)
	// sortArr := []int{3,1,2,6,5,9,1,2,5,4,10}
	// BubbleSort(sortArr)
	// fmt.Println("冒泡排序后：",sortArr);
	// InsertSort(sortArr)
	// fmt.Println("直接排序后：",sortArr);
	// QuickSort(0,len(sortArr)-1,sortArr)
	// fmt.Println("快速排序后：",sortArr);
	// SelectSort(sortArr)
	// fmt.Println("选择排序后：",sortArr);
	// HeapSort(sortArr)
	// fmt.Println("大堆排序后：",sortArr);
	// MergeSort(sortArr)
	// fmt.Println("归并排序后：",sortArr);
	user := User{1,"zhan",27}
	Info(&user)
	a,b := test(100)
	a()
	b()
	defer_call()
}

func defer_call() {
    defer func() {fmt.Println("打印前")}()
    defer func() {fmt.Println("打印中")}()
    defer func() {fmt.Println("打印后")}()
    panic("触发异常")
}
