<?php
require_once '../php_class/DataFetcher.php';

$dataFetcher1 = new DataFetcher('e:/datas/uplRatio/log/cashbal.txt');
$cashbal = $dataFetcher1->getCashbal();

header('Content-Type: application/octet-stream;charset=utf-8');
echo $cashbal;
?>