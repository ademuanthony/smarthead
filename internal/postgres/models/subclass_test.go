// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
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

func testSubclasses(t *testing.T) {
	t.Parallel()

	query := Subclasses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSubclassesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubclassesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Subclasses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubclassesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubclassSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubclassesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SubclassExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Subclass exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SubclassExists to return true, but got false.")
	}
}

func testSubclassesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	subclassFound, err := FindSubclass(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if subclassFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSubclassesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Subclasses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSubclassesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Subclasses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSubclassesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	subclassOne := &Subclass{}
	subclassTwo := &Subclass{}
	if err = randomize.Struct(seed, subclassOne, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}
	if err = randomize.Struct(seed, subclassTwo, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = subclassOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subclassTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subclasses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSubclassesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	subclassOne := &Subclass{}
	subclassTwo := &Subclass{}
	if err = randomize.Struct(seed, subclassOne, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}
	if err = randomize.Struct(seed, subclassTwo, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = subclassOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subclassTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testSubclassesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubclassesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(subclassColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubclassToManyStudents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subclass
	var b, c Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, studentDBTypes, false, studentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, studentDBTypes, false, studentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.SubclassID, a.ID)
	queries.Assign(&c.SubclassID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Students().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.SubclassID, b.SubclassID) {
			bFound = true
		}
		if queries.Equal(v.SubclassID, c.SubclassID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SubclassSlice{&a}
	if err = a.L.LoadStudents(ctx, tx, false, (*[]*Subclass)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Students); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Students = nil
	if err = a.L.LoadStudents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Students); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSubclassToManyTimetables(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subclass
	var b, c Timetable

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SubclassID = a.ID
	c.SubclassID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Timetables().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SubclassID == b.SubclassID {
			bFound = true
		}
		if v.SubclassID == c.SubclassID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SubclassSlice{&a}
	if err = a.L.LoadTimetables(ctx, tx, false, (*[]*Subclass)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Timetables); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Timetables = nil
	if err = a.L.LoadTimetables(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Timetables); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSubclassToManyAddOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subclass
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Student{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, studentDBTypes, false, strmangle.SetComplement(studentPrimaryKeyColumns, studentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Student{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStudents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.SubclassID) {
			t.Error("foreign key was wrong value", a.ID, first.SubclassID)
		}
		if !queries.Equal(a.ID, second.SubclassID) {
			t.Error("foreign key was wrong value", a.ID, second.SubclassID)
		}

		if first.R.Subclass != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Subclass != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Students[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Students[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Students().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testSubclassToManySetOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subclass
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Student{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, studentDBTypes, false, strmangle.SetComplement(studentPrimaryKeyColumns, studentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetStudents(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Students().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetStudents(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Students().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SubclassID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SubclassID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.SubclassID) {
		t.Error("foreign key was wrong value", a.ID, d.SubclassID)
	}
	if !queries.Equal(a.ID, e.SubclassID) {
		t.Error("foreign key was wrong value", a.ID, e.SubclassID)
	}

	if b.R.Subclass != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Subclass != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Subclass != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Subclass != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Students[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Students[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testSubclassToManyRemoveOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subclass
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Student{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, studentDBTypes, false, strmangle.SetComplement(studentPrimaryKeyColumns, studentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddStudents(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Students().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveStudents(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Students().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SubclassID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SubclassID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Subclass != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Subclass != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Subclass != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Subclass != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Students) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Students[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Students[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testSubclassToManyAddOpTimetables(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subclass
	var b, c, d, e Timetable

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Timetable{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Timetable{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTimetables(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SubclassID {
			t.Error("foreign key was wrong value", a.ID, first.SubclassID)
		}
		if a.ID != second.SubclassID {
			t.Error("foreign key was wrong value", a.ID, second.SubclassID)
		}

		if first.R.Subclass != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Subclass != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Timetables[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Timetables[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Timetables().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSubclassToOneClassUsingClass(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Subclass
	var foreign Class

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, classDBTypes, false, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ClassID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Class().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SubclassSlice{&local}
	if err = local.L.LoadClass(ctx, tx, false, (*[]*Subclass)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Class == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Class = nil
	if err = local.L.LoadClass(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Class == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSubclassToOneSetOpClassUsingClass(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subclass
	var b, c Class

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Class{&b, &c} {
		err = a.SetClass(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Class != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Subclasses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ClassID != x.ID {
			t.Error("foreign key was wrong value", a.ClassID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ClassID))
		reflect.Indirect(reflect.ValueOf(&a.ClassID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ClassID != x.ID {
			t.Error("foreign key was wrong value", a.ClassID, x.ID)
		}
	}
}

func testSubclassesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSubclassesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubclassSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSubclassesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subclasses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	subclassDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`, `ClassID`: `uuid`, `SchoolOrder`: `integer`, `Link`: `character varying`}
	_               = bytes.MinRead
)

func testSubclassesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(subclassPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(subclassAllColumns) == len(subclassPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSubclassesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(subclassAllColumns) == len(subclassPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Subclass{}
	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subclassDBTypes, true, subclassPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(subclassAllColumns, subclassPrimaryKeyColumns) {
		fields = subclassAllColumns
	} else {
		fields = strmangle.SetComplement(
			subclassAllColumns,
			subclassPrimaryKeyColumns,
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

	slice := SubclassSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSubclassesUpsert(t *testing.T) {
	t.Parallel()

	if len(subclassAllColumns) == len(subclassPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Subclass{}
	if err = randomize.Struct(seed, &o, subclassDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Subclass: %s", err)
	}

	count, err := Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, subclassDBTypes, false, subclassPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Subclass: %s", err)
	}

	count, err = Subclasses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
