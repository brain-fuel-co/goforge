package app

import (
	"fmt"
	"goforge/pkg/components/utils"
)

func New(name string) {
	if newAppErr := utils.ShouldNotCreateElement(name, utils.App); newAppErr != nil {
		fmt.Println(newAppErr)
	} else {
		utils.CreateNewElement(name, utils.App, utils.EPL2)
	}
}
