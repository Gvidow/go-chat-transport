package entity

type Segment struct {
	Time uint64 `json:"time" example:"1711902448"`
	Data string `json:"data"`
	Size int    `json:"size" example:"15"`
	Num  int    `json:"num" example:"4"`
}
