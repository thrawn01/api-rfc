package json

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context/ctxhttp"
	"net/http"
)

type TemplateVersion struct {
	Tag       string         `json:"tag"`
	Template  string         `json:"template,omitempty"`
	Engine    string         `json:"engine"`
	Comment   string         `json:"comment"`
}

type Resp struct {
	Message string   `json:"message"`
}

const endpoint = "http://mailgun.net/v3/example.com/templates/my-template/versions"

func CreateTemplate(ctx context.Context) {
	v := TemplateVersion{
		Tag:      "tag1",
		Template: "product-1",
		Engine:   "go",
		Comment:  "template for product-1",
	}

	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	resp, err := ctxhttp.Post(ctx, http.DefaultClient, endpoint,
		"json/application", bytes.NewBuffer(b))
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