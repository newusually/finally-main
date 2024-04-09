let currentPage = 1; // 默认为第一页

function fetchPage(page) {
    currentPage = page;
    updatetable();
}


function updatetable() {
    fetch(`static/php/get_buy_log_data.php`) // 请求所有的数据
        .then(response => response.json())
        .then(data => {
            let rows = '';
            data.buy_log_data.forEach(item => {
                rows += `<tr>
                  <td>${item.column1}</td>
                  <td>${item.column2}</td>
                  <td>${item.column3}</td>
                  <td>${item.column4}</td>
                  <td>${item.column5}</td>
                  <td>${item.column6}</td>
                  <td>${item.column7}</td>
                  <td>${item.column8}</td>
                </tr>`;
            });
            document.getElementById('buyLogTable').tBodies.item(0).innerHTML = rows;

            // 更新分页信息
            let totalPages = Math.ceil(data.lines / data.pageSize);
            let pagination = '';
            pagination += `<div class="page-item"><a href="#" onclick="fetchPage(1)">首页</a></div>`;
            for (let i = 2; i <= 6; i++) {
                pagination += `<div class="page-item ${i === data.page ? 'active' : ''}"><a href="#" onclick="fetchPage(${i})">${i}</a></div>`;
            }
            pagination += `<div class="page-item"><a href="#" onclick="fetchPage(${data.page + 1})">下一页 &raquo;</a></div>`;
            pagination += `<div class="page-item"><a href="#" onclick="fetchPage(${totalPages})">最后一页</a></div>`;
            document.getElementById('pagination').innerHTML = pagination;
        })
        .catch(error => console.error('Error fetching data: ', error));
}


function update() {

    fetch('static/php/get_cash_bal.php')
        .then(response => response.arrayBuffer())
        .then(buffer => {
            let decoder = new TextDecoder('utf-8');
            let data = decoder.decode(buffer);
            document.getElementById('cashbal').textContent = data;
        })
        .catch(error => console.error('Error fetching data: ', error));


    fetch('static/php/get_upl_ratio_data.php')
        .then(response => response.json())
        .then(data => {
            let rows = '';
            data.forEach(item => {
                rows += `<tr>
                      <td>${item.column1}</td>
                      <td>${item.column2}</td>
                      <td>${item.column3}</td>
                      <td>${item.column4}</td>
                      <td>${item.column5}</td>
                      <td>${item.column6}</td>
                      <td>${item.column7}</td>
                    </tr>`;
            });
            document.getElementById('uplRatioTable').tBodies.item(0).innerHTML = rows;
        })
        .catch(error => console.error('Error fetching data: ', error));


}

