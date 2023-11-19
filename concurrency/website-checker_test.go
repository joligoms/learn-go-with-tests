package concurrency

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://evrybadee.huuuuurts" {
		return false
	}

	return true
}

func getRandomMilisecondQuantity() int {
	min := 20
	max := 400

	randomNumber := rand.Intn(max-min+1) + min

	return randomNumber
}

func slowStubWebsiteChecker(_ string) bool {
	randomMilisecondNum := getRandomMilisecondQuantity()
	time.Sleep(time.Duration(randomMilisecondNum) * time.Millisecond)

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://evrybadee.huuuuurts",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://evrybadee.huuuuurts": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
