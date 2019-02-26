package tests

import (
	"github.com/revel/revel/testing"
)

type ZipCodeTest struct {
	testing.TestSuite
}

func (t *ZipCodeTest) Before() {
	println("Set up")
}

func (t *ZipCodeTest) TestThatZipCodeRouteWorks() {

	t.Get("/zipcode/04205050")
	t.AssertStatus(200)
}

func (t *ZipCodeTest) After() {
	println("Tear down")
}

