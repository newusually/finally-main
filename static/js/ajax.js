


function update() {

    fetch('static/php/get_cash_bal.php')
        .then(response => response.arrayBuffer())
        .then(buffer => {
            let decoder = new TextDecoder('utf-8');
            let data = decoder.decode(buffer);
            document.getElementById('cashbal').textContent = data;
        })
        .catch(error => console.error('Error fetching data: ', error));



}

