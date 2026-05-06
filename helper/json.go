package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&result)
	ErrorT(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	ErrorT(err)

}
