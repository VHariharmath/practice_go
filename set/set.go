package main

import "fmt"

type Set struct {
	integetMap map[int]bool
}

func (set *Set) AddElement(ele int) {

	if !set.ContainsElement(ele) {
		set.integetMap[ele] = true
	}
}

func (set *Set) DeleteElement(ele int) {
	delete(set.integetMap, ele)
}

func (set *Set) ContainsElement(ele int) bool {
	_, exists := set.integetMap[ele]

	return exists
}

func (set *Set) IntersectSet(anotherSet *Set) *Set {
	intersectSet := &Set{}
	intersectSet.New()

	for key, _ := range set.integetMap {
		if anotherSet.ContainsElement(key) {
			intersectSet.AddElement(key)
		}
	}

	return intersectSet
}

func (set *Set) UnionSet(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()

	for key, _ := range set.integetMap {
		unionSet.AddElement(key)
	}

	for key, _ := range anotherSet.integetMap {
		unionSet.AddElement(key)
	}

	return unionSet
}

func (set *Set) New() {
	set.integetMap = make(map[int]bool)
}

func main() {
	set := &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	set.AddElement(2)
	set.AddElement(3)
	fmt.Println("Set after adding")
	fmt.Println(set)
	set.DeleteElement(2)
	fmt.Println("set after deleting")
	fmt.Println(set)

	anotherSet := &Set{}
	anotherSet.New()

	anotherSet.AddElement(3)
	anotherSet.AddElement(4)
	fmt.Println("Another set")
	fmt.Println(anotherSet)

	fmt.Println("Intersection")
	fmt.Println(set.IntersectSet(anotherSet))

	fmt.Println("Union")
	fmt.Println(set.UnionSet(anotherSet))
}
