const ip = 'http://localhost:8000';

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
        data.push({
            time : Date.now(),
            value : parseFloat(result).toFixed(2)
        });
        if(data.length > 30){
            data.splice(0,1);
        }
        graph();
        console.log(data);
        document.getElementById("uso").innerHTML = `% de Uso: ${parseFloat(result).toFixed(2)}`;
    });      
}

getData();