# Balance Service in Go


## Build

make rebuild_app

## Run

make run

## Run tests

make test

## API requests

### Add balance

```
curl -X "POST" "http://localhost:8000/balances/" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
    "account_id": 1
}'
```
### List of balances

```
curl "http://localhost:8000/balances" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Make transaction

```
curl -x "POST" "http://localhost:8080/balances/transaction" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
     -d $'{
    "from_id": 1,
    "to_id": 2,
    "value": 4
}'
```


## CMD

### Make transaction

```
docker exec -it web ./transaction -t {transID} -r {resID} -v {value}
```

### Example

```
docker exec -it web ./transaction -t 1 -r 2 -v 5
```
