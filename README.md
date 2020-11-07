# Gonks

CLI app to display stonks

Gets stock data from [finnhub](http://finnhub.io)

```
SYMBOL                 PRICE              $ CHANGE              % CHANGE
AAPL                  118.69                 -0.34              -0.0029%
DIS                   127.46                  0.50               0.0039%
SPY                   350.16                 -0.08              -0.0002%
```

## Example config

Create a `.gonks.json` file in home directory
```
{
  "api_key": "abc-123",
  "stocks": [
    "AAPL",
    "SPY",
    "DIS"
  ]
}
```

## Build executable

Build with `go build` and copy the executable to a $PATH location

```
$ go build
# cp ./gonks /usr/local/bin/
```
