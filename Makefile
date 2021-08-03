.PHONY: mongods
mongods:
	go build -o ./mongods ./cmd/mongods/


.PHONY: cid2key
cid2key:
	go build -o ./cid2key ./cmd/cid2key/

