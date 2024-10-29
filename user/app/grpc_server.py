from concurrent import futures
from typing import Final

import grpc
from grpc_ import user_pb2, user_pb2_grpc

port: Final = 50051
from .dependencies import get_current_user


class AuthenticationService(user_pb2_grpc.AuthenticationServicer):
    def IsAuthenticated(self, request, context):
        token = request.token
        try:
            user = get_current_user(token)
            return user_pb2.AuthRes(userID=user.id)
        except Exception as e:
            return user_pb2.AuthRes()


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    print("starting server")
    authentication_service = AuthenticationService()
    user_pb2_grpc.add_AuthenticationServicer_to_server(authentication_service, server)
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    server.wait_for_termination()
