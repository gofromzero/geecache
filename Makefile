
.PHONY: proto
proto:
	protoc --go_out=. geecachepb/*.proto
