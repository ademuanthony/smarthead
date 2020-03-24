// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Periods", testPeriods)
	t.Run("Students", testStudents)
	t.Run("Subjects", testSubjects)
	t.Run("Teachers", testTeachers)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Periods", testPeriodsDelete)
	t.Run("Students", testStudentsDelete)
	t.Run("Subjects", testSubjectsDelete)
	t.Run("Teachers", testTeachersDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Periods", testPeriodsQueryDeleteAll)
	t.Run("Students", testStudentsQueryDeleteAll)
	t.Run("Subjects", testSubjectsQueryDeleteAll)
	t.Run("Teachers", testTeachersQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Periods", testPeriodsSliceDeleteAll)
	t.Run("Students", testStudentsSliceDeleteAll)
	t.Run("Subjects", testSubjectsSliceDeleteAll)
	t.Run("Teachers", testTeachersSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Periods", testPeriodsExists)
	t.Run("Students", testStudentsExists)
	t.Run("Subjects", testSubjectsExists)
	t.Run("Teachers", testTeachersExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Periods", testPeriodsFind)
	t.Run("Students", testStudentsFind)
	t.Run("Subjects", testSubjectsFind)
	t.Run("Teachers", testTeachersFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Periods", testPeriodsBind)
	t.Run("Students", testStudentsBind)
	t.Run("Subjects", testSubjectsBind)
	t.Run("Teachers", testTeachersBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Periods", testPeriodsOne)
	t.Run("Students", testStudentsOne)
	t.Run("Subjects", testSubjectsOne)
	t.Run("Teachers", testTeachersOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Periods", testPeriodsAll)
	t.Run("Students", testStudentsAll)
	t.Run("Subjects", testSubjectsAll)
	t.Run("Teachers", testTeachersAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Periods", testPeriodsCount)
	t.Run("Students", testStudentsCount)
	t.Run("Subjects", testSubjectsCount)
	t.Run("Teachers", testTeachersCount)
	t.Run("Users", testUsersCount)
}

func TestInsert(t *testing.T) {
	t.Run("Periods", testPeriodsInsert)
	t.Run("Periods", testPeriodsInsertWhitelist)
	t.Run("Students", testStudentsInsert)
	t.Run("Students", testStudentsInsertWhitelist)
	t.Run("Subjects", testSubjectsInsert)
	t.Run("Subjects", testSubjectsInsertWhitelist)
	t.Run("Teachers", testTeachersInsert)
	t.Run("Teachers", testTeachersInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PeriodToStudents", testPeriodToManyStudents)
	t.Run("StudentToPeriods", testStudentToManyPeriods)
	t.Run("StudentToSubjects", testStudentToManySubjects)
	t.Run("SubjectToTeachers", testSubjectToManyTeachers)
	t.Run("SubjectToStudents", testSubjectToManyStudents)
	t.Run("TeacherToSubjects", testTeacherToManySubjects)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PeriodToStudents", testPeriodToManyAddOpStudents)
	t.Run("StudentToPeriods", testStudentToManyAddOpPeriods)
	t.Run("StudentToSubjects", testStudentToManyAddOpSubjects)
	t.Run("SubjectToTeachers", testSubjectToManyAddOpTeachers)
	t.Run("SubjectToStudents", testSubjectToManyAddOpStudents)
	t.Run("TeacherToSubjects", testTeacherToManyAddOpSubjects)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("PeriodToStudents", testPeriodToManySetOpStudents)
	t.Run("StudentToPeriods", testStudentToManySetOpPeriods)
	t.Run("StudentToSubjects", testStudentToManySetOpSubjects)
	t.Run("SubjectToTeachers", testSubjectToManySetOpTeachers)
	t.Run("SubjectToStudents", testSubjectToManySetOpStudents)
	t.Run("TeacherToSubjects", testTeacherToManySetOpSubjects)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("PeriodToStudents", testPeriodToManyRemoveOpStudents)
	t.Run("StudentToPeriods", testStudentToManyRemoveOpPeriods)
	t.Run("StudentToSubjects", testStudentToManyRemoveOpSubjects)
	t.Run("SubjectToTeachers", testSubjectToManyRemoveOpTeachers)
	t.Run("SubjectToStudents", testSubjectToManyRemoveOpStudents)
	t.Run("TeacherToSubjects", testTeacherToManyRemoveOpSubjects)
}

func TestReload(t *testing.T) {
	t.Run("Periods", testPeriodsReload)
	t.Run("Students", testStudentsReload)
	t.Run("Subjects", testSubjectsReload)
	t.Run("Teachers", testTeachersReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Periods", testPeriodsReloadAll)
	t.Run("Students", testStudentsReloadAll)
	t.Run("Subjects", testSubjectsReloadAll)
	t.Run("Teachers", testTeachersReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Periods", testPeriodsSelect)
	t.Run("Students", testStudentsSelect)
	t.Run("Subjects", testSubjectsSelect)
	t.Run("Teachers", testTeachersSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Periods", testPeriodsUpdate)
	t.Run("Students", testStudentsUpdate)
	t.Run("Subjects", testSubjectsUpdate)
	t.Run("Teachers", testTeachersUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Periods", testPeriodsSliceUpdateAll)
	t.Run("Students", testStudentsSliceUpdateAll)
	t.Run("Subjects", testSubjectsSliceUpdateAll)
	t.Run("Teachers", testTeachersSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}