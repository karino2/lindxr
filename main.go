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

			// fmt.Printf("%d:%s\n", count, buf)
			// fmt.Printf("%d:%s\n", count, string(buf))
			// fmt.Printf("%d\n", count)
		}
    }

}

func indexing(arguments []string)  {
	indexCommand := flag.NewFlagSet("index", flag.ExitOnError)

	targetpat := indexCommand.String("target", "", "Trget pattern to index. Handled by Glob. ex: ../data/grants2012/ipg*.xml")
	indexdest := indexCommand.String("indexdest", "", "destination of index directory. ex: ../data/index")
	pat := indexCommand.String("pattern", "", `Pattern to index with regexp. ex."<doc-number>"`)

	indexCommand.Parse(arguments)

	
	if *targetpat == "" || *indexdest == "" || *pat == "" {
		fmt.Println(`Invalid argument for command "index"`)
		indexCommand.PrintDefaults()
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

func subsect(arguments []string)  {
	subCommand := flag.NewFlagSet("sub", flag.ExitOnError)
	subCommand.Parse(arguments)
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("usage: lindxr <command> [<args>]")
		fmt.Println("Commands are: ")
		fmt.Println(" index Start indexing.")
		fmt.Println(" sub  Retrieve part of files specified line range.")
		return
	}



	switch os.Args[1] {
	case "index":
		indexing(os.Args[2:])
	case "sub":
		subsect(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}


}