
package main

import (
  "github.com/jeffail/gabs"
  "fmt"
  "net/http"
  "errors"
  "io/ioutil"
)

type RabbitMQ struct {
  host string;
  port uint16;
  user string;
  pass string;
};

func (this RabbitMQ) Monitor(url string) (string, error) {
  data, err := this.QueueStats(url);
  if err != nil {
    return "", err
  }

  atollReport, err := this.statsToAtollReport(data);
  return atollReport, nil
}

func (this RabbitMQ) statsToAtollReport(data string) (string, error) {
  statsParsed, err := gabs.ParseJSON([]byte(data))
  if err != nil {
    return "", err
  }
  atollReport := gabs.New();
  atollReport.SetP("rabbitmq", "id");
  atollReport.SetP("RabbitMQ", "name");

  atollReport.ArrayP("report.items");

  var children,_ = statsParsed.S("result").Children();
  for i,child := range children {
      atollReport.ArrayAppendP(map[string]interface{}{}, "report.items");
      atollItem := atollReport.S("report").S("items").Index(i);
      atollItem.SetP(child.S("name").Data().(string), "name");
      atollItem.SetP([1]string{"queue"}, "classes");

      // Queue stats
      atollItem.SetP([2]string{"memory", "bytes"}, "stats.memoryUsage.classes");
      atollItem.SetP(child.S("memory").Data().(float64), "stats.memoryUsage.value");
      atollItem.SetP([2]string{"messages", "count"}, "stats.numMessages.classes");
      atollItem.SetP(child.S("messages").Data().(float64), "stats.numMessages.value");

      publishRate := child.S("message_stats").S("publish_details").S("rate");
      if publishRate.String() != "{}" {
        atollItem.SetP([3]string{"throughput", "per-second", "average"}, "stats.publishRate.classes");
        atollItem.SetP(publishRate.Data().(float64), "stats.publishRate.value");
      }
      deliverRate := child.S("message_stats").S("deliver_get_details").S("rate");
      if deliverRate.String() != "{}" {
        atollItem.SetP([3]string{"throughput", "per-second", "average"}, "stats.deliverRate.classes");
        atollItem.SetP(deliverRate.Data().(float64), "stats.deliverRate.value");
      }
  }

  return atollReport.String(), nil;
}

func (this RabbitMQ) QueueStats(url string) (string, error) {
  if len(url) == 0 {
    url = fmt.Sprintf("http://%s:%d/api/queues", this.host, this.port);
  }
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Accept", "application/json")
  req.SetBasicAuth(this.user, this.pass);
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  if resp.StatusCode != 200 {
    return "", errors.New("Invalid response from RabbitMQ server: " + resp.Status)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf(`{"result":%s}`, string(body)), nil
}
