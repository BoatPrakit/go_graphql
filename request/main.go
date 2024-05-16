package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	payload := `
	{
		"query": "# Welcome to Altair GraphQL Client.\n# You can send your request using CmdOrCtrl + Enter.\n\n# Enter your graphQL query here.\n\nquery {\n  hello\n}",
		"variables": {},
		"operationName": null
	}	
	`
	// req, err := http.NewRequest(http.MethodPost, "http://localhost:4000/graphql", strings.NewReader(payload))
	// if err != nil {
	// 	fmt.Errorf("Error: ", err)
	// }
	// req.Header.Add("Content-Type", "application/json")
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Errorf("Error: ", err)
	// }
	// defer res.Body.Close()
	// if err != nil {
	// 	fmt.Errorf("Error: %v", err)
	// }
	res, _ := http.Post("http://localhost:4000/graphql", "application/json", strings.NewReader(payload))
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
