:: This only work on my machine. 
:: bg is my script for running in bakground like & in bash
:: watchrun can be downloaded, 
:: clear is part of githubs suit
:: gchk just my script for running  2 extra linter, one error check and code examinator
@watchrun -care *.go -interval 1700ms bg "go vet ./... && clear & go test ./... && clear & gchk ./..."
