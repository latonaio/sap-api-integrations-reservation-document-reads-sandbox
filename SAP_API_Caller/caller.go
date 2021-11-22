package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetReservationDocument(Reservation, Product string) {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		c.Reservation(Reservation)
		wg.Done()
	}()
	
	go func() {
		c.ReservationProduct(Reservation, Product)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Reservation(Reservation string) {
	res, err := c.callReservationSrvAPIRequirementReservation("A_ReservationDocumentHeader", Reservation)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) ReservationProduct(Reservation, Product string) {
	res, err := c.callReservationSrvAPIRequirementProduct("A_ReservationDocumentItem", Reservation, Product)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)


func (c *SAPAPICaller) callReservationSrvAPIRequirementReservation(api, Reservation string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_RESERVATION_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Reservation")
	params.Add("$filter", fmt.Sprintf("Reservation eq '%s'", Reservation))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callReservationSrvAPIRequirementProduct(api, Reservation, Product string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_RESERVATION_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Reservation, Product")
	params.Add("$filter", fmt.Sprintf("Reservation eq '%s' and Product eq '%s'", Reservation, Product))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}