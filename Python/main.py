from flask import Flask
from flask import render_template
import redis

# connect to redis
client = redis.Redis(host='3.14.52.42', port=7001)

# creates a Flask application, named app
app = Flask(__name__)

# a route where we will display a welcome message via an HTML template
@app.route("/")
def hello():
    message = "The Flask Shop"
    return render_template('index.html', message=message)

@app.route("/rust")
def rust():
    message = "The Flask Shop"
    return render_template('rust.html', message=message)

@app.route("/ram")
def ram():    
    value = client.lrange("ram",0,10)
    print(value)
    return ''.join(str(x) + "-" for x in value)

@app.route("/cpu")
def cpu():    
    value = client.lrange("cpu",0,10)
    print(value)
    return ''.join(str(x) + "-" for x in value)

app.run(host ='0.0.0.0', port = 5001, debug = False)
