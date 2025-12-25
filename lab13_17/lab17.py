#17 lab
class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

def parse_tree(s):
    def helper():
        nonlocal i
        if i >= len(s) or not s[i].isdigit():
            return None
        num = ''
        while i < len(s) and s[i].isdigit():
            num += s[i]
            i += 1
        node = Node(int(num))
        if i < len(s) and s[i] == '(':
            i += 1
            node.left = helper()
            i += 1  # skip ')'
            if i < len(s) and s[i] == '(':
                i += 1
                node.right = helper()
                i += 1  # skip ')'
        return node
    i = 0
    return helper()

def tree_to_string(node):
    if not node:
        return ''
    left = tree_to_string(node.left)
    right = tree_to_string(node.right)
    if node.left or node.right:
        return f"{node.val}({left})({right})"
    return f"{node.val}"

def search(node, val):
    if not node:
        return False
    if val == node.val:
        return True
    elif val < node.val:
        return search(node.left, val)
    else:
        return search(node.right, val)

def insert(node, val):
    if not node:
        return Node(val)
    if val < node.val:
        node.left = insert(node.left, val)
    elif val > node.val:
        node.right = insert(node.right, val)
    return node

def delete(node, val):
    if not node:
        return None
    if val < node.val:
        node.left = delete(node.left, val)
    elif val > node.val:
        node.right = delete(node.right, val)
    else:
        if not node.left:
            return node.right
        if not node.right:
            return node.left
        # замена на минимальный в правом поддереве
        min_larger = node.right
        while min_larger.left:
            min_larger = min_larger.left
        node.val = min_larger.val
        node.right = delete(node.right, min_larger.val)
    return node

# Ввод дерева
tree_str = input("Введите дерево в линейно-скобочной форме: ").strip()
root = parse_tree(tree_str)

# Меню
while True:
    print("\nМеню: 1-Поиск  2-Добавить  3-Удалить  0-Выход")
    choice = input("Выбор: ").strip()
    if choice == '1':
        val = int(input("Значение для поиска: "))
        print("Найдено" if search(root, val) else "Не найдено")
    elif choice == '2':
        val = int(input("Значение для добавления: "))
        root = insert(root, val)
    elif choice == '3':
        val = int(input("Значение для удаления: "))
        root = delete(root, val)
    elif choice == '0':
        break
    else:
        print("Неверный выбор")

# Вывод дерева
print("Итоговое дерево:", tree_to_string(root))