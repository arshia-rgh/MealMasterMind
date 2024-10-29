from concurrent import futures
from typing import Final

import grpc
from grpc_ import meal_pb2_grpc

port: Final = 50051
from .dependencies import get_current_user


class AuthenticationService:
    def IsAuthenticated(self, request, context):
        token = request.token
        user = get_current_user(token)

        return user.id


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    authentication_service = AuthenticationService()
    meal_pb2_grpc.add_AuthenticationServicer_to_server(authentication_service, server)
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    server.wait_for_termination()
