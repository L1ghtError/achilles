.PHONY: run

run:
	seq 1 100 | xargs -n 1 -P 100 -I {} sh -c 'curl "http://localhost:8080/auction?sourceID=1&maxDuration=32"; echo'