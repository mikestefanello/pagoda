package components

type (
	Color int
	Size  int
)

const (
	ColorNone Color = iota
	ColorNeutral
	ColorPrimary
	ColorSecondary
	ColorAccent
	ColorInfo
	ColorSuccess
	ColorWarning
	ColorError
	ColorLink
)

const (
	SizeExtraSmall Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
	SizeExtraLarge
)
