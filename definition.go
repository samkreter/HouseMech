package main

import "encoding/xml"

type Searchresults struct {
	SearchResults 	xml.Name `xml:"SearchResults"`
    Request struct{
		Address			string			`xml:"address"`
		Citystatezip	string 			`xml:"citystatezip"`
	} 	`xml:"request"`
	Message struct{
		Text 			string 			`xml:"text"`
		Code 			int 			`xml:"code"`
	}	`xml:"message"`
	Response struct{
		Results 		[]Result		`xml:"results>result"`
	}	`xml:"response"`
}

type DeepSearchresults struct {
	SearchResults 		xml.Name 		`xml:"SearchResults"`
    Request struct{
		Address			string			`xml:"address"`
		Citystatezip	string 			`xml:"citystatezip"`
	} 	`xml:"request"`
	Message struct{
		Text 			string 			`xml:"text"`
		Code 			int 			`xml:"code"`
	}	`xml:"message"`
	Response struct{
		Results 		[]DeepResult	`xml:"results>result"`
	}	`xml:"response"`
}

type DeepResult struct{
	Zpid 				string 			`xml:"zpid"`
	Links				Links			`xml:"links"` //TODO - FIGURE THIS OUT 
	Address				Address			`xml:"address"`
	Zestimate			Zestimate		`xml:"zestimate"`
	LocalRealEstate		[]Region		`xml:"localRealEstate>region"`
	FIPScounty			int				`xml:"FIPScountry"`
    UseCode				string			`xml:"useCode"`
    TaxAssessmentYear	int				`xml:"taxAssessmentYear"`
    TaxAssessment		float64			`xml:"taxAssessment"`
    YearBuilt			int				`xml:"yearBuilt"`
    LotSizeSqFt			int				`xml:"lotSizeSqFt"`
    FinishedSqFt		int				`xml:"finishedSqFt"`
    Bathrooms			float64			`xml:"bathrooms"`
    Bedrooms			int				`xml:"bedrooms"`
    LastSoldDate		string			`xml:"lastSoldDate"`
	LastSoldPrice struct{
		Value 			int64			`xml:",chardata"`
		Currency		string			`xml:"currency,attr"`
	} `xml:"lastSoldPrice"`
}


type Result struct{
	Zpid 				string 			`xml:"zpid"`
	Links				Links			`xml:"links"`
	Address				Address			`xml:"address"`
	Zestimate			Zestimate		`xml:"zestimate"`
	LocalRealEstate		[]Region		`xml:"localRealEstate>region"`
}

type Links struct{
	Homedetails			string			`xml:"homedetails"`
	Graphsanddata		string			`xml:"graphsanddata"`
	Mapthishome			string			`xml:"mapthishome"`
	Comparables			string			`xml:"comparables"`
}

type Address struct{
	Street 				string			`xml:"street"`
	Zipcode				string			`xml:"zipcode"`
	City				string 			`xml:"city"`
	State				string 			`xml:"state"`
	Lat					string			`xml:"latitude"`
	Long				string			`xml:"longitude"`
}

type Zestimate struct{
	Amount	struct{
		Value			string			`xml:",chardata"`
		Currency		string			`xml:"currency,attr"`
	} `xml:"amount"`

	LastUpdated			string			`xml:"last-updated"` //TODO: change to date filed
	
	OneWeekChange struct {
		Deprecated		bool			`xml:"deprecated,attr"`
	} `xml:"oneWeekChange"` 
	
	Deprecated			bool			`xml:"deprecated"`

	ValueChange struct{
		Value			int				`xml:",chardata"`
		Duration		int				`xml:"duration,attr"`
		Currency		string			`xml:"currency,attr"`
	} `xml:"valueChange"` 	

	ValuationRange		ValuationRange	`xml:"valuationRange"`
	Percentile			int				`xml:"percentile"`
}

type ValuationRange struct{
	Low struct{
		Low				int				`xml:",chardata"`
		Currency		string			`xml:"currency"`
	} `xml:"low"`
	High struct{
		High			int				`xml:",chardata"`
		Currency		string			`xml:"currency"`
	} `xml:"high"`
}

type Region struct {
	Type				string			`xml:"type,attr"`
	Name				string			`xml:"name,attr"`
	Id					string			`xml:"id,attr"`
	ZindexValue			string			`xml:"zindexValue"`
	ZindexOneYearChange	string			`xml:"zindexOneYearChange"`
	Links struct{
		Overview		string			`xml:"overview"`
		ForSaleByOwner	string			`xml:"forSaleByOwner"`
		ForSale			string			`xml:"forSale"`
	} `xml:"links"`
}

type OneWeekChange struct {
	Deprecated			string			`xml:"deprecated,attr"`
}