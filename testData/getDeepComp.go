




type Comps struct {
	LastSoldDate	string	`xml:"response>properties>principal>lastSoldDate"`
	BedroomsCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>bedrooms"`
	LongitudeAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>longitude"`
	Amount	Amount	`xml:"response>properties>principal>zestimate>amount"`
	OneWeekChangeZestimateCompComparablesPropertiesResponse	[]OneWeekChange	`xml:"response>properties>comparables>comp>zestimate>oneWeekChange"`
	Code	string	`xml:"message>code"`
	Overview	string	`xml:"response>properties>principal>localRealEstate>region>links>overview"`
	High	High	`xml:"response>properties>principal>zestimate>valuationRange>high"`
	LastSoldPrice	LastSoldPrice	`xml:"response>properties>principal>lastSoldPrice"`
	Longitude	string	`xml:"response>properties>principal>address>longitude"`
	BathroomsCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>bathrooms"`
	AmountZestimateCompComparablesPropertiesResponse	[]Amount	`xml:"response>properties>comparables>comp>zestimate>amount"`
	TaxAssessmentCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>taxAssessment"`
	SchemaLocation	string	`xml:"schemaLocation,attr"`
	Text	string	`xml:"message>text"`
	Region	Region	`xml:"response>properties>principal>localRealEstate>region"`
	LatitudeAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>latitude"`
	LowValuationRangeZestimateCompComparablesPropertiesResponse	[]Low	`xml:"response>properties>comparables>comp>zestimate>valuationRange>low"`
	YearBuiltCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>yearBuilt"`
	Homedetails	string	`xml:"response>properties>principal>links>homedetails"`
	TaxAssessment	string	`xml:"response>properties>principal>taxAssessment"`
	YearBuilt	string	`xml:"response>properties>principal>yearBuilt"`
	TotalRooms	[]string	`xml:"response>properties>comparables>comp>totalRooms"`
	Zpid	string	`xml:"request>zpid"`
	Graphsanddata	string	`xml:"response>properties>principal>links>graphsanddata"`
	Bedrooms	string	`xml:"response>properties>principal>bedrooms"`
	Zipcode	string	`xml:"response>properties>principal>address>zipcode"`
	ForSaleLinksRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>links>forSale"`
	TaxAssessmentYearCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>taxAssessmentYear"`
	ForSaleByOwner	string	`xml:"response>properties>principal>localRealEstate>region>links>forSaleByOwner"`
	ValueChange	ValueChange	`xml:"response>properties>principal>zestimate>valueChange"`
	ZpidPrincipalPropertiesResponse	string	`xml:"response>properties>principal>zpid"`
	OneWeekChange	OneWeekChange	`xml:"response>properties>principal>zestimate>oneWeekChange"`
	LastSoldPriceCompComparablesPropertiesResponse	[]LastSoldPrice	`xml:"response>properties>comparables>comp>lastSoldPrice"`
	ForSaleByOwnerLinksRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>links>forSaleByOwner"`
	StreetAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>street"`
	Low	Low	`xml:"response>properties>principal>zestimate>valuationRange>low"`
	State	string	`xml:"response>properties>principal>address>state"`
	Last-UpdatedZestimateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>zestimate>last-updated"`
	GraphsanddataLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>graphsanddata"`
	FinishedSqFtCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>finishedSqFt"`
	MapthishomeLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>mapthishome"`
	OverviewLinksRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>links>overview"`
	ZipcodeAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>zipcode"`
	ForSale	string	`xml:"response>properties>principal>localRealEstate>region>links>forSale"`
	Bathrooms	string	`xml:"response>properties>principal>bathrooms"`
	LastSoldDateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>lastSoldDate"`
	HomedetailsLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>homedetails"`
	Comp	[]Comp	`xml:"response>properties>comparables>comp"`
	Mapthishome	string	`xml:"response>properties>principal>links>mapthishome"`
	TaxAssessmentYear	string	`xml:"response>properties>principal>taxAssessmentYear"`
	ZpidCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>zpid"`
	Xsi	string	`xml:"xsi,attr"`
	Count	string	`xml:"request>count"`
	CityAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>city"`
	StateAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>state"`
	LotSizeSqFt	string	`xml:"response>properties>principal>lotSizeSqFt"`
	ZindexValue	string	`xml:"response>properties>principal>localRealEstate>region>zindexValue"`
	HighValuationRangeZestimateCompComparablesPropertiesResponse	[]High	`xml:"response>properties>comparables>comp>zestimate>valuationRange>high"`
	Latitude	string	`xml:"response>properties>principal>address>latitude"`
	Street	string	`xml:"response>properties>principal>address>street"`
	ComparablesLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>comparables"`
	Last-Updated	string	`xml:"response>properties>principal>zestimate>last-updated"`
	Percentile	string	`xml:"response>properties>principal>zestimate>percentile"`
	City	string	`xml:"response>properties>principal>address>city"`
	ValueChangeZestimateCompComparablesPropertiesResponse	[]ValueChange	`xml:"response>properties>comparables>comp>zestimate>valueChange"`
	PercentileZestimateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>zestimate>percentile"`
	Comps	string	`xml:"Comps,attr"`
	Comparables	string	`xml:"response>properties>principal>links>comparables"`
	FinishedSqFt	string	`xml:"response>properties>principal>finishedSqFt"`
	RegionLocalRealEstateCompComparablesPropertiesResponse	[]Region	`xml:"response>properties>comparables>comp>localRealEstate>region"`
	ZindexValueRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>zindexValue"`
	LotSizeSqFtCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>lotSizeSqFt"`
}

type LastSoldPriceCompComparablesPropertiesResponse struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type Region struct {
	Type	string	`xml:"type,attr"`
	Name	string	`xml:"name,attr"`
	Id	string	`xml:"id,attr"`
}
type Low struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type High struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type AmountZestimateCompComparablesPropertiesResponse struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type OneWeekChangeZestimateCompComparablesPropertiesResponse struct {
	Deprecated	string	`xml:"deprecated,attr"`
}
type Comp struct {
	Score	string	`xml:"score,attr"`
}
type OneWeekChange struct {
	Deprecated	string	`xml:"deprecated,attr"`
}
type ValueChange struct {
	Text	string	`xml:",chardata"`
	Currency	string	`xml:"currency,attr"`
	Duration	string	`xml:"duration,attr"`
}
type Amount struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type HighValuationRangeZestimateCompComparablesPropertiesResponse struct {
	Text	string	`xml:",chardata"`
	Currency	string	`xml:"currency,attr"`
}
type RegionLocalRealEstateCompComparablesPropertiesResponse struct {
	Id	string	`xml:"id,attr"`
	Type	string	`xml:"type,attr"`
	Name	string	`xml:"name,attr"`
}
type LastSoldPrice struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type ValueChangeZestimateCompComparablesPropertiesResponse struct {
	Currency	string	`xml:"currency,attr"`
	Duration	string	`xml:"duration,attr"`
	Text	string	`xml:",chardata"`
}
type LowValuationRangeZestimateCompComparablesPropertiesResponse struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}

