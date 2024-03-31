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

            setInterval(function () {
                update();

            }, 60000); // 60000毫秒等于1分钟
        }
    </script>

</head>
<body>
<div class="content">
    <div class="text" style="right:30px  ">

        <h1>我们是守护者，也是一群时刻对抗危险和疯狂的可怜虫！！！" https://github.com/newusually/finally-main</h1>
    </div>

</div>

<br><br>
<h1>实时亏损盈利数据</h1><br><br><br>
<h1 id="cashbal" style="text-align: left; white-space: pre-wrap;"></h1><br>

<h1>实时亏损盈利曲线图</h1><br>

<iframe id="floating-iframe" src="static/php/line_stack.php" style="width: 600px;
height: 400px;"></iframe>


<br>


</body>
</html>