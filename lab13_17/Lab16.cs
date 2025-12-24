namespace LabsAsd;
using System.Collections.Generic;
using System.Text;

public class Lab16
{
    // Не рекурсивный прямой обход
    public class Node
    {
        public int Data;
        public Node Left, Right;
        public Node(int item)
        {
            Data = item;
            Left = Right = null;
        }
    }

    // Не рекурсивный прямой обход (Lab 16)
    public string IterativePreOrder(Node root)
    {
        if (root == null) return "Дерево пусто";

        StringBuilder result = new StringBuilder();
        Stack<Node> stack = new Stack<Node>();
        
        // Шаг 1: Кладем корень в стек
        stack.Push(root);

        while (stack.Count > 0)
        {
            // Шаг 2: Достаем верхний элемент
            Node current = stack.Pop();
            result.Append(current.Data + " ");

            // Шаг 3: Кладем в стек ПРАВОЕ, затем ЛЕВОЕ поддерево
            // (чтобы левое было сверху и обработалось первым)
            if (current.Right != null)
            {
                stack.Push(current.Right);
            }
            if (current.Left != null)
            {
                stack.Push(current.Left);
            }
        }

        return result.ToString().Trim();
    }
}
