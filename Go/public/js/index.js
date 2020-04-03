const ip = 'http://localhost:8000/';

var data =[];

var intervalID = setInterval(
    function(){
        getData();
    },5000
);

function graph(){
    document.getElementById('myfirstchart').innerHTML = "";
    new Morris.Line({
        element: 'myfirstchart',
        data: data,
        xkey: 'time',
        ykeys: ['value'],
        labels: ['Value']
      });
}

function getData(){
    $.get(`${ip}`, function (result) {
        var actual = JSON.parse(result);
        data.push({
            time : Date.now(),
            value : (actual.used/actual.total * 100).toFixed(2)
        });
        if(data.length > 30){
            data.splice(0,1);
        }
        graph();
        console.log(data);
        document.getElementById("uso").innerHTML = `% de Uso: ${(actual.used/actual.total * 100).toFixed(2)}`;
        document.getElementById("usoMb").innerHTML = `Uso(mb): ${actual.used}`;
        document.getElementById("totalMb").innerHTML = `Total(mb): ${actual.total}`;
    });      
}

getData();