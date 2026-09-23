package cloud_test

import (
	"testing"

	"github.com/Reflective-Technology/cloud"
)

func equal(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected Mode() to return %s, but got %s", expected, actual)
	}
}

func TestAWS(t *testing.T) {
	cloud.SetMode(cloud.AWS)
	equal(t, cloud.AWS, cloud.Mode())
}

func TestGCP(t *testing.T) {
	cloud.SetMode(cloud.GCP)
	equal(t, cloud.GCP, cloud.Mode())
}

func TestEmpty(t *testing.T) {
	cloud.SetMode("")
	equal(t, cloud.UNSPECIFIED, cloud.Mode())
}

func TestAzure(t *testing.T) {
	cloud.SetMode(cloud.AZURE)
	equal(t, cloud.AZURE, cloud.Mode())
}

func TestOnPremise(t *testing.T) {
	cloud.SetMode(cloud.ON_PREMISE)
	equal(t, cloud.ON_PREMISE, cloud.Mode())
}

func TestUnknown(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cloud.SetMode("unknown")
}
