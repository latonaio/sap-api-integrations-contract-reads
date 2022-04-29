package responses

type ToContractExternalPriceComponent struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
