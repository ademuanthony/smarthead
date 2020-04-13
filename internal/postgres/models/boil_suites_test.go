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
	t.Run("Accounts", testAccounts)
	t.Run("Classes", testClasses)
	t.Run("Deposits", testDeposits)
	t.Run("Periods", testPeriods)
	t.Run("Students", testStudents)
	t.Run("Subjects", testSubjects)
	t.Run("Subscriptions", testSubscriptions)
	t.Run("Teachers", testTeachers)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Classes", testClassesDelete)
	t.Run("Deposits", testDepositsDelete)
	t.Run("Periods", testPeriodsDelete)
	t.Run("Students", testStudentsDelete)
	t.Run("Subjects", testSubjectsDelete)
	t.Run("Subscriptions", testSubscriptionsDelete)
	t.Run("Teachers", testTeachersDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Classes", testClassesQueryDeleteAll)
	t.Run("Deposits", testDepositsQueryDeleteAll)
	t.Run("Periods", testPeriodsQueryDeleteAll)
	t.Run("Students", testStudentsQueryDeleteAll)
	t.Run("Subjects", testSubjectsQueryDeleteAll)
	t.Run("Subscriptions", testSubscriptionsQueryDeleteAll)
	t.Run("Teachers", testTeachersQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Classes", testClassesSliceDeleteAll)
	t.Run("Deposits", testDepositsSliceDeleteAll)
	t.Run("Periods", testPeriodsSliceDeleteAll)
	t.Run("Students", testStudentsSliceDeleteAll)
	t.Run("Subjects", testSubjectsSliceDeleteAll)
	t.Run("Subscriptions", testSubscriptionsSliceDeleteAll)
	t.Run("Teachers", testTeachersSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Classes", testClassesExists)
	t.Run("Deposits", testDepositsExists)
	t.Run("Periods", testPeriodsExists)
	t.Run("Students", testStudentsExists)
	t.Run("Subjects", testSubjectsExists)
	t.Run("Subscriptions", testSubscriptionsExists)
	t.Run("Teachers", testTeachersExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Classes", testClassesFind)
	t.Run("Deposits", testDepositsFind)
	t.Run("Periods", testPeriodsFind)
	t.Run("Students", testStudentsFind)
	t.Run("Subjects", testSubjectsFind)
	t.Run("Subscriptions", testSubscriptionsFind)
	t.Run("Teachers", testTeachersFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Classes", testClassesBind)
	t.Run("Deposits", testDepositsBind)
	t.Run("Periods", testPeriodsBind)
	t.Run("Students", testStudentsBind)
	t.Run("Subjects", testSubjectsBind)
	t.Run("Subscriptions", testSubscriptionsBind)
	t.Run("Teachers", testTeachersBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Classes", testClassesOne)
	t.Run("Deposits", testDepositsOne)
	t.Run("Periods", testPeriodsOne)
	t.Run("Students", testStudentsOne)
	t.Run("Subjects", testSubjectsOne)
	t.Run("Subscriptions", testSubscriptionsOne)
	t.Run("Teachers", testTeachersOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Classes", testClassesAll)
	t.Run("Deposits", testDepositsAll)
	t.Run("Periods", testPeriodsAll)
	t.Run("Students", testStudentsAll)
	t.Run("Subjects", testSubjectsAll)
	t.Run("Subscriptions", testSubscriptionsAll)
	t.Run("Teachers", testTeachersAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Classes", testClassesCount)
	t.Run("Deposits", testDepositsCount)
	t.Run("Periods", testPeriodsCount)
	t.Run("Students", testStudentsCount)
	t.Run("Subjects", testSubjectsCount)
	t.Run("Subscriptions", testSubscriptionsCount)
	t.Run("Teachers", testTeachersCount)
	t.Run("Users", testUsersCount)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Classes", testClassesInsert)
	t.Run("Classes", testClassesInsertWhitelist)
	t.Run("Deposits", testDepositsInsert)
	t.Run("Deposits", testDepositsInsertWhitelist)
	t.Run("Periods", testPeriodsInsert)
	t.Run("Periods", testPeriodsInsertWhitelist)
	t.Run("Students", testStudentsInsert)
	t.Run("Students", testStudentsInsertWhitelist)
	t.Run("Subjects", testSubjectsInsert)
	t.Run("Subjects", testSubjectsInsertWhitelist)
	t.Run("Subscriptions", testSubscriptionsInsert)
	t.Run("Subscriptions", testSubscriptionsInsertWhitelist)
	t.Run("Teachers", testTeachersInsert)
	t.Run("Teachers", testTeachersInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AccountToUserUsingBillingUser", testAccountToOneUserUsingBillingUser)
	t.Run("AccountToUserUsingSignupUser", testAccountToOneUserUsingSignupUser)
	t.Run("DepositToClassUsingClass", testDepositToOneClassUsingClass)
	t.Run("DepositToPeriodUsingPeriod", testDepositToOnePeriodUsingPeriod)
	t.Run("DepositToStudentUsingStudent", testDepositToOneStudentUsingStudent)
	t.Run("DepositToSubjectUsingSubject", testDepositToOneSubjectUsingSubject)
	t.Run("StudentToClassUsingClass", testStudentToOneClassUsingClass)
	t.Run("SubscriptionToClassUsingClass", testSubscriptionToOneClassUsingClass)
	t.Run("SubscriptionToDepositUsingDeposit", testSubscriptionToOneDepositUsingDeposit)
	t.Run("SubscriptionToPeriodUsingPeriod", testSubscriptionToOnePeriodUsingPeriod)
	t.Run("SubscriptionToStudentUsingStudent", testSubscriptionToOneStudentUsingStudent)
	t.Run("SubscriptionToSubjectUsingSubject", testSubscriptionToOneSubjectUsingSubject)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ClassToDeposits", testClassToManyDeposits)
	t.Run("ClassToStudents", testClassToManyStudents)
	t.Run("ClassToSubscriptions", testClassToManySubscriptions)
	t.Run("DepositToSubscriptions", testDepositToManySubscriptions)
	t.Run("PeriodToDeposits", testPeriodToManyDeposits)
	t.Run("PeriodToStudents", testPeriodToManyStudents)
	t.Run("PeriodToSubscriptions", testPeriodToManySubscriptions)
	t.Run("StudentToDeposits", testStudentToManyDeposits)
	t.Run("StudentToPeriods", testStudentToManyPeriods)
	t.Run("StudentToSubjects", testStudentToManySubjects)
	t.Run("StudentToSubscriptions", testStudentToManySubscriptions)
	t.Run("SubjectToDeposits", testSubjectToManyDeposits)
	t.Run("SubjectToTeachers", testSubjectToManyTeachers)
	t.Run("SubjectToStudents", testSubjectToManyStudents)
	t.Run("SubjectToSubscriptions", testSubjectToManySubscriptions)
	t.Run("TeacherToSubjects", testTeacherToManySubjects)
	t.Run("UserToBillingUserAccounts", testUserToManyBillingUserAccounts)
	t.Run("UserToSignupUserAccounts", testUserToManySignupUserAccounts)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AccountToUserUsingBillingUserAccounts", testAccountToOneSetOpUserUsingBillingUser)
	t.Run("AccountToUserUsingSignupUserAccounts", testAccountToOneSetOpUserUsingSignupUser)
	t.Run("DepositToClassUsingDeposits", testDepositToOneSetOpClassUsingClass)
	t.Run("DepositToPeriodUsingDeposits", testDepositToOneSetOpPeriodUsingPeriod)
	t.Run("DepositToStudentUsingDeposits", testDepositToOneSetOpStudentUsingStudent)
	t.Run("DepositToSubjectUsingDeposits", testDepositToOneSetOpSubjectUsingSubject)
	t.Run("StudentToClassUsingStudents", testStudentToOneSetOpClassUsingClass)
	t.Run("SubscriptionToClassUsingSubscriptions", testSubscriptionToOneSetOpClassUsingClass)
	t.Run("SubscriptionToDepositUsingSubscriptions", testSubscriptionToOneSetOpDepositUsingDeposit)
	t.Run("SubscriptionToPeriodUsingSubscriptions", testSubscriptionToOneSetOpPeriodUsingPeriod)
	t.Run("SubscriptionToStudentUsingSubscriptions", testSubscriptionToOneSetOpStudentUsingStudent)
	t.Run("SubscriptionToSubjectUsingSubscriptions", testSubscriptionToOneSetOpSubjectUsingSubject)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("AccountToUserUsingBillingUserAccounts", testAccountToOneRemoveOpUserUsingBillingUser)
	t.Run("AccountToUserUsingSignupUserAccounts", testAccountToOneRemoveOpUserUsingSignupUser)
	t.Run("StudentToClassUsingStudents", testStudentToOneRemoveOpClassUsingClass)
	t.Run("SubscriptionToPeriodUsingSubscriptions", testSubscriptionToOneRemoveOpPeriodUsingPeriod)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ClassToDeposits", testClassToManyAddOpDeposits)
	t.Run("ClassToStudents", testClassToManyAddOpStudents)
	t.Run("ClassToSubscriptions", testClassToManyAddOpSubscriptions)
	t.Run("DepositToSubscriptions", testDepositToManyAddOpSubscriptions)
	t.Run("PeriodToDeposits", testPeriodToManyAddOpDeposits)
	t.Run("PeriodToStudents", testPeriodToManyAddOpStudents)
	t.Run("PeriodToSubscriptions", testPeriodToManyAddOpSubscriptions)
	t.Run("StudentToDeposits", testStudentToManyAddOpDeposits)
	t.Run("StudentToPeriods", testStudentToManyAddOpPeriods)
	t.Run("StudentToSubjects", testStudentToManyAddOpSubjects)
	t.Run("StudentToSubscriptions", testStudentToManyAddOpSubscriptions)
	t.Run("SubjectToDeposits", testSubjectToManyAddOpDeposits)
	t.Run("SubjectToTeachers", testSubjectToManyAddOpTeachers)
	t.Run("SubjectToStudents", testSubjectToManyAddOpStudents)
	t.Run("SubjectToSubscriptions", testSubjectToManyAddOpSubscriptions)
	t.Run("TeacherToSubjects", testTeacherToManyAddOpSubjects)
	t.Run("UserToBillingUserAccounts", testUserToManyAddOpBillingUserAccounts)
	t.Run("UserToSignupUserAccounts", testUserToManyAddOpSignupUserAccounts)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("ClassToStudents", testClassToManySetOpStudents)
	t.Run("PeriodToStudents", testPeriodToManySetOpStudents)
	t.Run("PeriodToSubscriptions", testPeriodToManySetOpSubscriptions)
	t.Run("StudentToPeriods", testStudentToManySetOpPeriods)
	t.Run("StudentToSubjects", testStudentToManySetOpSubjects)
	t.Run("SubjectToTeachers", testSubjectToManySetOpTeachers)
	t.Run("SubjectToStudents", testSubjectToManySetOpStudents)
	t.Run("TeacherToSubjects", testTeacherToManySetOpSubjects)
	t.Run("UserToBillingUserAccounts", testUserToManySetOpBillingUserAccounts)
	t.Run("UserToSignupUserAccounts", testUserToManySetOpSignupUserAccounts)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("ClassToStudents", testClassToManyRemoveOpStudents)
	t.Run("PeriodToStudents", testPeriodToManyRemoveOpStudents)
	t.Run("PeriodToSubscriptions", testPeriodToManyRemoveOpSubscriptions)
	t.Run("StudentToPeriods", testStudentToManyRemoveOpPeriods)
	t.Run("StudentToSubjects", testStudentToManyRemoveOpSubjects)
	t.Run("SubjectToTeachers", testSubjectToManyRemoveOpTeachers)
	t.Run("SubjectToStudents", testSubjectToManyRemoveOpStudents)
	t.Run("TeacherToSubjects", testTeacherToManyRemoveOpSubjects)
	t.Run("UserToBillingUserAccounts", testUserToManyRemoveOpBillingUserAccounts)
	t.Run("UserToSignupUserAccounts", testUserToManyRemoveOpSignupUserAccounts)
}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Classes", testClassesReload)
	t.Run("Deposits", testDepositsReload)
	t.Run("Periods", testPeriodsReload)
	t.Run("Students", testStudentsReload)
	t.Run("Subjects", testSubjectsReload)
	t.Run("Subscriptions", testSubscriptionsReload)
	t.Run("Teachers", testTeachersReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Classes", testClassesReloadAll)
	t.Run("Deposits", testDepositsReloadAll)
	t.Run("Periods", testPeriodsReloadAll)
	t.Run("Students", testStudentsReloadAll)
	t.Run("Subjects", testSubjectsReloadAll)
	t.Run("Subscriptions", testSubscriptionsReloadAll)
	t.Run("Teachers", testTeachersReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Classes", testClassesSelect)
	t.Run("Deposits", testDepositsSelect)
	t.Run("Periods", testPeriodsSelect)
	t.Run("Students", testStudentsSelect)
	t.Run("Subjects", testSubjectsSelect)
	t.Run("Subscriptions", testSubscriptionsSelect)
	t.Run("Teachers", testTeachersSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Classes", testClassesUpdate)
	t.Run("Deposits", testDepositsUpdate)
	t.Run("Periods", testPeriodsUpdate)
	t.Run("Students", testStudentsUpdate)
	t.Run("Subjects", testSubjectsUpdate)
	t.Run("Subscriptions", testSubscriptionsUpdate)
	t.Run("Teachers", testTeachersUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Classes", testClassesSliceUpdateAll)
	t.Run("Deposits", testDepositsSliceUpdateAll)
	t.Run("Periods", testPeriodsSliceUpdateAll)
	t.Run("Students", testStudentsSliceUpdateAll)
	t.Run("Subjects", testSubjectsSliceUpdateAll)
	t.Run("Subscriptions", testSubscriptionsSliceUpdateAll)
	t.Run("Teachers", testTeachersSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
