package responses

type ToContractParty struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
