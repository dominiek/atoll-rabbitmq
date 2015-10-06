
package main

import (
  //"github.com/jeffail/gabs"
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

func (this RabbitMQ) Monitor() (string, error) {
  data, err := this.QueueStats("");
  if err != nil {
    return "", err
  }
  return data, nil;

  //atollReport, err := this.statsToAtollReport(data);
  //return atollReport, nil
}
/*
func (this RabbitMQ) statsToAtollReport(data string) (string, error) {
  statsParsed, err := gabs.ParseJSON([]byte(data))
  if err != nil {
    return "", err
  }
  atollReport := gabs.New();

  // Main state
  state := "error";
  esStatus := statsParsed.S("status").Data().(string);
  if esStatus == "green" {
    state = "ok";
  } else if esStatus == "yellow" {
    state = "warn";
  }
  atollReport.SetP(state, "report.status.state");

  // Num Nodes
  atollReport.SetP(statsParsed.S("nodes").S("count").S("total").Data().(float64), "report.stats.numberOfNodes.value");

  // Core caches
  atollReport.SetP(statsParsed.S("indices").S("fielddata").S("memory_size_in_bytes").Data().(float64)/1024/1024, "report.stats.fieldDataCacheSize.value");
  atollReport.SetP([2]string{"megabytes", "memory"}, "report.stats.fieldDataCacheSize.classes");
  atollReport.SetP(statsParsed.S("indices").S("filter_cache").S("memory_size_in_bytes").Data().(float64)/1024/1024, "report.stats.filterCacheSize.value");
  atollReport.SetP([2]string{"megabytes", "memory"}, "report.stats.filterCacheSize.classes");
  atollReport.SetP(statsParsed.S("indices").S("id_cache").S("memory_size_in_bytes").Data().(float64)/1024/1024, "report.stats.idCacheSize.value");
  atollReport.SetP([2]string{"megabytes", "memory"}, "report.stats.idCacheSize.classes");

  // JVM stats
  atollReport.SetP(statsParsed.S("nodes").S("jvm").S("mem").S("heap_used_in_bytes").Data().(float64)/1024/1024, "report.stats.jvmHeapUsage.value");
  atollReport.SetP([2]string{"megabytes", "memory"}, "report.stats.jvmHeapUsage.classes");
  atollReport.SetP(statsParsed.S("nodes").S("jvm").S("mem").S("heap_max_in_bytes").Data().(float64)/1024/1024, "report.stats.jvmHeapSize.value");
  atollReport.SetP([2]string{"megabytes", "memory"}, "report.stats.jvmHeapSize.classes");

  return atollReport.String(), nil;
}
*/
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
