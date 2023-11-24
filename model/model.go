package model

type M map[string]interface{}

type Weather struct {
	Water int `json:"water" form:"water" swagger:"description(water)"`
	Wind  int `json:"wind" form:"wind" swagger:"description(wind)"`
}
