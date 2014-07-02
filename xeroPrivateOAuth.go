package xero

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type xeroPrivateOAuth struct {
	privateKey     *rsa.PrivateKey
	certificate    *x509.Certificate
	consumerKey    string
	sharedSecret   string
	coreVersion    string
	payrollVersion string
}

func (x *xeroPrivateOAuth) doSecureRequest(method, requestURI string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest(method, requestURI, nil)
	if err != nil {
		return []byte{}, err
	}

	parsedURL, err := url.Parse(requestURI)
	if err != nil {
		return []byte{}, err
	}

	requestTime := time.Now()

	tsString := fmt.Sprintf("%d", requestTime.Unix())

	nonce := make([]byte, 32, 32)
	rand.Read(nonce)

	nonceString := base64.URLEncoding.EncodeToString(nonce)

	baseURL := parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.RequestURI()

	oaHeader := &OrderedPairs{}
	oaHeader.Add("oauth_consumer_key", x.consumerKey)
	oaHeader.Add("oauth_nonce", nonceString)
	oaHeader.Add("oauth_signature_method", "RSA-SHA1")
	oaHeader.Add("oauth_timestamp", tsString)
	oaHeader.Add("oauth_token", x.consumerKey)
	oaHeader.Add("oauth_version", "1.0")

	parts := make([]string, oaHeader.Len(), oaHeader.Len())

	for i, pair := range oaHeader.GetPairs() {
		parts[i] = percentEscapeLight(pair.K) + "=" + percentEscapeLight(pair.V)
	}

	paramString := strings.Join(parts, "&")

	baseString := method + "&" + percentEscapeLight(baseURL) + "&" + percentEscapeLight(paramString)

	signature, err := x.signRequest([]byte(baseString))
	if err != nil {
		return []byte{}, err
	}

	oaHeader.Add("oauth_signature", signature)

	sigStringParts := make([]string, oaHeader.Len(), oaHeader.Len())

	for i, pair := range oaHeader.GetPairs() {
		sigStringParts[i] = percentEscapeLight(pair.K) + "=\"" + percentEscapeLight(pair.V) + "\""
	}
	oauthHeaderString := "OAuth " + strings.Join(sigStringParts, ",")

	req.Header.Add("Authorization", oauthHeaderString)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	if body != nil {
		noOpCloser := ioutil.NopCloser(body)
		req.Body = noOpCloser
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bodyBytes, nil
}

func (x *xeroPrivateOAuth) signRequest(baseString []byte) (string, error) {
	h := crypto.SHA1.New()
	h.Write(baseString)
	sum := h.Sum(nil)

	sig, err := rsa.SignPKCS1v15(rand.Reader, x.privateKey, crypto.SHA1, sum)
	if err != nil {
		return "", fmt.Errorf("signing request: %s", err.Error())
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	if err != nil {
		return "", fmt.Errorf("signing request: %s", err.Error())
	}
	return sigStr, nil
}

func (x *xeroPrivateOAuth) DoGET(apiPath string) ([]byte, error) {
	resp, err := x.doSecureRequest("GET", "https://api.xero.com/api.xro/2.0/"+apiPath, nil)
	if err != nil {
		return []byte{}, err
	}
	return resp, nil
}
func (x *xeroPrivateOAuth) DoPOST(apiPath string, body []byte) ([]byte, error) {
	bodyReader := &bytes.Buffer{}
	bodyReader.Write(body)
	resp, err := x.doSecureRequest("POST", "https://api.xero.com/api.xro/2.0/"+apiPath, bodyReader)
	if err != nil {
		return []byte{}, err
	}
	return resp, nil
}
func (x *xeroPrivateOAuth) DoPUT(apiPath string, body []byte) ([]byte, error) {
	bodyReader := &bytes.Buffer{}
	bodyReader.Write(body)
	resp, err := x.doSecureRequest("POST", "https://api.xero.com/api.xro/2.0/"+apiPath, bodyReader)
	if err != nil {
		return []byte{}, err
	}
	return resp, nil
}
