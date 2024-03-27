<?php
require_once '../php_class/DataFetcher.php';
$dataFetcher = new DataFetcher('E:\datas\old_data\ETH-USDT-SWAP\ETH-USDT-SWAP-15min.csv');
$data = $dataFetcher->getContractData();
echo json_encode($data);
?>