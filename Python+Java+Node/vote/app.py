from flask import Flask, render_template, request, make_response, g
from redis import Redis
import os
import socket
import random
import json

# Load environment variables
option_a = os.getenv('OPTION_A', "Emacs")
option_b = os.getenv('OPTION_B', "Vi")
hostname = socket.gethostname()
version = 'v1'

app = Flask(__name__)

# Get Redis connection
def get_redis():
    if not hasattr(g, 'redis'):
        g.redis = Redis(
            host=os.getenv('REDIS_HOST', 'redis'),  # Ensure Redis host is 'redis'
            port=os.getenv('REDIS_PORT', 6379),    # Default Redis port
            db=0,
            socket_timeout=5
        )
    return g.redis

@app.route("/", methods=['POST', 'GET'])
def hello():
    # Get the voter ID from cookies, or create a new one
    voter_id = request.cookies.get('voter_id')
    if not voter_id:
        voter_id = hex(random.getrandbits(64))[2:-1]

    vote = None

    if request.method == 'POST':
        redis = get_redis()
        vote = request.form['vote']
        
        # Increment vote count in Redis based on the selected option
        if vote == 'a':
            redis.incr('option_a_count')  # Increment for option A
        elif vote == 'b':
            redis.incr('option_b_count')  # Increment for option B

        # Store vote data (voter_id and vote) into the 'votes' list in Redis
        data = json.dumps({'voter_id': voter_id, 'vote': vote})
        redis.rpush('votes', data)

    # Get the current vote counts from Redis
    redis = get_redis()
    option_a_count = int(redis.get('option_a_count') or 0)
    option_b_count = int(redis.get('option_b_count') or 0)

    # Render the HTML template with vote counts and other data
    resp = make_response(render_template(
        'index.html',
        option_a=option_a,
        option_b=option_b,
        hostname=hostname,
        vote=vote,
        option_a_count=option_a_count,
        option_b_count=option_b_count,
        version=version,
    ))
    
    # Set the voter_id cookie for future sessions
    resp.set_cookie('voter_id', voter_id)
    return resp

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80, debug=True, threaded=True)
