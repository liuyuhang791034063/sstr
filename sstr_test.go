package sstr

import (
	"testing"
)

func TestSstr_RemoveAllSpaces(t *testing.T) {
	s := Sstr{"  1  1 2 3 2  "}
	s.RemoveAllSpaces()
	if s.Val() != "11232" {
		t.Error("Not Remove All Spaces.")
	}
}

func TestSstr_Clean(t *testing.T) {
	s := Sstr{"123"}
	s.Clean()
	if s.Val() != "" {
		t.Error("Not Clean Strings.")
	}
}

func TestSstr_KeepNumbers(t *testing.T) {
	s := Sstr{"abc123"}
	s.KeepNumbers()
	if s.Val() != "123" {
		t.Error("Not Keep Numbers.")
	}
}

func TestSstr_KeepUpperChar(t *testing.T) {
	s := Sstr{"ASHdwadwa213"}
	s.KeepUpperChar()
	if s.Val() != "ASH" {
		t.Error("Not Keep Upper Char.")
	}
}

func TestSstr_KeepLowerChar(t *testing.T) {
	s := Sstr{"ASBabc123"}
	s.KeepLowerChar()
	if s.Val() != "abc" {
		t.Error("Not Keep Lower Char.")
	}
}

func TestSstr_KeepSpecificChar(t *testing.T) {
	s1 := Sstr{"ASBabc123"}
	s1.KeepSpecificChar(true,false,false)
	if s1.Val() != "123" {
		t.Error("Not Keep Specific Char.")
	}
	s2 := Sstr{"ASBabc123"}
	s2.KeepSpecificChar(false,true,false)
	if s2.Val() != "ASB" {
		t.Error("Not Keep Specific Char.")
	}
	s3 := Sstr{"ASBabc123"}
	s3.KeepSpecificChar(false,false,true)
	if s3.Val() != "abc" {
		t.Error("Not Keep Specific Char.")
	}
}