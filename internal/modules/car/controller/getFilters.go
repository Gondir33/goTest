package controller

import "net/url"

func getFilters(URL *url.URL) map[string]string {
	filter := make(map[string]string, 7)
	filter["regNum"] = URL.Query().Get("regNum")
	filter["mark"] = URL.Query().Get("mark")
	filter["model"] = URL.Query().Get("model")
	filter["year"] = URL.Query().Get("year")
	filter["name"] = URL.Query().Get("name")
	filter["surname"] = URL.Query().Get("surname")
	filter["patronymic"] = URL.Query().Get("patronymic")

	return filter
}
