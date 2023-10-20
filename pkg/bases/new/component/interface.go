package component

import (
	"fmt"
	"goforge/pkg/components/newcomponent"
)

func New(name string) {
	if newComponentErr := newcomponent.ShouldNotCreateComponent(name); newComponentErr != nil {
		fmt.Println(newComponentErr)
	} else {
		newcomponent.CreateNewComponent(name)
	}
}
