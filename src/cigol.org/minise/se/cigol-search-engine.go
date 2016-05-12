package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"cigol.org/minise/utils"
)

var config utils.Config = utils.GetConfig("config.ini")
var esUrl []string = getEsUrl()
var client utils.Client = utils.NewClient(esUrl)

func query(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles(config["page.query"])
		t.Execute(w, nil)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	// by default, the params are not to be parsed
	r.ParseForm()
	//	fmt.Println(r.Form["q"])
	if r.Method == "POST" {
		q := r.FormValue("q")
		ret := make(map[string]interface{})
		ret["Q"] = q
		hits, _ := searchES(q)
		ret["Hits"] = hits

		t, _ := template.ParseFiles(config["page.search"])
		t.Execute(w, ret)
	}
}

func getEsUrl() []string {
	// get the es url
	v, ok := config["es.urls"]
	if ok {
		return strings.Split(v, ",")
	}
	return []string{"http://127.0.0.1:9200"}
}

func searchES(q string) ([]map[string]interface{}, error) {
	//	data := map[string]interface{}{}
	// build the query
	// {"query":{"bool":{"must":[{"wildcard":{"title":"nginx"}}],"must_not":[],"should":[]}},"from":0,"size":10,"sort":[],"aggs":{}}:""

	//	d, err := json.Marshal(data)
	//	if err != nil {
	//		fmt.Println(err)
	//		return nil, err
	//	}

	//	d := `{"query":{"bool":{"must":[{"wildcard":{"title":"` + q + `"}},{"wildcard":{"metas.description":"` + q + `"}},{"wildcard":{"content":"` + q + `"}}]}}}`
	d := `{"query":{"bool":{"must":[{"wildcard":{"title":"*` + q + `*"}}]}}}`
	result, err := client.Get(config["es.index"]+"/_search", d)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(result)
	data := map[string]interface{}{}

	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ret := make([]map[string]interface{}, 0)
	hits, ok := data["hits"]
	if ok {
		outerHits, ok := hits.(map[string]interface{})
		if ok {
			innerHits, ok := outerHits["hits"]
			if ok {
				hs, ok := innerHits.([]interface{})
				if ok {
					for _, hit := range hs {
						obj, ok := hit.(map[string]interface{})
						if ok {
							href := obj["_id"]
							href_str, ok := href.(string)
							href_str = strings.Replace(href_str, "%2F", "/", -1)

							src := obj["_source"]
							_source, ok := src.(map[string]interface{})
							if ok {
								page := make(map[string]interface{})
								page["Title"] = _source["title"]

								var desc interface{} = ""
								metas := _source["metas"]
								de, ok := metas.(map[string]interface{})
								if ok {
									desc = de["description"]
								}
								page["Description"] = desc
								page["Content"] = _source["content"]
								page["Url"] = href_str

								ret = append(ret, page)
							}
						}
					}
				}
			}
		}
	}

	//	hits := make([]map[string]string, 10)
	//	for i := 0; i < 10; i++ {
	//		hits[i] = make(map[string]string)
	//		a := strconv.Itoa(i)
	//		hits[i]["Title"] = "hit title" + a
	//		hits[i]["Content"] = "hit content" + a
	//	}
	return ret, nil
}

func main() {
	http.HandleFunc("/query", query)
	http.HandleFunc("/search", search)

	static := http.Dir(config["static.dir"])
	staticHandler := http.FileServer(static)
	// 'static' must ended with a slash, or the files in the 'static' dir can not be accessed
	http.Handle("/static/", http.StripPrefix("/static", staticHandler))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
