<?php
require_once '../php_class/DataFetcher.php';
$pageSize = 30; // 每页显示的行数
$filePath = 'e:/datas/log/buylog.txt';
$dataFetcher3 = new DataFetcher($filePath);

$result = $dataFetcher3->getBuyLogData(1, $pageSize); // 先获取总页数
$totalPages = $result['totalPages']; // 获取总页数

$originalPage = isset($_GET['page']) && is_numeric($_GET['page']) && $_GET['page'] > 0 ? $_GET['page'] : 1; // 获取当前页码，如果没有或者不是一个大于0的整数则默认为第一页
$page = $totalPages - $originalPage + 1; // 进行倒序处理

$result = $dataFetcher3->getBuyLogData($page, $pageSize);
$buy_log_data = $result['data'];
$lines = $result['lines'];
$totalLines = $result['totalLines']; // 获取总行数

$response = array(
    'page' => $originalPage, // 返回的页码应为原始的 $_GET['page'] 值
    'pageSize' => $pageSize,
    'buy_log_data' => $buy_log_data,
    'lines' => $lines,
    'totalLines' => $totalLines, // 将总行数添加到响应中
    'totalPages' => $totalPages // 将总页数添加到响应中
);

header('Content-Type: application/json;charset=utf-8');
echo json_encode($response);


?>