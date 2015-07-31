import tornado.httpserver
import tornado.ioloop
import tornado.web

class MainHandler(tornado.web.RequestHandler):
  def get(self):
    self.write("Hello world!")

app = tornado.web.Application([(r"/", MainHandler),])

if __name__ == "__main__":
	server = tornado.httpserver.HTTPServer(app)
	server.bind(8080)
	server.start(0)  # autodetect number of cores and fork a process for each
	tornado.ioloop.IOLoop.instance().start()
