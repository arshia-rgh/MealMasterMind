from concurrent import futures

import grpc
from meal import meal_pb2_grpc
from meal.service import AuthenticationService


# TODO: Make this class more scalable
class GRPC:
    def __init__(self, host=None):
        self.host = host
        self.port = 50051

    def serve(self):
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        meal_pb2_grpc.add_AuthenticationServicer_to_server(AuthenticationService(), server)
        server.add_insecure_port(f"[::]:{self.port}")
        server.start()
        server.wait_for_termination()
