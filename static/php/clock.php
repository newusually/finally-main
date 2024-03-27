<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8"/>

    <!-- 引入刚刚下载的 ECharts 文件 -->

    <script src="../js/echarts.min.js"></script>


</head>
<body>
<!-- 为 ECharts 准备一个定义了宽高的 DOM -->
<div id="main" style="width: 300px;height:300px;"></div>


<script type="text/javascript">


    var chartDom = document.getElementById('main');
    var myChart = echarts.init(chartDom);
    var option;

    option = {
        series: [
            {
                name: 'hour',
                type: 'gauge',
                startAngle: 90,
                endAngle: -270,
                min: 0,
                max: 12,
                splitNumber: 12,
                clockwise: true,
                axisLine: {
                    lineStyle: {
                        width: 15,
                        color: [[1, 'rgba(0,0,0,0.7)']],
                        shadowColor: 'rgba(0, 0, 0, 0.5)',
                        shadowBlur: 15
                    }
                },
                splitLine: {
                    lineStyle: {
                        shadowColor: 'rgba(0, 0, 0, 0.3)',
                        shadowBlur: 3,
                        shadowOffsetX: 1,
                        shadowOffsetY: 2
                    }
                },

                pointer: {
                    icon: 'path://M2.9,0.7L2.9,0.7c1.4,0,2.6,1.2,2.6,2.6v115c0,1.4-1.2,2.6-2.6,2.6l0,0c-1.4,0-2.6-1.2-2.6-2.6V3.3C0.3,1.9,1.4,0.7,2.9,0.7z',
                    width: 12,
                    length: '55%',
                    offsetCenter: [0, '8%'],
                    itemStyle: {
                        color: '#C0911F',
                        shadowColor: 'rgba(0, 0, 0, 0.3)',
                        shadowBlur: 8,
                        shadowOffsetX: 2,
                        shadowOffsetY: 4
                    }
                },
                detail: {
                    show: false
                },
                title: {
                    offsetCenter: [0, '30%']
                },
                data: [
                    {
                        value: 0
                    }
                ]
            },
            {
                name: 'minute',
                type: 'gauge',
                startAngle: 90,
                endAngle: -270,
                min: 0,
                max: 60,
                clockwise: true,
                axisLine: {
                    show: false
                },
                splitLine: {
                    show: false
                },
                axisTick: {
                    show: false
                },
                axisLabel: {
                    show: false
                },
                pointer: {
                    icon: 'path://M2.9,0.7L2.9,0.7c1.4,0,2.6,1.2,2.6,2.6v115c0,1.4-1.2,2.6-2.6,2.6l0,0c-1.4,0-2.6-1.2-2.6-2.6V3.3C0.3,1.9,1.4,0.7,2.9,0.7z',
                    width: 8,
                    length: '70%',
                    offsetCenter: [0, '8%'],
                    itemStyle: {
                        color: '#C0911F',
                        shadowColor: 'rgba(0, 0, 0, 0.3)',
                        shadowBlur: 8,
                        shadowOffsetX: 2,
                        shadowOffsetY: 4
                    }
                },

                detail: {
                    show: false
                },
                title: {
                    offsetCenter: ['0%', '-40%']
                },
                data: [
                    {
                        value: 0
                    }
                ]
            },
            {
                name: 'second',
                type: 'gauge',
                startAngle: 90,
                endAngle: -270,
                min: 0,
                max: 60,
                animationEasingUpdate: 'bounceOut',
                clockwise: true,
                axisLine: {
                    show: false
                },
                splitLine: {
                    show: false
                },
                axisTick: {
                    show: false
                },
                axisLabel: {
                    show: false
                },
                pointer: {
                    icon: 'path://M2.9,0.7L2.9,0.7c1.4,0,2.6,1.2,2.6,2.6v115c0,1.4-1.2,2.6-2.6,2.6l0,0c-1.4,0-2.6-1.2-2.6-2.6V3.3C0.3,1.9,1.4,0.7,2.9,0.7z',
                    width: 4,
                    length: '85%',
                    offsetCenter: [0, '8%'],
                    itemStyle: {
                        color: '#C0911F',
                        shadowColor: 'rgba(0, 0, 0, 0.3)',
                        shadowBlur: 8,
                        shadowOffsetX: 2,
                        shadowOffsetY: 4
                    }
                },

                detail: {
                    show: false
                },
                title: {
                    offsetCenter: ['0%', '-40%']
                },
                data: [
                    {
                        value: 0
                    }
                ]
            }
        ]
    };
    setInterval(function () {
        var date = new Date();
        var second = date.getSeconds();
        var minute = date.getMinutes() + second / 60;
        var hour = (date.getHours() % 12) + minute / 60;
        option.animationDurationUpdate = 300;
        myChart.setOption({
            series: [
                {
                    name: 'hour',
                    animation: hour !== 0,
                    data: [{value: hour}]
                },
                {
                    name: 'minute',
                    animation: minute !== 0,
                    data: [{value: minute}]
                },
                {
                    animation: second !== 0,
                    name: 'second',
                    data: [{value: second}]
                }
            ]
        });
    }, 1000);

    option && myChart.setOption(option);

</script>
</body>
</html>