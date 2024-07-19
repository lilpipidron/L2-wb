package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// net/http для выполнения HTTP-запросов и получения HTML-страниц
// golang.org/x/net/html для парсинга HTML
// os и path/filepath для работы с фс и создания нужных файлов
// на основе этих пакетов и строится все решение
// скачиваем страницу, парсим ее, достаем все ссылки и парсим уже страницы, которые находятся за ними
func main() {
	url := "https://pkg.go.dev/golang.org/x/net/html#pkg-functions"
	dir := "./download"
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	if err := downloadPage(url, dir); err != nil {
		panic(err)
	}
}

func downloadPage(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	fp := filepath.Join(dir, "index.html")
	file, err := os.Create(fp)
	if err != nil {
		return err
	}

	defer file.Close()

	if err = html.Render(file, doc); err != nil {
		return err
	}

	downloadResources(doc, url, dir)
	return nil
}

func downloadResources(doc *html.Node, url, dir string) {
	if doc.Type == html.ElementNode {
		for _, attr := range doc.Attr {
			if attr.Key == "href" || attr.Key == "src" {
				resourceURL := attr.Val
				if !strings.HasPrefix(resourceURL, "http") {
					resourceURL = url + resourceURL
				}
				downloadResource(resourceURL, dir)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		downloadResources(c, url, dir)
	}
}

func downloadResource(url, dir string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed download resource: ", url, err)
		return
	}

	defer resp.Body.Close()
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]
	fp := filepath.Join(dir, filename)

	file, err := os.Create(fp)
	if err != nil {
		fmt.Println("Failed download resource: ", url, err)
		return
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		fmt.Println("Failed saving resource: ", url, err)
	}
}
