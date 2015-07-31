from tornado.httpclient import AsyncHTTPClient
import tornado
from datetime import datetime

URL = "http://localhost:8080"
REQUESTS = 10000

count = 0
error = 0


def handle_request(response):
    global count
    global error
    if response.error:
        error += 1
    else:
        count += 1
    if error + count == REQUESTS:
        end_time = datetime.now()
        total_time_delta = end_time - start_time
        result_time =  total_time_delta.seconds + (total_time_delta.microseconds / 1000000.0)
        print "count = " + str(count)
        print "error = " + str(error)
        print "total_time_delta = " + str(result_time)
        print "req/sec = " + str(REQUESTS/result_time)
        tornado.ioloop.IOLoop.instance().stop()


if __name__ == "__main__":
    global start_time
    http_client = AsyncHTTPClient()
    start_time = datetime.now()
    x = 0
    while x < REQUESTS:
        http_client.fetch(URL, handle_request)
        x += 1
    print "Now starting IOLoop"
    tornado.ioloop.IOLoop.instance().start()
