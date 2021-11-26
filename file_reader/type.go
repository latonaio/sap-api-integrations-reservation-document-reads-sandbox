package file_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Reservation   struct {
		Reservation                    int         `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		ResvnItmRequiredQtyInEntryUnit float64     `json:"quantity"`
		ResvnItmWithdrawnQtyInBaseUnit float64     `json:"picked_quantity"`
		Price                          float64     `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string `json:"api_schema"`
	MaterialCode            string `json:"material_code"`
	Plant_Supplier          string `json:"plant/supplier"`
	Stock                   float64 `json:"stock"`
	GoodsMovementType       string `json:"document_type"`
	Reservation             int    `json:"document_no"`
	ReservationDate         string `json:"planned_date"`
	ValidatedDate           string `json:"validated_date"`
	Deleted                 string `json:"deleted"`
}

type SDC struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	Reservation    struct {
		Reservation                  int    `json:"Reservation"`
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
		ReservationItem              struct {
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
			ResvnItmRequiredQtyInBaseUnit  string `json:"ResvnItmRequiredQtyInBaseUnit"`
			ReservationItemIsFinallyIssued string `json:"ReservationItemIsFinallyIssued"`
			ReservationItmIsMarkedForDeltn string `json:"ReservationItmIsMarkedForDeltn"`
			ResvnItmRequiredQtyInEntryUnit float64 `json:"ResvnItmRequiredQtyInEntryUnit"`
			ResvnItmWithdrawnQtyInBaseUnit float64 `json:"ResvnItmWithdrawnQtyInBaseUnit"`
			ResvnItmWithdrawnAmtInCCCrcy   float64 `json:"ResvnItmWithdrawnAmtInCCCrcy"`
			GoodsRecipientName             string `json:"GoodsRecipientName"`
			UnloadingPointName             string `json:"UnloadingPointName"`
			ReservationItemText            string `json:"ReservationItemText"`
		} `json:"ReservationItem"`
	} `json:"ReservationDocument"`
	APISchema           string `json:"api_schema"`
	Reservation         int    `json:"reservation_document"`
	Deleted             string `json:"deleted"`
}
