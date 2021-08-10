# tic-tac-toe-api

## Summary

`tic-tac-toe-api`  is an api service that exposes an endpoint that will always calculate the best move given a game state

## How to Run Locally
  - **Prerequisites**:  
    1. ✔ go go1.16.6 darwin/amd64. `brew install golang`
   
  - **Running**:
    1.  ✔  Clone this project: `git clone https://github.com/cosbor11/tic-tac-toe-api.git`
    2.  ✔  Navigate to the root of the project `cd tic-tac-toe-api`
    3.  ✔  run `go build` to install the dependencies
    4.  ✔  run: `go run server.go` to launch the application

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






