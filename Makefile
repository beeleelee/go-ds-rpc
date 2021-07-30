.PHONY: mongods
mongods:
	go build -o ./mongods ./cmd/mongods/

.PHONY: mongods2
mongods2:
	go build -o ./mongods2 ./cmd/mongods2/


.PHONY: badgerds
badgerds:
	go build -o ./badgerds ./cmd/badgerds/

