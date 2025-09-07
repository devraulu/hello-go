package main

import (
	"reflect"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {

}

func mockWebsiteChecker(url string) bool {
	return url == "http://www.nerv.org.jp"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"rrr://hostname.xxx",
		"http://google.com",
		"http://www.nerv.org.jp",
	}

	want := map[string]bool{
		"rrr://hostname.xxx":     false,
		"http://google.com":      false,
		"http://www.nerv.org.jp": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v but wanted %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
