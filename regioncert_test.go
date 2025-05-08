package awsregioncertificates

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"testing"
)

var (
	testIDs = map[[3]string]error{
		[3]string{"iid0.json", "iid0.sig", "us-east-1"}: nil,
		[3]string{"iid1.json", "iid1.sig", "us-east-1"}: nil,
	}
)

func ExampleRegionKeys_ValidateID() {
	regionKeys, err := GetRegionKeys()
	if err != nil {
		fmt.Println(fmt.Sprintf("error loading region keys: %v", err))
		return
	}
	iid := []byte(`{
  "accountId" : "975050371289",
  "architecture" : "x86_64",
  "availabilityZone" : "us-east-1b",
  "billingProducts" : null,
  "devpayProductCodes" : null,
  "marketplaceProductCodes" : null,
  "imageId" : "ami-0c7217cdde317cfec",
  "instanceId" : "i-0b02d936754a6d637",
  "instanceType" : "t2.micro",
  "kernelId" : null,
  "pendingTime" : "2024-02-15T14:12:11Z",
  "privateIp" : "172.31.17.154",
  "ramdiskId" : null,
  "region" : "us-east-1",
  "version" : "2017-09-30"
}`)
	iidSig := []byte(`OQTgfPTsc7hXR+3OWP7dk7qY1S1RNGvVvoVPzn/WogAqJpGBtei2pSx3OfZ7F1PDMpClQswDcF9N
iZmPD09xyJSrSRwYvx8SFoBWzUXS1hd9T1ZxpqWtloe/k//YmK7h9f7rjuT3/CxDDCWrbsKp8F0N
ck+YPKGzD+dtxEm6g1g=`)
	err = regionKeys.ValidateID("us-east-1", iid, iidSig)
	if err != nil {
		fmt.Println(fmt.Sprintf("error validating instance ID: %v", err))
		return
	}
	fmt.Println("validated OK")
	// Output:
	// validated OK
}

func TestValidateID(t *testing.T) {
	regionKeys, err := GetRegionKeys()
	if err != nil {
		t.Fatal(err)
	}
	_, ourFile, _, _ := runtime.Caller(0)
	testDataDir := path.Join(path.Dir(ourFile), "testdata")

	for testFiles, expectErr := range testIDs {

		f, err := os.Open(path.Join(testDataDir, testFiles[0]))
		if err != nil {
			t.Fatal(err)
		}
		iib, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		err = f.Close()
		if err != nil {
			t.Fatal(err)
		}
		f, err = os.Open(path.Join(testDataDir, testFiles[1]))
		if err != nil {
			t.Fatal(err)
		}
		var sig []byte
		sig, err = io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		err = regionKeys.ValidateID(testFiles[2], iib, sig)
		if err != expectErr {
			t.Fatalf("%s: %s (actual) != %s (expected)", testFiles[0], err, expectErr)
		}
	}
}
