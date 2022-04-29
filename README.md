# cs3probes

Die `cs3probes` implementieren API-Level Tests der CS3APIs als Nagiosproben. Zur Ausreißererkennung werden Perzentile verwendet.

## Kompilieren

### `make`

Alle Binaries nach `bin/`kompilieren:

`>$ make build`
 
Alle Dateien aus `bin/` löschen:

`>$ make clean`


Alle Dateien aus `bin/` löschen und neu kompilieren:

`>$ make rebuild`

### Windows

Unter Windows kann der Effekt von `make build` manuell erzeugt werden:

`go build -o bin ./cmd/...`


## Verwendung
Aktuell sind drei `core`-Proben verfügbar:

````
>$ ./bin/network                                                                                                                                                                                                                                        ST 2   main 

FLAGS:
-percentile int
    the percentile for outlier detection (default 90)
-target string
    [required] the target, [host]:[port]
-warnlimit int
    minimum number of logs for outlier detection (default 100)
````

````
>$ ./bin/fsoperations                                                                                                                                                                                                                               2  ST 2   main 

FLAGS:
  -pass string
        [required] the user password
  -percentile int
        the percentile for outlier detection (default 90)
  -target string
        [required] the target iop
  -user string
        [required] the username
  -warnlimit int
        minimum number of logs for outlier detection (default 100)
````
````
>$ ./bin/fsspeed  

FLAGS:
  -pass string
        [required] the user password
  -percentile int
        the percentile for outlier detection (default 90)
  -target string
        [required] the target iop
  -user string
        [required] the username
  -warnlimit int
        minimum number of logs for outlier detection (default 100)
````

### Histogramme
Histogramme erstellen mit `histograms/histograms.py`

`>$ python3 ./histograms.py /pfad/zur/datenbank.db [98]`

Das optionale Argument lässt ein Perzentil einstellen, default-Wert ist 90.