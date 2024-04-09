<?php
require_once '../php_class/DataFetcher.php';
$filePath ='e:/datas/uplRatio/log/cashbal.txt';
DataFetcher::init($filePath);
$cashbal = DataFetcher::getCashbal();

header('Content-Type: application/octet-stream;charset=utf-8');
echo $cashbal;
?>