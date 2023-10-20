package component

import "goforge/pkg/components/newcomponent"

func New(name string) {
	newcomponent.CreateNewComponent(name)
}
