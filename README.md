# distrbuted_system_lab


tidy packages
```
go mod tidy
```

run test:
-v: for verbose, to show log output
-count=1: for not to use cache
```
go test -race -v -count=1 ./...
```
