namespace LabsAsd;

public class Lab14
{
    // Дан текстовый файл с некоторым текстом на русском или английском языках
    // произвольной длины (организовать чтение).
    // Выбрав некоторую хеш-функцию, создать хеш-таблицу со списками
    
    // Массив списков (цепочек)
    private LinkedList<string>[] _buckets;
    private int _size;

    public Lab14(int size)
    {
        _size = size;
        _buckets = new LinkedList<string>[size];
            
        // Инициализируем каждый бакет (корзину)
        for (int i = 0; i < size; i++)
        {
            _buckets[i] = new LinkedList<string>();
        }
    }

    private int GetHash(string key)
    {
        return Math.Abs(key.GetHashCode()) % _size;
    }

    public void Insert(string key)
    {
        int index = GetHash(key);
            
        // Проверяем, нет ли уже такого слова в списке, чтобы избежать дубликатов
        if (!_buckets[index].Contains(key))
        {
            _buckets[index].AddLast(key);
        }
    }

    public void PrintTable()
    {
        for (int i = 0; i < _size; i++)
        {
            if (_buckets[i].Count > 0)
            {
                Console.Write($"[{i, 3}]: ");
                Console.WriteLine(string.Join(" -> ", _buckets[i]));
            }
        }
    }
}
