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

func testLessonStudents(t *testing.T) {
	t.Parallel()

	query := LessonStudents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLessonStudentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
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

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLessonStudentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LessonStudents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLessonStudentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LessonStudentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLessonStudentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LessonStudentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if LessonStudent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LessonStudentExists to return true, but got false.")
	}
}

func testLessonStudentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	lessonStudentFound, err := FindLessonStudent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if lessonStudentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLessonStudentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LessonStudents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLessonStudentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LessonStudents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLessonStudentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	lessonStudentOne := &LessonStudent{}
	lessonStudentTwo := &LessonStudent{}
	if err = randomize.Struct(seed, lessonStudentOne, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}
	if err = randomize.Struct(seed, lessonStudentTwo, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lessonStudentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lessonStudentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LessonStudents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLessonStudentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	lessonStudentOne := &LessonStudent{}
	lessonStudentTwo := &LessonStudent{}
	if err = randomize.Struct(seed, lessonStudentOne, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}
	if err = randomize.Struct(seed, lessonStudentTwo, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lessonStudentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lessonStudentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testLessonStudentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLessonStudentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(lessonStudentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLessonStudentToOneLessonUsingLesson(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LessonStudent
	var foreign Lesson

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.LessonID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Lesson().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LessonStudentSlice{&local}
	if err = local.L.LoadLesson(ctx, tx, false, (*[]*LessonStudent)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Lesson == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Lesson = nil
	if err = local.L.LoadLesson(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Lesson == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLessonStudentToOneStudentUsingStudent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LessonStudent
	var foreign Student

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, studentDBTypes, false, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StudentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Student().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LessonStudentSlice{&local}
	if err = local.L.LoadStudent(ctx, tx, false, (*[]*LessonStudent)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Student == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Student = nil
	if err = local.L.LoadStudent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Student == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLessonStudentToOneSetOpLessonUsingLesson(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LessonStudent
	var b, c Lesson

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lessonStudentDBTypes, false, strmangle.SetComplement(lessonStudentPrimaryKeyColumns, lessonStudentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, lessonDBTypes, false, strmangle.SetComplement(lessonPrimaryKeyColumns, lessonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, lessonDBTypes, false, strmangle.SetComplement(lessonPrimaryKeyColumns, lessonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Lesson{&b, &c} {
		err = a.SetLesson(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Lesson != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.LessonStudents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.LessonID != x.ID {
			t.Error("foreign key was wrong value", a.LessonID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LessonID))
		reflect.Indirect(reflect.ValueOf(&a.LessonID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.LessonID != x.ID {
			t.Error("foreign key was wrong value", a.LessonID, x.ID)
		}
	}
}
func testLessonStudentToOneSetOpStudentUsingStudent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LessonStudent
	var b, c Student

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lessonStudentDBTypes, false, strmangle.SetComplement(lessonStudentPrimaryKeyColumns, lessonStudentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, studentDBTypes, false, strmangle.SetComplement(studentPrimaryKeyColumns, studentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, studentDBTypes, false, strmangle.SetComplement(studentPrimaryKeyColumns, studentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Student{&b, &c} {
		err = a.SetStudent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Student != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.LessonStudents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StudentID != x.ID {
			t.Error("foreign key was wrong value", a.StudentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StudentID))
		reflect.Indirect(reflect.ValueOf(&a.StudentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StudentID != x.ID {
			t.Error("foreign key was wrong value", a.StudentID, x.ID)
		}
	}
}

func testLessonStudentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
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

func testLessonStudentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LessonStudentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLessonStudentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LessonStudents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	lessonStudentDBTypes = map[string]string{`ID`: `uuid`, `LessonID`: `uuid`, `StudentID`: `uuid`}
	_                    = bytes.MinRead
)

func testLessonStudentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(lessonStudentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(lessonStudentAllColumns) == len(lessonStudentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLessonStudentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(lessonStudentAllColumns) == len(lessonStudentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LessonStudent{}
	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lessonStudentDBTypes, true, lessonStudentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(lessonStudentAllColumns, lessonStudentPrimaryKeyColumns) {
		fields = lessonStudentAllColumns
	} else {
		fields = strmangle.SetComplement(
			lessonStudentAllColumns,
			lessonStudentPrimaryKeyColumns,
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

	slice := LessonStudentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLessonStudentsUpsert(t *testing.T) {
	t.Parallel()

	if len(lessonStudentAllColumns) == len(lessonStudentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := LessonStudent{}
	if err = randomize.Struct(seed, &o, lessonStudentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LessonStudent: %s", err)
	}

	count, err := LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, lessonStudentDBTypes, false, lessonStudentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LessonStudent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LessonStudent: %s", err)
	}

	count, err = LessonStudents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
