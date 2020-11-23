package main

import (
	"fmt"
	"strconv"
	"strings"
)

type spine struct {
	name        string
	connectedTo []string
}

type leaf spine

var leaves []leaf

func addLeaf(name string, spineNames []string, spines *[]spine) {
	newLeaf := leaf{name: name, connectedTo: spineNames}
	leaves = append(leaves, newLeaf)
	for i := 0; i < len(*spines); i++ {
		(*spines)[i].connectedTo = append((*spines)[i].connectedTo, name)
	}
}

func main() {
	var spines []spine
	var spineNames []string

	var location string
	fmt.Printf("Location name: ")
	fmt.Scanln(&location)

	var numSpines int
	fmt.Printf("number of spines: ")
	fmt.Scanln(&numSpines)

	var numLeaves int
	fmt.Printf("number of leaves: ")
	fmt.Scanln(&numLeaves)

	for i := 1; i < numSpines+1; i++ {
		spines = append(spines, spine{name: location + "-Spine-" + strconv.Itoa(i), connectedTo: nil})
	}

	for _, v := range spines {
		spineNames = append(spineNames, v.name)
	}

	for i := 1; i < numLeaves+1; i++ {
		addLeaf(location+"-Leaf-"+strconv.Itoa(i), spineNames, &spines)

	}

	// fmt.Println(leaves)

	for _, v := range spines {
		fmt.Println(v.name + " is connected to " + strings.Join(v.connectedTo, ", "))
	}

	// for _, v := range leaves {
	// 	fmt.Println(v.name + " is connected to: " + strings.Join(v.connectedTo, " and "))
	// }

	// fmt.Println(spineAa.name + " is connected to " + strings.Join(spineAa.connectedTo, ","))
	// fmt.Println(spineAb.name + " is connected to " + strings.Join(spineAb.connectedTo, ","))

}
