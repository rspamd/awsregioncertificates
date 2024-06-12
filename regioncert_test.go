package awsregioncertificates

import (
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
