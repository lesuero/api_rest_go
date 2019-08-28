package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyID         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}

func (country *Country) Get() *utils.ApiError{
	if country.ID == "" {
		return &utils.ApiError{
			Message: "Site ID empty",
			Status: http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s",utils.UrlCountries,country.ID)
	response,err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	//se puede crear la variable dentro del if aunque exista. se debe usarla en el if
	if err := json.Unmarshal(data, &country);err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	return nil
}

func (country *Country) GetWG(wg *sync.WaitGroup) *utils.ApiError{
	if country.ID == "" {
		return &utils.ApiError{
			Message: "Site ID empty",
			Status: http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s",utils.UrlCountries,country.ID)
	response,err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	//se puede crear la variable dentro del if aunque exista. se debe usarla en el if
	if err := json.Unmarshal(data, &country);err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	wg.Done()
	return nil
}

func (country *Country) GetCH(channel chan Result) *utils.ApiError{
	if country.ID == "" {
		return &utils.ApiError{
			Message: "Site ID empty",
			Status: http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s",utils.UrlCountries,country.ID)
	response,err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	//se puede crear la variable dentro del if aunque exista. se debe usarla en el if
	if err := json.Unmarshal(data, &country);err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	resultado := Result{
		User:nil,
		Site:nil,
		Country:country,
	}
	channel <- resultado
	return nil
}
