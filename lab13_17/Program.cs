using System;

namespace LabsAsd
{
    class Program
    {
        private static void Main(string[] args)
        {
            // Инициализация переменной для цикла в консоли
            var isWork = true;
            
            // Инициализация классов с сортировками
            var lab5 = new Lab5();
            var lab6 = new Lab6();
            var lab7 = new Lab7();
            var lab8 = new Lab8();
            var lab9 = new Lab9();
            var lab10 = new Lab10();
            var lab11 = new Lab11();
            var lab12 = new Lab12();
            var lab15 = new Lab15();
            var lab16 = new Lab16();
            
            // Инициализация списка для сортировки
            var numbers = new[] { 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6 };
            
            
            static void Verify(string path)
            {
                var numbers = File.ReadAllLines(path).Select(int.Parse).ToList();
                bool isSorted = true;
                for (int i = 0; i < numbers.Count - 1; i++)
                {
                    if (numbers[i] > numbers[i + 1]) isSorted = false;
                }
                Console.WriteLine(isSorted ? "ПРОВЕРКА: Успешно отсортировано!" : "ПРОВЕРКА: Ошибка сортировки!");
            }
            
            

            for (; isWork;)
            {
                Console.Clear();
                Console.WriteLine("Демонстрация лабораторных работ");
                Console.WriteLine(new string('-', Console.WindowWidth));
                Console.WriteLine("Выберите тип сортировки:");
                Console.WriteLine(" 5. Сортировка вставками          (Лаба 5)  [Сдано]");
                Console.WriteLine(" 6. Сортировка посредством выбора (Лаба 6)  [Сдано]");
                Console.WriteLine(" 7. Сортировка Шелла              (Лаба 7)  [Сдано]");
                Console.WriteLine(" 8. Поразрядная сортировка        (Лаба 8)  [Сдано]");
                Console.WriteLine(" 9. Пирамидальная сортировка      (Лаба 9)  [Сдано]");
                Console.WriteLine(" 10. Сортировка слиянием          (Лаба 10)");
                Console.WriteLine(" 11. Быстрая сортировка           (Лаба 11)");
                Console.WriteLine(" 12. Внешняя многофазовая         (Лаба 12)");
                Console.WriteLine(" 13. Хеш-таблица с наложением     (Лаба 13)");
                Console.WriteLine(" 14. Хеш-таблица со списками      (Лаба 14)");
                Console.WriteLine(" 15. Рекурсивные обходы           (Лаба 15)");
                Console.WriteLine(" 16. Не рекурсивный прямой обход  (Лаба 16)");
                Console.WriteLine(" 17. Операции над БНД             (Лаба 17)");
                Console.WriteLine("\n 0. Выход");
                Console.WriteLine(new string('-', Console.WindowWidth));
                
                var sortType = Console.ReadLine();

                switch (sortType)
                {
                    case "5":
                        Console.Clear();
                        Console.WriteLine("Демонстрация лабораторных работ -> Лабораторная работа 5 (Сортировка вставками)");
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.WriteLine("Исходная последовательность: 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6");
                        Console.WriteLine("Отсортированная: " + string.Join(", ", lab5.InsertionSort(numbers)));
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.ReadKey();
                        break;
                    
                    case "6":
                        Console.Clear();
                        Console.WriteLine("Демонстрация лабораторных работ -> Лабораторная работа 6 (Сортировка посредством выбора)");
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.WriteLine("Исходная последовательность: 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6");
                        Console.WriteLine("Отсортированная: " + string.Join(", ", lab6.SelectionSort(numbers)));
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.ReadKey();
                        break;
                    
                    case "7":
                        Console.Clear();
                        Console.WriteLine("Демонстрация лабораторных работ -> Лабораторная работа 7 (Сортировка Шелла)");
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.WriteLine("Исходная последовательность: 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6");
                        Console.WriteLine("Отсортированная: " + string.Join(", ", lab7.ShellSort(numbers)));
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.ReadKey();
                        break;
                    
                    case "8":
                        Console.Clear();
                        Console.WriteLine("Демонстрация лабораторных работ -> Лабораторная работа 8 (Поразрядная сортировка)");
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.WriteLine("Исходная последовательность: 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6");
                        Console.WriteLine("Отсортированная: " + string.Join(", ", lab8.RadixSort(numbers)));
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.ReadKey();
                        break;
                    
                    case "9":
                        Console.Clear();
                        Console.WriteLine("Демонстрация лабораторных работ -> Лабораторная работа 9 (Пирамидальная сортировка)");
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.WriteLine("Исходная последовательность: 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6");
                        Console.WriteLine("Отсортированная: " + string.Join(", ", lab9.HeapSort(numbers)));
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.ReadKey();
                        break;
                    
                    case "10":
                        Console.Clear();
                        Console.WriteLine("Демонстрация лабораторных работ -> Лабораторная работа 10 (Сортировка слиянием)");
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.WriteLine("Исходная последовательность: 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6");
                        lab10.MergeSort(numbers, 0, numbers.Length - 1);
                        Console.WriteLine("Отсортированная: " + string.Join(", ", numbers));
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        numbers = new[] { 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6 };
                        Console.ReadKey();
                        break;
                    
                    case "11":
                        Console.Clear();
                        Console.WriteLine("Демонстрация лабораторных работ -> Лабораторная работа 11 (Быстрая сортировка)");
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        Console.WriteLine("Исходная последовательность: 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6");
                        lab11.QuickSort(numbers, 0, numbers.Length - 1);
                        Console.WriteLine("Отсортированная: " + string.Join(", ", numbers));
                        Console.WriteLine(new string('-', Console.WindowWidth));
                        numbers = new[] { 1, 3, 3, 7, 5, 4, 4, 2, 8, 9, 6 };
                        Console.ReadKey();
                        break;
                    
                    case "12":
                        Console.Clear();
                        string input = "input.txt";
                        string output = "output.txt";

                        // 1. Создаем тестовый файл, если его нет (на основе ваших данных)
                        int[] startData = { 562, 662, 228, 826, 910, 22, 908, 519, 966, 610, 302, 697, 122, 731, 622, 836, 945, 547, 487, 439, 311, 684, 160, 883, 971, 816, 59, 944, 872, 534, 675, 661, 358, 234, 145, 752, 59, 784, 316, 632, 104, 273, 806, 528, 876, 849, 938, 805, 134, 895 };
                        File.WriteAllLines(input, startData.Select(x => x.ToString()));

                        Console.WriteLine("Начало внешней сортировки...");
                        
                        lab12.Sort(input, output);

                        Console.WriteLine($"Готово! Результат в файле: {output}");
                        Verify(output);
                        Console.ReadKey();
                        break;
                    
                    case "13":
                        Console.Clear();
                        
                        string filePath = "text.txt";
                        
                        if (!File.Exists(filePath))
                            File.WriteAllText(filePath, "Hello world! Привет мир! Это пример текста для хеш таблицы.");

                        try
                        {
                            // Читаем весь текст и разбиваем на слова
                            string content = File.ReadAllText(filePath);
                            string[] words = content.Split(new[] { ' ', '.', ',', '!', '?', '\r', '\n' }, StringSplitOptions.RemoveEmptyEntries);

                            // Создаем таблицу (размер чуть больше количества слов для уменьшения коллизий)
                            var hashTable = new Lab13(words.Length * 2);

                            foreach (string word in words)
                            {
                                hashTable.Insert(word.ToLower());
                            }

                            Console.WriteLine("Содержимое хеш-таблицы (индекс: значение):");
                            hashTable.PrintTable();
                        }
                        catch (Exception ex)
                        {
                            Console.WriteLine($"Ошибка: {ex.Message}");
                        }
                        Console.ReadKey();
                        break;
                    
                    case "14":
                        Console.Clear();
                        
                        string filePath2 = "text.txt";
            
                        if (!File.Exists(filePath2))
                            File.WriteAllText(filePath2, "Apple banana apple orange grape banana apple");

                        try
                        {
                            string content = File.ReadAllText(filePath2);
                            string[] words = content.Split(new[] { ' ', '.', ',', '!', '?' }, StringSplitOptions.RemoveEmptyEntries);

                            // Для списков размер таблицы может быть меньше количества слов
                            Lab14 hashTable = new Lab14(10); 

                            foreach (string word in words)
                            {
                                hashTable.Insert(word.ToLower());
                            }

                            Console.WriteLine("Хеш-таблица со списками (индекс: цепочка значений):");
                            hashTable.PrintTable();
                        }
                        catch (Exception ex)
                        {
                            Console.WriteLine($"Ошибка: {ex.Message}");
                        }
                        
                        Console.ReadKey();
                        break;
                    
                    case "15":
                        Console.Clear();
                        // Ручное создание дерева по схеме из задания
                        Lab15.Node root = new Lab15.Node(1);
                        root.Left = new Lab15.Node(2);
                        root.Right = new Lab15.Node(3);
        
                        root.Left.Right = new Lab15.Node(4);
                        root.Left.Right.Left = new Lab15.Node(7);
                        root.Left.Right.Right = new Lab15.Node(8);

                        root.Right.Left = new Lab15.Node(5);
                        root.Right.Right = new Lab15.Node(6);
                        root.Right.Left.Right = new Lab15.Node(9);
                        root.Right.Left.Right.Left = new Lab15.Node(11);
                        root.Right.Right.Left = new Lab15.Node(10);

                        Console.WriteLine("Лабораторная работа №15");
        
                        Console.Write("\nПрямой обход (Pre-order): ");
                        lab15.PreOrder(root); 
                        // Должно быть: 1 2 4 7 8 3 5 9 11 6 10 (как на картинке)

                        Console.Write("\nЦентральный обход (In-order): ");
                        lab15.InOrder(root);

                        Console.Write("\nКонцевой обход (Post-order): ");
                        lab15.PostOrder(root);
        
                        Console.ReadKey();
                        break;
                    
                    case "16":
                        Console.Clear();
                        
                        // Создаем то же дерево, что и в лабе 15
                        Lab16.Node root2 = new Lab16.Node(1);
                        root2.Left = new Lab16.Node(2);
                        root2.Right = new Lab16.Node(3);
                        root2.Left.Right = new Lab16.Node(4);
                        root2.Left.Right.Left = new Lab16.Node(7);
                        root2.Left.Right.Right = new Lab16.Node(8);
                        root2.Right.Left = new Lab16.Node(5);
                        root2.Right.Right = new Lab16.Node(6);
                        root2.Right.Left.Right = new Lab16.Node(9);
                        root2.Right.Left.Right.Left = new Lab16.Node(11);
                        root2.Right.Right.Left = new Lab16.Node(10);

                        Console.WriteLine("Лабораторная работа №16");
                        Console.WriteLine("Не рекурсивный прямой обход через Stack:");
        
                        string traversalResult = lab16.IterativePreOrder(root2);
        
                        Console.WriteLine($"Результат: {traversalResult}");
        
                        // Проверка на соответствие результату из методички
                        string expected = "1 2 4 7 8 3 5 9 11 6 10";
                        if (traversalResult == expected)
                            Console.WriteLine("\nУспех: Обход совпадает с ожидаемым!");
                        else
                            Console.WriteLine("\nОшибка: Есть расхождения с образцом.");
                        
                        Console.ReadKey();
                        break;
                    
                    case "17":
                        Console.Clear();
                        var bst = new Lab17.BinarySearchTree();
            
                        Console.WriteLine("Введите дерево в формате 8(3(1,6),10) или просто корень:");
                        string input2 = Console.ReadLine().Replace(" ", "");
                        int pos = 0;
                        bst.Root = bst.Parse(input2, ref pos);

                        bool running = true;
                        while (running)
                        {
                            Console.WriteLine("\n--- Меню ---");
                            Console.WriteLine("1. Поиск узла");
                            Console.WriteLine("2. Добавление узла");
                            Console.WriteLine("3. Удаление узла");
                            Console.WriteLine("4. Выход");
                            Console.Write("Выбор: ");

                            switch (Console.ReadLine())
                            {
                                case "1":
                                    Console.Write("Значение для поиска: ");
                                    int sKey = int.Parse(Console.ReadLine());
                                    Console.WriteLine(bst.Search(bst.Root, sKey) != null ? "Найдено" : "Не найдено");
                                    break;
                                case "2":
                                    Console.Write("Значение для добавления: ");
                                    int aKey = int.Parse(Console.ReadLine());
                                    bst.Root = bst.Insert(bst.Root, aKey);
                                    break;
                                case "3":
                                    Console.Write("Значение для удаления: ");
                                    int dKey = int.Parse(Console.ReadLine());
                                    bst.Root = bst.Delete(bst.Root, dKey);
                                    break;
                                case "4":
                                    running = false;
                                    break;
                            }
                        }

                        Console.WriteLine("\nИтоговое дерево (линейно-скобочная запись):");
                        Console.WriteLine(bst.ToBracketString(bst.Root));
                        Console.WriteLine("Нажмите любую клавишу для выхода...");
                        Console.ReadKey();
                        break;
                    
                    case "0":
                        isWork =  false;
                        break;
                }
            }
        }
    }   
}
