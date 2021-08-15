package communication

import (
	"encoding/json"
	"fmt"
	"steel-lang/misc"
	"sync"

	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type delegateDefault struct {
	localResources     misc.StringSet
	registry           resourceRegistry
	lockRegistry       sync.RWMutex
	waitingForRegistry chan string
	pendingUpdates     chan resourceRegistry
	haltUpdates        chan struct{}
}

func defaultDelegate(resources misc.StringSet) MemberlistDelegate {
	return &delegateDefault{
		localResources: resources,
	}
}

func (d *delegateDefault) Start(b BaseMembers) {
	d.registry = makeResourceRegistry(d.localResources, b.Config.Name, registrySize)
	d.waitingForRegistry = make(chan string)
	d.pendingUpdates = make(chan resourceRegistry)
	d.haltUpdates = make(chan struct{})
	go d.handleUpdates(b.List)
}

func (d *delegateDefault) Stop(b BaseMembers) {
	d.haltUpdates <- struct{}{}
	d.registry = nil
	d.waitingForRegistry = nil
	d.pendingUpdates = nil
	d.haltUpdates = nil
}

func (d *delegateDefault) FilterPartecipants(b BaseMembers, nodes []*memberlist.Node) []*memberlist.Node {
	res := make([]*memberlist.Node, 0, len(nodes))
	myName := b.List.LocalNode().Name
	for _, n := range nodes {
		if n.Name != myName {
			res = append(res, n)
		}
	}
	return res
}

func (d *delegateDefault) NodeMeta(b BaseMembers, limit int) []byte {
	return []byte{}
}

func (d *delegateDefault) NotifyMsg(b BaseMembers, m []byte) {
	var message messageUnion
	err := json.Unmarshal(m, &message)
	if err != nil {
		fmt.Println("error in message unmarshalling:", err.Error())
		return
	}
	switch message.Type {
	case "registry_request":
		select {
		case d.waitingForRegistry <- message.Sender.Name:
			fmt.Println("received registry request from", message.Sender.Name)
		default:
			fmt.Println("discarded incoming registry request from", message.Sender.Name)
		}
	case "registry_response":
		select {
		case d.pendingUpdates <- message.Registry:
			fmt.Println("received registry response from", message.Sender.Name)
		default:
			fmt.Println("discarded incoming registry response from", message.Sender.Name)
		}
	default:
		b.Logger.Error("unsupported message: "+message.Type,
			zap.String("act", "recv"),
			zap.String("obj", `"`+message.Type+`"`),
			zap.String("from", message.Sender.Name))
	}
}

func (d *delegateDefault) GetBroadcasts(b BaseMembers, overhead, limit int) [][]byte {
	return [][]byte{}
}

func (d *delegateDefault) LocalState(b BaseMembers, join bool) []byte {
	if join {
		d.lockRegistry.RLock()
		localRegistry, err := json.Marshal(d.registry)
		d.lockRegistry.RUnlock()
		if err != nil {
			fmt.Println("error in registry marshalling:", err.Error())
			return []byte{}
		}
		fmt.Println("join: sending registry")
		return localRegistry
	}
	d.lockRegistry.RLock()
	inventory := registryInventory{
		Sender:    b.List.LocalNode(),
		Inventory: d.registry.inventory(),
	}
	d.lockRegistry.RUnlock()
	inventoryMsg, err := json.Marshal(inventory)
	if err != nil {
		fmt.Println("error during inventory marshalling:", err.Error())
		return []byte{}
	}
	fmt.Println("sending inventory")
	return inventoryMsg
}

func (d *delegateDefault) MergeRemoteState(b BaseMembers, buf []byte, join bool) {
	if join {
		var remoteRegistry resourceRegistry
		err := json.Unmarshal(buf, &remoteRegistry)
		if err != nil {
			fmt.Println("join: error in registry unmarshalling:", err.Error())
			return
		}

		d.lockRegistry.RLock()
		size := len(d.registry)
		d.lockRegistry.RUnlock()
		if size == registrySize {
			fmt.Println("join: discarded received registry as already reached maximum registry size")
			return
		}
		select {
		case d.pendingUpdates <- remoteRegistry:
			fmt.Println("join: received registry")
		default:
			fmt.Println("join: discarded received registry")
		}
		return
	}
	var remoteInventory registryInventory
	err := json.Unmarshal(buf, &remoteInventory)
	if err != nil {
		fmt.Println("error in registry unmarshalling:", err.Error())
		return
	}
	d.lockRegistry.RLock()
	size := len(d.registry)
	d.lockRegistry.RUnlock()
	if size == registrySize {
		return
	}
	d.lockRegistry.RLock()
	inventory := d.registry.inventory()
	d.lockRegistry.RUnlock()
	if !inventory.ContainsSet(remoteInventory.Inventory) {
		message := messageUnion{
			Type:   "registry_request",
			Sender: b.List.LocalNode(),
		}
		marshalled, err := json.Marshal(message)
		if err != nil {
			fmt.Println("error in registry request marshalling:", err.Error())
			return
		}
		err = b.List.SendBestEffort(remoteInventory.Sender, marshalled)
		if err != nil {
			fmt.Println("error in sending registry request to", remoteInventory.Sender.Name, err.Error())
		}
		fmt.Println("sent registry request to", remoteInventory.Sender.Name)
	}
}

func (d *delegateDefault) NotifyJoin(b BaseMembers, node *memberlist.Node) {
	// do nothing
}

func (d *delegateDefault) NotifyLeave(b BaseMembers, node *memberlist.Node) {
	// do nothing
}

func (d *delegateDefault) NotifyUpdate(b BaseMembers, node *memberlist.Node) {
	// do nothing
}

func (d *delegateDefault) handleUpdates(list *memberlist.Memberlist) {
	for {
		select {
		case <-d.haltUpdates:
			return
		case remoteRegistry := <-d.pendingUpdates:
			for nodeName, resources := range remoteRegistry {
				d.lockRegistry.RLock()
				entry, present := d.registry[nodeName]
				d.lockRegistry.RUnlock()
				if !present || (entry != nil && resources == nil) {
					d.lockRegistry.Lock()
					if len(d.registry) == registrySize {
						d.lockRegistry.Unlock()
						break
					}
					d.registry[nodeName] = resources
					d.lockRegistry.Unlock()
				}
			}
		case destName := <-d.waitingForRegistry:
			for _, node := range list.Members() {
				if node.Name == destName {
					message := messageUnion{
						Type:     "registry_response",
						Sender:   list.LocalNode(),
						Registry: d.registry,
					}
					d.lockRegistry.RLock()
					localRegistry, err := json.Marshal(message)
					d.lockRegistry.RUnlock()
					if err != nil {
						fmt.Println("error in message marshalling:", err.Error())
						return
					}
					err = list.SendReliable(node, localRegistry)
					if err != nil {
						fmt.Println("error in sending registry response to", destName, err.Error())
						return
					}
					fmt.Println("sent registry response to", node.Name)
					break
				}
			}
		}
	}
}
