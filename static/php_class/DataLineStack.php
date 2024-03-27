<?php

class DataLineStack
{
    private string $dirPath;
    private array $date = [];
    private array $cashBal_history = [];
    private array $disEq_history = [];
    private array $frozenBal_history = [];
    private array $mgnRatio_history = [];
    private array $posdatacount_history = [];
    private array $upl_history = [];

    public function __construct($dirPath)
    {
        $this->dirPath = $dirPath;
    }

    public function getDate(): array
    {
        return $this->date;
    }

    public function readFiles(): void
    {
        // 获取文件名
        $fileName = basename($this->dirPath);

        // 创建一个空数组
        $files = [];

        // 根据文件名，设置对应的属性引用
        switch ($fileName) {
            case 'cashBal_history.txt':
                $files[$fileName] = &$this->cashBal_history;
                break;
            case 'disEq_history.txt':
                $files[$fileName] = &$this->disEq_history;
                break;
            case 'frozenBal_history.txt':
                $files[$fileName] = &$this->frozenBal_history;
                break;
            case 'mgnRatio_history.txt':
                $files[$fileName] = &$this->mgnRatio_history;
                break;
            case 'posdatacount_history.txt':
                $files[$fileName] = &$this->posdatacount_history;
                break;
            case 'upl_history.txt':
                $files[$fileName] = &$this->upl_history;
                break;
        }

        // 检查文件名是否在$files数组中
        if (isset($files[$fileName])) {
            $fileArray = &$files[$fileName];
            $filePath = $this->dirPath;
            $fileContent = file_get_contents($filePath);
            $lines = explode("\n", $fileContent);

            foreach ($lines as $line) {
                $columns = explode(',', $line);
                if (count($columns) >= 2) {
                    $date = trim($columns[0]); // 获取逗号前面的时间列数据
                    $this->date[] = $date; // 将时间数据保存到date数组中
                    $value = trim($columns[1]); // 使用trim函数去除\r字符
                    $fileArray[] = $value;
                }
            }
        }
    }

    public function getData(): array
    {
        // 获取文件名
        $fileName = basename($this->dirPath);

        switch ($fileName) {
            case 'cashBal_history.txt':
                return ['cashBal_history' => $this->cashBal_history];
            case 'disEq_history.txt':
                return ['disEq_history' => $this->disEq_history];
            case 'frozenBal_history.txt':
                return ['frozenBal_history' => $this->frozenBal_history];
            case 'mgnRatio_history.txt':
                return ['mgnRatio_history' => $this->mgnRatio_history];
            case 'posdatacount_history.txt':
                return ['posdatacount_history' => $this->posdatacount_history];
            case 'upl_history.txt':
                return ['upl_history' => $this->upl_history];
            default:
                return [];
        }
    }
}