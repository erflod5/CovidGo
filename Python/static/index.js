const ip = 'http://localhost:5000';

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
        var actual = result.split('-');
        for(let i = 0; i < actual.length - 1; i++){
            if(data[i] == undefined) data[i] = {};

            data[i].y = i;
            data[i].a = actual[i];
            data[i].b = 10;
        }
        graph();
        document.getElementById("ram").innerHTML = `% de Uso RAM: ${actual[0]}`;
    });      
}

getData();