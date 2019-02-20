package payugo

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
	"unicode/utf8"
)

func init() {
	fmt.Println("*** Payu Go Clint v0.4.3")
}

func TestPayuIPN(t *testing.T) {
	var o = Options{
		URL:      "",
		Merchant: "PALJZXGV",
		Secret:   "f*%J7z6_#|5]s7V4[g3]",
	}

	var date = time.Now().UTC().Format("2006-01-02 15:04:05")

	var array = map[string]string{
		"IpnPid":   "109652057",
		"IpnPname": "Lancome La Vie Est Belle Edp 100 Ml Kadın Parfümü",
		"IpnDate":  "20190220161357",
		"date":     date,
	}

	var hashString string
	hashString += strconv.Itoa(utf8.RuneCountInString(array["IpnPid"])) + array["IpnPid"]
	hashString += strconv.Itoa(utf8.RuneCountInString(array["IpnPname"])) + array["IpnPname"]
	hashString += strconv.Itoa(utf8.RuneCountInString(array["IpnDate"])) + array["IpnDate"]
	hashString += strconv.Itoa(utf8.RuneCountInString(array["date"])) + array["date"]

	var signature = signatureCalculate(o.Secret, hashString)

	var test = fmt.Sprintf("<EPAYMENT>%s|%s</EPAYMENT>", date, signature)

	var result = PayuIPN(o, array)

	if result != test {
		t.Errorf("PaymentThreeD testinde hata -> beklenen: %s, bulunan: %s", test, result)
	}
}

func TestPaymentThreeD(t *testing.T) {
	var o = Options{
		URL:      "https://secure.payu.com.tr/order/alu/v3",
		Merchant: "PALJZXGV",
		Secret:   "f*%J7z6_#|5]s7V4[g3]",
	}

	var request = map[string]string{
		"MERCHANT":                     o.Merchant,
		"LANGUAGE":                     "TR",
		"ORDER_REF":                    randomString(1000, 9999),
		"ORDER_DATE":                   time.Now().UTC().Format("2006-01-02 15:04:05"),
		"PAY_METHOD":                   "CCVISAMC",
		"BACK_REF":                     "http://85.98.179.163/payment",
		"PRICES_CURRENCY":              "TRY",
		"SELECTED_INSTALLMENTS_NUMBER": "1",
		"ORDER_SHIPPING":               "5",
		"CLIENT_IP":                    "127.0.0.1",

		"ORDER_PNAME[0]":      "Test Ürünü",
		"ORDER_PCODE[0]":      "Test Kodu",
		"ORDER_PINFO[0]":      "Test Açıklaması",
		"ORDER_PRICE[0]":      "5",
		"ORDER_VAT[0]":        "18",
		"ORDER_PRICE_TYPE[0]": "NET",
		"ORDER_QTY[0]":        "1",

		"ORDER_PNAME[1]":      "Test Ürünü-2",
		"ORDER_PCODE[1]":      "Test Kodu-2",
		"ORDER_PINFO[1]":      "Test Açıklaması-2",
		"ORDER_PRICE[1]":      "15",
		"ORDER_VAT[1]":        "24",
		"ORDER_PRICE_TYPE[1]": "GROSS",
		"ORDER_QTY[1]":        "3",

		"CC_NUMBER": "4355084355084358",
		"EXP_MONTH": "12",
		"EXP_YEAR":  "2019",
		"CC_CVV":    "000",
		"CC_OWNER":  "000",

		"BILL_FNAME":       "Nanoshop",
		"BILL_LNAME":       "E-Ticaret",
		"BILL_EMAIL":       "mail@mail.com",
		"BILL_PHONE":       "02129003711",
		"BILL_FAX":         "02129003711",
		"BILL_ADDRESS":     "Birinci Adres satırı",
		"BILL_ADDRESS2":    "İkinci Adres satırı",
		"BILL_ZIPCODE":     "34000",
		"BILL_CITY":        "ISTANBUL",
		"BILL_COUNTRYCODE": "TR",
		"BILL_STATE":       "Ayazağa",

		"DELIVERY_FNAME":       "Ad",
		"DELIVERY_LNAME":       "Soyad",
		"DELIVERY_EMAIL":       "mail@mail.com",
		"DELIVERY_PHONE":       "02129003711",
		"DELIVERY_COMPANY":     "PayU Ödeme Kuruluşu A.Ş.",
		"DELIVERY_ADDRESS":     "Birinci Adres satırı",
		"DELIVERY_ADDRESS2":    "İkinci Adres satırı",
		"DELIVERY_ZIPCODE":     "34000",
		"DELIVERY_CITY":        "ISTANBUL",
		"DELIVERY_STATE":       "TR",
		"DELIVERY_COUNTRYCODE": "Ayazağa",
	}

	var result = Payment(o, request)

	fmt.Println("--> Redirect URL: ", result.URL3Ds)

	if result.Status != "SUCCESS" {
		t.Errorf("PaymentThreeD testinde hata -> beklenen: %s, bulunan: %s", "SUCCESS", result.ReturnMessage)
	}

	if result.ReturnCode != "3DS_ENROLLED" {
		t.Errorf("PaymentThreeD testinde hata -> beklenen: %s, bulunan: %s", "SUCCESS", result.ReturnMessage)
	}
}

func TestPaymentNonThreeD(t *testing.T) {
	var o = Options{
		URL:      "https://secure.payu.com.tr/order/alu/v3",
		Merchant: "OPU_TEST",
		Secret:   "SECRET_KEY",
	}

	var request = map[string]string{
		"MERCHANT":                     "OPU_TEST",
		"LANGUAGE":                     "TR",
		"ORDER_REF":                    randomString(1000, 9999),
		"ORDER_DATE":                   time.Now().UTC().Format("2006-01-02 15:04:05"),
		"PAY_METHOD":                   "CCVISAMC",
		"BACK_REF":                     "http://www.backref.com.tr",
		"PRICES_CURRENCY":              "TRY",
		"SELECTED_INSTALLMENTS_NUMBER": "1",
		"ORDER_SHIPPING":               "5",
		"CLIENT_IP":                    "127.0.0.1",

		"ORDER_PNAME[0]":      "Test Ürünü",
		"ORDER_PCODE[0]":      "Test Kodu",
		"ORDER_PINFO[0]":      "Test Açıklaması",
		"ORDER_PRICE[0]":      "5",
		"ORDER_VAT[0]":        "18",
		"ORDER_PRICE_TYPE[0]": "NET",
		"ORDER_QTY[0]":        "1",

		"ORDER_PNAME[1]":      "Test Ürünü-2",
		"ORDER_PCODE[1]":      "Test Kodu-2",
		"ORDER_PINFO[1]":      "Test Açıklaması-2",
		"ORDER_PRICE[1]":      "15",
		"ORDER_VAT[1]":        "24",
		"ORDER_PRICE_TYPE[1]": "GROSS",
		"ORDER_QTY[1]":        "3",

		"CC_NUMBER": "4355084355084358",
		"EXP_MONTH": "12",
		"EXP_YEAR":  "2019",
		"CC_CVV":    "000",
		"CC_OWNER":  "000",

		"BILL_FNAME":       "Nanoshop",
		"BILL_LNAME":       "E-Ticaret",
		"BILL_EMAIL":       "mail@mail.com",
		"BILL_PHONE":       "02129003711",
		"BILL_FAX":         "02129003711",
		"BILL_ADDRESS":     "Birinci Adres satırı",
		"BILL_ADDRESS2":    "İkinci Adres satırı",
		"BILL_ZIPCODE":     "34000",
		"BILL_CITY":        "ISTANBUL",
		"BILL_COUNTRYCODE": "TR",
		"BILL_STATE":       "Ayazağa",

		"DELIVERY_FNAME":       "Ad",
		"DELIVERY_LNAME":       "Soyad",
		"DELIVERY_EMAIL":       "mail@mail.com",
		"DELIVERY_PHONE":       "02129003711",
		"DELIVERY_COMPANY":     "PayU Ödeme Kuruluşu A.Ş.",
		"DELIVERY_ADDRESS":     "Birinci Adres satırı",
		"DELIVERY_ADDRESS2":    "İkinci Adres satırı",
		"DELIVERY_ZIPCODE":     "34000",
		"DELIVERY_CITY":        "ISTANBUL",
		"DELIVERY_STATE":       "TR",
		"DELIVERY_COUNTRYCODE": "Ayazağa",
	}

	var result = Payment(o, request)

	if result.Status != "SUCCESS" {
		t.Errorf("PayThreeD testinde hata -> beklenen: %s, bulunan: %s", "SUCCESS", result.ReturnMessage)
	}
}

func TestBinNumber(t *testing.T) {
	var o = Options{
		URL:      "https://secure.payu.com.tr/api/card-info/v1/",
		Merchant: "OPU_TEST",
		Secret:   "SECRET_KEY",
	}
	var binNumber, _ = BinNumber(o, "454360")

	if statusCode := binNumber.Meta.Status.Code; statusCode != 0 {
		t.Errorf("BinNumber testinde hata -> beklenen: %s, bulunan: %s",
			"200", binNumber.Meta.Status.Message)
	}
}

func randomString(min int, max int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max-min) + min)
}
