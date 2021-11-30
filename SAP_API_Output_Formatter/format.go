package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-reservation-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToReservation(raw []byte, l *logger.Logger) *Reservation {
	pm := &responses.Reservation{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Reservation{
		Reservation                    data.Reservation,
		OrderID                        data.OrderID,
		GoodsMovementType              data.GoodsMovementType,
		CostCenter                     data.CostCenter,
		GoodsRecipientName             data.GoodsRecipientName,
		ReservationDate                data.ReservationDate,
		Customer                       data.Customer,
		WBSElement                     data.WBSElement,
		ControllingArea                data.ControllingArea,
		SalesOrder                     data.SalesOrder,
		SalesOrderItem                 data.SalesOrderItem,
		SalesOrderScheduleLine         data.SalesOrderScheduleLine,
		AssetNumber                    data.AssetNumber,
		AssetSubNumber                 data.AssetSubNumber,
		NetworkNumberForAcctAssgmt     data.NetworkNumberForAcctAssgmt,
		IssuingOrReceivingPlant        data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc   data.IssuingOrReceivingStorageLoc,
		ReservationItem                data.ReservationItem,
		RecordType                     data.RecordType,
		Product                        data.Product,
		RequirementType                data.RequirementType,
		MatlCompRequirementDate        data.MatlCompRequirementDate,
		Plant                          data.Plant,
		ManufacturingOrderOperation    data.ManufacturingOrderOperation,
		GoodsMovementIsAllowed         data.GoodsMovementIsAllowed,
		StorageLocation                data.StorageLocation,
		Batch                          data.Batch,
		DebitCreditCode                data.DebitCreditCode,
		BaseUnit                       data.BaseUnit,
		GLAccount                      data.GLAccount,
		GoodsMovementType              data.GoodsMovementType,
		EntryUnit                      data.EntryUnit,
		QuantityIsFixed                data.QuantityIsFixed,
		CompanyCodeCurrency            data.CompanyCodeCurrency,
		IssuingOrReceivingPlant        data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc   data.IssuingOrReceivingStorageLoc,
		PurchasingDocument             data.PurchasingDocument,
		PurchasingDocumentItem         data.PurchasingDocumentItem,
		Supplier                       data.Supplier,
		ResvnItmRequiredQtyInBaseUnit  data.ResvnItmRequiredQtyInBaseUnit,
		ReservationItemIsFinallyIssued data.ReservationItemIsFinallyIssued,
		ReservationItmIsMarkedForDeltn data.ReservationItmIsMarkedForDeltn,
		ResvnItmRequiredQtyInEntryUnit data.ResvnItmRequiredQtyInEntryUnit,
		ResvnItmWithdrawnQtyInBaseUnit data.ResvnItmWithdrawnQtyInBaseUnit,
		ResvnItmWithdrawnAmtInCCCrcy   data.ResvnItmWithdrawnAmtInCCCrcy,
		GoodsRecipientName             data.GoodsRecipientName,
		UnloadingPointName             data.UnloadingPointName,
		ReservationItemText            data.ReservationItemText,
	}
}
