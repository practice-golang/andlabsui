## GUI Practcie with Go and ui

Same as the example in andlabs/ui wiki.


## Build
```sh
go build -ldflags="-H windowsgui"
```


## Manifest
```sh
go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

.\bin\goversioninfo -icon="./pd.ico" -manifest="./ui.manifest" -platform-specific=true
```
