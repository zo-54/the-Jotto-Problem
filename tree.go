package main

const cmpZero uint32 = 0

type node struct {
	binaryWord uint32
	level      int
	parent     *node
	subnodes   []*node
}

func (n *node) addSubnode(bw uint32) {
	newSubnode := &node{
		binaryWord: bw,
		level:      n.level + 1,
		parent:     n,
	}

	if n.level <= 4 {
		newSubnode.subnodes = make([]*node, 0)
	}

	n.subnodes = append(n.subnodes, newSubnode)
}

func (n *node) checkBinaryWord(bw uint32, soln chan<- []uint32) {
	for _, sn := range n.subnodes {
		if bw&sn.binaryWord == cmpZero {
			if sn.level < 4 {
				sn.checkBinaryWord(bw, soln)
				sn.addSubnode(bw)
			} else {
				s := make([]uint32, 1, 5)

				s[0] = bw

				for currNode := sn; currNode.level > 0; currNode = currNode.parent {
					s = append(s, currNode.binaryWord)
				}

				soln <- s
			}
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

		parentNode.addSubnode(w)
	}
}
