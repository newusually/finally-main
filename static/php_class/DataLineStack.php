<?php

class DataLineStack
{
    private static string $dirPath;
    private static array $date = [];
    private static array $cashBal_history = [];
    private static array $disEq_history = [];
    private static array $frozenBal_history = [];
    private static array $mgnRatio_history = [];
    private static array $posdatacount_history = [];
    private static array $upl_history = [];

    public static function init($dirPath)
    {
        self::$dirPath = $dirPath;
    }

    public static function getDate(): array
    {
        return self::$date;
    }

    public static function readFiles(): void
    {
        // 获取文件名
        $fileName = basename(self::$dirPath);

        // 创建一个空数组
        $files = [];

        // 根据文件名，设置对应的属性引用
        switch ($fileName) {
            case 'cashBal_history.txt':
                $files[$fileName] = &self::$cashBal_history;
                break;
            case 'disEq_history.txt':
                $files[$fileName] = &self::$disEq_history;
                break;
            case 'frozenBal_history.txt':
                $files[$fileName] = &self::$frozenBal_history;
                break;
            case 'mgnRatio_history.txt':
                $files[$fileName] = &self::$mgnRatio_history;
                break;
            case 'posdatacount_history.txt':
                $files[$fileName] = &self::$posdatacount_history;
                break;
            case 'upl_history.txt':
                $files[$fileName] = &self::$upl_history;
                break;
        }

        // 检查文件名是否在$files数组中
        if (isset($files[$fileName])) {
            $fileArray = &$files[$fileName];
            $filePath = self::$dirPath;
            $fileContent = file_get_contents($filePath);
            $lines = explode("\n", $fileContent);

            foreach ($lines as $line) {
                $columns = explode(',', $line);
                if (count($columns) >= 2) {
                    $date = trim($columns[0]); // 获取逗号前面的时间列数据
                    self::$date[] = $date; // 将时间数据保存到date数组中
                    $value = trim($columns[1]); // 使用trim函数去除\r字符
                    $fileArray[] = $value;
                }
            }
        }
    }

    public static function getData(): array
    {
        // 获取文件名
        $fileName = basename(self::$dirPath);

        switch ($fileName) {
            case 'cashBal_history.txt':
                return ['cashBal_history' => self::$cashBal_history];
            case 'disEq_history.txt':
                return ['disEq_history' => self::$disEq_history];
            case 'frozenBal_history.txt':
                return ['frozenBal_history' => self::$frozenBal_history];
            case 'mgnRatio_history.txt':
                return ['mgnRatio_history' => self::$mgnRatio_history];
            case 'posdatacount_history.txt':
                return ['posdatacount_history' => self::$posdatacount_history];
            case 'upl_history.txt':
                return ['upl_history' => self::$upl_history];
            default:
                return [];
        }
    }
}
?>