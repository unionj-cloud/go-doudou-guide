package promsd

import (
	"context"
	"fmt"
	"github.com/unionj-cloud/go-doudou/svc/registry"
	"github.com/unionj-cloud/memberlist"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/log/level"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/prometheus/prometheus/util/strutil"
)

var (
	// addressLabel is the name for the label containing a target's address.
	addressLabel = model.MetaLabelPrefix + "go-doudou_address"
	// nodeLabel is the name for the label containing a target's node name.
	nodeLabel = model.MetaLabelPrefix + "go-doudou_node"
	// tagsLabel is the name of the label containing the tags assigned to the target.
	tagsLabel = model.MetaLabelPrefix + "go-doudou_tags"
	// serviceAddressLabel is the name of the label containing the (optional) service address.
	serviceAddressLabel = model.MetaLabelPrefix + "go-doudou_service_address"
	// servicePortLabel is the name of the label containing the service port.
	servicePortLabel = model.MetaLabelPrefix + "go-doudou_service_port"
	// serviceIDLabel is the name of the label containing the service ID.
	serviceIDLabel = model.MetaLabelPrefix + "go-doudou_service_id"
)

// Note: This is the struct with your implementation of the Discoverer interface (see Run function).
// Discovery retrieves target information from a Consul server and updates them via watches.
type discovery struct {
	refreshInterval time.Duration
}

func (d *discovery) parseServiceNodes(nodes []*memberlist.Node) ([]*targetgroup.Group, error) {
	sourceMap := make(map[string][]*registry.NodeInfo)
	for _, node := range nodes {
		nodeInfo := registry.Info(node)
		sourceMap[nodeInfo.SvcName] = append(sourceMap[nodeInfo.SvcName], &nodeInfo)
	}

	for key, value := range sourceMap {
		// TODO
	}

	tgroup := targetgroup.Group{
		Source: registry.Info(node).SvcName,
		Labels: make(model.LabelSet),
	}

	tgroup.Targets = make([]model.LabelSet, 0, len(nodes))

	for _, node := range nodes {
		// We surround the separated list with the separator as well. This way regular expressions
		// in relabeling rules don't have to consider tag positions.
		tags := "," + strings.Join(node.ServiceTags, ",") + ","

		// If the service address is not empty it should be used instead of the node address
		// since the service may be registered remotely through a different node.
		var addr string
		if node.ServiceAddress != "" {
			addr = net.JoinHostPort(node.ServiceAddress, fmt.Sprintf("%d", node.ServicePort))
		} else {
			addr = net.JoinHostPort(node.Address, fmt.Sprintf("%d", node.ServicePort))
		}

		target := model.LabelSet{model.AddressLabel: model.LabelValue(addr)}
		labels := model.LabelSet{
			model.AddressLabel:                   model.LabelValue(addr),
			model.LabelName(addressLabel):        model.LabelValue(node.Address),
			model.LabelName(nodeLabel):           model.LabelValue(node.Node),
			model.LabelName(tagsLabel):           model.LabelValue(tags),
			model.LabelName(serviceAddressLabel): model.LabelValue(node.ServiceAddress),
			model.LabelName(servicePortLabel):    model.LabelValue(strconv.Itoa(node.ServicePort)),
			model.LabelName(serviceIDLabel):      model.LabelValue(node.ServiceID),
		}
		tgroup.Labels = labels

		// Add all key/value pairs from the node's metadata as their own labels.
		for k, v := range node.NodeMeta {
			name := strutil.SanitizeLabelName(k)
			tgroup.Labels[model.LabelName(model.MetaLabelPrefix+name)] = model.LabelValue(v)
		}

		tgroup.Targets = append(tgroup.Targets, target)
	}
	return &tgroup, nil
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
			level.Error(d.logger).Log("msg", "Error parsing services nodes", "service", name, "err", err)
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

func NewDiscovery(interval time.Duration) (*discovery, error) {
	cd := &discovery{
		refreshInterval: interval,
	}
	return cd, nil
}
