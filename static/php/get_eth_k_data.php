<?php
require_once '../php_class/DataFetcher.php';
$filePath ='E:\datas\old_data\ETH-USDT-SWAP\ETH-USDT-SWAP-15min.csv';
DataFetcher::init($filePath);
$data = DataFetcher::getContractData();
echo json_encode($data);
?>