package main

import "testing"

func Test_options_builder(t *testing.T) {
	bigClass1 := NewBigClass()
	bigClass2 := NewBigClass(WithAge(20), WithName("小乌龟"))
	bigClass3 := NewBigClass(WithHeight(180), WithWidth(200), WithFieldA("aaa"), WithFieldB("bbb"))
	t.Logf("bigClass1: %+v", bigClass1)
	t.Logf("bigClass2: %+v", bigClass2)
	t.Logf("bigClass3: %+v", bigClass3)
}
