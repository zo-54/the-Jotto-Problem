package main

const cmpZero uint32 = 0

type node struct {
	binaryWord uint32
	level      int
	parent     *node
	subnodes   []*node
}

func (n *node) checkBinaryWord(bw uint32, soln chan<- []uint32) {
	if n.binaryWord&bw == cmpZero {
		if n.level == 4 {
			s := make([]uint32, 1, 5)

			s[0] = bw

			for currNode := n; currNode.level > 0; currNode = currNode.parent {
				s = append(s, currNode.binaryWord)
			}

			soln <- s

			return
		}

		newSubNode := &node{
			binaryWord: bw,
			level:      n.level + 1,
			parent:     n,
		}

		if n.subnodes == nil {
			n.subnodes = []*node{newSubNode}
		} else {
			for _, sn := range n.subnodes {
				sn.checkBinaryWord(bw, soln)
			}

			n.subnodes = append(n.subnodes, newSubNode)
		}
	}
}

func treeWorker(addWord <-chan uint32, solutions chan<- []uint32) {
	t := newTimer("condidate search")

	defer func() {
		close(solutions)
		t.end()
		return
	}()

	parentNode := &node{
		binaryWord: 0,
		level:      0,
		subnodes:   make([]*node, 0),
	}

	for {
		w, chanOpen := <-addWord

		if !chanOpen {
			return
		}

		parentNode.checkBinaryWord(w, solutions)
	}
}
