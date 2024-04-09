<?php
require_once '../php_class/DataFetcher.php';
$filePath ='e:/datas/uplRatio/log/uplRatio.txt';
DataFetcher::init($filePath);
$uplRatioData = DataFetcher::getUplRatioData();

// 检查$uplRatioData是否是数组或对象
if (!is_array($uplRatioData) && !is_object($uplRatioData)) {
    $uplRatioData = []; // 如果不是，设置为一个空数组
}

header('Content-Type: application/json;charset=utf-8');
echo json_encode($uplRatioData);

?>