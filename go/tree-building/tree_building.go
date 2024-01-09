package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func NewNode(id int) *Node {
	return &Node{ID: id}
}

func (n *Node) AddChild(child *Node) {
	n.Children=append(n.Children, child)
}

func (n *Node) GetNode(id int) *Node {
	if n.ID == id {
		return n
	}

	for i:=range n.Children {
		child:=n.Children[i]
		if child.ID == id {
			return child
		} else {
			if found:=child.GetNode(id); found != nil {
				return found
			}
		}
	}

	return nil
}

func Build(records []Record) (*Node, error) {
	var result *Node
	var err error

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if len(records) == 1 && records[0].ID != 0 {
		return nil, errors.New("duplicated record")
	}

	lastId:=0
	for i:=range records {
		record:=records[i]
		if record.ID > 0 && result != nil {
			if node:=result.GetNode(record.Parent); node != nil {
				if result.GetNode(record.ID) != nil {
					err=errors.New("duplicated record")
					break
				}

				if (record.ID - lastId) > 1 {
					err=errors.New("not continuous")
					break
				} else {
					lastId=record.ID
				}

				node.AddChild(NewNode(record.ID))
			} else {
				if record.Parent == 0 && result == nil {
					err=errors.New("parent not found")
					break
				}

				if record.ID == record.Parent {
					err=errors.New("circular reference")
					break
				}

				result.AddChild(NewNode(record.ID))
			}
		} else if record.ID == 0 {
			if record.Parent > 0 {
				err=errors.New("root can't have parent")
				break
			}

			if result != nil {
				err=errors.New("root has been created already")
			}
			
			result=NewNode(0)
		}
	}

	return result, err
}
