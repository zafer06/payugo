package payugo

import (
	"encoding/xml"
)

// Options struct
type Options struct {
	URL      string
	Merchant string
	Secret   string
}

// BinNumberRes struct
type BinNumberRes struct {
	Meta struct {
		Status struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"status"`
		Response struct {
			HTTPCode    int    `json:"httpCode"`
			HTTPMessage string `json:"httpMessage"`
		} `json:"response"`
	} `json:"meta"`
	CardBinInfo struct {
		BinType       string        `json:"binType"`
		BinIssuer     string        `json:"binIssuer"`
		CardType      string        `json:"cardType"`
		CardProfile   string        `json:"cardProfile"`
		Country       string        `json:"country"`
		Program       string        `json:"program"`
		Installments  []interface{} `json:"installments"`
		PaymentMethod string        `json:"paymentMethod"`
	} `json:"cardBinInfo"`
}

// PaymentRes struct
type PaymentRes struct {
	RefNo           string `xml:"REFNO"`
	Alias           string `xml:"ALIAS"`
	Status          string `xml:"STATUS"`
	ReturnCode      string `xml:"RETURN_CODE"`
	ReturnMessage   string `xml:"RETURN_MESSAGE"`
	Date            string `xml:"DATE"`
	URL3Ds          string `xml:"URL_3DS"`
	Amount          string `xml:"AMOUNT"`
	Currency        string `xml:"CURRENCY"`
	InstallmentsNo  int    `xml:"INSTALLMENTS_NO"`
	CardProgramName string `xml:"CARD_PROGRAM_NAME"`
	OrderRef        string `xml:"ORDER_REF"`
	AuthCode        string `xml:"AUTH_CODE"`
	Rrn             string `xml:"RRN"`
	ErrorMessage    string `xml:"ERRORMESSAGE"`
	ProcReturnCode  string `xml:"PROCRETURNCODE"`
	BankMerchantID  string `xml:"BANK_MERCHANT_ID"`
	Pan             string `xml:"PAN"`
	ExpYear         string `xml:"EXPYEAR"`
	ExpMonth        string `xml:"EXPMONTH"`
	ClientID        string `xml:"CLIENTID"`
	HostRefNum      string `xml:"HOSTREFNUM"`
	Oid             string `xml:"OID"`
	Response        string `xml:"RESPONSE"`
	TerminalBank    string `xml:"TERMINAL_BANK"`
	MdStatus        string `xml:"MDSTATUS"`
	MdErrorMsg      string `xml:"MDERRORMSG"`
	TxStatus        string `xml:"TXSTATUS"`
	Xid             string `xml:"XID"`
	Eci             string `xml:"ECI"`
	Cavv            string `xml:"CAVV"`
	TransID         string `xml:"TRANSID"`
	Hash            string `xml:"HASH"`
}

// Parse function
func (p *PaymentRes) Parse(data string) (PaymentRes, error) {
	var payment PaymentRes
	err := xml.Unmarshal([]byte(data), &payment)

	return payment, err
}
