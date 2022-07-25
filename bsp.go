package bsp

import (
	"fmt"
	"strings"
)

type BSPGenerationResult int8

const (
	BSPGenerationResult_INVALID BSPGenerationResult = 0
	BSPGenerationResult_IN      BSPGenerationResult = 1
	BSPGenerationResult_UP      BSPGenerationResult = 2
	BSPGenerationResult_BOTTOM  BSPGenerationResult = 3
	BSPGenerationResult_CUT     BSPGenerationResult = 4
)

type BSPItem interface {
	Valid() bool
	String() string
}

type BSP[T BSPItem] interface {
	String(space ...int) string
}

type bspBranch[T BSPItem] struct {
	items  []T
	up     BSP[T]
	bottom BSP[T]
}

type bspLeaf[T BSPItem] struct {
	items []T
}

// TODO Checks if the default item is valid
func GenerateBSP[T BSPItem](items []T, f func(T, T) (BSPGenerationResult, T, T)) BSP[T] {
	if len(items) == 0 {
		return &bspLeaf[T]{}
	}

	var in, up, bottom []T

	// Takes an item from the list
	// It becomes the cut in the space
	defaultItem := items[0]

	in = append(in, defaultItem)

	for _, item := range items[1:] {
		if !item.Valid() {
			fmt.Printf("item %s is invalid\n", item)
			continue
		}
		// iUp & iBottom are defined iff res = CUT
		res, iUp, iBottom := f(defaultItem, item)
		switch res {
		case BSPGenerationResult_UP:
			fmt.Printf("UP:\t%s -- %s\n", defaultItem, item)
			up = append(up, item)
		case BSPGenerationResult_BOTTOM:
			fmt.Printf("BOTTOM:\t%s -- %s\n", defaultItem, item)
			bottom = append(bottom, item)
		case BSPGenerationResult_IN:
			fmt.Printf("IN:\t%s -- %s\n", defaultItem, item)
			in = append(in, item)
		case BSPGenerationResult_CUT:
			fmt.Printf("CUT:\t%s -- %s\n", defaultItem, item)
			if iUp.Valid() {
				up = append(up, iUp)
			} else {
				fmt.Printf("generated item %s is invalid\n", iUp.String())
			}
			if iBottom.Valid() {
				bottom = append(bottom, iBottom)
			} else {
				fmt.Printf("generated item %s is invalid\n", iBottom.String())
			}
		}
	}
	fmt.Println(up)
	fmt.Println(bottom)
	if len(up)+len(bottom) == 0 {
		return &bspLeaf[T]{
			items: in,
		}
	} else {
		return &bspBranch[T]{
			items:  in,
			up:     GenerateBSP(up, f),
			bottom: GenerateBSP(bottom, f),
		}
	}
}

func (this bspBranch[T]) String(spaces ...int) string {
	if len(spaces) == 0 {
		return this.string(0)
	} else {
		return this.string(spaces[0])
	}
}

func (this bspLeaf[T]) String(spaces ...int) string {
	if len(spaces) == 0 {
		return this.string(0)
	} else {
		return this.string(spaces[0])
	}
}

func (this bspBranch[T]) string(spaces int) string {
	sSpaces := strings.Repeat("  ", spaces)
	return fmt.Sprintf("Branch: %s\n%s- up: %s\n%s- bottom: %s",
		stringAllItems(this.items),
		sSpaces, this.up.String(spaces+1),
		sSpaces, this.bottom.String(spaces+1))

}

func (this bspLeaf[T]) string(spaces int) string {
	return fmt.Sprintf("Leaf: %s", stringAllItems(this.items))
}

func stringAllItems[T BSPItem](items []T) string {
	sItems := make([]string, len(items))
	for i, item := range items {
		sItems[i] = item.String()
	}
	return fmt.Sprintf("[%s]", strings.Join(sItems, ", "))
}
