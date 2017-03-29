package main

import (
	"github.com/tokillamockingbird/searchblox-golang"
)

func main() {
	c := searchblox.Client{Host: "http://searchblox:80"}
	customCollection := searchblox.CustomCollection{
		ApiKey: "25B213BA03FAB750790FC63FD1C6B301",
		Document: searchblox.Document{
			Colname: "Test",
		},
	}
	err := c.CreateCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
	indexCustomCollection := searchblox.CustomCollection{
		ApiKey: "25B213BA03FAB750790FC63FD1C6B301",
		Document: searchblox.Document{
			Colname:      "Test",
			Url:          "http://www.searchblox.com",
			Uid:          "http://www.searchblox.com",
			Location:     "http://www.searchblox.com",
			Alpha:        "string",
			Size:         "44244",
			Title:        "Text",
			Keywords:     "keywords",
			Description:  "SearchBlox Content Search Software",
			Content:      "content",
			LastModified: "14 January 2015 06:19:42 GMT",
			ContentType:  "HTML",
			Meta: searchblox.Meta{
				Location: "San Francisco",
				Temp:     "23",
				Weather:  "sunny",
			},
		},
	}
	err = c.IndexDocumentCustomCollection(indexCustomCollection)
	if err != nil {
		panic(err)
	}
	err = c.ClearCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
	err = c.DeleteCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
}