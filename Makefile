.PHONY: crash

crash:
	go test -bench=. -cpu=2 2>&1 |head
