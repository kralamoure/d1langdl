package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
)

const (
	programName        = "d1langdl"
	programDescription = "d1langdl is a program for downloading Dofus 1 lang files."
	programMoreInfo    = "https://github.com/kralamoure/d1langdl"
)

var (
	printHelp  bool
	debug      bool
	dataUrlStr string
	languages  []string
)

var (
	outDir       = "out"
	allLanguages = []string{"de", "en", "es", "fr", "it", "nl", "pt"}
	client       = &http.Client{}
	wg           *sync.WaitGroup
	dataUrl      url.URL
	downloadErr  uint32 // laziness
)

var (
	flagSet *pflag.FlagSet
	logger  *zap.SugaredLogger
)

func main() {
	l := log.New(os.Stderr, "", 0)

	initFlagSet()
	err := flagSet.Parse(os.Args)
	if err != nil {
		l.Println(err)
		os.Exit(2)
	}

	if printHelp {
		fmt.Println(help(flagSet.FlagUsages()))
		return
	}

	if debug {
		tmp, err := zap.NewDevelopment()
		if err != nil {
			l.Println(err)
			os.Exit(1)
		}
		logger = tmp.Sugar()
	} else {
		tmp, err := zap.NewProduction()
		if err != nil {
			l.Println(err)
			os.Exit(1)
		}
		logger = tmp.Sugar()
	}

	rand.Seed(time.Now().UnixNano())

	err = run()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// laziness
	if downloadErr != 0 {
		os.Exit(1)
	}
}

func run() error {
	defer logger.Sync()

	dataUrlP, err := url.Parse(dataUrlStr)
	if err != nil {
		return err
	}
	dataUrlP.RawQuery = ""
	dataUrl = *dataUrlP

	wg = &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		download("lang/versions.swf", false)
	}()

	for _, language := range languages {
		filename := fmt.Sprintf("lang/versions_%s.txt", language)
		wg.Add(1)
		go func() {
			defer wg.Done()
			download(filename, true)
		}()
	}

	wg.Wait()

	return nil
}

func download(filename string, langVersion bool) {
	var err error
	defer func() {
		if err != nil {
			atomic.StoreUint32(&downloadErr, 1)
			logger.Errorw(errors.Wrap(err, "could not download file").Error(),
				"file_name", filename,
			)
		}
	}()

	dir := filepath.Dir(filename)
	err = os.MkdirAll(filepath.Join(outDir, dir), 0775)
	if err != nil {
		err = errors.Wrap(err, "could not create directory")
		return
	}

	u := dataUrl
	u.Path = path.Join(u.Path, filename)
	u.RawQuery = fmt.Sprintf("v=%d", rand.Int31())

	data, err := get(u.String())
	if err != nil {
		err = errors.Wrap(err, "could not get url")
	}

	err = ioutil.WriteFile(filepath.Join(outDir, filename), data, 0664)
	if err != nil {
		err = errors.Wrap(err, "could not write to file")
	}

	logger.Infow("downloaded file",
		"url", u.String(),
	)

	if !langVersion {
		return
	}

	query, err := url.ParseQuery(string(data))
	if err != nil {
		logger.Errorw("could not parse query",
			"filename", filename,
		)
		return
	}

	s := query.Get("f")
	sli := strings.Split(s, "|")
	for _, langFilename := range sli {
		if langFilename == "" {
			continue
		}
		sli := strings.Split(langFilename, ",")
		if len(sli) != 3 {
			err = errors.New("malformed lang file name")
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			download(fmt.Sprintf("lang/swf/%s_%s_%s.swf", sli[0], sli[1], sli[2]), false)
		}()
	}
}

func get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Shockwave Flash")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	p, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func help(flagUsages string) string {
	buf := &buffer.Buffer{}
	fmt.Fprintf(buf, "%s\n\n", programDescription)
	fmt.Fprintf(buf, "Find more information at: %s\n\n", programMoreInfo)
	fmt.Fprint(buf, "Options:\n")
	fmt.Fprintf(buf, "%s\n", flagUsages)
	fmt.Fprintf(buf, "Usage: %s [options]", programName)
	return buf.String()
}

func initFlagSet() {
	flagSet = pflag.NewFlagSet("d1login", pflag.ContinueOnError)
	flagSet.BoolVarP(&printHelp, "help", "h", false, "Print usage information")
	flagSet.BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
	flagSet.StringVarP(&dataUrlStr, "url", "u", "https://dofusretro.cdn.ankama.com/", "Data URL")
	flagSet.StringSliceVarP(&languages, "languages", "l", allLanguages, "Language codes, separated by comma")
	flagSet.SortFlags = false
}
