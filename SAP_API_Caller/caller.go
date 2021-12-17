package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-reservation-document-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
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

func (c *SAPAPICaller) AsyncGetReservationDocument(reservation, recordType, product string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(reservation)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(reservation, recordType, product)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(reservation string) {
	data, err := c.callReservationDocumentSrvAPIRequirementHeader("A_ReservationDocumentHeader", reservation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callReservationDocumentSrvAPIRequirementHeader(api, reservation string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_RESERVATION_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, reservation)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(reservation, recordType, product string) {
	data, err := c.callReservationDocumentSrvAPIRequirementItem("A_ReservationDocumentItem", reservation, recordType, product)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callReservationDocumentSrvAPIRequirementItem(api, reservation, recordType, product string) (*sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_RESERVATION_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, reservation, recordType, product)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, reservation string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Reservation eq '%s'", reservation))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, reservation, recordType, product string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Reservation eq '%s' and RecordType eq '%s' and Product eq '%s'", reservation, recordType, product))
	req.URL.RawQuery = params.Encode()
}
