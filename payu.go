package payugo

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// PayuIPN function
func PayuIPN(o Options, array map[string]string) string {
	var hashString string
	hashString += strconv.Itoa(utf8.RuneCountInString(array["IpnPid"])) + array["IpnPid"]
	hashString += strconv.Itoa(utf8.RuneCountInString(array["IpnPname"])) + array["IpnPname"]
	hashString += strconv.Itoa(utf8.RuneCountInString(array["IpnDate"])) + array["IpnDate"]
	hashString += strconv.Itoa(utf8.RuneCountInString(array["date"])) + array["date"]

	var signature = signatureCalculate(o.Secret, hashString)

	return fmt.Sprintf("<EPAYMENT>%s|%s</EPAYMENT>", array["date"], signature)
}

// Payment function
func Payment(o Options, request map[string]string) PaymentRes {
	var hashString string
	for _, v := range sortParamList(request) {
		hashString += strconv.Itoa(utf8.RuneCountInString(v)) + v
	}

	var signature = signatureCalculate(o.Secret, hashString)
	request["ORDER_HASH"] = signature

	var v = url.Values{}
	for k, p := range request {
		v.Add(k, p)
	}

	var result = connect("POST", o.URL, v.Encode())

	var payment PaymentRes
	params, _ := payment.Parse(result)

	return params
}

// BinNumber function
func BinNumber(o Options, bin string) (BinNumberRes, error) {
	var timestamp = strconv.FormatInt(time.Now().UTC().Unix(), 10)

	var mac = hmac.New(sha256.New, []byte(o.Secret))
	mac.Write([]byte(o.Merchant + timestamp))
	var signature = hex.EncodeToString(mac.Sum(nil))

	var request = fmt.Sprintf("%s%s?merchant=%s&timestamp=%s&signature=%s",
		o.URL, bin, o.Merchant, timestamp, signature)

	var result = connect("GET", request, "")

	var b BinNumberRes
	err := json.Unmarshal([]byte(result), &b)

	return b, err
}

func connect(method string, url string, request string) string {
	payload := strings.NewReader(request)

	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func signatureCalculate(secret string, hashString string) string {
	var mac = hmac.New(md5.New, []byte(secret))
	mac.Write([]byte(hashString))
	return hex.EncodeToString(mac.Sum(nil))
}

func sortParamList(p map[string]string) []string {
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var ret []string
	for _, k := range keys {
		ret = append(ret, p[k])
	}
	return ret
}
