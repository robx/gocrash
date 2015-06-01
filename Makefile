.PHONY: crash

crash:
	go test -bench=. -cpu=4 2>&1 |head
