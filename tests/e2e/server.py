import BaseHTTPServer
import os

PORT_NUMBER = 8080

if "CUSTOM_SERVER_MESSAGE_FILE" not in os.environ:
    raise ValueError("Missing required environment variable 'CUSTOM_SERVER_MESSAGE_FILE'")

with open(os.environ["CUSTOM_SERVER_MESSAGE_FILE"], "r") as f:
    SERVER_MESSAGE = f.read()


class MyHandler(BaseHTTPServer.BaseHTTPRequestHandler):
    def do_GET(s):
        print("Received HTTP request")
        s.send_response(200)
        s.send_header("Content-type", "text/plain")
        s.end_headers()
        s.wfile.write(SERVER_MESSAGE)


if __name__ == '__main__':
    server_class = BaseHTTPServer.HTTPServer
    print("Starting HTTP server on port %d" % PORT_NUMBER)
    httpd = server_class(("0.0.0.0", PORT_NUMBER), MyHandler)
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        pass
    httpd.server_close()
