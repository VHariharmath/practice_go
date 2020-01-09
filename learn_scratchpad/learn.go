package main

import (
	"fmt"
)

func sockMerchant(n int, ar []int) int {
	var max int
	for i := 0; i < n; i++ {
		if max < ar[i] {
			max = ar[i]
		}
	}

	set := make([]int, max+1)
	var count int

	for i := 0; i < n; i++ {
		set[ar[i]]++
	}

	for i := 0; i < n; i++ {
		count += set[ar[i]] / 2
		set[ar[i]] = 0
	}
	return count
}

func main() {
	ar := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}

	fmt.Println(sockMerchant(len(ar), ar))
}

/*
func isValidIP(ipAddr string) bool {

	strSlice := strings.Split(ipAddr, ".")

	if len(strSlice) < 4 {
		fmt.Println("only", len(strSlice), "present")
		return false
	}

	for _, x := range strSlice {
		a, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println(x, "is not a number")
			return false
		}

		if a >= 0 && a <= 255 {
			continue
		}

	}

	return true
}

func main() {

	addresses := []string{"192.168.0.1", "192.168.0", "goh", "", "10.0.5.25", "125.512.100.abc"}

	for _, ipAddr := range addresses {

		if isValidIP(ipAddr) {
			fmt.Println("Valid IP")
		} else {
			fmt.Println("Invalid IP")
		}

	}
}

/*
type headNode struct {
	hn *node
}

type node struct {
	data int
	next *node
}

func allocNode(item int) *node {
	return &node{data: item, next: nil}
}

func (head *headNode) insertRear(item int) *headNode {
	new := allocNode(item)
	if new == nil {
		fmt.Println("Allocation failed")
		return nil
	}

	if head.hn == nil {
		head.hn = new
		return head
	}

	temp := head.hn
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = new

	return head
}

func (head *headNode) insertFront(item int) *headNode {
	new := allocNode(item)
	if new == nil {
		fmt.Println("Allocation failed")
		return nil
	}

	if head == nil {
		return nil
	} else {
		new.next = head.hn
		head.hn = new
		return head
	}
}

func (head *headNode) printList() {
	temp := head.hn
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (head *headNode) deleteRear() *headNode {
	temp := head.hn
	var prev *node

	for temp.next != nil {
		prev = temp
		temp = temp.next
	}

	prev.next = nil

	return head
}

func (head *headNode) deleteFront() *headNode {
	temp := head.hn
	head.hn = temp.next

	return head
}

func main() {
	head := &headNode{nil}
	head = head.insertFront(3)
	head = head.insertFront(2)
	head = head.insertFront(1)

	head = head.insertRear(4)
	head = head.insertRear(5)
	head.printList()

	log.Println("Delete from Rear")
	head = head.deleteRear()
	head.printList()

	log.Println("Delete from Front")
	head = head.deleteFront()
	head.printList()
}

/*
package main

import (
	"fmt"
)


func addOneToElements(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
}

func deleteElementFromSlice(slice []int) []int {
	return slice[0 : len(slice)-1]
}

func extendSlice(slice []int ) []int {
	return slice[0: len(slice)+1]
}

func main() {
	var buffer [256]int
	slice := buffer[20:30]

	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	addOneToElements(slice)
	fmt.Println(len(slice))

	newSlice := deleteElementFromSlice(slice)
	fmt.Println(len(newSlice))

	newSlice = extendSlice(newSlice)
	fmt.Println(len(newSlice))
}

/*
func main() {

	resp, err := http.Get("https://www.amazon.in/")
	if err != nil {
		fmt.Println("Failed to get")
		return
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to get")
		return
	}

	scanner := bufio.NewScanner(resp.Body)
	resp.Body.Close()
	fmt.Println(bs)
}

/*
func deleteOneElementFromSlice(slice *[]byte) {
	sliceptr := *slice
	*slice = sliceptr[0 : len(sliceptr)-1]
}
func addOneToElements(slice []byte) {

	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
}
func main() {
	var buffer [256]byte
	slice := buffer[20:30]

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println(slice)

	addOneToElements(slice)

	fmt.Println(slice)

	fmt.Println("Before:\t", slice, len(slice))
	deleteOneElementFromSlice(&slice)
	fmt.Println("After:\t", slice, len(slice))

}

/*
func main() {
	x := []int{3, 1, 7, 5, 0, 45, 34, 28, 89}
	s := []string{"Virat", "Rohit", "AB", "Ajinkya", "MS"}
	sort.Ints(x)
	fmt.Println(x)
	sort.Strings(s)
	fmt.Println(s)
}

/*
import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	x := []person{{"Virat", 30}, {"AB", 31}, {"Ajinkya", 29}}

	sort.Sort(ByAge(x))

	fmt.Println(x)

}

/*
import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	x := []person{{"Virat", "Kohli", 29},
		{"Ajinkya", "Rahane", 29},
	}

	bs, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))

	var y [10]person
	z := y[:2]

	err = json.Unmarshal(bs, &z)
	if err != nil {
		panic(err)
	}

	fmt.Println(z)

}

/*
func main() {
	s := "My experience and knowledge need a “LANGUAGE” upgradation"
	fmt.Println(s)
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Replace(s, " ", "-", -1))

}

/*
type rect struct {
	l float64
	b float64
}

type circ struct {
	r float64
}

type measure interface {
	area() float64
}

func geometry(m measure) float64 {
	return m.area()
}

func (a rect) area() float64 {
	return a.l * a.b
}

func (b circ) area() float64 {
	return math.Pi * b.r * b.r
}

func main() {
	x := rect{10, 20}
	y := circ{10}
	fmt.Println(geometry(x))
	fmt.Println(geometry(y))
}

/*
func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

/*
func sum(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func even(f func(args ...int) int, args ...int) int {
	xi := []int{}
	for _, v := range args {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}

	return f(xi...)
}

func odd(f func(args ...int) int, args ...int) int {
	xi := []int{}
	for _, v := range args {
		if v%2 == 1 {
			xi = append(xi, v)
		}
	}

	return f(xi...)
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := sum(x...)
	fmt.Println(s1)
	s2 := even(sum, x...)
	fmt.Println(s2)
	s3 := odd(sum, x...)
	fmt.Println(s3)

}

/*
type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(args []int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}

	return sum
}

func server(req chan *Request) {
	request := <-req

	for _, v := range request.args {
		fmt.Println(v)
	}
	sum := request.f(request.args)
	req <- request
	request.resultChan <- sum
}

func main() {
	ChannelRequests := make(chan *Request)
	request := &Request{[]int{10, 20, 30}, sum, make(chan int)}

	ChannelRequests <- request

	go server(ChannelRequests)

	<-ChannelRequests
	fmt.Println("Reply: ", <-request.resultChan)

}

/*
type homePageSize struct {
	url  string
	size int
}

func main() {
	urls := []string{
		"https://www.amazon.in",
		"https://www.apple.com",
		"https://www.google.com",
		"https://www.microsoft.com",
	}

	results := make(chan homePageSize)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Error")
				panic(err)
			}

			results <- homePageSize{
				url:  url,
				size: len(bs),
			}
		}(url)
	}

	fmt.Println("routines are scheduled... ")
	var biggest homePageSize

	for i, _ := range urls {
		result := <-results
		fmt.Println(i, "results:", result)
		if result.size > biggest.size {
			biggest = result
		}
	}

	fmt.Println("Biggest gome page is: ", biggest.url)
}

/*
import (
	"fmt"
)

func main() {
	m := map[int]map[string]int{
		1: map[string]int{
			"Vinayak": 29,
			"Pallavi": 28,
		},
		2: map[string]int{
			"Josh":  21,
			"merry": 35,
		},
	}

	if el, ok := m[1]; ok {
		fmt.Println(el["Pallavi"], el["Vinayak"])
	}
}

/*
type Node struct {
	data int
	next *Node
}

func (head *Node) insert(key int) *Node {

	item := new(Node)
	item.data = key
	var prev *Node

	if head == nil {
		return item
	} else {
		for i := head; i != nil; i = i.next {
			prev = i
		}
		prev.next = item
		return head
	}
}

func printList(head *Node) {
	for i := head; i != nil; i = i.next {
		fmt.Println(i.data)
	}
}
func main() {
	head := new(Node)
	head = head.insert(1)
	head = head.insert(2)
	head = head.insert(3)
	head = head.insert(4)
	printList(head)
}

/*
func main() {
	var x list.List
	x.PushBack(12)
	x.PushFront(11)
	x.PushBack(13)
	x.PushFront(10)

	var el *list.Element
	count := 0

	for i := x.Front(); i != nil; i = i.Next() {
		count++
		if count == 3 {
			el = i
		}
		fmt.Println(i.Value)
	}

	fmt.Println()
	x.InsertAfter(20, el)

	for i := x.Front(); i != nil; i = i.Next() {

		fmt.Println(i.Value)
	}
}

/*
type Node struct {
	data int
	next *Node
}

func insert(key int, head *Node) Node {

	item := Node{key, nil}
	var prev *Node

	if head == nil {
		return item
	} else {
		for i := head; i != nil; i = i.next {
			prev = i
		}
		prev.next = &item
		return *head
	}
}

func printList(head *Node) {
	for i := head; i != nil; i = i.next {
		fmt.Println(i.data)
	}
}
func main() {
	head := Node{}
	head = insert(1, &head)
	head = insert(2, &head)
	printList(&head)
}

/*
func check(slice ...int) {

	fmt.Printf("%T \n", slice)
	fmt.Printf("%v \n", len(slice))
	fmt.Printf("%v \n", cap(slice))
	fmt.Printf("%v \n", &slice[0])

}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := x[:5]

	fmt.Printf("%v \n", &slice[0])
	check(slice...)

}

/*
func pinger(c chan string) {

	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {

	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	var msg string
	for {
		msg = <-c
		fmt.Println(msg)
	}
}
func main() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}

/*
func Append(slice []int, items ...int) []int {

	n := len(slice)

	total := n + len(items)

	if n < total {
		newSlice := make([]int, n, total*2+1)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:total]
	copy(slice[n:], items)
	return slice
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := []int{11, 12, 13, 15, 16, 17, 18, 19, 20}
	slice := x[:5]

	fmt.Println(slice, len(slice), cap(slice))

	slice = Append(slice, y...)

	fmt.Println(slice, len(slice), cap(slice))
}

/*
type path []byte

func (s *path) truncateFromLast() {
	fmt.Printf("%s \n", *s)
	i := bytes.LastIndex(*s, []byte("/"))

	if i >= 0 {
		*s = (*s)[0:i]
	}
}

func (s *path) replace() {
	for i, b := range *s {
		if b == 'U' {
			(*s)[i] = '^'
		}
	}
}
func (s *path) toUpper() {
	for i, b := range *s {
		if b >= 'a' && b <= 'z' {
			(*s)[i] = b + 'A' - 'a'
		}
	}
}
func main() {
	s := path("/usr/bin/local")

	s.truncateFromLast()
	s.toUpper()
	fmt.Printf("%s \n", s)
	s.replace()
	fmt.Printf("%s \n", s)
}

/*
func deleteOneElementFromSlice(slice *[]byte) {
	sliceptr := *slice
	*slice = sliceptr[0 : len(sliceptr)-1]
}
func addOneToElements(slice []byte) {

	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
}
func main() {
	var buffer [256]byte
	slice := buffer[20:30]

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println(slice)

	addOneToElements(slice)

	fmt.Println(slice)

	fmt.Println("Before:\t", slice, len(slice))
	deleteOneElementFromSlice(&slice)
	fmt.Println("After:\t", slice, len(slice))

}

/*
func main() {

	var x list.List
	var e *list.Element

	x.PushBack(12)
	x.PushBack(24)
	x.PushBack(48)

	var p int

	for i := x.Front(); i != nil; i = i.Next() {

		if p == 2 {
			e = i
		}
		p++
		fmt.Println(i.Value.(int))
	}

	x.InsertAfter(36, e)

	for i := x.Front(); i != nil; i = i.Next() {
		fmt.Println()
		fmt.Println(i.Value)
	}

}

/*
func main() {

	dir, err := os.Open("../.")

	if err != nil {
		fmt.Println("failed to open dir")
		return
	}

	defer dir.Close()

	contents, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("failed to read dir")
		return
	}

	for _, v := range contents {
		fmt.Println(v)

	}
}

/*
func main () {

	s := "Grit and Persistence"
	fmt.Println(s)
	fmt.Printf("after: %v\n", strings.Replace(s, " ", "-", -1))
}
/*
func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}

	_, err = file.WriteString("Good morning Pune !")
	if err != nil {
		fmt.Println("Failed to write to file")
		return
	}

	file.Close()

	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("Failed to open file")
		return
	}

	defer file.Close()

	bs, err := ioutil.ReadFile("test.txt")

	stat, err := file.Stat()

		if err != nil {
			fmt.Println("Failed to stat file")
			return
		}

		bs := make([]byte, stat.Size())

		_, err = file.Read(bs)
		if err != nil {
			fmt.Println("Failed to read file", err)
			return
		}
	s := string(bs)

	fmt.Println(s)

}

/*
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%v \t\t %b \n", kb, kb)
	fmt.Printf("%v \t\t %b \n", mb, mb)
	fmt.Printf("%v \t\t %b \n", gb, gb)
}

/*
var (
	a int     = 2
	b float64 = 3.0
	c string  = "hello"
)

var p int = 12

const l int = 90
const (
	x int    = 4
	y string = "H"
)

func main() {

}

/*
func replace(arr *string) {

	for i := 0; i < len(*arr); i++ {
		if arr[i] == 'a' {
			arr[i] = '1'
		}
	}

}

func main() {

	s := "vinayak"
	replace(&s)
	fmt.Println(s)
}

/*
type hotdog int

var a hotdog = 90

func main() {

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	a = 42
	fmt.Println(a)

	var b int
	fmt.Println(int(a) == b)
	b = int(a)

	fmt.Println(b)

}

/*

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v \n", x, y, z)
	fmt.Println(s)

}

/*
type hotdog int

var a hotdog = 90

func main() {
	var b int

	fmt.Println(b)
	b = int(a)

	fmt.Println(b)

}

/*
func main() {
	// Example 1: Basic
	myText := "Welcome to GoLangCode.com"
	myText = strings.Replace(myText, "Welcome", "Willkommen", -1)

	// Output: Willkommen to GoLangCode.com
	fmt.Println(myText)

	// Example 2: Change first occurance
	myText = "The sound sounds sound"
	myText = strings.Replace(myText, "sound", "car", 1)

	// Output: The car sounds sound
	fmt.Println(myText)

	// Example 3: Replacing quotes (double backslash needed)
	myText = "vinayak"
	myText = strings.Replace(myText, "a", "1", -1)

	// Output: I \'quote\' this text
	fmt.Println(myText)
}

/*
func fibonacci(n int, c chan int) {

	x , y := 0,1

	for i := ran
}
func main() {

}
/*
func sum(s []int, c chan int) {
	var sum int = 0
	for _, i := range s {
		sum = sum + i
	}

	c <- sum
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	c := make(chan int)
	go sum(arr[:len(arr)/2], c)
	//x := <-c
	//fmt.Println(<-c)
	go sum(arr[len(arr)/2:], c)
	//y := <-c
	//fmt.Println(<-c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

/*
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*

/*
func main() {
	list := []int{1, 2, 9, 20, 21, 31, 45, 63, 70, 100}

	fmt.Printf("%v", BinarySearch(list, 31))
}

func BinarySearch(list []int, key int) bool {

	var low, mid int = 0, 0
	high := len(list) - 1

	for low <= high {

		mid = (low + high) / 2

		if list[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	fmt.Printf("list[%v] = %v \n", low, list[low])
	if low == len(list) || list[low] != key {
		return false
	}

	return true

}

/*
func LinearSearch(list []int, key int) bool {
	for _, item := range list {
		if item == key {
			return true
		}
	}

	return false
}

func main() {
	list := []int{12, 34, 56, 13, 54, 14, 35, 78, 98, 90}

	fmt.Printf("%v", LinearSearch(list, 34))
}

/*
func main() {
	var num, rev, rem int = 0, 0, 0

	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)

	orig := num

	for {
		rem = num % 10
		rev = rev*10 + rem
		num /= 10

		if num == 0 {
			break
		}
	}

	fmt.Printf("orig = %v reversed = %v \n", orig, rev)
	if orig == rev {
		fmt.Print("palindrome")
	} else {
		fmt.Print("Not palindrome")
	}

}

/*
func main() {
	var num1, num2 = 23, 78

	fmt.Printf("num1 = %v and num2 = %v \n", num1, num2)
	num1 = num1 - num2
	num2 = num1 + num2
	num1 = num2 - num1

	fmt.Printf("num1 = %v and num2 = %v \n", num1, num2)
}

/*
func main() {
	fmt.Println(strings.Compare("India", "Indian"))
	fmt.Println(strings.Compare("India", "India"))
	fmt.Println(strings.Compare("Indian", "India"))
}

/*
func main() {
	var arr [10]int
	var num, sum int = 0, 0
	fmt.Print("Enter the number of elements: ")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Print("Enter the element: ")
		fmt.Scanln(&arr[i])
		sum = sum + arr[i]
	}

	fmt.Printf("avg = %v \n", sum/num)
}

/*
func main() {
	var string1, string2 string
	fmt.Print("Enter the first string:")
	fmt.Scanln(&string1)
	fmt.Print("Enter the seconcd string:")
	fmt.Scanln(&string2)

	fmt.Println(string1 + " " + string2)
}

/*
func main() {
	var country map[int]string
	country = make(map[int]string)

	country[1] = "US"
	country[2] = "China"
	country[3] = "Japan"
	country[4] = "Germany"
	country[5] = "UK"
	country[6] = "India"

	for i, value := range country {
		fmt.Printf("country[%v] = %v \n", i, value)
	}
}

/*
func main() {
	var count [128]int
	// search_string := "aaavvvvddd ppalldk"
	var search_string string
	var index, max int = 0, 0

	fmt.Print("Enter the string:")
	fmt.Scanln(&search_string)

	for i := 0; i < len(search_string); i++ {
		count[search_string[i]]++
	}

	for i := 0; i < 128; i++ {
		if count[i] > max {
			max = count[i]
			index = i
		}
	}

	fmt.Printf("%c appeared %v times \n", index, max)
}

/*
func main() {
	var arr [100]int
	var numEle, big, index int = 0, 0, 0

	fmt.Print("Enter number of elements \n")
	fmt.Scanln(&numEle)

	for i := 0; i < numEle; i++ {
		fmt.Printf("Enter %v the number:", i)
		fmt.Scanln(&arr[i])
	}

	for i := 0; i < numEle; i++ {
		if arr[i] > big {
			big = arr[i]
			index = i
		}
	}

	fmt.Printf("Largest element in the array is arr[%v] = %v \n", index, big)
}

/*
type Person struct {
	Name string
	Age  int
}

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("%v : %v: %v \n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

/*
type IPAddr [4]byte

func (p Person) string() string {
	return fmt.Sprint("%v (%v years)", p.Name, p.Age)
}

func main() {

	hosts := map[string]IPAddr{
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v : %v \n", name, ip)
	}

	a := Person{"Harman", 27}
	b := Person{"Seagate", 28}

	fmt.Println(a, b)
}

*/

/*
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
*/
