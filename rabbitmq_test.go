package main

import (
  "testing"
  "fmt"
  "github.com/stretchr/testify/assert"
  "net/http/httptest"
  "net/http"
  "github.com/jeffail/gabs"
)

func TestRabbitMQReport(t *testing.T) {
  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `
[{"memory":7736,"messages":0,"messages_details":{"rate":0.0},"messages_ready":0,"messages_ready_details":{"rate":0.0},"messages_unacknowledged":0,"messages_unacknowledged_details":{"rate":0.0},"idle_since":"2015-10-05 22:14:57","consumer_utilisation":"","policy":"","exclusive_consumer_tag":"","consumers":0,"recoverable_slaves":"","state":"running","messages_ram":0,"messages_ready_ram":0,"messages_unacknowledged_ram":0,"messages_persistent":0,"message_bytes":0,"message_bytes_ready":0,"message_bytes_unacknowledged":0,"message_bytes_ram":0,"message_bytes_persistent":0,"disk_reads":0,"disk_writes":0,"backing_queue_status":{"q1":0,"q2":0,"delta":["delta","undefined",0,"undefined"],"q3":0,"q4":0,"len":0,"target_ram_count":"infinity","next_seq_id":0,"avg_ingress_rate":0.0,"avg_egress_rate":0.0,"avg_ack_ingress_rate":0.0,"avg_ack_egress_rate":0.0},"name":"my-new-queue","vhost":"/","durable":false,"auto_delete":false,"arguments":{},"node":"rabbit@localhost"},{"memory":11256,"message_stats":{"deliver_get":3,"deliver_get_details":{"rate":0.0},"get":1,"get_details":{"rate":0.0},"get_no_ack":2,"get_no_ack_details":{"rate":0.0},"publish":3,"publish_details":{"rate":0.0}},"messages":1,"messages_details":{"rate":0.0},"messages_ready":1,"messages_ready_details":{"rate":0.0},"messages_unacknowledged":0,"messages_unacknowledged_details":{"rate":0.0},"idle_since":"2015-10-05 22:14:58","consumer_utilisation":"","policy":"","exclusive_consumer_tag":"","consumers":0,"recoverable_slaves":"","state":"running","messages_ram":1,"messages_ready_ram":1,"messages_unacknowledged_ram":0,"messages_persistent":0,"message_bytes":12,"message_bytes_ready":12,"message_bytes_unacknowledged":0,"message_bytes_ram":12,"message_bytes_persistent":0,"disk_reads":0,"disk_writes":0,"backing_queue_status":{"q1":0,"q2":0,"delta":["delta","undefined",0,"undefined"],"q3":0,"q4":1,"len":1,"target_ram_count":"infinity","next_seq_id":3,"avg_ingress_rate":0.02539109711017337,"avg_egress_rate":0.012695759639972454,"avg_ack_ingress_rate":0.012695337470200914,"avg_ack_egress_rate":0.0},"name":"myQueue","vhost":"/","durable":true,"auto_delete":false,"arguments":{},"node":"rabbit@localhost"}]`)
  })
  ts := httptest.NewServer(handler)
  defer ts.Close();

  rabbitmq := RabbitMQ{"localhost", 15672, "guest", "guest"};
  data, err := rabbitmq.Monitor(ts.URL);
  assert.Equal(t, err, nil)

  t.Logf("Report: %v", data)

  jsonParsed, err := gabs.ParseJSON([]byte(data))
  assert.Equal(t, err, nil)

  children, _ := jsonParsed.S("report").S("items").Children();
  assert.Equal(t, len(children), 2)
}

func TestRabbitMQQueueStats(t *testing.T) {
  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `
[{"memory":7736,"messages":0,"messages_details":{"rate":0.0},"messages_ready":0,"messages_ready_details":{"rate":0.0},"messages_unacknowledged":0,"messages_unacknowledged_details":{"rate":0.0},"idle_since":"2015-10-05 22:14:57","consumer_utilisation":"","policy":"","exclusive_consumer_tag":"","consumers":0,"recoverable_slaves":"","state":"running","messages_ram":0,"messages_ready_ram":0,"messages_unacknowledged_ram":0,"messages_persistent":0,"message_bytes":0,"message_bytes_ready":0,"message_bytes_unacknowledged":0,"message_bytes_ram":0,"message_bytes_persistent":0,"disk_reads":0,"disk_writes":0,"backing_queue_status":{"q1":0,"q2":0,"delta":["delta","undefined",0,"undefined"],"q3":0,"q4":0,"len":0,"target_ram_count":"infinity","next_seq_id":0,"avg_ingress_rate":0.0,"avg_egress_rate":0.0,"avg_ack_ingress_rate":0.0,"avg_ack_egress_rate":0.0},"name":"my-new-queue","vhost":"/","durable":false,"auto_delete":false,"arguments":{},"node":"rabbit@localhost"},{"memory":11256,"message_stats":{"deliver_get":3,"deliver_get_details":{"rate":0.0},"get":1,"get_details":{"rate":0.0},"get_no_ack":2,"get_no_ack_details":{"rate":0.0},"publish":3,"publish_details":{"rate":0.0}},"messages":1,"messages_details":{"rate":0.0},"messages_ready":1,"messages_ready_details":{"rate":0.0},"messages_unacknowledged":0,"messages_unacknowledged_details":{"rate":0.0},"idle_since":"2015-10-05 22:14:58","consumer_utilisation":"","policy":"","exclusive_consumer_tag":"","consumers":0,"recoverable_slaves":"","state":"running","messages_ram":1,"messages_ready_ram":1,"messages_unacknowledged_ram":0,"messages_persistent":0,"message_bytes":12,"message_bytes_ready":12,"message_bytes_unacknowledged":0,"message_bytes_ram":12,"message_bytes_persistent":0,"disk_reads":0,"disk_writes":0,"backing_queue_status":{"q1":0,"q2":0,"delta":["delta","undefined",0,"undefined"],"q3":0,"q4":1,"len":1,"target_ram_count":"infinity","next_seq_id":3,"avg_ingress_rate":0.02539109711017337,"avg_egress_rate":0.012695759639972454,"avg_ack_ingress_rate":0.012695337470200914,"avg_ack_egress_rate":0.0},"name":"myQueue","vhost":"/","durable":true,"auto_delete":false,"arguments":{},"node":"rabbit@localhost"}]`)
  })
  ts := httptest.NewServer(handler)
  defer ts.Close();

  rabbitmq := RabbitMQ{"localhost", 15672, "guest", "guest"};
  data, err := rabbitmq.QueueStats(ts.URL);
  assert.Equal(t, err, nil)

  t.Logf("Report: %v", data)

  jsonParsed, err := gabs.ParseJSON([]byte(data))
  assert.Equal(t, err, nil)

  children, _ := jsonParsed.S("result").Children();
  memory, _ := children[1].S("memory").Data().(float64);
  assert.Equal(t, float64(11256), memory)
}
