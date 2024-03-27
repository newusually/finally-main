<!DOCTYPE html>
<html lang="ch">
<head>
    <meta charset="utf-8">
    <title>收益率和买入合约的条件</title>
    <link href="static/css/style.css" rel="stylesheet">
    <script src="static/js/ajax.js"></script>
    <script>
        window.onload = function () {
            update();
            fetchPage(1);
            setInterval(function () {
                update();
                fetchPage(1);
            }, 60000); // 60000毫秒等于1分钟
        }
    </script>

</head>
<body>
<div class="content">
    <div class="text">
        <h1>公开云代码网址：
            https://github.com/newusually/heyiran</h1><br><br>
        <h1>我们是守护者，也是一群时刻对抗危险和疯狂的可怜虫！！！"</h1>

        <iframe id="floating-iframe" src="static/php/clock.php" style="width: 300px;
    height: 300px;"></iframe>

    </div>
    <div class="image">
        <iframe id="floating-iframe" src="static/php/eth_k.php" style="width: 600px;
    height: 400px;"></iframe>
    </div>
</div>

<br><br>


<h1>实时亏损盈利数据</h1><br>
<h1 id="cashbal" style="text-align: left; white-space: pre-wrap;"></h1><br>
<h1>实时亏损盈利曲线图</h1><br>
<div class="image">
    <iframe id="floating-iframe" src="static/php/line_stack.php" style="width: 600px;
    height: 400px;"></iframe>
</div>
<h1>盈利数据表</h1><br>

<div class="scrollable-table">
    <table class="colorful" id="uplRatioTable">

        <thead>
        <tr>
            <th>合约名称</th>
            <th>未实现收益率</th>
            <th>现价</th>
            <th>开仓均价</th>
            <th>保证金</th>
            <th>杠杆倍数</th>
            <th>总计亏损金额</th>
        </tr>
        </thead>
        <tbody>
        <!-- 这里将通过AJAX填充数据 -->
        </tbody>

    </table>
</div>


<br>

<h1>买入合约的条件</h1><br>

<div class="scrollable-table">
    <table class="colorful" id="buyLogTable">
        <thead>
        <tr>
            <th>买入时间</th>
            <th>买入的合约名称</th>
            <th>当前收盘价/开盘价的比例</th>
            <th>前一个交易时间段的收盘价/开盘价的比例</th>
            <th>当前最高价/收盘价的比例</th>
            <th>前一个交易时间段的最高价/收盘价的比例</th>
            <th>当前的macd值</th>
            <th>当前的macd值/前一个交易时间段的macd值的比例</th>
            <th>diff值</th>
            <th>dea值</th>
            <th>买入交易的分钟数(5分钟,15分钟)</th>
        </tr>
        </thead>
        <tbody>
        <!-- 这里将通过AJAX填充数据 -->
        </tbody>

    </table>
</div>
<div class="pagination" id="pagination">

</div>


</body>
</html>