INTERFACES_DIR = ./internal/services/
INTERFACES_EXT_DIR = ./pkg/zincsearch/
MOCK_DIR = ./tests/mock/

.PHONY = grpc-generate grpc-clear mock-generate
mock:
	files=`ls -p $(INTERFACES_DIR) | grep -v /`; for mockfile in $$files ; do mockgen -source $(INTERFACES_DIR)$$mockfile -destination $(MOCK_DIR)$$mockfile -package mock ; done
	files=`ls -p $(INTERFACES_EXT_DIR) | grep -v /`; for mockfile in $$files ; do mockgen -source $(INTERFACES_EXT_DIR)$$mockfile -destination $(MOCK_DIR)$$mockfile -package mock ; done

run:
	export ZINC_EXCHANGE_NAME=actionplan
	export ZINC_ROUTING_KEY=actionplan.logzincsearch
	export ZINC_QUEUE=actionplan.logqueue
	export ZINC_BROKER_URL=amqp://admin:admin@localhost:5672/
	export ZINC_URL=http://localhost:4080/api/%s/_doc
	export ZINC_INDEX=actionplan
	export ZINC_USER=admin
	export ZINC_PASS=admin
	export ZINC_LOG_LEVEL=-1
	export ZINC_VERSION=0.0.1
	go run cmd/main.go
