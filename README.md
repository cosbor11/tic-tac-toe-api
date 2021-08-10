# tic-tac-toe-api

## Summary

`tic-tac-toe-api`  is an api service that exposes an endpoint that will always calculate the best move given a game state

## How to Run Locally
  - **Prerequisites**:  
    1. ✔ go go1.16.6 darwin/amd64
    2. ✔ Flutter 2.2.3
    3. ✔ Dart 2.13.4
    4. ✔ Chrome browser
   
  - **Running**:
    1.  ✔  Clone the `tic-tac-toe-api` project and launch the backend service. See the tic-tac-toe-api/README.md for details
    2.  ✔  run `go build` to install the dependencies
    3.  ✔  run: `go run server.go` to launch the application

  - **User Interface**:
     -  see: https://github.com/cosbor11/tic_tac_toe_ui

## Endpoints

  - **POST /next-move**:

Example Request:
```
curl --location --request POST 'http://localhost:5500/next-move' \
--header 'Content-Type: application/json' \
--data-raw '{
    "player": "O",
    "moves": [{
        "player": "X",
        "cell": "A1"
    },{
        "player": "O",
        "cell": "B3"
    }]
}'
```

Example Response:
```
{
    "player": "O",
    "cell": "A2"
}
```


@author Chris Osborn






