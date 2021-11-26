package sap_api_caller

type ReservationDocument struct {
	 ConnectionKey       string `json:"connection_key"`
	 Result              bool   `json:"result"`
	 RedisKey            string `json:"redis_key"`
	 Filepath            string `json:"filepath"`
	 APISchema           string `json:"api_schema"`
	 Reservation         int    `json:"reservation_document"`
	 Deleted             string `json:"deleted"`
}

type ReservationDocumentHeader struct {
	 Reservation                  int    `json:"reservation_document"`     
     OrderID                      string `json:"OrderID"`
	 GoodsMovementType            string `json:"GoodsMovementType"`
	 CostCenter                   string `json:"CostCenter"`
	 GoodsRecipientName           string `json:"GoodsRecipientName"`
	 ReservationDate              string `json:"ReservationDate"`
	 Customer                     string `json:"Customer"`
	 WBSElement                   string `json:"WBSElement"`
	 ControllingArea              string `json:"ControllingArea"`
	 SalesOrder                   string `json:"SalesOrder"`
	 SalesOrderItem               string `json:"SalesOrderItem"`
	 SalesOrderScheduleLine       string `json:"SalesOrderScheduleLine"`
	 AssetNumber                  string `json:"AssetNumber"`
	 AssetSubNumber               string `json:"AssetSubNumber"`
	 NetworkNumberForAcctAssgmt   string `json:"NetworkNumberForAcctAssgmt"`
	 IssuingOrReceivingPlant      string `json:"IssuingOrReceivingPlant"`
	 IssuingOrReceivingStorageLoc string `json:"IssuingOrReceivingStorageLoc"`
}	

type ReservationDocumentItem struct {
	 Reservation                    int    `json:"reservation_document"`
     ReservationItem                int    `json:"ReservationItem"`
     RecordType                     string `json:"RecordType"`
     Product                        string `json:"Product"`
     RequirementType                string `json:"RequirementType"`
     MatlCompRequirementDate        string `json:"MatlCompRequirementDate"`
     Plant                          string `json:"Plant"`
     ManufacturingOrderOperation    int    `json:"ManufacturingOrderOperation"`
     GoodsMovementIsAllowed         string `json:"GoodsMovementIsAllowed"`
     StorageLocation                string `json:"StorageLocation"`
     Batch                          string `json:"Batch"`
     DebitCreditCode                string `json:"DebitCreditCode"`
     BaseUnit                       string `json:"BaseUnit"`
     GLAccount                      string `json:"GLAccount"`
     GoodsMovementType              string `json:"GoodsMovementType"`
     EntryUnit                      string `json:"EntryUnit"`
     QuantityIsFixed                string `json:"QuantityIsFixed"`
     CompanyCodeCurrency            string `json:"CompanyCodeCurrency"`
     IssuingOrReceivingPlant        string `json:"IssuingOrReceivingPlant"`
     IssuingOrReceivingStorageLoc   string `json:"IssuingOrReceivingStorageLoc"`
     PurchasingDocument             string `json:"PurchasingDocument"`
     PurchasingDocumentItem         int    `json:"PurchasingDocumentItem"`
     Supplier                       string `json:"Supplier"`
     ResvnItmRequiredQtyInBaseUnit  float64 `json:"ResvnItmRequiredQtyInBaseUnit"`
     ReservationItemIsFinallyIssued string `json:"ReservationItemIsFinallyIssued"`
     ReservationItmIsMarkedForDeltn string `json:"ReservationItmIsMarkedForDeltn"`
     ResvnItmRequiredQtyInEntryUnit float64 `json:"ResvnItmRequiredQtyInEntryUnit"`
     ResvnItmWithdrawnQtyInBaseUnit float64 `json:"ResvnItmWithdrawnQtyInBaseUnit"`
     ResvnItmWithdrawnAmtInCCCrcy   float64 `json:"ResvnItmWithdrawnAmtInCCCrcy"`
     GoodsRecipientName             string `json:"GoodsRecipientName"`
     UnloadingPointName             string `json:"UnloadingPointName"`
     ReservationItemText            string `json:"ReservationItemText"`
}
