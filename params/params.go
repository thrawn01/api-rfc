package params

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context/ctxhttp"
	"net/http"
	"net/url"
)

type Resp struct {
	Message string   `json:"message"`
}

const endpoint = "http://mailgun.net/v3/example.com/templates/my-template/versions"

func CreateTemplate(ctx context.Context) {
	p := url.Values{}
	p.Add("tag", "tag1")
	p.Add("template", "product-1")
	p.Add("engine", "go")
	p.Add("comment", "template for product-1")

	resp, err := ctxhttp.Post(ctx, http.DefaultClient, endpoint,
		"", bytes.NewBufferString(p.Encode()))
	if err != nil {
		panic(err)
	}

	var r Resp
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(r); err != nil {
		panic(err)
	}
	fmt.Printf("message: %s\n", r.Message)
}
