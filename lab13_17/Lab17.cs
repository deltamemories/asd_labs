namespace LabsAsd;
using System.Collections.Generic;
using System.Text;

public class Lab17
{
    public class Node
    {
        public int Data;
        public Node Left;
        public Node Right;

        public Node(int data)
        {
            Data = data;
            Left = null;
            Right = null;
        }
    }

    public class BinarySearchTree
    {
        public Node Root;

        // 1. Поиск в БДП
        public Node Search(Node root, int key)
        {
            if (root == null || root.Data == key)
                return root;

            if (key < root.Data)
                return Search(root.Left, key);

            return Search(root.Right, key);
        }

        // 2. Добавление в БДП
        public Node Insert(Node root, int key)
        {
            if (root == null) return new Node(key);

            if (key < root.Data)
                root.Left = Insert(root.Left, key);
            else if (key > root.Data)
                root.Right = Insert(root.Right, key);

            return root;
        }

        // 3. Удаление из БДП
        public Node Delete(Node root, int key)
        {
            if (root == null) return root;

            if (key < root.Data)
                root.Left = Delete(root.Left, key);
            else if (key > root.Data)
                root.Right = Delete(root.Right, key);
            else
            {
                // Узел с одним потомком или без них
                if (root.Left == null) return root.Right;
                if (root.Right == null) return root.Left;

                // Узел с двумя потомками: ищем минимальный в правом поддереве
                root.Data = MinValue(root.Right);
                root.Right = Delete(root.Right, root.Data);
            }
            return root;
        }

        private int MinValue(Node root)
        {
            int minv = root.Data;
            while (root.Left != null)
            {
                minv = root.Left.Data;
                root = root.Left;
            }
            return minv;
        }

        // 4. Парсинг линейно-скобочной записи: 8(3(1,6),10)
        public Node Parse(string s, ref int pos)
        {
            if (pos >= s.Length || s[pos] == ')' || s[pos] == ',') return null;

            // Читаем число
            StringBuilder valStr = new StringBuilder();
            while (pos < s.Length && char.IsDigit(s[pos]))
            {
                valStr.Append(s[pos]);
                pos++;
            }

            Node node = new Node(int.Parse(valStr.ToString()));

            if (pos < s.Length && s[pos] == '(')
            {
                pos++; // пропускаем '('
                node.Left = Parse(s, ref pos); // Левый потомок

                if (pos < s.Length && s[pos] == ',')
                {
                    pos++; // пропускаем ','
                    node.Right = Parse(s, ref pos); // Правый потомок
                }

                if (pos < s.Length && s[pos] == ')')
                    pos++; // пропускаем ')'
            }
            return node;
        }

        // 5. Вывод в линейно-скобочном виде
        public string ToBracketString(Node node)
        {
            if (node == null) return "";
            string s = node.Data.ToString();
            if (node.Left != null || node.Right != null)
            {
                s += "(" + ToBracketString(node.Left) + "," + ToBracketString(node.Right) + ")";
            }
            return s;
        }
    }
}
