package base

import (
	"fmt"
	"goforge/pkg/components/newutils"
)

func New(name string) {
	if newBaseErr := newutils.ShouldNotCreateElement(name, newutils.Base); newBaseErr != nil {
		fmt.Println(newBaseErr)
	} else {
		newutils.CreateNewElement(name, newutils.Base)
	}
}
