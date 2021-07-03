package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	var list list.List

	// Push value ke list
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(50)
	list.PushBack(60)

	// Print semua isi list
	// Logic : For dimulai dari nilai pertama dari list
	// Ketika pada pointer memori tertentu nilai i tidak kosong
	// Maka lanjut. Kalau kosong end loop

	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(int))
		fmt.Print(" ")

	}
}
