namespace LabsAsd;

public class Lab13
{
    // Дан текстовый файл с некоторым текстом на русском или английском языках
    // произвольной длины (организовать чтение).
    // Выбрав некоторую хеш-функцию, создать хеш-таблицу с наложением
    
    private string[] _table;
    private int _size;

    public Lab13(int size)
    {
        _size = size;
        _table = new string[size];
    }

    // Хеш-функция
    private int GetHash(string key)
    {
        // Используем Math.Abs, чтобы избежать отрицательных индексов
        return Math.Abs(key.GetHashCode()) % _size;
    }

    // Вставка с линейным наложением
    public void Insert(string key)
    {
        int index = GetHash(key);

        // Линейное пробирование в случае коллизии
        while (_table[index] != null)
        {
            if (_table[index] == key) return; // Слово уже есть
            index = (index + 1) % _size;   // Переходим к следующей ячейке
        }

        _table[index] = key;
    }

    public void PrintTable()
    {
        for (int i = 0; i < _size; i++)
        {
            if (_table[i] != null)
                Console.WriteLine($"[{i}]: {_table[i]}");
        }
    }
}
