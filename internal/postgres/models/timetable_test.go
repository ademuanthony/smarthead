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

func testTimetables(t *testing.T) {
	t.Parallel()

	query := Timetables()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTimetablesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
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

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTimetablesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Timetables().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTimetablesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TimetableSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTimetablesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TimetableExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Timetable exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TimetableExists to return true, but got false.")
	}
}

func testTimetablesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	timetableFound, err := FindTimetable(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if timetableFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTimetablesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Timetables().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTimetablesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Timetables().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTimetablesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	timetableOne := &Timetable{}
	timetableTwo := &Timetable{}
	if err = randomize.Struct(seed, timetableOne, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}
	if err = randomize.Struct(seed, timetableTwo, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = timetableOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = timetableTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Timetables().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTimetablesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	timetableOne := &Timetable{}
	timetableTwo := &Timetable{}
	if err = randomize.Struct(seed, timetableOne, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}
	if err = randomize.Struct(seed, timetableTwo, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = timetableOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = timetableTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTimetablesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTimetablesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(timetableColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTimetableToManyLessons(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Timetable
	var b, c Lesson

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.TimetableID = a.ID
	c.TimetableID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Lessons().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.TimetableID == b.TimetableID {
			bFound = true
		}
		if v.TimetableID == c.TimetableID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TimetableSlice{&a}
	if err = a.L.LoadLessons(ctx, tx, false, (*[]*Timetable)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Lessons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Lessons = nil
	if err = a.L.LoadLessons(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Lessons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTimetableToManyAddOpLessons(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Timetable
	var b, c, d, e Lesson

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Lesson{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, lessonDBTypes, false, strmangle.SetComplement(lessonPrimaryKeyColumns, lessonColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Lesson{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLessons(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TimetableID {
			t.Error("foreign key was wrong value", a.ID, first.TimetableID)
		}
		if a.ID != second.TimetableID {
			t.Error("foreign key was wrong value", a.ID, second.TimetableID)
		}

		if first.R.Timetable != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Timetable != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Lessons[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Lessons[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Lessons().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTimetableToOnePeriodUsingPeriod(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Timetable
	var foreign Period

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, periodDBTypes, false, periodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Period struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PeriodID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Period().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TimetableSlice{&local}
	if err = local.L.LoadPeriod(ctx, tx, false, (*[]*Timetable)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Period == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Period = nil
	if err = local.L.LoadPeriod(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Period == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTimetableToOneSubclassUsingSubclass(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Timetable
	var foreign Subclass

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, subclassDBTypes, false, subclassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subclass struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SubclassID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subclass().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TimetableSlice{&local}
	if err = local.L.LoadSubclass(ctx, tx, false, (*[]*Timetable)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Subclass == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Subclass = nil
	if err = local.L.LoadSubclass(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Subclass == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTimetableToOneSubjectUsingSubject(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Timetable
	var foreign Subject

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, subjectDBTypes, false, subjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subject struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SubjectID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subject().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TimetableSlice{&local}
	if err = local.L.LoadSubject(ctx, tx, false, (*[]*Timetable)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Subject = nil
	if err = local.L.LoadSubject(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTimetableToOneUserUsingTeacher(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Timetable
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TeacherID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Teacher().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TimetableSlice{&local}
	if err = local.L.LoadTeacher(ctx, tx, false, (*[]*Timetable)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Teacher == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Teacher = nil
	if err = local.L.LoadTeacher(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Teacher == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTimetableToOneSetOpPeriodUsingPeriod(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Timetable
	var b, c Period

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, periodDBTypes, false, strmangle.SetComplement(periodPrimaryKeyColumns, periodColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, periodDBTypes, false, strmangle.SetComplement(periodPrimaryKeyColumns, periodColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Period{&b, &c} {
		err = a.SetPeriod(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Period != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Timetables[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PeriodID != x.ID {
			t.Error("foreign key was wrong value", a.PeriodID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PeriodID))
		reflect.Indirect(reflect.ValueOf(&a.PeriodID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PeriodID != x.ID {
			t.Error("foreign key was wrong value", a.PeriodID, x.ID)
		}
	}
}
func testTimetableToOneSetOpSubclassUsingSubclass(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Timetable
	var b, c Subclass

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, subclassDBTypes, false, strmangle.SetComplement(subclassPrimaryKeyColumns, subclassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Subclass{&b, &c} {
		err = a.SetSubclass(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subclass != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Timetables[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubclassID != x.ID {
			t.Error("foreign key was wrong value", a.SubclassID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubclassID))
		reflect.Indirect(reflect.ValueOf(&a.SubclassID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubclassID != x.ID {
			t.Error("foreign key was wrong value", a.SubclassID, x.ID)
		}
	}
}
func testTimetableToOneSetOpSubjectUsingSubject(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Timetable
	var b, c Subject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, subjectDBTypes, false, strmangle.SetComplement(subjectPrimaryKeyColumns, subjectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, subjectDBTypes, false, strmangle.SetComplement(subjectPrimaryKeyColumns, subjectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Subject{&b, &c} {
		err = a.SetSubject(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Timetables[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubjectID != x.ID {
			t.Error("foreign key was wrong value", a.SubjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubjectID))
		reflect.Indirect(reflect.ValueOf(&a.SubjectID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubjectID != x.ID {
			t.Error("foreign key was wrong value", a.SubjectID, x.ID)
		}
	}
}
func testTimetableToOneSetOpUserUsingTeacher(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Timetable
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetTeacher(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Teacher != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TeacherTimetables[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TeacherID != x.ID {
			t.Error("foreign key was wrong value", a.TeacherID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TeacherID))
		reflect.Indirect(reflect.ValueOf(&a.TeacherID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TeacherID != x.ID {
			t.Error("foreign key was wrong value", a.TeacherID, x.ID)
		}
	}
}

func testTimetablesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
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

func testTimetablesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TimetableSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTimetablesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Timetables().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	timetableDBTypes = map[string]string{`ID`: `uuid`, `SubclassID`: `uuid`, `SubjectID`: `uuid`, `TeacherID`: `character`, `PeriodID`: `uuid`, `Day`: `integer`}
	_                = bytes.MinRead
)

func testTimetablesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(timetablePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(timetableAllColumns) == len(timetablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTimetablesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(timetableAllColumns) == len(timetablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Timetable{}
	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, timetableDBTypes, true, timetablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(timetableAllColumns, timetablePrimaryKeyColumns) {
		fields = timetableAllColumns
	} else {
		fields = strmangle.SetComplement(
			timetableAllColumns,
			timetablePrimaryKeyColumns,
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

	slice := TimetableSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTimetablesUpsert(t *testing.T) {
	t.Parallel()

	if len(timetableAllColumns) == len(timetablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Timetable{}
	if err = randomize.Struct(seed, &o, timetableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Timetable: %s", err)
	}

	count, err := Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, timetableDBTypes, false, timetablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Timetable: %s", err)
	}

	count, err = Timetables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
