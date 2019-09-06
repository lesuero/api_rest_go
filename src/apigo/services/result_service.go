package services

import (
	"../utils"
	"../domains"
	"strconv"
	"net/http"
	"time"
)

func GetResult(userId string) (*domains.Result,*utils.ApiError) {
	userChan := make(chan domains.User)
	userErrChan := make(chan utils.ApiError)
	errChan := make(chan utils.ApiError)


	go func() {
		time.Sleep(time.Second*5)
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

	go user.GetChan(userChan, userErrChan)


	SELECT:
	select {
		case error := <- errChan:
			return nil, &error
		case errores := <- userErrChan:
				return nil,&errores
		case <- userChan:
			break SELECT

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
		select {
		case error := <- errChan:
			return nil, &error
		case resultados:= <- chanel:
			if resultados.Site == nil {
				resultados2.Country = resultados.Country
			}
			if resultados.Country== nil {
				resultados2.Site = resultados.Site
			}
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


