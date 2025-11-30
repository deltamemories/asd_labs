package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"sort"
	"strconv"
)

const (
	// chunkSize определяет, сколько чисел мы можем отсортировать в памяти за раз.
	// В реальном приложении это значение было бы значительно больше.
	chunkSize = 100
	// totalNumbers - общее количество чисел для генерации в исходном файле.
	totalNumbers = 1000
	// numTempFiles - количество временных файлов для слияния.
	// Классическая полифазная сортировка часто использует 3.
	numTempFiles = 3
)

// fileInfo хранит информацию о временном файле.
type fileInfo struct {
	name      string
	runCount  int
	dummyRuns int // Фиктивные прогоны для выравнивания
}

// main - точка входа, которая управляет всем процессом сортировки.
func main() {
	// --- Шаг 1: Создание исходного файла ---
	sourceFile := "source_data.txt"
	fmt.Printf("1. Создание исходного файла '%s' с %d случайными числами...\n", sourceFile, totalNumbers)
	if err := createSourceFile(sourceFile, totalNumbers, totalNumbers*10); err != nil {
		panic(fmt.Sprintf("не удалось создать исходный файл: %v", err))
	}
	fmt.Println("   Исходный файл успешно создан.")

	// --- Шаг 2: Создание начальных отсортированных прогонов ---
	fmt.Printf("\n2. Разделение '%s' на отсортированные части (прогоны)...\n", sourceFile)
	initialRuns, err := createInitialRuns(sourceFile, chunkSize)
	if err != nil {
		panic(fmt.Sprintf("не удалось создать начальные прогоны: %v", err))
	}
	fmt.Printf("   Создано %d начальных прогонов.\n", len(initialRuns))

	// --- Шаг 3: Полифазное слияние ---
	fmt.Println("\n3. Выполнение полифазной сортировки слиянием...")
	sortedFile, err := polyphaseMerge(initialRuns, numTempFiles)
	if err != nil {
		panic(fmt.Sprintf("ошибка во время полифазного слияния: %v", err))
	}
	fmt.Printf("   Сортировка завершена. Результат в файле: '%s'\n", sortedFile)

	// --- Шаг 4: Очистка ---
	fmt.Println("\n4. Очистка временных файлов...")
	os.Remove(sourceFile)
	for _, runFile := range initialRuns {
		os.Remove(runFile)
	}
	os.Remove(sortedFile) // Удалить также конечный отсортированный файл
	fmt.Println("   Очистка завершена.")
}

// createSourceFile создает файл с заданным количеством случайных чисел.
func createSourceFile(name string, count, maxVal int) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for i := 0; i < count; i++ {
		fmt.Fprintln(w, rand.IntN(maxVal))
	}
	return w.Flush()
}

// createInitialRuns читает исходный файл, разбивает его на части,
// сортирует каждую часть в памяти и записывает в отдельный временный файл.
func createInitialRuns(sourceFile string, size int) ([]string, error) {
	f, err := os.Open(sourceFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var runFiles []string
	scanner := bufio.NewScanner(f)
	runIndex := 0
	for {
		chunk := make([]int, 0, size)
		for i := 0; i < size && scanner.Scan(); i++ {
			num, _ := strconv.Atoi(scanner.Text())
			chunk = append(chunk, num)
		}

		if len(chunk) == 0 {
			break
		}

		sort.Ints(chunk)

		runFileName := fmt.Sprintf("initial_run_%d.tmp", runIndex)
		if err := writeChunkToFile(runFileName, chunk); err != nil {
			return nil, err
		}
		runFiles = append(runFiles, runFileName)
		runIndex++
	}
	return runFiles, scanner.Err()
}

// writeChunkToFile записывает срез чисел в файл.
func writeChunkToFile(name string, chunk []int) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, num := range chunk {
		fmt.Fprintln(w, num)
	}
	return w.Flush()
}

// polyphaseMerge реализует основной алгоритм полифазного слияния.
func polyphaseMerge(runFiles []string, k int) (string, error) {
	if len(runFiles) <= 1 {
		return runFiles[0], nil
	}

	// Инициализация k временных файлов.
	files := make([]fileInfo, k)
	for i := 0; i < k; i++ {
		files[i] = fileInfo{name: fmt.Sprintf("merge_tape_%d.tmp", i)}
	}
	defer func() {
		for _, f := range files {
			os.Remove(f.name)
		}
	}()

	// Распределение начальных прогонов по k-1 файлам
	// с использованием чисел Фибоначчи для идеального распределения.
	distributeInitialRuns(runFiles, files)

	// Основной цикл слияния
	for {
		// Находим файл для вывода (тот, у которого 0 прогонов)
		outputIndex := -1
		for i := range files {
			if files[i].runCount == 0 && files[i].dummyRuns == 0 {
				outputIndex = i
				break
			}
		}

		// Если нет выходного файла, значит что-то пошло не так
		if outputIndex == -1 {
			// This can happen in the final stage, let's find the file with runs.
			// Re-check logic for completion.
			totalRuns := 0
			lastFile := ""
			for _, f := range files {
				totalRuns += f.runCount
				if f.runCount > 0 {
					lastFile = f.name
				}
			}
			if totalRuns == 1 {
				finalName := "sorted_result.txt"
				os.Rename(lastFile, finalName)
				return finalName, nil
			}
			return "", fmt.Errorf("не удалось найти пустой файл для вывода, состояние: %+v", files)
		}


		// Определяем, сколько прогонов нужно слить.
		// Это минимальное количество прогонов среди всех входных файлов.
	runsToMerge := -1
	hasInput := false
		for i := range files {
			if i == outputIndex {
				continue
			}
			// Only consider files that actually have runs
			if files[i].runCount > 0 {
				hasInput = true
				if runsToMerge == -1 || files[i].runCount < runsToMerge {
					runsToMerge = files[i].runCount
				}
			}
		}

		// If there are no runs to merge, we are done.
		if !hasInput {
			// This means the last output is the final result.
			// Find the file with content.
			lastFile := ""
			for _, f := range files {
				info, err := os.Stat(f.name)
				if err == nil && info.Size() > 0 {
					lastFile = f.name
					break
				}
			}
			if lastFile == "" {
				return "", fmt.Errorf("логическая ошибка: слияние завершено, но не найден финальный файл")
			}
			finalName := "sorted_result.txt"
			os.Rename(lastFile, finalName)
			return finalName, nil
		}
		
		fmt.Printf("   Слияние %d прогонов в '%s'...\n", runsToMerge, files[outputIndex].name)
		
		var inputFiles []*os.File
		var inputFileNames []string
		for i, f := range files {
			if i != outputIndex && f.runCount > 0 {
				file, err := os.Open(f.name)
				if err != nil { return "", err } // Corrected: removed defer file.Close() here as it would close all files prematurely
				// defer file.Close() // This defer is incorrect here, it would close the file immediately after opening in the first iteration.
				inputFiles = append(inputFiles, file)
				inputFileNames = append(inputFileNames, f.name)
			}
		}
		// Corrected: Ensure all opened input files are closed after mergeKWay finishes or if an error occurs.
		defer func() {
			for _, file := range inputFiles {
				file.Close()
			}
		}()
		
		if err := mergeKWay(inputFiles, files[outputIndex].name, runsToMerge); err != nil {
			return "", err
		}
		
		// Обновляем информацию о файлах и очищаем ставшие пустыми входные файлы
		files[outputIndex].runCount += runsToMerge
		for i := range files {
			if i != outputIndex && files[i].runCount > 0 {
				files[i].runCount -= runsToMerge
				if files[i].runCount == 0 {
					// Очищаем файл, так как все его прогоны были слиты
					os.Truncate(files[i].name, 0)
				}
			}
		}
	}
}


// distributeInitialRuns распределяет начальные прогоны по временным файлам.
func distributeInitialRuns(runFiles []string, files []fileInfo) {
	for i, runFile := range runFiles {
		targetFileIndex := i % (len(files) - 1)
		
		src, _ := os.Open(runFile)
		dst, _ := os.OpenFile(files[targetFileIndex].name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		io.Copy(dst, src)
		src.Close()
		dst.Close()
		
		files[targetFileIndex].runCount++
	}
	fmt.Println("   Начальные прогоны распределены по временным файлам.")
}


// mergeKWay выполняет k-стороннее слияние заданного числа прогонов.
func mergeKWay(inputFiles []*os.File, outputFile string, numRuns int) error {
	out, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil { return err }
	defer out.Close()
	writer := bufio.NewWriter(out)

	scanners := make([]*bufio.Scanner, len(inputFiles))
	for i, f := range inputFiles {
		scanners[i] = bufio.NewScanner(f)
	}

	pq := &MinHeap{}
	heap.Init(pq)

	for i, s := range scanners {
		if s.Scan() {
			num, _ := strconv.Atoi(s.Text())
			heap.Push(pq, HeapItem{Value: num, FileIndex: i})
		}
	}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(HeapItem)
		fmt.Fprintln(writer, item.Value)

		scanner := scanners[item.FileIndex]
		if scanner.Scan() {
			num, _ := strconv.Atoi(scanner.Text())
			heap.Push(pq, HeapItem{Value: num, FileIndex: item.FileIndex})
		}
	}
	
	return writer.Flush()
}

// --- Код для Min-Heap ---
type HeapItem struct {
	Value    int
	FileIndex int
}
type MinHeap []HeapItem
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(HeapItem)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
