package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	ContractCollection struct {
		ObjectID                                              string `json:"ObjectID"`
		ID                                                    string `json:"ID"`
		BuyerID                                               string `json:"BuyerID"`
		MainBusinessProcessVariantTypeCode                    string `json:"MainBusinessProcessVariantTypeCode"`
		ProcessingTypeCode                                    string `json:"ProcessingTypeCode"`
		ContractDeterminationCoveredObjectsRelevanceCode      string `json:"ContractDeterminationCoveredObjectsRelevanceCode"`
		TemplateIndicator                                     bool   `json:"TemplateIndicator"`
		Name                                                  string `json:"Name"`
		CreationDateTime                                      string `json:"CreationDateTime"`
		CreationIdentityUUID                                  string `json:"CreationIdentityUUID"`
		LastChangeDateTime                                    string `json:"LastChangeDateTime"`
		LastChangeIdentityUUID                                string `json:"LastChangeIdentityUUID"`
		Date                                                  string `json:"Date"`
		ValidityStartDate                                     string `json:"ValidityStartDate"`
		ValidityEndDate                                       string `json:"ValidityEndDate"`
		TimeZoneCode                                          string `json:"TimeZoneCode"`
		DocumentLanguageCode                                  string `json:"DocumentLanguageCode"`
		IncludeBusinessPartnerHierarchyIndicator              bool   `json:"IncludeBusinessPartnerHierarchyIndicator"`
		IncludeReleaseAuthorisedPartiesIndicator              bool   `json:"IncludeReleaseAuthorisedPartiesIndicator"`
		ServiceLevelObjectiveID                               string `json:"ServiceLevelObjectiveID"`
		ServiceLevelObjectiveUUID                             string `json:"ServiceLevelObjectiveUUID"`
		ServiceLevelObjectiveName                             string `json:"ServiceLevelObjectiveName"`
		CancellationReasonCode                                string `json:"CancellationReasonCode"`
		InvoiceScheduleStartDate                              int    `json:"InvoiceScheduleStartDate"`
		InvoiceScheduleEndDate                                int    `json:"InvoiceScheduleEndDate"`
		InvoiceScheduleInvoicingInAdvanceIndicator            bool   `json:"InvoiceScheduleInvoicingInAdvanceIndicator"`
		InvoiceScheduleHorizonDateCalculationFunctionCode     string `json:"InvoiceScheduleHorizonDateCalculationFunctionCode"`
		InvoiceScheduleNextInvoiceDateCalculationFunctionCode string `json:"InvoiceScheduleNextInvoiceDateCalculationFunctionCode"`
		ContractLifeCycleStatusCode                           string `json:"ContractLifeCycleStatusCode"`
		ConsistencyStatusCode                                 string `json:"ConsistencyStatusCode"`
		ReplicationProcessingStatusCode                       string `json:"ReplicationProcessingStatusCode"`
		ApprovalStatusCode                                    string `json:"ApprovalStatusCode"`
		ExternalContractLifeCycleStatusCode                   string `json:"ExternalContractLifeCycleStatusCode"`
		ExternalContractReferenceProcessingStatusCode         string `json:"ExternalContractReferenceProcessingStatusCode"`
		ExternalContractInvoiceProcessingStatusCode           string `json:"ExternalContractInvoiceProcessingStatusCode"`
		BuyerPartyID                                          string `json:"BuyerPartyID"`
		BuyerPartyUUID                                        string `json:"BuyerPartyUUID"`
		BuyerPartyName                                        string `json:"BuyerPartyName"`
		BuyerPartyMainContactPartyID                          string `json:"BuyerPartyMainContactPartyID"`
		BuyerPartyMainContactPartyUUID                        string `json:"BuyerPartyMainContactPartyUUID"`
		BuyerPartyMainContactPartyName                        string `json:"BuyerPartyMainContactPartyName"`
		ProductRecipientPartyID                               string `json:"ProductRecipientPartyID"`
		ProductRecipientPartyUUID                             string `json:"ProductRecipientPartyUUID"`
		ProductRecipientPartyName                             string `json:"ProductRecipientPartyName"`
		BillToPartyID                                         string `json:"BillToPartyID"`
		BillToPartyUUID                                       string `json:"BillToPartyUUID"`
		BillToPartyName                                       string `json:"BillToPartyName"`
		PayerPartyID                                          string `json:"PayerPartyID"`
		PayerPartyUUID                                        string `json:"PayerPartyUUID"`
		PayerPartyName                                        string `json:"PayerPartyName"`
		SellerPartyID                                         string `json:"SellerPartyID"`
		SellerPartyUUID                                       string `json:"SellerPartyUUID"`
		SellerPartyName                                       string `json:"SellerPartyName"`
		AdministratorPartyID                                  string `json:"AdministratorPartyID"`
		AdministratorPartyEmployeeID                          string `json:"AdministratorPartyEmployeeID"`
		AdministratorPartyUUID                                string `json:"AdministratorPartyUUID"`
		AdministratorPartyName                                string `json:"AdministratorPartyName"`
		EmployeeResponsiblePartyID                            string `json:"EmployeeResponsiblePartyID"`
		EmployeeResponsiblePartyEmployeeID                    string `json:"EmployeeResponsiblePartyEmployeeID"`
		EmployeeResponsiblePartyUUID                          string `json:"EmployeeResponsiblePartyUUID"`
		EmployeeResponsiblePartyName                          string `json:"EmployeeResponsiblePartyName"`
		ContractingUnitPartyID                                string `json:"ContractingUnitPartyID"`
		ContractingUnitPartyUUID                              string `json:"ContractingUnitPartyUUID"`
		ContractingUnitPartyName                              string `json:"ContractingUnitPartyName"`
		SalesUnitPartyID                                      string `json:"SalesUnitPartyID"`
		SalesUnitPartyUUID                                    string `json:"SalesUnitPartyUUID"`
		SalesUnitPartyName                                    string `json:"SalesUnitPartyName"`
		SalesOrganisationID                                   string `json:"SalesOrganisationID"`
		SalesOrganisationUUID                                 string `json:"SalesOrganisationUUID"`
		SalesOrganisationName                                 string `json:"SalesOrganisationName"`
		DistributionChannelCode                               string `json:"DistributionChannelCode"`
		DivisionCode                                          string `json:"DivisionCode"`
		SalesOfficeID                                         string `json:"SalesOfficeID"`
		SalesOfficeUUID                                       string `json:"SalesOfficeUUID"`
		SalesOfficeName                                       string `json:"SalesOfficeName"`
		SalesGroupID                                          string `json:"SalesGroupID"`
		SalesGroupUUID                                        string `json:"SalesGroupUUID"`
		SalesGroupName                                        string `json:"SalesGroupName"`
		TerritoryID                                           string `json:"TerritoryID"`
		TerritoryUUID                                         string `json:"TerritoryUUID"`
		TerritoryName                                         string `json:"TerritoryName"`
		DeliveryPriorityCode                                  string `json:"DeliveryPriorityCode"`
		IncotermsClassificationCode                           string `json:"IncotermsClassificationCode"`
		IncotermsTransferLocationName                         string `json:"IncotermsTransferLocationName"`
		CurrencyCode                                          string `json:"CurrencyCode"`
		CashDiscountTermsCode                                 string `json:"CashDiscountTermsCode"`
		PriceDateTime                                         string `json:"PriceDateTime"`
		ExternalPriceDocumentCalculationStatusCode            string `json:"ExternalPriceDocumentCalculationStatusCode"`
		ExternalPriceDocumentPricingProcedureCode             string `json:"ExternalPriceDocumentPricingProcedureCode"`
		GrossAmount                                           string `json:"GrossAmount"`
		GrossAmountCurrencyCode                               string `json:"GrossAmountCurrencyCode"`
		NetAmount                                             string `json:"NetAmount"`
		NetAmountCurrencyCode                                 string `json:"NetAmountCurrencyCode"`
		InvoiceScheduleNetAmount                              string `json:"InvoiceScheduleNetAmount"`
		InvoiceScheduleNetAmountCurrencyCode                  string `json:"InvoiceScheduleNetAmountCurrencyCode"`
		RequestExternalPricing                                bool   `json:"RequestExternalPricing"`
		Transfer                                              bool   `json:"Transfer"`
		EntityLastChangedOn                                   string `json:"EntityLastChangedOn"`
		ContractExternalPriceComponent                        struct {
			ObjectID                     string `json:"ObjectID"`
			ParentObjectID               string `json:"ParentObjectID"`
			ContractID                   string `json:"ContractID"`
			MajorLevelOrdinalNumberValue int    `json:"MajorLevelOrdinalNumberValue"`
			MinorLevelOrdinalNumberValue int    `json:"MinorLevelOrdinalNumberValue"`
			TypeCode                     string `json:"TypeCode"`
			Description                  string `json:"Description"`
			CategoryCode                 string `json:"CategoryCode"`
			RateDecimalValue             string `json:"RateDecimalValue"`
			RateMeasureUnitCode          string `json:"RateMeasureUnitCode"`
			RateCurrencyCode             string `json:"RateCurrencyCode"`
			RateBaseDecimalValue         string `json:"RateBaseDecimalValue"`
			RateBaseMeasureUnitCode      string `json:"RateBaseMeasureUnitCode"`
			CalculatedAmount             string `json:"CalculatedAmount"`
			CalculatedAmountCurrencyCode string `json:"CalculatedAmountCurrencyCode"`
			EffectiveIndicator           bool   `json:"EffectiveIndicator"`
			ManuallyChangedIndicator     bool   `json:"ManuallyChangedIndicator"`
			InactivityReasonCode         string `json:"InactivityReasonCode"`
			OriginCode                   string `json:"OriginCode"`
			ContractItem                 struct {
				ObjectID                                              string `json:"ObjectID"`
				ParentObjectID                                        string `json:"ParentObjectID"`
				ContractID                                            string `json:"ContractID"`
				ItemID                                                string `json:"ID"`
				ParentItemID                                          string `json:"ParentItemID"`
				ParentItemUUID                                        string `json:"ParentItemUUID"`
				ProcessingTypeCode                                    string `json:"ProcessingTypeCode"`
				DeterminationRelevanceCode                            string `json:"DeterminationRelevanceCode"`
				ReleaseControlCode                                    string `json:"ReleaseControlCode"`
				Description                                           string `json:"Description"`
				ValidityStartDate                                     string `json:"ValidityStartDate"`
				ValidityEndDate                                       string `json:"ValidityEndDate"`
				WarrantyGoodwillCode                                  string `json:"WarrantyGoodwillCode"`
				InvoiceScheduleStartDate                              string `json:"InvoiceScheduleStartDate"`
				InvoiceScheduleEndDate                                string `json:"InvoiceScheduleEndDate"`
				InvoiceScheduleInvoicingInAdvanceIndicator            bool   `json:"InvoiceScheduleInvoicingInAdvanceIndicator"`
				InvoiceScheduleNextInvoiceDateCalculationFunctionCode string `json:"InvoiceScheduleNextInvoiceDateCalculationFunctionCode"`
				InvoiceScheduleHorizonDateCalculationFunctionCode     string `json:"InvoiceScheduleHorizonDateCalculationFunctionCode"`
				CancellationReasonCode                                string `json:"CancellationReasonCode"`
				ContractLifeCycleStatusCode                           string `json:"ContractLifeCycleStatusCode"`
				ConsistencyStatusCode                                 string `json:"ConsistencyStatusCode"`
				ExternalContractLifeCycleStatusCode                   string `json:"ExternalContractLifeCycleStatusCode"`
				ExternalContractReferenceProcessingStatusCode         string `json:"ExternalContractReferenceProcessingStatusCode"`
				ExternalContractInvoiceProcessingStatusCode           string `json:"ExternalContractInvoiceProcessingStatusCode"`
				ProductID                                             string `json:"ProductID"`
				ProductUUID                                           string `json:"ProductUUID"`
				ProductDescription                                    string `json:"ProductDescription"`
				Quantity                                              string `json:"Quantity"`
				QuantityUnitCode                                      string `json:"QuantityUnitCode"`
				TargetQuantity                                        string `json:"TargetQuantity"`
				TargetQuantityUnitCode                                string `json:"TargetQuantityUnitCode"`
				TotalPeriodReleaseQuantity                            string `json:"TotalPeriodReleaseQuantity"`
				TotalPeriodReleaseQuantityUnitCode                    string `json:"TotalPeriodReleaseQuantityUnitCode"`
				TotalPeriodOpenQuantity                               string `json:"TotalPeriodOpenQuantity"`
				TotalPeriodOpenQuantityUnitCode                       string `json:"TotalPeriodOpenQuantityUnitCode"`
				TargetStatusCode                                      string `json:"TargetStatusCode"`
				DeliveryPriorityCode                                  string `json:"DeliveryPriorityCode"`
				IncotermsClassificationCode                           string `json:"IncotermsClassificationCode"`
				IncotermsTransferLocationName                         string `json:"IncotermsTransferLocationName"`
				CashDiscountTermsCode                                 string `json:"CashDiscountTermsCode"`
				PriceDateTime                                         string `json:"PriceDateTime"`
				NetAmount                                             string `json:"NetAmount"`
				NetAmountCurrencyCode                                 string `json:"NetAmountCurrencyCode"`
				NetPriceAmount                                        string `json:"NetPriceAmount"`
				NetPriceAmountCurrencyCode                            string `json:"NetPriceAmountCurrencyCode"`
				InvoiceScheduleNetAmount                              string `json:"InvoiceScheduleNetAmount"`
				InvoiceScheduleNetAmountCurrencyCode                  string `json:"InvoiceScheduleNetAmountCurrencyCode"`
				ProductRecipientPartyID                               string `json:"ProductRecipientPartyID"`
				ProductRecipientPartyUUID                             string `json:"ProductRecipientPartyUUID"`
				ProductRecipientPartyFormattedName                    string `json:"ProductRecipientPartyFormattedName"`
				BillToPartyID                                         string `json:"BillToPartyID"`
				BillToPartyUUID                                       string `json:"BillToPartyUUID"`
				BillToPartyFormattedName                              string `json:"BillToPartyFormattedName"`
				PayerPartyID                                          string `json:"PayerPartyID"`
				PayerPartyUUID                                        string `json:"PayerPartyUUID"`
				PayerPartyFormattedName                               string `json:"PayerPartyFormattedName"`
				EntityLastChangedOn                                   string `json:"EntityLastChangedOn"`
				ContractParty                                         struct {
					ObjectID                   string `json:"ObjectID"`
					ParentObjectID             string `json:"ParentObjectID"`
					ContractID                 string `json:"ContractID"`
					PartyID                    string `json:"PartyID"`
					PartyEmployeeID            string `json:"PartyEmployeeID"`
					PartyUUID                  string `json:"PartyUUID"`
					PartyName                  string `json:"PartyName"`
					PartyTypeCode              string `json:"PartyTypeCode"`
					RoleCode                   string `json:"RoleCode"`
					MainIndicator              bool   `json:"MainIndicator"`
					CountryCode                string `json:"CountryCode"`
					StateCode                  string `json:"StateCode"`
					CareOfName                 string `json:"CareOfName"`
					AddressLine1               string `json:"AddressLine1"`
					AddressLine2               string `json:"AddressLine2"`
					HouseNumber                string `json:"HouseNumber"`
					Street                     string `json:"Street"`
					AddressLine4               string `json:"AddressLine4"`
					AddressLine5               string `json:"AddressLine5"`
					District                   string `json:"District"`
					City                       string `json:"City"`
					DifferentCity              string `json:"DifferentCity"`
					StreetPostalCode           string `json:"StreetPostalCode"`
					County                     string `json:"County"`
					CompanyPostalCode          string `json:"CompanyPostalCode"`
					POBoxIndicator             bool   `json:"POBoxIndicator"`
					POBox                      string `json:"POBox"`
					POBoxPostalCode            string `json:"POBoxPostalCode"`
					POBoxDeviatingCountryCode  string `json:"POBoxDeviatingCountryCode"`
					POBoxDeviatingStateCode    string `json:"POBoxDeviatingStateCode"`
					POBoxDeviatingCity         string `json:"POBoxDeviatingCity"`
					TimeZoneCode               string `json:"TimeZoneCode"`
					Latitude                   string `json:"Latitude"`
					Longitude                  string `json:"Longitude"`
					Building                   string `json:"Building"`
					Floor                      string `json:"Floor"`
					Room                       string `json:"Room"`
					Phone                      string `json:"Phone"`
					Mobile                     string `json:"Mobile"`
					Fax                        string `json:"Fax"`
					Email                      string `json:"Email"`
					WebSite                    string `json:"WebSite"`
					CorrespondenceLanguageCode string `json:"CorrespondenceLanguageCode"`
					BestReachedByCode          string `json:"BestReachedByCode"`
				} `json:"ContractParty"`
			} `json:"ContractItem"`
		} `json:"ContractExternalPriceComponent"`
	} `json:"ContractCollection"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	ContractCode string   `json:"contract_code"`
	Deleted      bool     `json:"deleted"`
}
