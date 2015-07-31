import grequests
from datetime import datetime, timedelta


NUMBER_REQUESTS = 10000

def exception_handler(request, exception):
    print "Request Failed"

urls = []
x = 0
start_time = datetime.now()
while x < NUMBER_REQUESTS:
    urls.append("http://10.12.1.30:6060")
    x += 1
print "Now requesting!"
rs = (grequests.get(u) for u in urls)
results = grequests.map(rs, size=4)
end_time = datetime.now()
total_time = end_time - start_time
print "Total time: " + str(total_time)
result_time =  total_time.seconds + (total_time.microseconds / 1000000.0)
print "Req/Sec: " + str(NUMBER_REQUESTS/result_time)
print "Done!"
