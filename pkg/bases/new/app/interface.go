package app

import (
	"fmt"
	"goforge/pkg/components/newutils"
)

func New(name string) {
	if newAppErr := newutils.ShouldNotCreateElement(name, newutils.App); newAppErr != nil {
		fmt.Println(newAppErr)
	} else {
		newutils.CreateNewElement(name, newutils.App)
	}
}
