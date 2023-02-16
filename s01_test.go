package structs

import (
	"reflect"
	"testing"
)

type fakeUser struct {
	fName string
}

func (f fakeUser) SetFirstName(_ string) {

}

func (f fakeUser) SetLastName(_ string) {

}

func (f fakeUser) FullName() string {
	return f.fName
}

func TestNew(t *testing.T) {
	u := New()
	u2 := User{}
	if reflect.TypeOf(u) != reflect.TypeOf(u2) {
		t.Error("Constructor New() expects to create User object")
	}
}

func TestUser_NoExported(t *testing.T) {
	u := New()
	v := reflect.ValueOf(u)
	exported := false
	for _, v := range reflect.VisibleFields(v.Type()) {
		exported = exported || v.IsExported()
	}
	if exported {
		t.Error("User struct expected to have no exported fields")
	}
}

func TestIsUser_True(t *testing.T) {
	u := New()
	ok := IsUser(u)
	if !ok {
		t.Error("IsUser must TRUE if User was passed")
	}
}

func TestIsUser_False(t *testing.T) {
	u := 10
	ok := IsUser(u)
	if ok {
		t.Error("IsUser must FALSE if NOT User was passed")
	}
}

func TestUser_FullName(t *testing.T) {
	u := New()
	u.SetFirstName("qq")
	u.SetLastName("ww")
	if u.FullName() != "ww qq" {
		t.Error("FullName expected to return 'LastName FirstName' string")
	}
}

func TestResetUser(t *testing.T) {
	u := New()
	u.SetFirstName("qq")
	u.SetLastName("ww")
	ResetUser(&u)
	if u.FullName() != " " {
		t.Error("FullName expected to return ' ' for reset user string")
	}
}

func TestProcessUser(t *testing.T) {
	u := fakeUser{fName: "faked user"}
	if ProcessUser(u) != u.fName {
		t.Error("Process user must accept UserInterface object")
	}
}
