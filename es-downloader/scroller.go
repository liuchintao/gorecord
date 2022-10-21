package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

func search(host, index, scrollID string, query []byte) ([]byte, error) {
	url := "http://" + host + "/" + index + "/_search?scroll=1h"
	if scrollID != "" {
		url = "http://" + host + "/_search/scroll"
		query = []byte(fmt.Sprintf(`{"scroll_id": "%s", "scroll": "1h"}`, scrollID))
	}
	log.Println(url, string(query))
	body := bytes.NewBuffer(query)
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("unexpected args number", os.Args)
	}
	host, id := os.Args[1], os.Args[2]
	file, err := os.OpenFile(id+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatal("open file", id+".log", err)
	}
	query := `{
		"size": 5000,
		"query": {
			"term": {
				"kubernetes_labels_id.keyword": {
					"value": "` + id + `"
				}
			}
		}
	}`
	resp, err := search(host, "cambricon-paas-search", "", []byte(query))
	if err != nil {
		log.Fatal("get scroll id", err)
	}
	gjson.GetBytes(resp, "hits.hits.#._source").ForEach(func(key, value gjson.Result) bool {
		if _, err = file.WriteString(value.Raw + "\n"); err != nil {
			log.Println("append value", value, "[error]", err)
		}
		return true
	})

	scrollID := gjson.GetBytes(resp, "_scroll_id")
	for {
		resp, err = search(host, "cambricon-paas-search", scrollID.String(), nil)
		if err != nil {
			log.Fatal("get scroll id", err)
		}
		count := 0
		gjson.GetBytes(resp, "hits.hits.#._source").ForEach(func(key, value gjson.Result) bool {
			count++
			if _, err = file.WriteString(value.Raw + "\n"); err != nil {
				log.Println("append value", value, "[error]", err)
			}
			return true
		})
		log.Println("record", count, "logs, finished?", count == 0)
		if count == 0 {
			return
		}
	}
}
