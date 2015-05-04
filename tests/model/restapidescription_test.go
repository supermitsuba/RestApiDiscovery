package model

import (
	"RestApiDiscovery/libs/data"
	"RestApiDiscovery/libs/helpers"
	"RestApiDiscovery/libs/model"
	"testing"
)

func TestReplace_MultipleRecord_ReplaceFirstIndex(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/three_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))
	var item = model.RestApiDescription{
		"localhost:8066",
		"Test derp.",
		"Test",
		"test",
		"EAST",
		true,
		"a9aa2783-e3c7-447f-80e0-f08fee96c14e",
	}
	var result = model.Replace(d, item)

	if result == nil {
		t.Errorf("Should be an array.  %v", result)
	}

	if len(result) != 3 {
		t.Errorf("Should be an array of 3.  %v", result)
	}
}

func TestReplace_OneRecord_ArrayOfTwo(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var item = model.RestApiDescription{
		"localhost:8066",
		"Test derp.",
		"Test",
		"test",
		"EAST",
		true,
		"a9be2783-e3c7-457f-80e0-f08fee96c14e",
	}
	var result = model.Replace(d, item)

	if result == nil {
		t.Errorf("Should be an array.  %v", result)
	}

	if len(result) != 2 {
		t.Errorf("Should be an array of 2.  %v", result)
	}

	var index = model.Find(result, "a9be2783-e3c7-457f-80e0-f08fee96c14e")
	if index != 0 {
		t.Errorf("Should not have found the record.  %v Index: %v", result, index)
	}

	if result[index].Id != item.Id &&
		result[index].Description != item.Description &&
		result[index].Environment != item.Environment &&
		result[index].IsActive != item.IsActive &&
		result[index].Location != item.Location &&
		result[index].Name != item.Name &&
		result[index].Url != item.Url {
		t.Errorf("Should have the same record. Expected: %v, Actual: %v", item, result)
	}
}

func TestReplace_NoRecord_ReturnUnModified(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var item = model.RestApiDescription{
		"localhost:8066",
		"Test derp.",
		"Test",
		"test",
		"EAST",
		true,
		"z",
	}
	var result = model.Replace(d, item)

	if result == nil {
		t.Errorf("Should be an array.  %v", result)
	}

	if len(result) != 2 {
		t.Errorf("Should be an array of 2.  %v", result)
	}

	if model.Find(result, "z") != -1 {
		t.Errorf("Should not have found the record.  %v", result)
	}
}

func TestReplace_NilParameters_ReturnEmptyArray(t *testing.T) {
	var item = model.RestApiDescription{
		"localhost:8066",
		"Test derp.",
		"Test",
		"test",
		"EAST",
		true,
		"a9aa2783-e3c7-447f-80e0-f08fee96c14e",
	}
	var result = model.Replace(nil, item)

	if result != nil {
		t.Errorf("Should be an array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be an array of 0.  %v", result)
	}
}

func TesRemove_OneRecord_ReturnListWithFirstRemoved(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Remove(d, 0)

	if result == nil {
		t.Errorf("Should be the list with first index removed.  %v", result)
	}
}

func TestRemove_NoRecord_RemoveNothing(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Remove(d, -1)

	if result == nil {
		t.Errorf("Should be the same list.  %v", result)
	}
}

func TestRemove_NilParameters_ReturnEmptyList(t *testing.T) {
	var result = model.Remove(nil, 0)

	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be empty array.  %v", result)
	}
}

func TestFind_FindMultipleRecord_ReturnFirstIndex(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Find(d, "a9be2783-e3c7-457f-80e0-f08fee96c14e")

	if result != 0 {
		t.Errorf("Should be 0 index.  %v", result)
	}
}

func TestFind_OneRecord_IndexZero(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Find(d, "a9be2783-e3c7-457f-80e0-f08fee96c14e")

	if result != 0 {
		t.Errorf("Should be 0 index.  %v", result)
	}
}

func TestFind_NoRecord_NegativeOne(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Find(d, "")

	if result != -1 {
		t.Errorf("Should be -1.  %v", result)
	}
}

func TestFind_NilParameters_NegativeOne(t *testing.T) {
	var result = model.Find(nil, "")

	if result != -1 {
		t.Errorf("Should be -1.  %v", result)
	}
}

func TestTake_TakeTooMany_ReturnAll(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Take(d, 3)

	if result == nil {
		t.Errorf("Should be an array.  %v", result)
	}

	if len(result) != 2 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestTake_OneTake_ReturnOne(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Take(d, 1)

	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 1 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestTake_NegativeTake_ReturnEmptyArray(t *testing.T) {
	var result = model.Take(nil, -1)
	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestTake_NilParameters_ReturnEmptyArray(t *testing.T) {
	var result = model.Take(nil, 0)
	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestSkip_NilParameters_ReturnEmptyArray(t *testing.T) {
	var result = model.Skip(nil, 0)
	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestSkip_NegativeSkip_EmptyArray(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Skip(d, -1)

	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestSkip_ZeroSkip_ReturnArray(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Skip(d, 0)

	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 2 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestSkip_OneSkip_ReturnArrayOfOne(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Skip(d, 1)

	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 1 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestSkip_SkipTooMany_ReturnEmptyArray(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Skip(d, 5)

	if result == nil {
		t.Errorf("Should be empty array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be the array.  %v", result)
	}
}

func TestWhere_NilParameters_ReturnNil(t *testing.T) {
	var result = model.Where(nil, nil)
	if result != nil {
		t.Errorf("Should be empty array.  %v", result)
	}
}

func TestWhere_NilFunc_ReturnArray(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Where(d, nil)
	if result == nil {
		t.Errorf("Should be the array.  %v", result)
	}

	if len(result) != 2 {
		t.Errorf("Should be the array.  %v", result)
	}

	if result[0].Id != d[0].Id &&
		result[0].Description != d[0].Description &&
		result[0].Environment != d[0].Environment &&
		result[0].IsActive != d[0].IsActive &&
		result[0].Location != d[0].Location &&
		result[0].Name != d[0].Name &&
		result[0].Url != d[0].Url {
		t.Errorf("Should have the same record. Expected: %v, Actual: %v", d, result)
	}
}

func TestWhere_NilData_ReturnEmptyArray(t *testing.T) {
	var result = model.Where(nil, func(item model.RestApiDescription) bool { return item.Url == "localhost:8088" })
	if result == nil {
		t.Errorf("Should be the array.  %v", result)
	}

	if len(result) != 0 {
		t.Errorf("Should be empty array.  %v", result)
	}
}

func TestWhere_SearchFunc_ReturnArrayOfOne(t *testing.T) {
	var dataLayer = data.File_access{FileLocation: "../../documentation/testdata/two_file.json"}
	var d = helpers.ConvertToThing(dataLayer.Load(""))

	var result = model.Where(d, func(item model.RestApiDescription) bool { return item.Url == "localhost:8088" })

	if result == nil {
		t.Errorf("Should be the array.  %v", result)
	}

	if len(result) != 1 {
		t.Errorf("Should be the array.  %v", result)
	}

	if result[0].Id != d[0].Id &&
		result[0].Description != d[0].Description &&
		result[0].Environment != d[0].Environment &&
		result[0].IsActive != d[0].IsActive &&
		result[0].Location != d[0].Location &&
		result[0].Name != d[0].Name &&
		result[0].Url != d[0].Url {
		t.Errorf("Should have the same record. Expected: %v, Actual: %v", d, result)
	}
}
