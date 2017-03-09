//
// DISCLAIMER
//
// Copyright 2017 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package test

import (
	"context"
	"testing"

	driver "github.com/arangodb/go-driver"
)

// TestImportEdgesWithKeys imports documents and then checks that it exists.
func TestImportEdgesWithKeys(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []RouteEdgeWithKey{
		RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		RouteEdgeWithKey{
			"edge2",
			from.ID.String(),
			to.ID.String(),
			50,
		},
		RouteEdgeWithKey{
			"edge3",
			from.ID.String(),
			to.ID.String(),
			60,
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, nil)
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != int64(len(docs)) {
			t.Errorf("Expected %d created documents, got %d (json %s)", len(docs), stats.Created, string(raw))
		}
		if stats.Errors != 0 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 0, stats.Errors, string(raw))
		}
		if stats.Empty != 0 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 0, stats.Empty, string(raw))
		}
	}
}

// TestImportEdgesWithoutKeys imports documents and then checks that it exists.
func TestImportEdgesWithoutKeys(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_withhoutKeys_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_withhoutKeys_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []RouteEdgeWithKey{
		RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		RouteEdgeWithKey{
			"edge2",
			from.ID.String(),
			to.ID.String(),
			50,
		},
		RouteEdgeWithKey{
			"edge3",
			from.ID.String(),
			to.ID.String(),
			60,
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, nil)
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != int64(len(docs)) {
			t.Errorf("Expected %d created documents, got %d (json %s)", len(docs), stats.Created, string(raw))
		}
		if stats.Errors != 0 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 0, stats.Errors, string(raw))
		}
		if stats.Empty != 0 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 0, stats.Empty, string(raw))
		}
	}
}

// TestImportEdgesEmptyEntries imports documents and then checks that it exists.
func TestImportEdgesEmptyEntries(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_emptyEntries_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_emptyEntries_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []*RouteEdgeWithKey{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		&RouteEdgeWithKey{
			"edge2",
			from.ID.String(),
			to.ID.String(),
			50,
		},
		nil,
		&RouteEdgeWithKey{
			"edge3",
			from.ID.String(),
			to.ID.String(),
			60,
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, nil)
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != int64(len(docs))-1 {
			t.Errorf("Expected %d created documents, got %d (json %s)", len(docs)-1, stats.Created, string(raw))
		}
		if stats.Errors != 0 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 0, stats.Errors, string(raw))
		}
		if stats.Empty != 1 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 1, stats.Empty, string(raw))
		}
	}
}

// TestImportEdgesInvalidEntries imports documents and then checks that it exists.
func TestImportEdgesInvalidEntries(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_invalidEntries_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_invalidEntries_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		[]string{"array", "is", "invalid"},
		&RouteEdgeWithKey{
			"edge2",
			from.ID.String(),
			to.ID.String(),
			50,
		},
		"string is not valid",
		nil,
		&RouteEdgeWithKey{
			"edge3",
			from.ID.String(),
			to.ID.String(),
			60,
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, nil)
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != int64(len(docs))-3 {
			t.Errorf("Expected %d created documents, got %d (json %s)", len(docs)-3, stats.Created, string(raw))
		}
		if stats.Errors != 2 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 2, stats.Errors, string(raw))
		}
		if stats.Empty != 1 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 1, stats.Empty, string(raw))
		}
	}
}

// TestImportEdgesDuplicateEntries imports documents and then checks that it exists.
func TestImportEdgesDuplicateEntries(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_duplicateEntries_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_duplicateEntries_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []*RouteEdgeWithKey{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, nil)
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != 1 {
			t.Errorf("Expected %d created documents, got %d (json %s)", 1, stats.Created, string(raw))
		}
		if stats.Errors != 1 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 1, stats.Errors, string(raw))
		}
		if stats.Empty != 0 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 0, stats.Empty, string(raw))
		}
		if stats.Updated != 0 {
			t.Errorf("Expected %d updated documents, got %d (json %s)", 0, stats.Updated, string(raw))
		}
		if stats.Ignored != 0 {
			t.Errorf("Expected %d ignored documents, got %d (json %s)", 0, stats.Ignored, string(raw))
		}
	}
}

// TestImportEdgesDuplicateEntriesComplete imports documents and then checks that it exists.
func TestImportEdgesDuplicateEntriesComplete(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_duplicateEntriesComplete_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_duplicateEntriesComplete_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []*RouteEdgeWithKey{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		nil,
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	if _, err := col.ImportDocuments(ctx, docs, &driver.ImportDocumentOptions{
		Complete: true,
	}); !driver.IsConflict(err) {
		t.Errorf("Expected ConflictError, got %s", describe(err))
	}
}

// TestImportEdgesDuplicateEntriesUpdate imports documents and then checks that it exists.
func TestImportEdgesDuplicateEntriesUpdate(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_duplicateEntriesUpdate_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_duplicateEntriesUpdate_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		map[string]interface{}{
			"_key":  "edge1",
			"_from": to.ID.String(),
			"_to":   from.ID.String(),
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, &driver.ImportDocumentOptions{
		OnDuplicate: driver.ImportOnDuplicateUpdate,
	})
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != 1 {
			t.Errorf("Expected %d created documents, got %d (json %s)", 1, stats.Created, string(raw))
		}
		if stats.Errors != 0 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 0, stats.Errors, string(raw))
		}
		if stats.Empty != 0 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 0, stats.Empty, string(raw))
		}
		if stats.Updated != 1 {
			t.Errorf("Expected %d updated documents, got %d (json %s)", 1, stats.Updated, string(raw))
		}
		if stats.Ignored != 0 {
			t.Errorf("Expected %d ignored documents, got %d (json %s)", 0, stats.Ignored, string(raw))
		}

		var edge RouteEdgeWithKey
		if _, err := col.ReadDocument(nil, "edge1", &edge); err != nil {
			t.Errorf("ReadDocument failed: %s", describe(err))
		} else {
			if edge.From != to.ID.String() {
				t.Errorf("Expected From to be '%s', got '%s'", to, edge.From)
			}
			if edge.Distance != 40 {
				t.Errorf("Expected Distance to be 40, got %d", edge.Distance)
			}
		}
	}
}

// TestImportEdgesDuplicateEntriesReplace imports documents and then checks that it exists.
func TestImportEdgesDuplicateEntriesReplace(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_duplicateEntriesReplace_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_duplicateEntriesReplace_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		map[string]interface{}{
			"_key":  "edge1",
			"_from": to.ID.String(),
			"_to":   from.ID.String(),
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, &driver.ImportDocumentOptions{
		OnDuplicate: driver.ImportOnDuplicateReplace,
	})
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != 1 {
			t.Errorf("Expected %d created documents, got %d (json %s)", 1, stats.Created, string(raw))
		}
		if stats.Errors != 0 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 0, stats.Errors, string(raw))
		}
		if stats.Empty != 0 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 0, stats.Empty, string(raw))
		}
		if stats.Updated != 1 {
			t.Errorf("Expected %d updated documents, got %d (json %s)", 1, stats.Updated, string(raw))
		}
		if stats.Ignored != 0 {
			t.Errorf("Expected %d ignored documents, got %d (json %s)", 0, stats.Ignored, string(raw))
		}

		var edge RouteEdgeWithKey
		if _, err := col.ReadDocument(nil, "edge1", &edge); err != nil {
			t.Errorf("ReadDocument failed: %s", describe(err))
		} else {
			if edge.From != to.ID.String() {
				t.Errorf("Expected From to be '%s', got '%s'", to, edge.From)
			}
			if edge.Distance != 0 {
				t.Errorf("Expected Distance to be 0, got %d", edge.Distance)
			}
		}
	}
}

// TestImportEdgesDuplicateEntriesIgnore imports documents and then checks that it exists.
func TestImportEdgesDuplicateEntriesIgnore(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_duplicateEntriesIgnore_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_duplicateEntriesIgnore_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		map[string]interface{}{
			"_key":  "edge1",
			"_from": to.ID.String(),
			"_to":   from.ID.String(),
		},
	}

	var raw []byte
	ctx = driver.WithRawResponse(ctx, &raw)
	stats, err := col.ImportDocuments(ctx, docs, &driver.ImportDocumentOptions{
		OnDuplicate: driver.ImportOnDuplicateIgnore,
	})
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != 1 {
			t.Errorf("Expected %d created documents, got %d (json %s)", 1, stats.Created, string(raw))
		}
		if stats.Errors != 0 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 0, stats.Errors, string(raw))
		}
		if stats.Empty != 0 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 0, stats.Empty, string(raw))
		}
		if stats.Updated != 0 {
			t.Errorf("Expected %d updated documents, got %d (json %s)", 0, stats.Updated, string(raw))
		}
		if stats.Ignored != 1 {
			t.Errorf("Expected %d ignored documents, got %d (json %s)", 1, stats.Ignored, string(raw))
		}

		var edge RouteEdgeWithKey
		if _, err := col.ReadDocument(nil, "edge1", &edge); err != nil {
			t.Errorf("ReadDocument failed: %s", describe(err))
		} else {
			if edge.From != from.ID.String() {
				t.Errorf("Expected From to be '%s', got '%s'", to, edge.From)
			}
			if edge.Distance != 40 {
				t.Errorf("Expected Distance to be 0, got %d", edge.Distance)
			}
		}
	}
}

// TestImportEdgesDetails imports documents and then checks that it exists.
func TestImportEdgesDetails(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_details_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_details_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"_key": "venlo", "name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"_key": "lb", "name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdgeWithKey{
			"edge1",
			from.ID.String(),
			to.ID.String(),
			40,
		},
		map[string]interface{}{
			"_key":  "edge1",
			"_from": to.ID.String(),
			"_to":   from.ID.String(),
		},
	}

	var raw []byte
	var details []string
	ctx = driver.WithImportDetails(driver.WithRawResponse(ctx, &raw), &details)
	stats, err := col.ImportDocuments(ctx, docs, nil)
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != 1 {
			t.Errorf("Expected %d created documents, got %d (json %s)", 1, stats.Created, string(raw))
		}
		if stats.Errors != 1 {
			t.Errorf("Expected %d error documents, got %d (json %s)", 1, stats.Errors, string(raw))
		}
		if stats.Empty != 0 {
			t.Errorf("Expected %d empty documents, got %d (json %s)", 0, stats.Empty, string(raw))
		}
		if stats.Updated != 0 {
			t.Errorf("Expected %d updated documents, got %d (json %s)", 0, stats.Updated, string(raw))
		}
		if stats.Ignored != 0 {
			t.Errorf("Expected %d ignored documents, got %d (json %s)", 0, stats.Ignored, string(raw))
		}

		detailsExpected := `at position 1: creating document failed with error 'unique constraint violated', offending document: {"_from":"state/lb","_key":"edge1","_to":"city/venlo"}`
		if len(details) != 1 {
			t.Errorf("Expected 1 details, to %d", len(details))
		} else if details[0] != detailsExpected {
			t.Errorf("Expected details[0] to be '%s', got '%s'", detailsExpected, details[0])
		}
	}
}

// TestImportEdgesOverwriteYes imports documents and then checks that it exists.
func TestImportEdgesOverwriteYes(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_overwriteYes_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_overwriteYes_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdge{
			from.ID.String(),
			to.ID.String(),
			40,
		},
		map[string]interface{}{
			"_from": to.ID.String(),
			"_to":   from.ID.String(),
		},
	}

	for i := 0; i < 3; i++ {
		var raw []byte
		var details []string
		ctx := driver.WithImportDetails(driver.WithRawResponse(nil, &raw), &details)
		stats, err := col.ImportDocuments(ctx, docs, &driver.ImportDocumentOptions{
			Overwrite: true,
		})
		if err != nil {
			t.Fatalf("Failed to import documents: %s", describe(err))
		} else {
			if stats.Created != 2 {
				t.Errorf("Expected %d created documents, got %d (json %s)", 2, stats.Created, string(raw))
			}
		}

		countExpected := int64(2)
		if count, err := col.Count(nil); err != nil {
			t.Errorf("Failed to count documents: %s", describe(err))
		} else if count != countExpected {
			t.Errorf("Expected count to be %d in round %d, got %d", countExpected, i, count)
		}
	}
}

// TestImportEdgesOverwriteNo imports documents and then checks that it exists.
func TestImportEdgesOverwriteNo(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_overwriteNo_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_overwriteNo_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	from := createDocument(ctx, cities, map[string]interface{}{"name": "Venlo"}, t)
	to := createDocument(ctx, states, map[string]interface{}{"name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdge{
			from.ID.String(),
			to.ID.String(),
			40,
		},
		map[string]interface{}{
			"_from": to.ID.String(),
			"_to":   from.ID.String(),
		},
	}

	for i := 0; i < 3; i++ {
		var raw []byte
		var details []string
		ctx := driver.WithImportDetails(driver.WithRawResponse(nil, &raw), &details)
		stats, err := col.ImportDocuments(ctx, docs, &driver.ImportDocumentOptions{
			Overwrite: false,
		})
		if err != nil {
			t.Fatalf("Failed to import documents: %s", describe(err))
		} else {
			if stats.Created != 2 {
				t.Errorf("Expected %d created documents, got %d (json %s)", 2, stats.Created, string(raw))
			}
		}

		countExpected := int64(2 * (i + 1))
		if count, err := col.Count(nil); err != nil {
			t.Errorf("Failed to count documents: %s", describe(err))
		} else if count != countExpected {
			t.Errorf("Expected count to be %d in round %d, got %d", countExpected, i, count)
		}
	}
}

// TestImportEdgesPrefix imports documents and then checks that it exists.
func TestImportEdgesPrefix(t *testing.T) {
	ctx := context.Background()
	c := createClientFromEnv(t, true)
	db := ensureDatabase(ctx, c, "import_edges_prefix_test", nil, t)
	g := ensureGraph(ctx, db, "import_edges_prefix_test", nil, t)
	col := ensureEdgeCollection(ctx, g, "citiesPerState", []string{"city"}, []string{"state"}, t)
	cities := ensureCollection(ctx, db, "city", nil, t)
	states := ensureCollection(ctx, db, "state", nil, t)
	createDocument(ctx, cities, map[string]interface{}{"_key": "venlo", "name": "Venlo"}, t)
	createDocument(ctx, states, map[string]interface{}{"_key": "lb", "name": "Limburg"}, t)

	docs := []interface{}{
		&RouteEdge{
			"venlo",
			"lb",
			40,
		},
		map[string]interface{}{
			"_from": "venlo",
			"_to":   "lb",
		},
	}

	var raw []byte
	var details []string
	ctx = driver.WithImportDetails(driver.WithRawResponse(ctx, &raw), &details)
	stats, err := col.ImportDocuments(ctx, docs, &driver.ImportDocumentOptions{
		FromPrefix: "city",
		ToPrefix:   "state",
	})
	if err != nil {
		t.Fatalf("Failed to import documents: %s", describe(err))
	} else {
		if stats.Created != 2 {
			t.Errorf("Expected %d created documents, got %d (json %s)", 2, stats.Created, string(raw))
		}
	}

	countExpected := int64(2)
	if count, err := col.Count(nil); err != nil {
		t.Errorf("Failed to count documents: %s", describe(err))
	} else if count != countExpected {
		t.Errorf("Expected count to be %d, got %d", countExpected, count)
	}
}