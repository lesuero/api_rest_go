package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type User struct {
	ID               int    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	Address          struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
	UserType         string      `json:"user_type"`
	Tags             []string    `json:"tags"`
	Logo             interface{} `json:"logo"`
	Points           int         `json:"points"`
	SiteID           string      `json:"site_id"`
	Permalink        string      `json:"permalink"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
}


func (user *User) Get() *utils.ApiError{
	//id := strconv.FormatInt(int64(user.ID), 10)
	id := strconv.Itoa(user.ID)
	if id == "" {
		return &utils.ApiError{
			Message: "Site ID empty",
			Status: http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s",utils.UrlUsers,id)
	//fmt.Println(url)
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
	if err := json.Unmarshal(data, &user);err != nil {
		//fmt.Println(response)

		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}
func (user *User) GetChan(usersChan chan User,errorChan chan utils.ApiError){
	//id := strconv.FormatInt(int64(user.ID), 10)
	id := strconv.Itoa(user.ID)
	if id == "" {
		errorChan <-  utils.ApiError{
			Message: "Site ID empty",
			Status: http.StatusBadRequest,
		}
		return
	}
	url := fmt.Sprintf("%s%s",utils.UrlUsers,id)
	//fmt.Println(url)
	response,err := http.Get(url)
	if err != nil {
		errorChan <- utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorChan <- utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
		return
	}

	//se puede crear la variable dentro del if aunque exista. se debe usarla en el if
	if err := json.Unmarshal(data, &user);err != nil {
		//fmt.Println(response)

		errorChan <- utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	usuarioOk := User{
	}
	usersChan <- usuarioOk

	return
}