package provider

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/config/env"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Obj     interface{} `json:"obj"`
}

type AddInboundsParam struct {
	UserId     string `json:"userId"`
	OrderId    string `json:"orderId"`
	ExpiryTime int64  `json:"expiryTime"`
}

type DelInboundsParam struct {
	UserIds []string `json:"userIds"`
}

// todo:test
var (
	TestAddInboundsApi = "http://%s:55555/mises/add_inbounds"
	TestDelInboundsApi = "http://%s:55555/mises/del_inbounds"
)

type MisesXuiClient struct{}

func (cc *MisesXuiClient) AddInbounds(userId, orderId, server string, expiryTime int64) (string, error) {
	param := new(AddInboundsParam)
	param.UserId = userId
	param.OrderId = orderId
	param.ExpiryTime = expiryTime
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return "", err
	}

	bs, err := requestXui(paramBytes, fmt.Sprintf(TestAddInboundsApi, server))
	if err != nil {
		return "", err
	}
	res := new(ApiResponse)
	err = json.Unmarshal(bs, res)
	if err != nil {
		return "", err
	}
	if res.Obj == nil {
		return "", errors.New("null obj")
	}
	link := res.Obj.(string)
	if link == "" {
		return "", errors.New("empty vpn link")
	}
	return link, nil
}

func (cc *MisesXuiClient) DelInbounds(userIds []string, server string) error {
	param := new(DelInboundsParam)
	param.UserIds = userIds
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return err
	}

	bs, err := requestXui(paramBytes, fmt.Sprintf(TestDelInboundsApi, server))
	if err != nil {
		return err
	}
	res := new(ApiResponse)
	err = json.Unmarshal(bs, res)
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("DelInbounds error: %s", res.Msg)
	}
	return nil
}

func requestXui(paramBytes []byte, api string) ([]byte, error) {
	if env.Envs == nil || env.Envs.MisesVpnPrivateKey == "" {
		return nil, errors.New("config error")
	}
	client := &http.Client{Transport: &http.Transport{Proxy: setProxy()}}
	client.Timeout = time.Second * 60
	sig, err := generateSignature(paramBytes)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", api, bytes.NewReader(paramBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Signature", sig)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func generateSignature(messageBytes []byte) (sig string, err error) {
	// check
	if len(messageBytes) == 0 {
		err = errors.New("empty params")
		return
	}

	// Convert the private key from hex to bytes
	privateKeyBytes, err := hex.DecodeString(env.Envs.MisesVpnPrivateKey)
	if err != nil {
		logrus.Error("Error decoding private key:", err)
		return
	}

	// recover the pkcs8 format private key
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	if err != nil {
		logrus.Error("Error parsing private key:", err)
		return
	}

	// Signing the message using the private key
	hash := sha256.Sum256(messageBytes)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash[:])
	if err != nil {
		logrus.Error("Error creating signature:", err)
		return
	}
	sig = base64.StdEncoding.EncodeToString(signature)

	return
}

func setProxy() func(*http.Request) (*url.URL, error) {
	return func(_ *http.Request) (*url.URL, error) {
		return nil, nil
		//return url.Parse("http://127.0.0.1:7078")
	}
}
