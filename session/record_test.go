package session

import "testing"

func TestRecord(t *testing.T) {
	s := NewTestSession()
	t.Log(s.Insert(&User{Name: "Jack", Age: 18}))
	t.Log(s.Insert(&User{Name: "Bob", Age: 25}))
}
