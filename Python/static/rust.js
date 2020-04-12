const ip = 'http://3.14.52.42:8000';

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
    $.get(`${ip}/api/cpu`, function (result) {
        console.log(result);
        var data = result.split('"');
        data.push({
            time : Date.now(),
            value : parseFloat(data[1]).toFixed(2)
        });
        if(data.length > 30){
            data.splice(0,1);
        }
        graph();
        console.log(data);
        document.getElementById("uso").innerHTML = `% de Uso: ${parseFloat(data[1]).toFixed(2)}`;
    });      
}

getData();
