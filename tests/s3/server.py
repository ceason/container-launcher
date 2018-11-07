import BaseHTTPServer
import argparse

PORT_NUMBER = 8080

parser = argparse.ArgumentParser()
parser.add_argument(
    '--custom_server_message_file', action='store', required=True,
    help='The content of this file will be provided to HTTP GET requests.')

args = parser.parse_args()

with open(args.custom_server_message_file, "r") as f:
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
