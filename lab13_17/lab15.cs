namespace LabsAsd;

public class Lab15
{
    // Рекурсивные обходы (прямой, центральный, кольцевой)
    // Класс узла дерева
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

    // 1. Прямой обход (Pre-order)
    public void PreOrder(Node node)
    {
        if (node == null) return;
        Console.Write(node.Data + " "); // Посещаем корень
        PreOrder(node.Left);           // Рекурсия влево
        PreOrder(node.Right);          // Рекурсия вправо
    }

    // 2. Центральный обход (In-order)
    public void InOrder(Node node)
    {
        if (node == null) return;
        InOrder(node.Left);            // Рекурсия влево
        Console.Write(node.Data + " "); // Посещаем корень
        InOrder(node.Right);           // Рекурсия вправо
    }

    // 3. Концевой обход (Post-order)
    public void PostOrder(Node node)
    {
        if (node == null) return;
        PostOrder(node.Left);          // Рекурсия влево
        PostOrder(node.Right);         // Рекурсия вправо
        Console.Write(node.Data + " "); // Посещаем корень
    }
}
