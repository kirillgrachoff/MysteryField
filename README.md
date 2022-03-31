# Mystery Field
This is client-server application to play _mystery field_ game.
Player should guess whole secret word using 1-character requests.
If the character exists in secret word, it becomes visible.

## Requirements
`go 1.18+`

## Use
Modify `PORT_`, `HOST_`, `SECRET_WORD_` (must contain only ASCII-characters), `FUEL_BUDGET_` values before start.

Then type `make run_server` on `$HOST_` machine.

Type `make run_client` on your local machine and play the game.

## Implementation details
This is application written in pure Go. It uses gRPC framework and Protobuf.
