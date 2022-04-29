package responses

type ContractCollection struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ContractExternalPriceComponent"`
			ContractItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ContractItem"`
			ContractParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ContractParty"`
		} `json:"results"`
	} `json:"d"`
}
