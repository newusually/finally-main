package models

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// 定义实时收益率记录的数据结构
type UplRatioRecord struct {
	Symbol      string
	UplRatio    string
	NotionalUsd string
	AvgPx       string
	Imr         string
	Pos         string
	OnlyImr     string
}

// 定义买入合约条件的数据结构
type Data struct {
	Columns []string
	Rows    [][]string
}

type Row struct {
	Time    time.Time
	Columns []string
}

type SafeMap struct {
	sync.RWMutex
	m map[string]bool
}

// parseUplRatioFile 解析实时收益率文件
func ParseUplRatioFile(filename string) ([]UplRatioRecord, error) {
	// 使用 defer 语句来捕获 panic
	defer func() {
		if r := recover(); r != nil {
			// 如果发生了 panic，将 panic 的值转换为 error 并返回
			fmt.Errorf("panic: %v", r)
		}
	}()
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records []UplRatioRecord
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 7 {
			continue // Skip lines that don't have exactly 7 fields
		}
		record := UplRatioRecord{
			Symbol:      strings.TrimPrefix(fields[0], "symbol--->>>"),
			UplRatio:    strings.TrimPrefix(fields[1], "未实现收益率--->>>"),
			NotionalUsd: strings.TrimPrefix(fields[2], "现价--->>>"),
			AvgPx:       strings.TrimPrefix(fields[3], "开仓均价--->>>"),
			Imr:         strings.TrimPrefix(fields[4], "保证金--->>>"),
			Pos:         strings.TrimPrefix(fields[5], "杠杆倍数--->>>"),
			OnlyImr:     strings.TrimPrefix(fields[6], "总计亏损金额--->>>"),
		}
		records = append(records, record)
	}
	return records, scanner.Err()
}
func ParseFile(filename string) (*Data, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("panic: %v", r)
		}
	}()
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := &Data{}
	columnsSet := &SafeMap{
		m: make(map[string]bool),
	}

	// Use a buffered reader to reduce disk I/O operations
	reader := bufio.NewReader(file)

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		wg.Add(1)
		go func(line string) {
			defer wg.Done()

			var row []string
			processLine(data, line, columnsSet, &row)
		}(line)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return data, nil
}

func (sm *SafeMap) Load(key string) (bool, bool) {
	sm.RLock()
	defer sm.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

func (sm *SafeMap) Store(key string, value bool) {
	sm.Lock()
	defer sm.Unlock()
	sm.m[key] = value
}

func processLine(data *Data, line string, columnsSet *SafeMap, row *[]string) {
	// 使用 defer 语句来捕获 panic
	defer func() {
		if r := recover(); r != nil {
			// 如果发生了 panic，将 panic 的值转换为 error 并返回
			fmt.Errorf("panic: %v", r)
		}
	}()
	line = strings.ReplaceAll(line, "\n", "")
	parts := strings.Split(line, ",")
	for _, part := range parts {
		keyValue := strings.Split(part, "--->>")
		if len(keyValue) == 2 {
			key := strings.TrimSpace(strings.Trim(keyValue[0], ","))
			value := strings.TrimSpace(keyValue[1])
			if _, ok := columnsSet.Load(key); !ok && len(data.Columns) < 11 {
				data.Columns = append(data.Columns, key)
				columnsSet.Store(key, true)
			}
			*row = append(*row, fmt.Sprintf("%s--->>%s", key, value))
		}
	}
	for len(*row) < 11 {
		*row = append(*row, "")
	}
	data.Rows = append(data.Rows, *row)
}

// loadData 从所有 .txt 文件中加载买入合约条件数据
func LoadData(startPage, rowsPerPage, maxPages int) (*Data, int, error) {
	// 使用 defer 语句来捕获 panic
	defer func() {
		if r := recover(); r != nil {
			// 如果发生了 panic，将 panic 的值转换为 error 并返回
			fmt.Errorf("panic: %v", r)
		}
	}()

	files, err := filepath.Glob("../datas/log/*.txt")
	if err != nil {
		return nil, 0, err
	}

	allData := &Data{}
	var allRows [][]string
	totalRows := 0
	totalPages := 0
	for _, file := range files {
		if totalPages < startPage-1 {
			totalPages++
			continue
		}
		fileData, err := ParseFile(file)
		if err != nil {
			log.Println("Error parsing file:", err)
			continue
		}
		if len(allData.Columns) == 0 {
			allData.Columns = fileData.Columns // Take the column headers from the first file
		}

		// Reverse the rows before appending them to allRows
		sort.SliceStable(fileData.Rows, func(i, j int) bool {
			return i > j
		})
		allRows = append(allRows, fileData.Rows...)
		totalRows += len(fileData.Rows)
		totalPages++
		if totalPages >= startPage+maxPages-1 {
			break
		}
	}

	startRow := (startPage - 1) * rowsPerPage
	endRow := startRow + rowsPerPage
	if endRow > totalRows {
		endRow = totalRows
	}

	allData.Rows = allRows[startRow:endRow]

	totalPages = totalRows / rowsPerPage
	if totalRows%rowsPerPage != 0 {
		totalPages++
	}

	return allData, totalPages, nil

}

func ReadFileContent(filename string) (string, error) {
	// 使用 defer 语句来捕获 panic
	defer func() {
		if r := recover(); r != nil {
			// 如果发生了 panic，将 panic 的值转换为 error 并返回
			fmt.Errorf("panic: %v", r)
		}
	}()
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
