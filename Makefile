
export GOPATH=$(shell pwd)
export GOBIN=$(shell pwd)/bin
export RABBITMQSBIN="/usr/local/Cellar/rabbitmq/3.5.4/sbin"

all:
	mkdir -p $(GOBIN)
	go get github.com/stretchr/testify/assert
	go get
	go install

clean:
	rm -rf bin

test:
	go test *.go

test.verbose:
	go test *.go -v

simulate:
	cd $(RABBITMQSBIN); ./rabbitmqadmin declare queue name="myQueue" durable=true auto_delete=false
	cd $(RABBITMQSBIN); ./rabbitmqadmin declare exchange name="myExchange" type="direct" auto_delete=false internal=false durable=true
	cd $(RABBITMQSBIN); ./rabbitmqadmin declare binding source="myExchange" destination_type="queue" destination="myQueue" routing_key="myMessage"
	cd $(RABBITMQSBIN); ./rabbitmqadmin publish exchange="myExchange" routing_key="myMessage" payload="hello, world"
	cd $(RABBITMQSBIN); ./rabbitmqadmin get queue="myQueue" requeue=true
