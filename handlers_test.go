// handlers_test
package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet_Request_Ok(t *testing.T) {
	locationOfFile = "test/two_file.json"
	var expectedData = GetFileOfRestApiDescriptions(locationOfFile)
	var response = httptest.NewRecorder()

	Get_RestApiRecords_Impl(nil, nil, nil, nil, nil, nil, response)

	if response.Code != 200 {
		t.Errorf("Should have 200 status code.")
	}

	var actualData = GetBody(response.Body.Bytes(), t)

	AssertArrayMatches(expectedData, actualData, t)
}

func TestGet_NilParameters(t *testing.T) {
	locationOfFile = "test/two_file.json"
	Get_RestApiRecords_Impl(nil, nil, nil, nil, nil, nil, nil)
}

func TestPut_Request_Ok(t *testing.T) {
	locationOfFile = "test/two_file_put.json"
	WriteRestApiDescriptionsToFile(GetFileOfRestApiDescriptions("test/two_file.json"), locationOfFile)
	var updateRecordString = "{\"url\":\"localhost:8088\",\"description\":\"This service is for discovering Apis.\",\"name\":\"Rest Api Discovery Service\",\"environment\":\"prod\",\"location\":\"EAST\",\"active\":true,\"id\":\"a9be2783-e3c7-457f-80e0-f08fee96c14e\"}"
	var updateRecord = new(RestApiDescription)
	var response = httptest.NewRecorder()

	json.Unmarshal([]byte(updateRecordString), updateRecord)

	Put_RestApiRecords_Impl(updateRecord, response)

	if response.Code != 200 {
		t.Errorf("Should have 200 status code. Status Code: %v", response.Code)
	}

	if !strings.Contains(response.Body.String(), "{ 'result':'OK' }") {
		t.Errorf("Should have an ok json message. Body: %v", response.Body.String())
	}
}

func TestPut_NilParameters(t *testing.T) {
	locationOfFile = "test/two_file.json"
	Put_RestApiRecords_Impl(nil, nil)
}

func TestPost_NilParameters(t *testing.T) {
	locationOfFile = "test/two_file.json"
	Post_RestApiRecords_Impl(nil, nil)
}

func TestPost_Request_Ok(t *testing.T) {
	locationOfFile = "test/two_file_post.json"
	WriteRestApiDescriptionsToFile(GetFileOfRestApiDescriptions("test/two_file.json"), locationOfFile)
	var updateRecordString = "{\"url\":\"localhost:8088\",\"description\":\"This service is for discovering Apis.\",\"name\":\"Rest Api Discovery Service\",\"environment\":\"prod\",\"location\":\"EAST\",\"active\":true,\"id\":\"a9be2783-e3c7-457f-80e0-f08fee96c14e\"}"
	var updateRecord = new(RestApiDescription)
	var response = httptest.NewRecorder()

	json.Unmarshal([]byte(updateRecordString), updateRecord)

	Post_RestApiRecords_Impl(updateRecord, response)

	if response.Code != 200 {
		t.Errorf("Should have 200 status code. Status Code: %v", response.Code)
	}

	if !strings.Contains(response.Body.String(), "{ 'result':'OK' }") {
		t.Errorf("Should have an ok json message. Body: %v", response.Body.String())
	}
}

func TestDelete_Request_Ok(t *testing.T) {
	locationOfFile = "test/two_file_delete.json"
	WriteRestApiDescriptionsToFile(GetFileOfRestApiDescriptions("test/two_file.json"), locationOfFile)

	var response = httptest.NewRecorder()

	Delete_RestApiRecord_Impl("a9be2783-e3c7-457f-80e0-f08fee96c14e", response)

	if response.Code != 200 {
		t.Errorf("Should have 200 status code. Status Code: %v", response.Code)
	}

	if !strings.Contains(response.Body.String(), "{ 'result':'OK' }") {
		t.Errorf("Should have an ok json message. Body: %v", response.Body.String())
	}
}

func TestDelete_NilParameters(t *testing.T) {
	locationOfFile = "test/two_file.json"
	Delete_RestApiRecord_Impl("", nil)
}

/*
Helper methods
*/

func AssertArrayMatches(expected []RestApiDescription, actual []RestApiDescription, t *testing.T) {
	for j := 0; j < len(expected)-1; j++ {
		if actual[j].Id != expected[j].Id &&
			actual[j].Description != expected[j].Description &&
			actual[j].Environment != expected[j].Environment &&
			actual[j].IsActive != expected[j].IsActive &&
			actual[j].Location != expected[j].Location &&
			actual[j].Name != expected[j].Name &&
			actual[j].Url != expected[j].Url {

			t.Errorf("Should have matched the records.  Actual: %v\nExpected: %v\n", actual, expected)
		}
	}
}

func GetBody(body []byte, t *testing.T) []RestApiDescription {
	var actualData = []RestApiDescription{}
	if err := json.Unmarshal(body, &actualData); err != nil {
		t.Errorf("Could not convert body: %v.", err)
	}
	return actualData
}
