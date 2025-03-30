/*
Binance Cfutures API

OpenAPI specification for Binance cryptocurrency exchange - Cfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
	"bytes"
	"net/url"
)

type KeyType string

const (
	KeyTypeHMAC     KeyType = "HMAC"
	KeyTypeRSA              = "RSA"
	KeyTypeED25519          = "ED25519"
)

func NewAuth(apiKey string) *Auth {
	return &Auth{
		APIKey:  apiKey,
	}
}

// Auth provides Binance API key based authentication to a request passed via context using ContextBinanceAuth
type Auth struct {
	APIKey           string
    KeyType          KeyType
    PrivateKeyPath   string    // The path to the private key.
    PrivateKeyReader io.Reader // provide the APIKey using the types which implement io.Reader interface.
    Passphrase       string    // The passphrase to decrypt the private key, if the key is encrypted.
    privateKey       crypto.PrivateKey
    secretKey        string
}

func (b *Auth) SetSecretKey(secretKey string) {
	b.secretKey = secretKey
	b.KeyType = KeyTypeHMAC
}

// SetPrivateKey accepts a private key string and sets it.
func (b *Auth) SetPrivateKey(privateKey []byte) error {
	return b.parsePrivateKey(privateKey)
}

// ContextWithValue validates the Auth configuration parameters and returns a context
// suitable for HTTP signature. An error is returned if the Auth configuration parameters
// are invalid.
func (b *Auth) ContextWithValue(ctx context.Context) (context.Context, error) {
	if b.KeyType == "" {
		if (len(b.PrivateKeyPath) == 0 && b.PrivateKeyReader == nil) && b.privateKey == nil {
			return nil, fmt.Errorf("private key path must be specified")
		}
		if len(b.PrivateKeyPath) > 0 && b.PrivateKeyReader != nil {
			return nil, fmt.Errorf("Specify only one of PrivateKeyPath or PrivateKeyReader")
		}
		if err := b.loadPrivateKey(); err != nil {
			return nil, err
		}
	}
	return context.WithValue(ctx, ContextBinanceAuth, *b), nil
}

// loadPrivateKey reads the private key from the file specified in the Binance Auth.
// The key is loaded only when privateKey is not already set.
func (b *Auth) loadPrivateKey() (err error) {
	if b.privateKey != nil {
		return nil
	}
	var priv []byte
	keyReader := b.PrivateKeyReader
	if keyReader == nil {
		var file *os.File
		file, err = os.Open(b.PrivateKeyPath)
		if err != nil {
			return fmt.Errorf("cannot load private key '%s'. Error: %v", b.PrivateKeyPath, err)
		}
		keyReader = file
		defer func() {
			err = file.Close()
		}()
	}
	priv, err = io.ReadAll(keyReader)
	if err != nil {
		return err
	}
	return b.parsePrivateKey(priv)
}

// parsePrivateKey decodes privateKey byte array to crypto.PrivateKey type.
func (b *Auth) parsePrivateKey(priv []byte) error {
	pemBlock, _ := pem.Decode(priv)
	if pemBlock == nil {
		// No PEM data has been found.
		return fmt.Errorf("file '%s' does not contain PEM data", b.PrivateKeyPath)
	}
	var privKey []byte
	var err error
	if x509.IsEncryptedPEMBlock(pemBlock) {
		// The PEM data is encrypted.
		privKey, err = x509.DecryptPEMBlock(pemBlock, []byte(b.Passphrase))
		if err != nil {
			// Failed to decrypt PEM block. Because of deficiencies in the encrypted-PEM format,
			// it's not always possible to detect an incorrect password.
			return err
		}
	} else {
		privKey = pemBlock.Bytes
	}
	switch pemBlock.Type {
	case "RSA PRIVATE KEY":
		if b.privateKey, err = x509.ParsePKCS1PrivateKey(privKey); err != nil {
			return err
		}
		b.KeyType = KeyTypeRSA
	case "EC PRIVATE KEY", "PRIVATE KEY":
		// https://tools.ietf.org/html/rfc5915 section 4.
		if b.privateKey, err = x509.ParsePKCS8PrivateKey(privKey); err != nil {
			return err
		}
		if _, ok := b.privateKey.(*rsa.PrivateKey); ok {
			b.KeyType = KeyTypeRSA
		} else if _, ok := b.privateKey.(ed25519.PrivateKey); ok {
			b.KeyType = KeyTypeED25519
		} else {
			return fmt.Errorf("private key '%s' is not supported", pemBlock.Type)
		}
	default:
		return fmt.Errorf("key '%s' is not supported", pemBlock.Type)
	}
	return nil
}

func (b *Auth) Sign(r *http.Request) (err error) {
	q := r.URL.Query()
	queryString := q.Encode()
	// FIXME: binance uses @ for '%40' in the query string, so we need to replace it with %40 as a workaround
	queryString = strings.ReplaceAll(queryString, "%40", "@")
	var dataToSign []byte
	if r.Body != nil {
		reader, err := r.GetBody()
		if err != nil {
			return fmt.Errorf("failed to get request body: %s", err)
		}
		bodyData, err := io.ReadAll(reader)
		if err != nil {
			return fmt.Errorf("failed to read request body: %s", err)
		}
		// FIXME: binance uses @ for '%40' in the request body, so we need to replace it with %40 as a workaround
		bodyData = bytes.ReplaceAll(bodyData, []byte("%40"), []byte("@"))
		// If we have both query parameters and body data, concatenate them properly
		if queryString != "" {
			dataToSign = []byte(queryString)
			if len(bodyData) > 0 {
				dataToSign = append(dataToSign, bodyData...)
			}
		} else {
			dataToSign = bodyData
		}
	} else {
		// No body, just use query parameters
		dataToSign = []byte(queryString)
	}
	var signature string
	switch b.KeyType {
	case KeyTypeHMAC:
		mac := hmac.New(sha256.New, []byte(b.secretKey))
		_, err = mac.Write(dataToSign)
		if err != nil {
			return err
		}
		signature = fmt.Sprintf("%x", (mac.Sum(nil)))
	case KeyTypeRSA:
		privKey, ok := b.privateKey.(*rsa.PrivateKey)
		if !ok {
			return fmt.Errorf("invalid RSA privateKey")
		}
		hashed := sha256.Sum256(dataToSign)
		bs, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])
		if err != nil {
			return err
		}
		signature = base64.StdEncoding.EncodeToString(bs)
	case KeyTypeED25519:
		privKey, ok := b.privateKey.(ed25519.PrivateKey)
		if !ok {
			return fmt.Errorf("invalid Ed25519 privateKey")
		}
		signature = base64.StdEncoding.EncodeToString(ed25519.Sign(privKey, dataToSign))
	default:
	    return fmt.Errorf("unsupported binance key type: %s", b.KeyType)
	}
	// Make sure signature is added to the end of the query string
	v := url.Values{}
	v.Add("signature", signature)
	if queryString != "" {
		r.URL.RawQuery = fmt.Sprintf("%s&%s", queryString, v.Encode())
	} else {
		r.URL.RawQuery = v.Encode()
	}
	return nil
}
