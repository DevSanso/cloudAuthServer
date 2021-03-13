import subprocess
import sys


if sys.argv[1] == "protoc":
    #exist output folder set bug
    subprocess.call(["protoc","--go-grpc_out=./proto", "./api/proto/api.proto"])
    subprocess.call(["protoc","--go_out=./proto", "./api/proto/api.proto"])
elif sys.argv[1] == "server":
    subprocess.call(["go","build"])
else:
    raise Exception("plz input argc protoc, or server")