.PHONY: mongods
mongods:
	go build -o ./mongods ./cmd/mongods/


.PHONY: badgerds
badgerds:
	go build -o ./badgerds ./cmd/badgerds/

