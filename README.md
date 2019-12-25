# crypto-monitor
## Description
**crypto-monitor** is a simple service to monitor cryptocurrencies conversion rates on different markets.
## How to run
### Build Docker image
`docker build -t crypto-monitor .`
### Setup work directory
1. Create work directory `mkdir path/to/work/dir`
2. Open work directory `cd path/to/work/dir`
3. Create config.json
```
echo '{
    "DbPath": "/work/rates.db",
    "RatesUpdateFreq": 10,
    "ObservedSymbols": [
        "ETH/BTC",
        "BTC/USDT",
        "BTC/ETH"
    ]
}' > config.json
```
### Run Docker image
```
docker run \
-p 11000:80 \
-v $(pwd):/work \
-it --rm --name crypto-monitor \
crypto-monitor \
-config /work/config.json
```
### Test
Open `http://localhost:11000/get-rates` in your browser or run `curl http://localhost:11000/get-rates`
You should see the correct response
