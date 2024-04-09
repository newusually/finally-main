<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8"/>

    <!-- 引入刚刚下载的 ECharts 文件 -->

    <script src="../js/echarts.min.js"></script>
    <script>
        window.onload = function () {
            line_stack();
            setInterval(function () {

                location.reload();
            }, 10000); // 60000毫秒等于1分钟
        }
    </script>

</head>
<body>
<!-- 为 ECharts 准备一个定义了宽高的 DOM -->
<div id="main" style="width: 600px;height:400px;"></div>

<?php
require_once '../php_class/DataLineStack.php';

DataLineStack::init('E:\datas\uplRatio\log\cashBal_history.txt');
DataLineStack::readFiles();
$data1 = DataLineStack::getData();

# 获取时间数据
$date = DataLineStack::getDate();

#实际未结算盈亏总额
DataLineStack::init('E:\datas\uplRatio\log\frozenBal_history.txt');
DataLineStack::readFiles();
$data2 = DataLineStack::getData();

# 获取时间数据
$date = DataLineStack::getDate();
#USDT币种余额
$cashBal_history = $data1['cashBal_history'];

#保证金金额
$frozenBal_history = $data2['frozenBal_history'];
?>

<script type="text/javascript">


    function line_stack() {
        var chartDom = document.getElementById('main');
        var myChart = echarts.init(chartDom, 'dark', {
            renderer: 'canvas',
            useDirtyRect: false
        });
        var option;

        option = {
            title: {
                text: '实时亏损盈利'
            },
            tooltip: {
                trigger: 'axis'
            },
            legend: {
                data: ['USDT余额', '保证金金额']
            },
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            toolbox: {
                feature: {
                    saveAsImage: {}
                }
            },
            xAxis: {
                type: 'category',
                boundaryGap: false,
                data: <?php echo json_encode($date);?>
            },
            yAxis: {
                type: 'value'
            },

            series: [
                {
                    name: 'USDT币种余额',
                    type: 'line',
                    stack: 'Total',
                    data: <?php echo json_encode($cashBal_history);?>
                },

                {
                    name: '保证金金额',
                    type: 'line',
                    stack: 'Total',
                    data: <?php echo json_encode($frozenBal_history);?>
                }
            ]
        };

        option && myChart.setOption(option);
    }
</script>
</body>
</html>