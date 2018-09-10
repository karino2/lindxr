package main

import (
	"flag"
	"fmt"
    "os"
	"path/filepath"
	"regexp"
	"bufio"
)

func indexOne(fpath string, idxpath string, pat string, rexp *regexp.Regexp) {
    fh, err := os.Open(fpath)

    if err != nil {
		fmt.Printf("Can't open file, skip: %s\n", fpath)
		return
    }

	defer fh.Close()

	fout, err := os.Create(idxpath)
    if err != nil {

		fmt.Printf("Can't create index file, skip: %s\n", idxpath)
		return
    }
	defer fout.Close()


    fr := bufio.NewReader(fh)
	fw := bufio.NewWriter(fout)

	buf := make([]byte, 1024)
	count := 0
    for {
		buf, _ , err = fr.ReadLine()
		if err != nil {
			fmt.Printf("Read %d on %s\n", count, fpath)
			fw.Flush()
			return
		}
		count++

		if rexp.Match(buf) {
			fmt.Fprintf(fw, "%d:%s\n", count, buf)
		}
    }

}


func main() {
	targetpat := flag.String("target", "", "Trget pattern to index. Handled by Glob. ex: ../data/grants2012/ipg*.xml")
	indexdest := flag.String("indexdest", "", "destination of index directory. ex: ../data/index")
	pat := flag.String("pattern", "", `Pattern to index with regexp. ex."<doc-number>"`)

	flag.Parse()

	
	if *targetpat == "" || *indexdest == "" || *pat == "" {
		fmt.Println(`Usave : lindxr`)
		flag.PrintDefaults()
		os.Exit(2)
	}

	rexp := regexp.MustCompile(*pat)
	if rexp == nil {
		fmt.Println(`Invalid regexp for "index"`)
		os.Exit(2)
		
	}

	err := os.MkdirAll(*indexdest, 0744)

	if err != nil {
		fmt.Printf("Can't create index destination: %s\n", *indexdest)
		return
	}


	files, err := filepath.Glob(*targetpat)
    if err != nil {
        panic(err)
    }
    for _, file := range files {
		indexpath := *indexdest +"/" + filepath.Base(file) + ".idx"
		indexOne(file, indexpath, *pat, rexp)
    }

}