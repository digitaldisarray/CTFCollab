package tests

import (
	"github.com/go-resty/resty/v2"
)

var (
	client *resty.Client
)

func setup() {
	client = resty.New()
}
