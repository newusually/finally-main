package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

// FuncMap 用于在模板中添加数学运算的辅助函数
var funcMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
}

// parseUplRatioFile 解析实时收益率文件
func parseUplRatioFile(filename string) ([]UplRatioRecord, error) {
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

func parseFile(filename string) (*Data, error) {
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

type SafeMap struct {
	sync.RWMutex
	m map[string]bool
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
func loadData(startPage, rowsPerPage, maxPages int) (*Data, int, error) {
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
		fileData, err := parseFile(file)
		if err != nil {
			log.Println("Error parsing file:", err)
			continue
		}
		if len(allData.Columns) == 0 {
			allData.Columns = fileData.Columns // Take the column headers from the first file
		}
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

var columnColors = []string{
	"#ffdddd", "#ddffdd", "#ddddff",
	"#ffddff", "#ddffff", "#ffffdd",
	"#d3d3d3", "#ffa07a", "#20b2aa",
}

func readFileContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		page := 1
		if p := r.URL.Query().Get("page"); p != "" {
			if num, err := strconv.Atoi(p); err == nil && num > 0 {
				page = num
			}
		}

		// 加载买入合约条件数据
		maxPages := page + 10
		contractData, totalPages, err := loadData(page, 30, maxPages)
		if err != nil {
			http.Error(w, "Error loading contract data", http.StatusInternalServerError)
			return
		}

		// 解析实时收益率数据
		uplRatioRecords, err := parseUplRatioFile("../datas/uplRatio/log/uplRatio.txt")
		if err != nil {
			http.Error(w, "Error reading uplRatio file", http.StatusInternalServerError)
			return
		}

		// 计算总计亏损金额和保证金的总和
		var totalOnlyImr, totalImr float64
		for _, record := range uplRatioRecords {
			onlyImr, err := strconv.ParseFloat(record.OnlyImr, 64)
			if err != nil {
				http.Error(w, "Error parsing OnlyImr", http.StatusInternalServerError)
				return
			}
			totalOnlyImr += onlyImr

			imr, err := strconv.ParseFloat(record.Imr, 64)
			if err != nil {
				http.Error(w, "Error parsing Imr", http.StatusInternalServerError)
				return
			}
			totalImr += imr
		}

		// 格式化为保留两位小数的字符串
		totalOnlyImrStr := fmt.Sprintf("%.2f", totalOnlyImr)
		totalImrStr := fmt.Sprintf("%.2f", totalImr)

		// 将字符串转回为浮点数
		totalOnlyImr, _ = strconv.ParseFloat(totalOnlyImrStr, 64)
		totalImr, _ = strconv.ParseFloat(totalImrStr, 64)

		// 读取现金余额文件
		// Read the content of the cashbal.txt file
		cashbalContent, err := readFileContent("../datas/uplRatio/log/cashbal.txt")
		if err != nil {
			http.Error(w, "Error reading cashbal file", http.StatusInternalServerError)
			return
		}

		// 创建模板并添加FuncMap
		t, err := template.New("template.html").Funcs(funcMap).ParseFiles("template.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			// Handle the error in a way that makes sense for your application.
			// For example, you might want to send an HTTP error response.
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		var pages []int
		startPage := page
		if startPage < 1 {
			startPage = 1
		}
		endPage := startPage + 6
		if endPage > totalPages {
			endPage = totalPages
		}
		for i := startPage; i < endPage; i++ {
			pages = append(pages, i)
		}
		data := struct {
			UplRatioData []UplRatioRecord
			Data         *Data
			Colors       []string
			Page         int
			Pages        []int
			TotalPages   int
			TotalOnlyImr float64
			TotalImr     float64
			Cashbal      string
		}{
			UplRatioData: uplRatioRecords,
			Data:         contractData,
			Colors:       columnColors,
			Page:         page,
			Pages:        pages,
			TotalPages:   totalPages,
			TotalOnlyImr: totalOnlyImr,
			TotalImr:     totalImr,
			Cashbal:      cashbalContent,
		}
		//fmt.Println(t.DefinedTemplates()) // 打印模板内容
		//fmt.Printf("%+v\n", data)         // 打印传入的数据

		err = t.Execute(w, data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}

	})

	log.Fatal(http.ListenAndServe(":80", mux))
}
