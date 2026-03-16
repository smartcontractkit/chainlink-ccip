package shared

import (
	"context"
	"fmt"
	"strings"

	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
)

type NodeLookup struct {
	nodesByName map[string]*nodev1.Node
}

func NewNodeLookup(nodes []*nodev1.Node) (*NodeLookup, error) {
	lookup := &NodeLookup{
		nodesByName: make(map[string]*nodev1.Node),
	}
	for _, node := range nodes {
		key := strings.ToLower(node.Name)
		if existing, ok := lookup.nodesByName[key]; ok && existing.Id != node.Id {
			return nil, fmt.Errorf("duplicate node name %q: node IDs %s and %s both claim this name", node.Name, existing.Id, node.Id)
		}
		lookup.nodesByName[key] = node
	}
	return lookup, nil
}

func (l *NodeLookup) FindByName(name string) (*nodev1.Node, bool) {
	node, ok := l.nodesByName[strings.ToLower(name)]
	return node, ok
}

func NodeIDsToSet(nodeIDs []string) map[string]bool {
	if len(nodeIDs) == 0 {
		return nil
	}
	set := make(map[string]bool, len(nodeIDs))
	for _, id := range nodeIDs {
		set[id] = true
	}
	return set
}

func FetchNodeLookup(ctx context.Context, jdClient JDClient, nodeIDs []string) (*NodeLookup, error) {
	if len(nodeIDs) == 0 {
		return nil, fmt.Errorf("nodeIDs must be specified - refusing to fetch all nodes for security reasons")
	}

	filter := &nodev1.ListNodesRequest_Filter{
		Ids: nodeIDs,
	}

	nodesResp, err := jdClient.ListNodes(ctx, &nodev1.ListNodesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list nodes: %w", err)
	}

	lookup, err := NewNodeLookup(nodesResp.Nodes)
	if err != nil {
		return nil, fmt.Errorf("failed to build node lookup: %w", err)
	}
	return lookup, nil
}
