package promsd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/prometheus/prometheus/util/strutil"
	"github.com/unionj-cloud/cast"
	"github.com/unionj-cloud/go-doudou/framework/memberlist"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"io/ioutil"
	"net"
	"os"
	"time"

	"github.com/go-kit/log/level"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"
)

var (
	serviceIDLabel = "service_id"
	hostnameLabel  = "hostname"
	baseUrlLabel   = "base_url"
	buildTimeLabel = "build_time"
	buildUserLabel = "build_user"
	goVerLabel     = "go_ver"
	godoudouLabel  = "godoudou_ver"
)

// Note: This is the struct with your implementation of the Discoverer interface (see Run function).
// Discovery retrieves target information from a Consul server and updates them via watches.
type discovery struct {
	refreshInterval time.Duration
	logger          log.Logger
	sdFile          string
}

type customSD struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

func (d *discovery) parseServiceNodes(nodes []*memberlist.Node) ([]*targetgroup.Group, error) {
	var groups []*targetgroup.Group
	for _, item := range nodes {
		node := registry.Info(item)
		tgroup := &targetgroup.Group{
			Source: node.Hostname,
			Labels: make(model.LabelSet),
		}
		addr := net.JoinHostPort(node.Host, fmt.Sprintf("%d", node.SvcPort))
		target := model.LabelSet{
			model.AddressLabel: model.LabelValue(addr),
		}
		tgroup.Targets = []model.LabelSet{target}

		labels := model.LabelSet{
			model.LabelName(serviceIDLabel): model.LabelValue(node.SvcName),
			model.LabelName(hostnameLabel):  model.LabelValue(node.Hostname),
			model.LabelName(baseUrlLabel):   model.LabelValue(node.BaseUrl),
			model.LabelName(buildTimeLabel): model.LabelValue(node.BuildTime),
			model.LabelName(buildUserLabel): model.LabelValue(node.BuildUser),
			model.LabelName(goVerLabel):     model.LabelValue(node.GoVer),
			model.LabelName(godoudouLabel):  model.LabelValue(node.GddVer),
		}
		tgroup.Labels = labels
		// Add all key/value pairs from the node's metadata as their own labels.
		for k, v := range node.Data {
			name := strutil.SanitizeLabelName(k)
			tgroup.Labels[model.LabelName(model.MetaLabelPrefix+name)] = model.LabelValue(cast.ToString(v))
		}
		groups = append(groups, tgroup)
	}

	groupMap := make(map[string]*targetgroup.Group)
	for _, item := range groups {
		groupMap[item.Source] = item
	}

	f, err := os.Open(d.sdFile)
	if err != nil {
		return nil, err
	}
	old, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var sd []customSD
	_ = json.Unmarshal(old, &sd)

	for _, item := range sd {
		source := item.Labels[hostnameLabel]
		if _, exists := groupMap[source]; !exists {
			groups = append(groups, &targetgroup.Group{
				Source: source,
			})
		}
	}
	return groups, nil
}

// Run Note: you must implement this function for your discovery implementation as part of the
// Discoverer interface. Here you should query your SD for it's list of known targets, determine
// which of those targets you care about (for example, which of Consuls known services do you want
// to scrape for metrics), and then send those targets as a target.TargetGroup to the ch channel.
func (d *discovery) Run(ctx context.Context, ch chan<- []*targetgroup.Group) {
	for c := time.Tick(d.refreshInterval); ; {
		nodes, _ := registry.AllNodes()
		tgs, err := d.parseServiceNodes(nodes)
		if err != nil {
			level.Error(d.logger).Log("msg", "Error parsing services nodes", "err", err)
			break
		}
		// We're returning all Consul services as a single targetgroup.
		ch <- tgs
		// Wait for ticker or exit when ctx is closed.
		select {
		case <-c:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func NewDiscovery(interval time.Duration, logger log.Logger, sdFile string) (*discovery, error) {
	cd := &discovery{
		refreshInterval: interval,
		logger:          logger,
		sdFile:          sdFile,
	}
	return cd, nil
}
