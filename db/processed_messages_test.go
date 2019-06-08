// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testProcessedMessages(t *testing.T) {
	t.Parallel()

	query := ProcessedMessages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProcessedMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessedMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ProcessedMessages().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessedMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProcessedMessageSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessedMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProcessedMessageExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ProcessedMessage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProcessedMessageExists to return true, but got false.")
	}
}

func testProcessedMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	processedMessageFound, err := FindProcessedMessage(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if processedMessageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProcessedMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ProcessedMessages().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testProcessedMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ProcessedMessages().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProcessedMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	processedMessageOne := &ProcessedMessage{}
	processedMessageTwo := &ProcessedMessage{}
	if err = randomize.Struct(seed, processedMessageOne, processedMessageDBTypes, false, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, processedMessageTwo, processedMessageDBTypes, false, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = processedMessageOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = processedMessageTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProcessedMessages().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProcessedMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	processedMessageOne := &ProcessedMessage{}
	processedMessageTwo := &ProcessedMessage{}
	if err = randomize.Struct(seed, processedMessageOne, processedMessageDBTypes, false, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, processedMessageTwo, processedMessageDBTypes, false, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = processedMessageOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = processedMessageTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func processedMessageBeforeInsertHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageAfterInsertHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageAfterSelectHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageBeforeUpdateHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageAfterUpdateHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageBeforeDeleteHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageAfterDeleteHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageBeforeUpsertHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func processedMessageAfterUpsertHook(e boil.Executor, o *ProcessedMessage) error {
	*o = ProcessedMessage{}
	return nil
}

func testProcessedMessagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &ProcessedMessage{}
	o := &ProcessedMessage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, processedMessageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage object: %s", err)
	}

	AddProcessedMessageHook(boil.BeforeInsertHook, processedMessageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	processedMessageBeforeInsertHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.AfterInsertHook, processedMessageAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	processedMessageAfterInsertHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.AfterSelectHook, processedMessageAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	processedMessageAfterSelectHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.BeforeUpdateHook, processedMessageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	processedMessageBeforeUpdateHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.AfterUpdateHook, processedMessageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	processedMessageAfterUpdateHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.BeforeDeleteHook, processedMessageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	processedMessageBeforeDeleteHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.AfterDeleteHook, processedMessageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	processedMessageAfterDeleteHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.BeforeUpsertHook, processedMessageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	processedMessageBeforeUpsertHooks = []ProcessedMessageHook{}

	AddProcessedMessageHook(boil.AfterUpsertHook, processedMessageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	processedMessageAfterUpsertHooks = []ProcessedMessageHook{}
}

func testProcessedMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProcessedMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(processedMessageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProcessedMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testProcessedMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProcessedMessageSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testProcessedMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProcessedMessages().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	processedMessageDBTypes = map[string]string{`ID`: `INTEGER`, `MessageID`: `VARCHAR(200)`, `Processed`: `BOOL`}
	_                       = bytes.MinRead
)

func testProcessedMessagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(processedMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(processedMessageColumns) == len(processedMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProcessedMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(processedMessageColumns) == len(processedMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProcessedMessage{}
	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessedMessages().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, processedMessageDBTypes, true, processedMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessedMessage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(processedMessageColumns, processedMessagePrimaryKeyColumns) {
		fields = processedMessageColumns
	} else {
		fields = strmangle.SetComplement(
			processedMessageColumns,
			processedMessagePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ProcessedMessageSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
