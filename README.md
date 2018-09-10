# lindxr
Large file simple line indexer intended to use from slow LL.

Match specifierd pattern line by line and print matched line with linenumber like grep -n.

## Usage example:

```
lindxr -indexdest index/grants2012 -pattern "<doc-number>" -target "../data/grants2012/ipg12011*.xml"
```

### Result index example:

```
8:<doc-number>D0651376</doc-number>
16:<doc-number>29390372</doc-number>
38:<doc-number>D11495</doc-number>
50:<doc-number>D50715</doc-number>
62:<doc-number>D74119</doc-number>
```

### Python sample

Build index.

```
import subprocess

def build_index(indexdest, pattern, targetpat):
    subprocess.call(["lindxr", "index", "-indexdest", indexdest, "-pattern", pattern, "-target", targetpat])
```

Result handling.

```
def collect_lines(fpath):
    with open(fpath, "r") as f:
        return [int(l.split(":", 1)[0]) for l in f]

def collect_line_match_tupple(indexfile):
    with open(indexfile, "r") as f:
        return [(int(lnumstr), match) for (lnumstr, match) in (l.rstrip("\n").split(":", 1) for l in f)]
```






