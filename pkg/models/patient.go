package models

//{"ConfirmDate":"2021-05-04","No":null,"Age":51,"Gender":"หญิง","GenderEn":"Female","Nation":null,"NationEn":"China","Province":"Phrae","ProvinceId":46,"District":null,"ProvinceEn":"Phrae","StatQuarantine":5}
type Patient struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             string `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	ProvinceId     int    `json:"ProvinceId"`
	Province       string `json:"Province"`
	ProvinceEn     string `json:"ProvinceEn"`
	District       string `json:"District"`
	StatQuarantine int    `json:"StatQuarantine"`
}

type AgeSummary struct {
	Thirty    int `json:"0-30"`
	Sixty     int `json:"31-60"`
	SixtyPlus int `json:"61+"`
	Na        int `json:"N/A"`
}

type Provinces struct {
	Province map[string]int
}

type Summary struct {
	Province Provinces  `json:"Province"`
	AgeGroup AgeSummary `json:"AgeGroup"`
}
