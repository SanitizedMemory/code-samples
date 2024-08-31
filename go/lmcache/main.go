package main

import (
    "fmt"
    "math"
)

type ContinuationRadixTrieNode struct {
    content []byte
    children []ContinuationRadixTrieNode
    continuation *[]byte
}

type NodeTuple struct {
    idx int
    prefixLength uint
}

func hasPrefix(sequence []byte, prefix []byte) bool {
    if len(sequence) < len(prefix) {
        return false
    }

    if len(prefix) == 0 {
        return true
    }

    for idx, value := range prefix {
        if sequence[idx] != value {
            return false
        }
    }
    
    return true
}

func (node *ContinuationRadixTrieNode) GetContinuation(sequence []byte) *[]byte {
    // If the sequence is empty, we don't know how to continue.
    if len(sequence) == 0 {
        return nil
    }

    // Is the sequence invalid, or can we continue down the Trie?
    if !hasPrefix(sequence, node.content) {
        return nil
    }

    nextSequence := sequence[len(node.content):]

    // If there is more to the sequence, check if any children have continuations.
    if len(nextSequence) > 0 {
        for _, child := range node.children {
            if childContinuation := child.GetContinuation(nextSequence); childContinuation != nil {
                return node.continuation
            }
        }
    }

    // If no child reports a continuation, check if this node can contribute a continuation.
    continuation := *(node.continuation)
    if continuation != nil {
        // Is the continuation prefixed by the next sequence portion?
        if !hasPrefix(continuation, nextSequence) {
            return nil
        }

        remainingContinuation := continuation[len(nextSequence):]

        // Is there any content left to continue?
        if len(remainingContinuation) == 0 {
            return nil
        }

        // We got a hit!
        return &remainingContinuation
    }

    return nil
}

func sequencePrefixLength(sequenceA []byte, sequenceB []byte) uint {
    minLength := uint(math.Min(float64(len(sequenceA)), float64(len(sequenceB))))
    if minLength == 0 {
        return 0
    }
    
    var idx uint
    for idx = 0; idx < minLength; idx++ {
        if sequenceA[idx] != sequenceB[idx] {
            return idx
        }
    }

    return minLength
}

func (node *ContinuationRadixTrieNode) splitInsert(prefixLength uint, sequence []byte, continuation []byte) {
    prefixContent := node.content[:prefixLength]
    oldSuffixContent := node.content[prefixLength:]

    oldChildren := node.children
    oldContinuation := node.continuation
    node.children = []ContinuationRadixTrieNode{
        ContinuationRadixTrieNode{
            content: oldSuffixContent,
            children: oldChildren,
            continuation: oldContinuation,
        },
    }
    node.continuation = nil

    var childOld ContinuationRadixTrieNode

    node.content = prefixContent
    node.children = append(node.children, childOld)

    if len(sequence) > 0 {
        newSuffixContent := sequence
        childNew := ContinuationRadixTrieNode{
            content: newSuffixContent,
            continuation: &continuation,
        }
        node.children = append(node.children, childNew)
    } else {
        node.continuation = &continuation
    }
}

func (node *ContinuationRadixTrieNode) InsertContinuationIterative(sequence []byte, continuation []byte) {
    currentNode := node
    for {
        var nextNodeData *NodeTuple

        for childIdx := 0; childIdx < len(currentNode.children); childIdx++ {
            child := currentNode.children[childIdx]
            if len(child.content) == 0 {
                panic("Found empty node as non-root in Trie.")
            }

            prefixLength := sequencePrefixLength(sequence, child.content) 
            if prefixLength > 0 {
                nextNodeData = &NodeTuple{childIdx, prefixLength}
                break
            }
        }

        if nextNodeData == nil {
            currentNode.children = append(currentNode.children, ContinuationRadixTrieNode{
                content: sequence,
                continuation: &continuation,
            })
            break
        }

        nextNodeIdx := nextNodeData.idx
        nextNodePrefixLength := nextNodeData.prefixLength
        nextNode := &currentNode.children[nextNodeIdx]
        sequence = sequence[nextNodePrefixLength:]

        if int(nextNodePrefixLength) < len(nextNode.content) {
            nextNode.splitInsert(nextNodePrefixLength, sequence, continuation)
            break
        }

        if len(sequence) == 0 {
            nextNode.continuation = &continuation
            break
        }

        currentNode = nextNode
    }
}

func (node *ContinuationRadixTrieNode) ToGraphviz() string {
    graphviz := "digraph RadixTrie {\n    node [shape=record];\n    graph [rankdir=LR splines=ortho];\n"
    node.graphvizHelper(&graphviz, 0)
    graphviz = graphviz + "}\n"
    return graphviz
}

func (node *ContinuationRadixTrieNode) graphvizHelper(graphviz *string, nodeId uint) uint {
    contentStr := string(node.content)
    continuationStr := ""
    if node.continuation != nil {
        continuationStr = string(*node.continuation)
    }
    *graphviz = *graphviz + fmt.Sprintf("    node%d [label=\"'%s'|%s\"];\n", nodeId, contentStr, continuationStr)
    nextNodeId := nodeId + 1
    for _, child := range node.children {
        childNodeId := nextNodeId
        *graphviz = *graphviz + fmt.Sprintf("    node%d -> node%d;\n", nodeId, childNodeId)
        nextNodeId = child.graphvizHelper(graphviz, childNodeId)
    }
    return nextNodeId
}

func main() {
    node := ContinuationRadixTrieNode{}
    node.InsertContinuationIterative([]byte("abc"), []byte("de"))
    node.InsertContinuationIterative([]byte("abcd"), []byte("efg"))
    node.InsertContinuationIterative([]byte("ab"), []byte("xyz"))

    graph := node.ToGraphviz()
    fmt.Println(graph)
}
