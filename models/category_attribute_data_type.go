package models

type CategoryAttributeDataType int

const (
	text CategoryAttributeDataType = iota
	number
	boolean
	selectType
	mutliSelect
	color
)
