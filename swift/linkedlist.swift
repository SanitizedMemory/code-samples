public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init() { self.val = 0; self.next = nil; }
    public init(_ val: Int) { self.val = val; self.next = nil; }
    public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
}

extension ListNode {
    static func create(_ arr: [Int]) -> ListNode? {
        var head: ListNode? = nil
        var node: ListNode? = nil
        for i in arr {
            let newNode = ListNode(i, nil)
            node?.next = newNode
            node = newNode
            if head == nil {
                head = node
            }
        }
        return head
    }
    
    func printValues() {
        print(val)
        next?.printValues()
    }
}

func mergeTwoLists(_ list1: ListNode?, _ list2: ListNode?) -> ListNode? {
    if list1 == nil && list2 == nil {
        return nil
    }
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    var list1Head = list1
    var list2Head = list2
    let list3: ListNode? = ListNode(-101, nil)
    var list3Prev: ListNode? = nil
    var list3Head = list3!
    while list1Head != nil || list2Head != nil {
        let l1Value = list1Head?.val ?? 101
        let l2Value = list2Head?.val ?? 101
        print((l1Value, l2Value))
        if l2Value > l1Value {
            list1Head = list1Head!.next
            list3Head.val = l1Value
        } else {
            list2Head = list2Head!.next
            list3Head.val = l2Value
        }
        let nextVal = ListNode(-101, nil)
        list3Head.next = nextVal
        list3Prev = list3Head
        list3Head = nextVal
    }
    list3Prev!.next = nil
    return list3
}

let list1 = ListNode.create([1, 2, 4])
let list2 = ListNode.create([1, 3, 4])

let list3 = mergeTwoLists(list1, list2)
list3?.printValues()
