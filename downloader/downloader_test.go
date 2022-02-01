// downloader_test.go
package downloader

func TestGet(t *testing.T) {
	urlString := "https://velog.io/@kineo2k"

	dl := New(urlString)
	err := dl.Get()
	if err != nil {
		t.Fail()
	}
}
