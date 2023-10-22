package internalBase

import "goforge/pkg/components/utils"

func New(fileName string, elementName string, elementType utils.ElementType) {
	utils.CreateInternalFile(fileName, elementName, elementType, utils.EPL2)
}
