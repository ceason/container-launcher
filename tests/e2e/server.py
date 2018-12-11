import BaseHTTPServer
import os

PORT_NUMBER = 8080

SERVER_MESSAGE = None

if "CUSTOM_SERVER_MESSAGE_FILE" in os.environ and "CUSTOM_SERVER_MESSAGE" in os.environ:
    raise ValueError("May only specify one of 'CUSTOM_SERVER_MESSAGE_FILE' or 'CUSTOM_SERVER_MESSAGE'")
elif "CUSTOM_SERVER_MESSAGE_FILE" in os.environ:
    with open(os.environ["CUSTOM_SERVER_MESSAGE_FILE"], "r") as f:
        SERVER_MESSAGE = f.read()
elif "CUSTOM_SERVER_MESSAGE" in os.environ:
    SERVER_MESSAGE = os.environ["CUSTOM_SERVER_MESSAGE"]
else:
    raise ValueError("Must specify at least one of 'CUSTOM_SERVER_MESSAGE_FILE' or 'CUSTOM_SERVER_MESSAGE'")


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
