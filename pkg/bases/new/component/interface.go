package component

import (
	"fmt"
	"goforge/pkg/components/newutils"
)

func New(name string) {
	if newComponentErr := newutils.ShouldNotCreateElement(name, newutils.Component); newComponentErr != nil {
		fmt.Println(newComponentErr)
	} else {
		newutils.CreateNewElement(name, newutils.Component)
	}
}
