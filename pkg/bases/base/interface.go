package base

import (
	"fmt"
	"goforge/pkg/components/utils"
)

func New(name string) {
	if newBaseErr := utils.ShouldNotCreateElement(name, utils.Base); newBaseErr != nil {
		fmt.Println(newBaseErr)
	} else {
		utils.CreateNewElement(name, utils.Base, utils.EPL2)
	}
}

func Remove(name string) {
	if removeBaseErr := utils.RemoveElement(name, utils.Base); removeBaseErr != nil {
		fmt.Println(removeBaseErr)
	}
}
