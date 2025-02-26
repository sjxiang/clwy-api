package handlers

import (
	"fmt"
	"net/http"
    "strconv"
	"errors"

	"github.com/go-chi/chi/v5"
)


func getInt64FromPathParam(r *http.Request) (int64, error) {

	idParam := chi.URLParam(r, "id")
	if len(idParam) == 0 {
		return 0, errors.New("input missing id field")
	}

	id, okAssert := strconv.ParseInt(idParam, 10, 64)
	if okAssert != nil {
		return 0, errors.New("input id in wrong format")
	} 
	
	return id, nil
}


func parseQueryParamToInt64(r *http.Request, key string, defaultValue int64) (int64, error) {
    valueStr := r.URL.Query().Get(key)
    if valueStr == "" {
        return defaultValue, nil
    }

	value, err := strconv.ParseInt(valueStr, 10, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid query parameter '%s': %v", key, err)
    }

    return value, nil
}

func readQueryParam(r *http.Request, key string) string {
    value := r.URL.Query().Get(key)
	return value
}