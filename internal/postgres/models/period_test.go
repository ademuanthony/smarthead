// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPeriods(t *testing.T) {
	t.Parallel()

	query := Periods()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPeriodsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
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

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeriodsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Periods().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeriodsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeriodSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeriodsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PeriodExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Period exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PeriodExists to return true, but got false.")
	}
}

func testPeriodsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	periodFound, err := FindPeriod(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if periodFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPeriodsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Periods().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPeriodsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Periods().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPeriodsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	periodOne := &Period{}
	periodTwo := &Period{}
	if err = randomize.Struct(seed, periodOne, periodDBTypes, false, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}
	if err = randomize.Struct(seed, periodTwo, periodDBTypes, false, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = periodOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = periodTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Periods().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPeriodsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	periodOne := &Period{}
	periodTwo := &Period{}
	if err = randomize.Struct(seed, periodOne, periodDBTypes, false, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}
	if err = randomize.Struct(seed, periodTwo, periodDBTypes, false, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = periodOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = periodTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPeriodsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeriodsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(periodColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeriodToManyStudents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Period
	var b, c Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
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

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"students_periods\" (\"period_id\", \"student_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"students_periods\" (\"period_id\", \"student_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Students().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeriodSlice{&a}
	if err = a.L.LoadStudents(ctx, tx, false, (*[]*Period)(&slice), nil); err != nil {
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

func testPeriodToManySubscriptions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Period
	var b, c Subscription

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
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

	b.PeriodID = a.ID
	c.PeriodID = a.ID

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
		if v.PeriodID == b.PeriodID {
			bFound = true
		}
		if v.PeriodID == c.PeriodID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeriodSlice{&a}
	if err = a.L.LoadSubscriptions(ctx, tx, false, (*[]*Period)(&slice), nil); err != nil {
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

func testPeriodToManyAddOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Period
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, periodDBTypes, false, strmangle.SetComplement(periodPrimaryKeyColumns, periodColumnsWithoutDefault)...); err != nil {
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

		if first.R.Periods[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Periods[0] != &a {
			t.Error("relationship was not added properly to the slice")
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

func testPeriodToManySetOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Period
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, periodDBTypes, false, strmangle.SetComplement(periodPrimaryKeyColumns, periodColumnsWithoutDefault)...); err != nil {
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

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Periods) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Periods) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Periods[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Periods[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Students[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Students[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPeriodToManyRemoveOpStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Period
	var b, c, d, e Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, periodDBTypes, false, strmangle.SetComplement(periodPrimaryKeyColumns, periodColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.Periods) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Periods) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Periods[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Periods[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
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

func testPeriodToManyAddOpSubscriptions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Period
	var b, c, d, e Subscription

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, periodDBTypes, false, strmangle.SetComplement(periodPrimaryKeyColumns, periodColumnsWithoutDefault)...); err != nil {
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

		if a.ID != first.PeriodID {
			t.Error("foreign key was wrong value", a.ID, first.PeriodID)
		}
		if a.ID != second.PeriodID {
			t.Error("foreign key was wrong value", a.ID, second.PeriodID)
		}

		if first.R.Period != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Period != &a {
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

func testPeriodsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
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

func testPeriodsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeriodSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPeriodsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Periods().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	periodDBTypes = map[string]string{`ID`: `uuid`, `StartHour`: `integer`, `StartMinute`: `integer`, `EndHour`: `integer`, `EndMinute`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_             = bytes.MinRead
)

func testPeriodsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(periodPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(periodAllColumns) == len(periodPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, periodDBTypes, true, periodPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPeriodsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(periodAllColumns) == len(periodPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Period{}
	if err = randomize.Struct(seed, o, periodDBTypes, true, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, periodDBTypes, true, periodPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(periodAllColumns, periodPrimaryKeyColumns) {
		fields = periodAllColumns
	} else {
		fields = strmangle.SetComplement(
			periodAllColumns,
			periodPrimaryKeyColumns,
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

	slice := PeriodSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPeriodsUpsert(t *testing.T) {
	t.Parallel()

	if len(periodAllColumns) == len(periodPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Period{}
	if err = randomize.Struct(seed, &o, periodDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Period: %s", err)
	}

	count, err := Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, periodDBTypes, false, periodPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Period: %s", err)
	}

	count, err = Periods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
