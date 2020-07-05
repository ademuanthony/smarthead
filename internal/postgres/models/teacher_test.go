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

func testTeachers(t *testing.T) {
	t.Parallel()

	query := Teachers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTeachersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
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

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeachersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Teachers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeachersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeacherSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeachersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TeacherExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Teacher exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TeacherExists to return true, but got false.")
	}
}

func testTeachersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	teacherFound, err := FindTeacher(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if teacherFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTeachersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Teachers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTeachersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Teachers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTeachersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	teacherOne := &Teacher{}
	teacherTwo := &Teacher{}
	if err = randomize.Struct(seed, teacherOne, teacherDBTypes, false, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}
	if err = randomize.Struct(seed, teacherTwo, teacherDBTypes, false, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = teacherOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teacherTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Teachers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTeachersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	teacherOne := &Teacher{}
	teacherTwo := &Teacher{}
	if err = randomize.Struct(seed, teacherOne, teacherDBTypes, false, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}
	if err = randomize.Struct(seed, teacherTwo, teacherDBTypes, false, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = teacherOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teacherTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTeachersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeachersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(teacherColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeacherToManySubjects(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Teacher
	var b, c Subject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, subjectDBTypes, false, subjectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, subjectDBTypes, false, subjectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"specialization\" (\"teacher_id\", \"subject_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"specialization\" (\"teacher_id\", \"subject_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Subjects().All(ctx, tx)
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

	slice := TeacherSlice{&a}
	if err = a.L.LoadSubjects(ctx, tx, false, (*[]*Teacher)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Subjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Subjects = nil
	if err = a.L.LoadSubjects(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Subjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTeacherToManyAddOpSubjects(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Teacher
	var b, c, d, e Subject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teacherDBTypes, false, strmangle.SetComplement(teacherPrimaryKeyColumns, teacherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Subject{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, subjectDBTypes, false, strmangle.SetComplement(subjectPrimaryKeyColumns, subjectColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Subject{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSubjects(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Teachers[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Teachers[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Subjects[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Subjects[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Subjects().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTeacherToManySetOpSubjects(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Teacher
	var b, c, d, e Subject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teacherDBTypes, false, strmangle.SetComplement(teacherPrimaryKeyColumns, teacherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Subject{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, subjectDBTypes, false, strmangle.SetComplement(subjectPrimaryKeyColumns, subjectColumnsWithoutDefault)...); err != nil {
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

	err = a.SetSubjects(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Subjects().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSubjects(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Subjects().Count(ctx, tx)
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
	// if len(b.R.Teachers) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Teachers) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Teachers[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Teachers[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Subjects[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Subjects[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testTeacherToManyRemoveOpSubjects(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Teacher
	var b, c, d, e Subject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teacherDBTypes, false, strmangle.SetComplement(teacherPrimaryKeyColumns, teacherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Subject{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, subjectDBTypes, false, strmangle.SetComplement(subjectPrimaryKeyColumns, subjectColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSubjects(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Subjects().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSubjects(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Subjects().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Teachers) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Teachers) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Teachers[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Teachers[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Subjects) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Subjects[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Subjects[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testTeachersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
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

func testTeachersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeacherSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTeachersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Teachers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	teacherDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`, `CV`: `character varying`, `YearsOfExperience`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_              = bytes.MinRead
)

func testTeachersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(teacherPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(teacherAllColumns) == len(teacherPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTeachersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(teacherAllColumns) == len(teacherPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Teacher{}
	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teacherDBTypes, true, teacherPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(teacherAllColumns, teacherPrimaryKeyColumns) {
		fields = teacherAllColumns
	} else {
		fields = strmangle.SetComplement(
			teacherAllColumns,
			teacherPrimaryKeyColumns,
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

	slice := TeacherSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTeachersUpsert(t *testing.T) {
	t.Parallel()

	if len(teacherAllColumns) == len(teacherPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Teacher{}
	if err = randomize.Struct(seed, &o, teacherDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Teacher: %s", err)
	}

	count, err := Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, teacherDBTypes, false, teacherPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Teacher struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Teacher: %s", err)
	}

	count, err = Teachers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
