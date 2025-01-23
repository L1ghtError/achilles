.PHONY: run

run:
	seq 1 200 | xargs -n 1 -P 100 -I {} sh -c 'curl "http://localhost:8080/auction?sourceID=0&maxDuration=0"; echo'
