package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"
)

func TestFindEntries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
		if r.URL.String() == "/" {
			w.Write([]byte(`
				<table summary="作者データ">
				<tr>
					<td>分類</td>
					<td>著者</td>
				</tr>
				<tr>
					<td>作家名</td>
					<td><a href="">芥川 龍之介</a></td>
				</tr>
				</table>
				<ol>
				<li><a href="../cards/999999/card001.html">テスト書籍001</a></li>
				</ol>
			`))
		} else {
			pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+)\.html$`)
			token := pat.FindStringSubmatch(r.URL.String())
			w.Write([]byte(fmt.Sprintf(`
				<table summary="作者データ">
				<tr>
					<td>分類</td>
					<td>著者</td>
				</tr>
				<tr>
					<td>作家名</td>
					<td><a href="">芥川 龍之介</a></td>
				</tr>
				</table>
				<table summary="ダウンロードデータ">
				<tr>
					<td><a href="./files/%[1]s_%[2]s.zip">%[1]s_%[2]s.zip</a></td>
				</tr>
				</table>
			`, token[1], token[2])))
		}
	}))
	defer ts.Close()

	tmp := pageURLFormat
	pageURLFormat = ts.URL + "/cards/%s/card%s.html"
	defer func() {
		pageURLFormat = tmp
	}()

	got, err := findEntries(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}

	want := []Entry{
		{
			AuthorID: "999999",
			Author:   "テスト 太郎",
			TitleID:  "001",
			Title:    "テスト書籍001",
			SiteURL:  ts.URL,
			ZipURL:   ts.URL + "/cards/999999/files/999999_001.zip",
		},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %+v, but got %+v", want, got)
	}
}
