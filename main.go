package main

import (
	sap_api_caller "sap-api-integrations-reservation-document-reads/SAP_API_Caller"
	"sap-api-integrations-reservation-document-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Reservation_Document_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetReservationDocument(
		inoutSDC.Reservation.Reservation,
		inoutSDC.Reservation.ReservationItem.RecordType,
		inoutSDC.Reservation.ReservationItem.Product,
	)
}
