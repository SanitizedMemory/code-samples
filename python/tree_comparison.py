class Node:
    def __init__(self, id, quantity):
        self.id = id
        self.quantity = quantity
        self.left = None
        self.right = None

def compare_trees(tree1, tree2):
    # This function compares two trees and returns the differences in nodes.
    # We use a pre-order traversal for comparison.
    differences = []

    def traverse(node1, node2):
        if node1 and node2:
            if node1.id != node2.id or node1.quantity != node2.quantity:
                differences.append((node1, node2))

            traverse(node1.left, node2.left)
            traverse(node1.right, node2.right)
        elif node1 or node2:  # One of the nodes is None
            differences.append((node1, node2))

    traverse(tree1, tree2)
    return differences

def list_difference(list1, list2):
    # Assuming the lists are of the same length.
    # If not, we should handle that case as well.
    all_differences = []
    for tree1, tree2 in zip(list1, list2):
        differences = compare_trees(tree1, tree2)
        if differences:
            all_differences.append(differences)
    return all_differences

# List 1
tree1_list1 = Node(1, 10)
tree1_list1.left = Node(2, 20)
tree1_list1.right = Node(3, 30)

tree2_list1 = Node(4, 40)
tree2_list1.left = Node(5, 50)
tree2_list1.right = Node(6, 60)

# List 2
tree1_list2 = Node(1, 10)
tree1_list2.left = Node(2, 25)  # Different quantity
tree1_list2.right = Node(3, 30)

tree2_list2 = Node(4, 40)
tree2_list2.left = Node(5, 55)  # Different quantity
tree2_list2.right = Node(7, 60) # Different ID

# Creating the lists
list1 = [tree1_list1, tree2_list1]
list2 = [tree1_list2, tree2_list2]

# Testing the algorithm
differences = list_difference(list1, list2)
print("Differences between the tree lists:")
for diff in differences:
    for d in diff:
        print(f"Tree 1 Node: ID={d[0].id}, Quantity={d[0].quantity} | Tree 2 Node: ID={d[1].id if d[1] else 'None'}, Quantity={d[1].quantity if d[1] else 'None'}")

