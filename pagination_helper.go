package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getPagination(r *http.Request) (limit int32, skip int32) {
	query := r.URL.Query()
	limit = 10
	skip = 0
	if li := query.Get("limit"); li != "" {
		newLimit, err := strconv.Atoi(li)
		if err != nil {
			fmt.Println("Unable to parse Limit:", li)
		} else {
			limit = int32(newLimit)
		}
	}
	if sk := query.Get("skip"); sk != "" {
		newSkip, err := strconv.Atoi(sk)
		if err != nil {
			fmt.Println("Unable to parse Offset:", sk)
		} else {
			skip = int32(newSkip)
		}
	}
	return limit, skip
}
