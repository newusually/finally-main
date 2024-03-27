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
            }, 60000); // 60000毫秒等于1分钟
        }
    </script>

</head>
<body>
<!-- 为 ECharts 准备一个定义了宽高的 DOM -->
<div id="main" style="width: 600px;height:400px;"></div>

<?php
require_once '../php_class/DataLineStack.php';

$dataLineStack1 = new DataLineStack('E:\datas\uplRatio\log\cashBal_history.txt');
$dataLineStack1->readFiles();
$data1 = $dataLineStack1->getData();

# 获取时间数据
$date = $dataLineStack1->getDate();


#美金层面币种折算权益
$dataLineStack2 = new DataLineStack('E:\datas\uplRatio\log\disEq_history.txt');
$dataLineStack2->readFiles();
$data2 = $dataLineStack2->getData();

#USDT保证金金额
$dataLineStack3 = new DataLineStack('E:\datas\uplRatio\log\frozenBal_history.txt');
$dataLineStack3->readFiles();
$data3 = $dataLineStack3->getData();

#USDT保证金率
$dataLineStack4 = new DataLineStack('E:\datas\uplRatio\log\mgnRatio_history.txt');
$dataLineStack4->readFiles();
$data4 = $dataLineStack4->getData();

#合约订单数量
$dataLineStack5 = new DataLineStack('E:\datas\uplRatio\log\posdatacount_history.txt');
$dataLineStack5->readFiles();
$data5 = $dataLineStack5->getData();

#实际未结算盈亏总额
$dataLineStack6 = new DataLineStack('E:\datas\uplRatio\log\upl_history.txt');
$dataLineStack6->readFiles();
$data6 = $dataLineStack6->getData();

# 获取时间数据
$date = $dataLineStack1->getDate();
#USDT币种余额
$cashBal_history = $data1['cashBal_history'];
#USDT保证金金额
$disEq_history = $data2['disEq_history'];
#美金层面币种折算权益
$frozenBal_history = $data3['frozenBal_history'];
#USDT保证金率
$mgnRatio_history = $data4['mgnRatio_history'];
#合约订单数量
$posdatacount_history = $data5['posdatacount_history'];
#实际未结算盈亏总额
$upl_history = $data6['upl_history'];
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
                text: '实时亏损盈利曲线图'
            },
            tooltip: {
                trigger: 'axis'
            },
            legend: {
                data: ['余额', '折算权益', '保证金金额', '保证金率', '订单数量', '未结算盈亏']
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
                    name: '美金层面币种折算权益',
                    type: 'line',
                    stack: 'Total',
                    data: <?php echo json_encode($disEq_history);?>
                },
                {
                    name: 'USDT保证金金额',
                    type: 'line',
                    stack: 'Total',
                    data: <?php echo json_encode($frozenBal_history);?>
                },
                {
                    name: 'USDT保证金率',
                    type: 'line',
                    stack: 'Total',
                    data: <?php echo json_encode($mgnRatio_history);?>
                },
                {
                    name: '合约订单数量',
                    type: 'line',
                    stack: 'Total',
                    data: <?php echo json_encode($posdatacount_history);?>
                },
                {
                    name: '实际未结算盈亏总额',
                    type: 'line',
                    stack: 'Total',
                    data: <?php echo json_encode($upl_history);?>
                }
            ]
        };

        option && myChart.setOption(option);
    }
</script>
</body>
</html>