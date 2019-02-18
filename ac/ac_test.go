package ac

import "testing"

func TestAc_ALL(t *testing.T) {
	ac := New()

	_ = ac.Add("foo", 1)
	_ = ac.Add("bar", 1)

	ac.BuildFailPointer()

	m := ac.Match("fooobbar")
	t.Logf("%v",m)
	t.Logf("%s",string(42))
}