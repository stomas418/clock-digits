package main

import "testing"

func Test1(t *testing.T) {
	result := solution(1, 8, 3, 2)
	expected := 6
	if result != expected {
		t.Errorf("Expected (%d) but returned (%d)", expected, result)
	} else {
		t.Log("TEST1 SUCCESS")
	}
}

func Test2(t *testing.T) {
	result := solution(2, 3, 3, 2)
	expected := 3
	if result != expected {
		t.Errorf("Expected (%d) but returned (%d)", expected, result)
	} else {
		t.Log("TEST2 SUCCESS")
	}
}

func Test3(t *testing.T) {
	result := solution(6, 2, 4, 7)
	expected := 0
	if result != expected {
		t.Errorf("Expected (%d) but returned (%d)", expected, result)
	} else {
		t.Log("TEST3 SUCCESS")
	}

}

func Test4(t *testing.T) {
	result := solution(1, 1, 1, 11)
	expected := 0
	if result != expected {
		t.Errorf("Expected (%d) but returned (%d)", expected, result)
	} else {
		t.Log("TEST4 SUCCESS")
	}

}
