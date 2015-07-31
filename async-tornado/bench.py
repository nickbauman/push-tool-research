from Queue import Queue
import datetime
from tornado import gen
from tornado.httpclient import AsyncHTTPClient
from threading import Thread
from tornado.ioloop import IOLoop

QUEUE_SIZE = 32
NUMBER_REQUESTS = 1000
URL = "http://localhost:8080"

q = Queue(maxsize=QUEUE_SIZE)
http = AsyncHTTPClient()


def request_worker():
    print "starting requests"
    for i in range(NUMBER_REQUESTS):
        q.put(http.fetch(URL))
    print "completed making requests"


def response_worker():
    results_dict = {}
    while q.empty():
        continue
    while not q.empty():
        future_response = q.get()
        while not future_response.done():
            continue
        response = future_response.result()
        status_code = response.code
        if results_dict.has_key(status_code):
            results_dict[status_code] = 1 + results_dict[status_code]
        else:
            results_dict[status_code] = 1
        q.task_done()
    print results_dict

    IOLoop.instance().stop()

if __name__ == '__main__':
    start_time = datetime.datetime.now()
    request_thread = Thread(target=request_worker)
    request_thread.setDaemon(True)
    request_thread.start()

    response_thread = Thread(target=response_worker)
    response_thread.setDaemon(True)
    response_thread.start()

    IOLoop.instance().start()

    end_time = datetime.datetime.now()

    duration = end_time - start_time
    print "Took {}".format(duration)
