package main

import (
  "testing"
  "fmt"
  "github.com/stretchr/testify/assert"
  "net/http/httptest"
  "net/http"
  "github.com/jeffail/gabs"
)
/*
func TestElasticsearchReport(t *testing.T) {
  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"timestamp":1444099819856,"cluster_name":"elasticsearch_dominiek","status":"green","indices":{"count":0,"shards":{},"docs":{"count":0,"deleted":0},"store":{"size_in_bytes":0,"throttle_time_in_millis":0},"fielddata":{"memory_size_in_bytes":0,"evictions":0},"filter_cache":{"memory_size_in_bytes":0,"evictions":0},"id_cache":{"memory_size_in_bytes":0},"completion":{"size_in_bytes":0},"segments":{"count":0,"memory_in_bytes":0,"index_writer_memory_in_bytes":0,"index_writer_max_memory_in_bytes":0,"version_map_memory_in_bytes":0,"fixed_bit_set_memory_in_bytes":0},"percolate":{"total":0,"time_in_millis":0,"current":0,"memory_size_in_bytes":-1,"memory_size":"-1b","queries":0}},"nodes":{"count":{"total":1,"master_only":0,"data_only":0,"master_data":1,"client":0},"versions":["1.7.2"],"os":{"available_processors":4,"mem":{"total_in_bytes":8589934592},"cpu":[{"vendor":"Intel","model":"MacBook8,1","mhz":1100,"total_cores":4,"total_sockets":4,"cores_per_socket":16,"cache_size_in_bytes":256,"count":1}]},"process":{"cpu":{"percent":0},"open_file_descriptors":{"min":156,"max":156,"avg":156}},"jvm":{"max_uptime_in_millis":49103653,"versions":[{"version":"1.8.0_60","vm_name":"Java HotSpot(TM) 64-Bit Server VM","vm_version":"25.60-b23","vm_vendor":"Oracle Corporation","count":1}],"mem":{"heap_used_in_bytes":67845896,"heap_max_in_bytes":1038876672},"threads":46},"fs":{"total_in_bytes":249678528512,"free_in_bytes":91290705920,"available_in_bytes":91028561920,"disk_reads":0,"disk_writes":0,"disk_io_op":0,"disk_read_size_in_bytes":0,"disk_write_size_in_bytes":0,"disk_io_size_in_bytes":0},"plugins":[]}}`)
  })
  ts := httptest.NewServer(handler)
  defer ts.Close();

  rabbitmq := RabbitMQ{"localhost", 9200};
  data, err := rabbitmq.Monitor();
  assert.Equal(t, err, nil)

  t.Logf("Report: %v", data)

  jsonParsed, err := gabs.ParseJSON([]byte(data))
  assert.Equal(t, err, nil)

  status, _ := jsonParsed.S("report").S("status").S("state").Data().(string);
  assert.Equal(t, status, "ok")
}
*/
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
