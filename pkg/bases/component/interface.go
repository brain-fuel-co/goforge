package component

import (
	"fmt"
	"goforge/pkg/components/utils"
)

func New(name string) {
	if newComponentErr := utils.ShouldNotCreateElement(name, utils.Component); newComponentErr != nil {
		fmt.Println(newComponentErr)
	} else if newElementErr := utils.CreateNewElement(name, utils.Component, utils.EPL2); newElementErr != nil {
		fmt.Println(newElementErr)
	}
}

func Remove(name string) {
	if removeComponentErr := utils.RemoveElement(name, utils.Component); removeComponentErr != nil {
		fmt.Println(removeComponentErr)
	}
}
