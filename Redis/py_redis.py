import redis

# connect to redis
client = redis.Redis(host='192.168.1.187', port=7001)

# set a key
client.set('test-key', 'test-val')

# get a value
value = client.get('test-key')
print(value)