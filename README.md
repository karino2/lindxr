# lindxr
Large file simple line indexer intended t used from slow LL.


## Usage example:

### Indexing mached lines

```
lindxr index -indexdest index/grants2012 -pattern "<doc-number>" -target "../data/grants2012/ipg12011*.xml"
```

### For subsecting large file

```
lindxr sub -start 1 -end 5 -input "../data/grants2012/ipg120110.xml"
```

### Python sample

```
import subprocess

def build_index(indexdest, pattern, targetpat):
    subprocess.call(["lindxr", "index", "-indexdest", indexdest, "-pattern", pattern, "-target", targetpat])
```

