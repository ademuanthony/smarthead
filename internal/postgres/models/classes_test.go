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

func testClasses(t *testing.T) {
	t.Parallel()

	query := Classes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClassesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
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

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClassesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Classes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClassesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClassSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClassesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClassExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Class exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClassExists to return true, but got false.")
	}
}

func testClassesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	classFound, err := FindClass(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if classFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClassesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Classes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClassesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Classes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClassesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	classOne := &Class{}
	classTwo := &Class{}
	if err = randomize.Struct(seed, classOne, classDBTypes, false, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}
	if err = randomize.Struct(seed, classTwo, classDBTypes, false, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = classOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = classTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Classes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClassesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	classOne := &Class{}
	classTwo := &Class{}
	if err = randomize.Struct(seed, classOne, classDBTypes, false, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}
	if err = randomize.Struct(seed, classTwo, classDBTypes, false, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = classOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = classTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testClassesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClassesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(classColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClassToManyDeposits(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c Deposit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, depositDBTypes, false, depositColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, depositDBTypes, false, depositColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ClassID = a.ID
	c.ClassID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Deposits().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ClassID == b.ClassID {
			bFound = true
		}
		if v.ClassID == c.ClassID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClassSlice{&a}
	if err = a.L.LoadDeposits(ctx, tx, false, (*[]*Class)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Deposits); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Deposits = nil
	if err = a.L.LoadDeposits(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Deposits); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testClassToManyStudents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
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

	queries.Assign(&b.ClassID, a.ID)
	queries.Assign(&c.ClassID, a.ID)
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
		if queries.Equal(v.ClassID, b.ClassID) {
			bFound = true
		}
		if queries.Equal(v.ClassID, c.ClassID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClassSlice{&a}
	if err = a.L.LoadStudents(ctx, tx, false, (*[]*Class)(&slice), nil); err != nil {
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

func testClassToManySubclasses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c Subclass

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ClassID = a.ID
	c.ClassID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Subclasses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ClassID == b.ClassID {
			bFound = true
		}
		if v.ClassID == c.ClassID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClassSlice{&a}
	if err = a.L.LoadSubclasses(ctx, tx, false, (*[]*Class)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Subclasses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Subclasses = nil
	if err = a.L.LoadSubclasses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Subclasses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testClassToManySubscriptions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c Subscription

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ClassID = a.ID
	c.ClassID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Subscriptions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ClassID == b.ClassID {
			bFound = true
		}
		if v.ClassID == c.ClassID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClassSlice{&a}
	if err = a.L.LoadSubscriptions(ctx, tx, false, (*[]*Class)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Subscriptions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Subscriptions = nil
	if err = a.L.LoadSubscriptions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Subscriptions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testClassToManyAddOpDeposits(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c, d, e Deposit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Deposit{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, depositDBTypes, false, strmangle.SetComplement(depositPrimaryKeyColumns, depositColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Deposit{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDeposits(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClassID {
			t.Error("foreign key was wrong value", a.ID, first.ClassID)
		}
		if a.ID != second.ClassID {
			t.Error("foreign key was wrong value", a.ID, second.ClassID)
		}

		if first.R.Class != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Class != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Deposits[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Deposits[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Deposits().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testClassToManyAddOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
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

		if !queries.Equal(a.ID, first.ClassID) {
			t.Error("foreign key was wrong value", a.ID, first.ClassID)
		}
		if !queries.Equal(a.ID, second.ClassID) {
			t.Error("foreign key was wrong value", a.ID, second.ClassID)
		}

		if first.R.Class != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Class != &a {
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

func testClassToManySetOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
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

	if !queries.IsValuerNil(b.ClassID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ClassID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ClassID) {
		t.Error("foreign key was wrong value", a.ID, d.ClassID)
	}
	if !queries.Equal(a.ID, e.ClassID) {
		t.Error("foreign key was wrong value", a.ID, e.ClassID)
	}

	if b.R.Class != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Class != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Class != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Class != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Students[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Students[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testClassToManyRemoveOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
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

	if !queries.IsValuerNil(b.ClassID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ClassID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Class != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Class != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Class != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Class != &a {
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

func testClassToManyAddOpSubclasses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c, d, e Subclass

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Subclass{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Subclass{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSubclasses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClassID {
			t.Error("foreign key was wrong value", a.ID, first.ClassID)
		}
		if a.ID != second.ClassID {
			t.Error("foreign key was wrong value", a.ID, second.ClassID)
		}

		if first.R.Class != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Class != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Subclasses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Subclasses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Subclasses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testClassToManyAddOpSubscriptions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Class
	var b, c, d, e Subscription

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classDBTypes, false, strmangle.SetComplement(classPrimaryKeyColumns, classColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Subscription{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, subscriptionDBTypes, false, strmangle.SetComplement(subscriptionPrimaryKeyColumns, subscriptionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Subscription{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSubscriptions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClassID {
			t.Error("foreign key was wrong value", a.ID, first.ClassID)
		}
		if a.ID != second.ClassID {
			t.Error("foreign key was wrong value", a.ID, second.ClassID)
		}

		if first.R.Class != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Class != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Subscriptions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Subscriptions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Subscriptions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testClassesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
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

func testClassesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClassSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClassesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Classes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	classDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`, `SchoolOrder`: `integer`}
	_            = bytes.MinRead
)

func testClassesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(classPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(classAllColumns) == len(classPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, classDBTypes, true, classPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClassesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(classAllColumns) == len(classPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Class{}
	if err = randomize.Struct(seed, o, classDBTypes, true, classColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, classDBTypes, true, classPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(classAllColumns, classPrimaryKeyColumns) {
		fields = classAllColumns
	} else {
		fields = strmangle.SetComplement(
			classAllColumns,
			classPrimaryKeyColumns,
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

	slice := ClassSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClassesUpsert(t *testing.T) {
	t.Parallel()

	if len(classAllColumns) == len(classPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Class{}
	if err = randomize.Struct(seed, &o, classDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Class: %s", err)
	}

	count, err := Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, classDBTypes, false, classPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Class struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Class: %s", err)
	}

	count, err = Classes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
