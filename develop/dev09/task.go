package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

/*
	Реализовать утилиту Wget с возможностью скачивать сайты целиком.
	https://www.gnu.org/software/wget/manual/wget.html
*/

// Args is a struct that holds args for Wget
type Args struct {
	o         []string
	maxDepth  int
	addresses []string
}

// GetArgs is a function that returns parsed args
func GetArgs() (*Args, error) {
	o := flag.String("O", "", "new filename")
	maxDepth := flag.Int("depth", 1, "sets the maximum depth for recursively loading the entire site")

	flag.Parse()

	var files []string

	if *maxDepth < 1 {
		return nil, errors.New("depth should be positive")
	}

	if len(*o) > 0 {
		files = strings.Split(*o, " ")
	}

	args := &Args{
		o:        files,
		maxDepth: *maxDepth,
	}

	if len(flag.Args()) < 1 {
		return nil, errors.New("wrong address")
	}

	args.addresses = append(args.addresses, flag.Args()...)

	return args, nil
}

// GetFileName is a function that sets filename and suffix for existed one
func GetFileName(filename, address, path string) string {
	var counter int

	if filename == "" {
		if strings.HasSuffix(address, "/") {
			filename = "index.html"
		} else {
			filename = filepath.Base(address)
			if !strings.Contains(filename, ".") {
				filename = fmt.Sprintf("%s.html", filename)
			}
		}
	}

	originalName := filename

	for {
		_, err := os.Stat(fmt.Sprintf("%s/%s", path, filename))

		if errors.Is(err, os.ErrNotExist) {
			break
		}

		counter++
		suffix := fmt.Sprintf(".%d", counter)

		newFilename := strings.Builder{}
		newFilename.Grow(len(originalName) + len(suffix))
		newFilename.WriteString(originalName)
		newFilename.WriteString(suffix)

		filename = newFilename.String()
	}

	return filename
}

// SaveToFile is a function that saving body to filename
func SaveToFile(body []byte, filename, address string) error {
	parsed, err := url.Parse(address)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s%s", parsed.Host, parsed.Path)
	path = filepath.Dir(path)

	err = os.MkdirAll(path, 0644)
	if err != nil && os.IsNotExist(err) {
		return err
	}

	filename = GetFileName(filename, address, path)

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", path, filename), body, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("%s - '%s'saved\n\n", time.Now().Format("01/02/06 15:04:05"), filename)

	return nil
}

// GetLinks is a function that returns all links from page
func GetLinks(address string) map[string]bool {
	links := make(map[string]bool)

	parsed, _ := url.Parse(address)
	host := parsed.Hostname()

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(address)
	if err != nil || resp == nil {
		return nil
	}
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	document.Find("a").Each(func(index int, element *goquery.Selection) {
		link, _ := element.Attr("href")
		parsed, err := url.Parse(link)
		if err != nil || parsed.Path == "" {
			return
		}

		linkHost := parsed.Hostname()
		if linkHost != "" && linkHost != host {
			return
		}

		scheme := "https"
		if parsed.Scheme != "" {
			scheme = parsed.Scheme
		}

		newLink := fmt.Sprintf("%s://%s%s", scheme, host, parsed.Path)

		if !uniqueLinks[newLink] {
			links[newLink] = true
			uniqueLinks[newLink] = true
		}
	})

	return links
}

// Download is a function that downloads whole site with goquery tool
func Download(address, filename string, maxDepth int) error {
	if maxDepth < 1 {
		return nil
	}

	fmt.Printf("--%s--  %s\n", time.Now().Format("01/02/06 15:04:05"), address)
	fmt.Printf("Sending http request to %s\n", address)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(address)
	if err != nil || resp == nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("response status: %s", resp.Status)
	}

	fmt.Printf("Response status: %s - Connection established\n", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = SaveToFile(body, filename, address)
	if err != nil {
		return err
	}

	if maxDepth-1 > 1 {
		links := GetLinks(address)

		for link := range links {
			err = Download(link, filename, maxDepth-1)
			if err != nil {
				continue
			}
		}
	}

	return nil
}

var uniqueLinks map[string]bool

// Wget is a function that Download websites/pages
func Wget() error {
	if len(os.Args) < 2 {
		return errors.New("you need to specify a webaddress")
	}

	args, err := GetArgs()
	if err != nil {
		return err
	}

	uniqueLinks = make(map[string]bool)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGSEGV)

	go func() {
		<-sigs
		fmt.Println("Stopped by exceeded time")
		os.Exit(1)
	}()

	for i, address := range args.addresses {
		var filename string

		if i < len(args.o) {
			filename = args.o[i]
		}

		err = Download(address, filename, args.maxDepth)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := Wget()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
