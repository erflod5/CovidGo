const ip = 'http://localhost:5001';

var data =[];

var intervalID = setInterval(
    function(){
        getData();
    },5000
);

function graph(){
    document.getElementById('myfirstchart').innerHTML = "";
    Morris.Bar({
        barGap:2,
        barSizeRatio:0.55,
        element: 'myfirstchart',
        data: data,
        xkey: 'y',
        ykeys: ['a', 'b'],
        labels: ['A', 'B'],
        barColors: ['#0B62A4','#f75b68'],
        hideHover: 'auto'
      });
}

function getData(){
    $.get(`${ip}/ram`, function (result) {
        var actual = result.split("b").join("");
        actual = actual.split("'").join("");
        actual = actual.split('-');
        for(let i = 0; i < actual.length - 1; i++){
            if(data[i] == undefined) data[i] = {};

            data[i].y = i;
            data[i].a = actual[i];
        }
        graph();
        document.getElementById("ram").innerHTML = `% de Uso RAM: ${actual[0]}`;
    });      

    $.get(`${ip}/cpu`, function (result) {
        var actual = result.split("b").join("");
        actual = actual.split("'").join("");
        actual = actual.split('-');
        for(let i = 0; i < actual.length - 1; i++){
            if(data[i] == undefined) data[i] = {};
            data[i].y = i;
            data[i].b = parseFloat(actual[i]).toFixed(2)
        }
        graph();
        document.getElementById("cpu").innerHTML = `% de Uso CPU: ${parseFloat(actual[0]).toFixed(2)}`;
    });
}

getData();