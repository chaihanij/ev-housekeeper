package reports

type B2B17Header struct {
	DocType        string //4 chars
	SalesOrg       string //4 chars
	DC             string //2 chars
	DV             string //2 chars
	Contract       string //10 chars
	RefNo          string //10 chars
	SoldTo         string //10 chars
	ShipTo         string //8 chars
	Plant          string //4 chars
	SalesRep       string //8 chars
	PONumber       string //35 chars
	ReqDeliverDate string //10 chars
	PricingDate    string //10 chars
	PricingTime    string //5 chars
	ProcessFlag    string // 1 chars
	DOCDate        string //10 chars, 01.12.2020
	SentDate       string //10 chars
	Time           string //5 chars
	PaymentTerm    string //4 chars
	PODate         string //10 chars
	//Ref string
	PaymentMethod string //1 chars
	SalesDistrict string //6 chars
}
