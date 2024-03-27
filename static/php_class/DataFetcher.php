<?php
class DataFetcher {
    #文件路径
    private $filePath;

    #构造函数 传入文件路径
    public function __construct($filePath) {
        $this->filePath = $filePath;
    }

    public function getFilePath() {
        return $this->filePath;
    }

    #获取盈亏列表
    public function getCashbal() {
        return file_get_contents($this->filePath);
    }


    #获取实时亏损盈利记录 循环表格打印
    public function getUplRatioData() {
        $fileContent = file_get_contents($this->filePath);
        // 将文件内容转换为UTF-8编码
        $fileContent = mb_convert_encoding($fileContent, 'UTF-8');
        $lines = explode("\n", $fileContent);
        $data = [];

        foreach ($lines as $line) {
            $line = trim($line);
            if (empty($line)) {
                continue; // 跳过空行
            }
            $columns = explode(',', $line); // 根据逗号分隔每一行的数据
            if (count($columns) != 7) {
                continue; // 如果一行的数据列数不是7，那么跳过这一行
            }
            $item = [];
            foreach ($columns as $index => $column) {
                $item["column" . ($index + 1)] = $column; // 使用 "column1"、"column2" 等作为键
            }
            $data[] = $item;
        }

        return $data;
    }

    /***
     * 这段代码首先读取文件的每一行到 $lines 数组中，然后遍历每一行。
     * 对于每一行，它首先移除行的开始和结束处的空白字符，然后检查行是否为空。
     * 如果行为空，那么就跳过这一行，继续处理下一行。
     * 然后，它使用 explode 函数根据逗号分隔行的数据到 $columns 数组中，
     * 然后检查 $columns 数组的长度是否为7。如果 $columns 数组的长度不是7，
     * 那么就跳过这一行，继续处理下一行。最后，它遍历 $columns 数组，
     * 将每一列的数据添加到 $item 数组中，
     * 然后将 $item 数组添加到 $rows 数组中。最后，它返回 $rows 数组
     */


    public function getBuyLogData($page, $pageSize) {
        $fileContent = file_get_contents($this->filePath);
        $lines = explode("\n", $fileContent); // 倒序读取
        $totalLines = count($lines); // 计算总行数
        $totalPages = ceil($totalLines / $pageSize); // 计算总页数

        $start = max(0, ($page - 1) * $pageSize); // 计算开始行
        $end = min($totalLines, $page * $pageSize); // 计算结束行


        $data = [];
        for ($i = $start; $i < $end; $i++) {
            if (isset($lines[$i])) {
                $line = trim($lines[$i]);
                if (empty($line)) {
                    continue; // 跳过空行
                }
                $columns = explode(',', $line); // 根据逗号分隔每一行的数据
                $item = [];
                foreach ($columns as $index => $column) {
                    $item["column" . ($index + 1)] = $column; // 使用 "column1"、"column2" 等作为键
                }
                $data[] = $item;
            }
        }
        return ['data' => $data, 'lines' => $lines, 'totalLines' => $totalLines, 'totalPages' => $totalPages, 'originalPage' => $page];
    }
    #获取合约15分钟数据

    /***
     * 在这段代码中，file() 函数读取文件的每一行到 $lines 数组中。
     * 然后，array_slice() 函数获取 $lines 数组的最后50个元素到
     * $last50Lines 数组中。最后，我们遍历 $last50Lines 数
     * 组，使用 str_getcsv() 函数将每一行的数据解析为 CSV 格式，
     * 然后将解析后的数据添加到 $data 数组中。最后，我们返回 $data 数组。
     */
    public function getContractData() {
        $lines = file($this->filePath, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
        $last50Lines = array_slice($lines, -50);
        $data = [];
        foreach ($last50Lines as $line) {
            $data[] = str_getcsv($line);
        }
        return $data;
    }
}
?>