from concurrent import futures

import grpc

from mailer.mails.mails_pb2_grpc import add_SendMailServicer_to_server
from mailer.send_mail import SendMail


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    add_SendMailServicer_to_server(SendMail(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
