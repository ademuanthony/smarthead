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

func testLessons(t *testing.T) {
	t.Parallel()

	query := Lessons()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLessonsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
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

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLessonsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Lessons().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLessonsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LessonSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLessonsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LessonExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Lesson exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LessonExists to return true, but got false.")
	}
}

func testLessonsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	lessonFound, err := FindLesson(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if lessonFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLessonsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Lessons().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLessonsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Lessons().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLessonsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	lessonOne := &Lesson{}
	lessonTwo := &Lesson{}
	if err = randomize.Struct(seed, lessonOne, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}
	if err = randomize.Struct(seed, lessonTwo, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lessonOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lessonTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Lessons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLessonsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	lessonOne := &Lesson{}
	lessonTwo := &Lesson{}
	if err = randomize.Struct(seed, lessonOne, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}
	if err = randomize.Struct(seed, lessonTwo, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lessonOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lessonTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testLessonsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLessonsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(lessonColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLessonToManyLessonStudents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lesson
	var b, c LessonStudent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, lessonStudentDBTypes, false, lessonStudentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.LessonID = a.ID
	c.LessonID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LessonStudents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LessonID == b.LessonID {
			bFound = true
		}
		if v.LessonID == c.LessonID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LessonSlice{&a}
	if err = a.L.LoadLessonStudents(ctx, tx, false, (*[]*Lesson)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LessonStudents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LessonStudents = nil
	if err = a.L.LoadLessonStudents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LessonStudents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testLessonToManyAddOpLessonStudents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lesson
	var b, c, d, e LessonStudent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lessonDBTypes, false, strmangle.SetComplement(lessonPrimaryKeyColumns, lessonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*LessonStudent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, lessonStudentDBTypes, false, strmangle.SetComplement(lessonStudentPrimaryKeyColumns, lessonStudentColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*LessonStudent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLessonStudents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LessonID {
			t.Error("foreign key was wrong value", a.ID, first.LessonID)
		}
		if a.ID != second.LessonID {
			t.Error("foreign key was wrong value", a.ID, second.LessonID)
		}

		if first.R.Lesson != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Lesson != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LessonStudents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LessonStudents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LessonStudents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testLessonToOneUserUsingTeacher(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Lesson
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.TeacherID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Teacher().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LessonSlice{&local}
	if err = local.L.LoadTeacher(ctx, tx, false, (*[]*Lesson)(&slice), nil); err != nil {
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

func testLessonToOneTimetableUsingTimetable(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Lesson
	var foreign Timetable

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, lessonDBTypes, false, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, timetableDBTypes, false, timetableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Timetable struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TimetableID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Timetable().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LessonSlice{&local}
	if err = local.L.LoadTimetable(ctx, tx, false, (*[]*Lesson)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Timetable == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Timetable = nil
	if err = local.L.LoadTimetable(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Timetable == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLessonToOneSetOpUserUsingTeacher(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lesson
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lessonDBTypes, false, strmangle.SetComplement(lessonPrimaryKeyColumns, lessonColumnsWithoutDefault)...); err != nil {
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

		if x.R.TeacherLessons[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.TeacherID, x.ID) {
			t.Error("foreign key was wrong value", a.TeacherID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TeacherID))
		reflect.Indirect(reflect.ValueOf(&a.TeacherID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.TeacherID, x.ID) {
			t.Error("foreign key was wrong value", a.TeacherID, x.ID)
		}
	}
}

func testLessonToOneRemoveOpUserUsingTeacher(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lesson
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lessonDBTypes, false, strmangle.SetComplement(lessonPrimaryKeyColumns, lessonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTeacher(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTeacher(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Teacher().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Teacher != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.TeacherID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.TeacherLessons) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testLessonToOneSetOpTimetableUsingTimetable(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lesson
	var b, c Timetable

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lessonDBTypes, false, strmangle.SetComplement(lessonPrimaryKeyColumns, lessonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, timetableDBTypes, false, strmangle.SetComplement(timetablePrimaryKeyColumns, timetableColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Timetable{&b, &c} {
		err = a.SetTimetable(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Timetable != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Lessons[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TimetableID != x.ID {
			t.Error("foreign key was wrong value", a.TimetableID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TimetableID))
		reflect.Indirect(reflect.ValueOf(&a.TimetableID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TimetableID != x.ID {
			t.Error("foreign key was wrong value", a.TimetableID, x.ID)
		}
	}
}

func testLessonsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
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

func testLessonsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LessonSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLessonsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Lessons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	lessonDBTypes = map[string]string{`ID`: `uuid`, `TimetableID`: `uuid`, `Date`: `bigint`, `StartDate`: `bigint`, `EndDate`: `bigint`, `TeacherID`: `character`, `TeacherAttendanceDate`: `bigint`}
	_             = bytes.MinRead
)

func testLessonsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(lessonPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(lessonAllColumns) == len(lessonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLessonsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(lessonAllColumns) == len(lessonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Lesson{}
	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lessonDBTypes, true, lessonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(lessonAllColumns, lessonPrimaryKeyColumns) {
		fields = lessonAllColumns
	} else {
		fields = strmangle.SetComplement(
			lessonAllColumns,
			lessonPrimaryKeyColumns,
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

	slice := LessonSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLessonsUpsert(t *testing.T) {
	t.Parallel()

	if len(lessonAllColumns) == len(lessonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Lesson{}
	if err = randomize.Struct(seed, &o, lessonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Lesson: %s", err)
	}

	count, err := Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, lessonDBTypes, false, lessonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Lesson struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Lesson: %s", err)
	}

	count, err = Lessons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
