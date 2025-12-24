def merge(left,right):
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result

def external_sort(arr, chunk_size = 3):
    #делим массив на куски
    chunks = [sorted(arr[i:i + chunk_size]) for i in range(0, len(arr), chunk_size)]

    #постепенно сливаем куски
    while len(chunks) > 1:
        new_chunks = []
        for i in range(0, len(chunks), 2):
            if i+1 < len(chunks):
                new_chunks.append(merge(chunks[i], chunks[i+1]))
            else:
                new_chunks.append(chunks[i])
        chunks = new_chunks
    return chunks[0]

user_input = input("Введите числа через пробел: ")
numbers = list(map(int, user_input.split()))

sorted_numbers = external_sort(numbers)
print("Отсортированный массив", sorted_numbers)
