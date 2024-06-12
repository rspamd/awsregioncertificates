// awsregioncertificates provides embedded region certificates for AWS
// and can be used to validate instance IDs
package awsregioncertificates

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

var (
	// ErrNoKey is returned if an error is encountered processing embedded data
	ErrNoKey         = errors.New("failed to extract key")

	// ErrRegionUnknown is returned if an unrecognised region was specified
	ErrRegionUnknown = errors.New("region was unrecognized")
)

// RegionKeys is a map of regions (as strings) to *rsa.PublicKey
type RegionKeys struct {
	keys map[string]*rsa.PublicKey
}

// ValidateID validates a given signature against a given instance ID using the key of the specified region.
func (r *RegionKeys) ValidateID(region string, iib []byte, sig []byte) error {
	key, ok := r.keys[region]
	if !ok {
		return ErrRegionUnknown
	}
	decodedSig := make([]byte, base64.StdEncoding.DecodedLen(len(sig)))
	_, err := base64.StdEncoding.Decode(decodedSig, sig)
	if err != nil {
		return err
	}
	decodedSig, _ = bytes.CutSuffix(decodedSig, []byte{00}) // 🤔
	h := sha256.Sum256(iib)
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, h[:], decodedSig)
	if err != nil {
		return err
	}
	return nil
}


// GetRegionKeys returns a map of regions to *rsa.PublicKey
// or an error in case the embedded data could not be processed.
func GetRegionKeys() (RegionKeys, error) {
	r := RegionKeys{
		keys: make(map[string]*rsa.PublicKey),
	}
	for region, certStr := range getRegionCertStrings() {
		block, _ := pem.Decode([]byte(certStr))
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return r, err
		}
		key, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return r, ErrNoKey
		}
		r.keys[region] = key
	}
	return r, nil
}

//go:generate go run generator/main.go
//go:generate go fmt regioncert_gen.go
