package data_test

import (
	data "RestApiDiscovery/libs/data"
	"RestApiDiscovery/libs/helpers"
	model "RestApiDiscovery/libs/model"

	"encoding/json"
	"testing"
)

func TestReadFile_NoFile_NilArray(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/empty_file.json"}
	var result = helpers.ConvertToThing(dataLayer.Load(""))

	if result != nil {
		t.Errorf("Should have been null.")
	}
}

/*
[
	{
		"url":"localhost:8088",
		"description":"This service is for discovering Apis.",
		"name":"Rest Api Discovery Service",
		"environment":"prod",
		"location":"EAST",
		"active":true,
		"id":"a9be2783-e3c7-457f-80e0-f08fee96c14e"
	}
]
*/
func TestReadFile_OneRecord_ArrayOfOne(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/one_file.json"}
	var result = helpers.ConvertToThing(dataLayer.Load(""))

	if result == nil {
		t.Errorf("Should have read one record.")
	}

	if len(result) != 1 {
		t.Errorf("Should only be one record. Count: %v", len(result))
	}

	if result[0].Id != "a9be2783-e3c7-457f-80e0-f08fee96c14e" &&
		result[0].Description != "This service is for discovering Apis." &&
		result[0].Environment != "prod" &&
		result[0].IsActive != true &&
		result[0].Location != "EAST" &&
		result[0].Name != "Rest Api Discovery Service" && result[0].Url != "localhost:8088" {
		t.Errorf("Should have deserialized the file correctly.")
	}
}

/*
[
	{
		"url":"localhost:8088",
		"description":"This service is for discovering Apis.",
		"name":"Rest Api Discovery Service",
		"environment":"prod",
		"location":"EAST",
		"active":true,"
		id":"a9be2783-e3c7-457f-80e0-f08fee96c14e"
	},{
		"url":"localhost:8050",
		"description":"This service is for testing.",
		"name":"testing Discovery Service",
		"environment":"beta",
		"location":"EAST",
		"active":true,
		"id":"18696a36-8f4a-4f46-8496-c3a34352a946"
	}
]
*/
func TestReadFile_TwoRecords_ArrayOf2(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var result = helpers.ConvertToThing(dataLayer.Load(""))

	if result == nil {
		t.Errorf("Should have read one record.")
	}

	if len(result) != 2 {
		t.Errorf("Should only be two records. Count: %v", len(result))
	}

	if result[0].Id != "a9be2783-e3c7-457f-80e0-f08fee96c14e" &&
		result[0].Description != "This service is for discovering Apis." &&
		result[0].Environment != "prod" &&
		result[0].IsActive != true &&
		result[0].Location != "EAST" &&
		result[0].Name != "Rest Api Discovery Service" &&
		result[0].Url != "localhost:8088" {
		t.Errorf("Should have deserialized the first record correctly.")
	}

	if result[1].Id != "18696a36-8f4a-4f46-8496-c3a34352a946" &&
		result[1].Description != "This service is for discovering Apis." &&
		result[1].Environment != "beta" &&
		result[1].IsActive != true &&
		result[1].Location != "EAST" &&
		result[1].Name != "testing Discovery Service" &&
		result[1].Url != "localhost:8050" {
		t.Errorf("Should have deserialized the second record correctly.")
	}
}

func TestWriteFile_EmptyFile_WriteNil(t *testing.T) {
	var d = new([]model.RestApiDescription)
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/output/TestWriteFile_EmptyFile.json"}
	var myData, _ = json.Marshal(*d)
	dataLayer.Save("", string(myData))

	var result = helpers.ConvertToThing(dataLayer.Load(""))

	if result != nil {
		t.Errorf("Should have nil value.")

	}
}

func TestWriteFile_ArrayOfOne_WriteOneRecord(t *testing.T) {
	var d = []model.RestApiDescription{model.RestApiDescription{
		"localhost:8088",
		"This service is for discovering Apis.",
		"Rest Api Discovery Service",
		"prod",
		"EAST",
		true,
		"a9be2783-e3c7-457f-80e0-f08fee96c14e",
	}}
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/output/TestWriteFile_OneRecordFile.json"}

	var myData, _ = json.Marshal(d)
	dataLayer.Save("", string(myData))

	var result = helpers.ConvertToThing(dataLayer.Load(""))

	if result == nil {
		t.Errorf("Should not be nil value.")
	}

	if len(result) != 1 {
		t.Errorf("Should have one record.")
	}

	if result[0].Id != d[0].Id &&
		result[0].Description != d[0].Description &&
		result[0].Environment != d[0].Environment &&
		result[0].IsActive != d[0].IsActive &&
		result[0].Location != d[0].Location &&
		result[0].Name != d[0].Name &&
		result[0].Url != d[0].Url {
		t.Errorf("Should have deserialized the first record correctly.")
	}
}

func TestWriteFile_ArrayOfTwo_WriteTwoRecord(t *testing.T) {
	var d = []model.RestApiDescription{model.RestApiDescription{
		"localhost:8088",
		"This service is for discovering Apis.",
		"Rest Api Discovery Service",
		"prod",
		"EAST",
		true,
		"a9be2783-e3c7-457f-80e0-f08fee96c14e",
	},
		model.RestApiDescription{
			"localhost:8066",
			"Test derp.",
			"Test",
			"test",
			"EAST",
			true,
			"a9aa2783-e3c7-447f-80e0-f08fee96c14e",
		}}

	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/output/TestWriteFile_OneRecordFile.json"}
	var myData, _ = json.Marshal(d)
	dataLayer.Save("", string(myData))

	var result = helpers.ConvertToThing(dataLayer.Load(""))

	if result == nil {
		t.Errorf("Should not be nil value.")
	}

	if len(result) != 2 {
		t.Errorf("Should have two record.")
	}

	for j := 0; j < len(result); j++ {
		if result[j].Id != d[j].Id &&
			result[j].Description != d[j].Description &&
			result[j].Environment != d[j].Environment &&
			result[j].IsActive != d[j].IsActive &&
			result[j].Location != d[j].Location &&
			result[j].Name != d[j].Name &&
			result[j].Url != d[j].Url {
			t.Errorf("Should have deserialized the 2 record correctly.")
		}
	}
}
