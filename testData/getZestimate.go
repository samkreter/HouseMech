type Zestimate struct {
	Text	string	`xml:"message>text"`
	Code	string	`xml:"message>code"`
	ForSaleByOwner	string	`xml:"response>localRealEstate>region>links>forSaleByOwner"`
	Homedetails	string	`xml:"response>links>homedetails"`
	SchemaLocation	string	`xml:"schemaLocation,attr"`
	Percentile	string	`xml:"response>zestimate>percentile"`
	Region	Region	`xml:"response>localRealEstate>region"`
	Longitude	string	`xml:"response>address>longitude"`
	State	string	`xml:"response>address>state"`
	Low	Low	`xml:"response>zestimate>valuationRange>low"`
	Zpid	string	`xml:"request>zpid"`
	ZindexValue	string	`xml:"response>localRealEstate>region>zindexValue"`
	Zipcode	string	`xml:"response>address>zipcode"`
	Latitude	string	`xml:"response>address>latitude"`
	Zestimate	string	`xml:"Zestimate,attr"`
	City-Id	string	`xml:"response>regions>city-id"`
	Mapthishome	string	`xml:"response>links>mapthishome"`
	OneWeekChange	OneWeekChange	`xml:"response>zestimate>oneWeekChange"`
	ValueChange	ValueChange	`xml:"response>zestimate>valueChange"`
	ForSale	string	`xml:"response>localRealEstate>region>links>forSale"`
	State-Id	string	`xml:"response>regions>state-id"`
	Zipcode-Id	string	`xml:"response>regions>zipcode-id"`
	ZpidResponse	string	`xml:"response>zpid"`
	Street	string	`xml:"response>address>street"`
	City	string	`xml:"response>address>city"`
	Amount	Amount	`xml:"response>zestimate>amount"`
	County-Id	string	`xml:"response>regions>county-id"`
	Comparables	string	`xml:"response>links>comparables"`
	Last-Updated	string	`xml:"response>zestimate>last-updated"`
	Xsi	string	`xml:"xsi,attr"`
	Graphsanddata	string	`xml:"response>links>graphsanddata"`
	High	High	`xml:"response>zestimate>valuationRange>high"`
	Overview	string	`xml:"response>localRealEstate>region>links>overview"`
}

type ValueChange struct {
	Currency	string	`xml:"currency,attr"`
	Duration	string	`xml:"duration,attr"`
	Text	string	`xml:",chardata"`
}
type Low struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type High struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type Region struct {
	Id	string	`xml:"id,attr"`
	Type	string	`xml:"type,attr"`
	Name	string	`xml:"name,attr"`
}
type Amount struct {
	Text	string	`xml:",chardata"`
	Currency	string	`xml:"currency,attr"`
}
type OneWeekChange struct {
	Deprecated	string	`xml:"deprecated,attr"`
}

