package dto

type GeneralListDto struct {
	Offset int
	Limit int
	Order string
	Q map[string]interface{}
}
