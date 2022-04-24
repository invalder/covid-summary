package controllers

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/invalder/covid-summary/pkg/models"
)

type covidCase struct {
	Data []struct {
		ConfirmDate    string      `json:"ConfirmDate"`
		No             interface{} `json:"No"`
		Age            *int        `json:"Age,omitempty"`
		Gender         string      `json:"Gender"`
		GenderEn       string      `json:"GenderEn"`
		Nation         interface{} `json:"Nation"`
		NationEn       string      `json:"NationEn"`
		Province       string      `json:"Province"`
		ProvinceID     int         `json:"ProvinceId"`
		District       interface{} `json:"District"`
		ProvinceEn     string      `json:"ProvinceEn"`
		StatQuarantine int         `json:"StatQuarantine"`
	} `json:"Data"`
}

var url = "https://static.wongnai.com/devinterview/covid-cases.json"
var Patients = models.Patient{}
var Summary = models.Summary{}

func InitializePatients() {
	cases := fetchData()
	Patients = cases.Data
	Summary = fetchSummry(cases)
}

func fetchData() covidCase {

	//Since the Cert is not valid, therefore bypassing is needed
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.Header.Set(http.Header{}, "Content-Type", "application/json")
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	covidCases := covidCase{}

	err = json.Unmarshal(body, &covidCases)
	if err != nil {
		log.Fatal(err)
	}

	return covidCases
}

func fetchSummry(cvdCase covidCase) models.Summary {
	Patients = cvdCase.Data

	PInfo := map[int]string{}
	Provinces := models.Provinces{
		Province: map[string]int{},
	}
	AgeSummary := models.AgeSummary{
		Thirty:    0,
		Sixty:     0,
		SixtyPlus: 0,
		Na:        0,
	}

	for _, patient := range Patients {
		// add province to PInfo
		_, ok := PInfo[patient.ProvinceID]
		if !ok {
			PInfo[patient.ProvinceID] = patient.Province
		}
		// if Province is blank string, update if new item is not blank
		if PInfo[patient.ProvinceID] == "" {
			if patient.Province != PInfo[patient.ProvinceID] {
				PInfo[patient.ProvinceID] = patient.Province
			}
		}

		// add Province to Summary
		pvName := ""
		if patient.Province == "" {
			pvName = PInfo[patient.ProvinceID]
		} else {
			pvName = patient.Province
		}

		_, ok = Provinces.Province[pvName]
		if ok {
			Provinces.Province[pvName]++
		} else {
			Provinces.Province[pvName] = 1
		}

		// Age group procesing
		if patient.Age != nil {
			switch {
			case *patient.Age <= 30:
				AgeSummary.Thirty++
			case *patient.Age <= 60:
				AgeSummary.Sixty++
			case *patient.Age > 60:
				AgeSummary.SixtyPlus++
			default:
				AgeSummary.Na++
			}
		} else {
			AgeSummary.Na++
		}

	}

	return models.Summary{
		Province: Provinces,
		AgeGroup: AgeSummary,
	}
}

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func GetAllPatients(c *gin.Context) {
	c.JSON(200, gin.H{
		"paients": &Patients,
	})
}

func GetSummary(c *gin.Context) {
	c.JSON(200, gin.H{
		"summary": &Summary,
	})
}
