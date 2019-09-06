package circuitbraker

import (
	"time"
	"net/http"
	"io/ioutil"
	"../utils"
	"fmt"
)

var(
	CircuitResult = CircuitBreaker{
		State:"Closed",
		Canterrores:0,
	}
)


type CircuitBreaker struct {
	State string
	Url string
	Canterrores int
}

/*
func (circuit *CircuitBreaker) StartCircuit(url string){
	circuit.State = "Closed"
	circuit.Canterrores = 0
	circuit.Url = url
}*/

func (circ *CircuitBreaker) AddError() {
	circ.Canterrores++

	if circ.Canterrores == 3{

		go func(){

			FOR:
			for {
				circ.State = "Open"
				time.Sleep(time.Second*7)
				circ.State = "HalfOpen"
				response,_ := http.Get(utils.UrlSitePing)
				response1,_:= http.Get(utils.UrlCountriesPing)
				response2,_ := http.Get(utils.UrlUsersPing)
				data,_ := ioutil.ReadAll(response.Body)
				data1,_ := ioutil.ReadAll(response1.Body)
				data2,_ := ioutil.ReadAll(response2.Body)
				fmt.Println(string(data2))
				if string(data2) == string(data1) && string(data1) == string(data) && string(data) == "Pong" {
					circ.State = "Closed"
					circ.Canterrores = 0
					break FOR
				}
			}
		}()
	}

}

