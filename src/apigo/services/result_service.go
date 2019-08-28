package services

import (
	"../utils"
	"../domains"
	"strconv"
	"net/http"
	"time"
	"fmt"
)

func GetResult(userId string) (*domains.Result,*utils.ApiError) {


	respChan := make(chan domains.Result)
	errChan := make(chan utils.ApiError)
	timeout := time.After(5*time.Second)

	go func() {
		time.Sleep(time.Second*10)
		errChan <- utils.ApiError{
			"TIEMPO AGOTADO",
			http.StatusGatewayTimeout,
		}
		}()


	chanel := make(chan domains.Result,2)


	//var wg sync.WaitGroup
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return &domains.Result{
		}, &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}



	user := domains.User{
		ID: userIdInt,
	}

	errores := user.Get()

	if errores != nil {
		return nil,errores
	}

	country := domains.Country{
		ID: user.CountryID,
	}

	site := domains.Site{
		ID: user.SiteID,
	}



	go country.GetCH(chanel)
	go site.GetCH(chanel)
	resultados2 := domains.Result{
		User:nil,
		Site:nil,
		Country:nil,
	}



	for i:=0;i<2;i++{
		resultados:= <- chanel
		if resultados.Site == nil {
			resultados2.Country = resultados.Country
		}
		if resultados.Country== nil {
			resultados2.Site = resultados.Site
		}
	}




	//wg.Add(2)
	//go country.GetWG(&wg)

	/*
	errores = country.Get()
	if errores != nil {
		return nil,errores
	}*/




	//go site.GetWG(&wg)


	/*
	errores = site.Get()
	if errores != nil {
		return nil,errores
	}*/

	//wg.Wait()


	resp := domains.Result{
		User: &user,
		Country: resultados2.Country,
		Site: resultados2.Site,
	}






	return &resp,nil
}


