# lindxr
Large file simple line indexer intended to use from slow LL.

Match specifierd pattern line by line and print matched line with linenumber like grep -n.

## Usage example:

```
lindxr -indexdest index/grants2012 -pattern "<doc-number>" -target "../data/grants2012/ipg12011*.xml"
```


### Python sample

```
import subprocess

def build_index(indexdest, pattern, targetpat):
    subprocess.call(["lindxr", "index", "-indexdest", indexdest, "-pattern", pattern, "-target", targetpat])

```

