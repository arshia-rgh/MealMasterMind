# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import warnings

import grpc
from meal import meal_pb2 as meal_dot_meal__pb2

GRPC_GENERATED_VERSION = "1.67.0"
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower

    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f"The grpc package installed is at version {GRPC_VERSION},"
        + f" but the generated code in meal/meal_pb2_grpc.py depends on"
        + f" grpcio>={GRPC_GENERATED_VERSION}."
        + f" Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}"
        + f" or downgrade your generated code using grpcio-tools<={GRPC_VERSION}."
    )


class AuthenticationStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.IsAuthenticated = channel.unary_unary(
            "/meal.Authentication/IsAuthenticated",
            request_serializer=meal_dot_meal__pb2.AuthReq.SerializeToString,
            response_deserializer=meal_dot_meal__pb2.AuthRes.FromString,
            _registered_method=True,
        )


class AuthenticationServicer(object):
    """Missing associated documentation comment in .proto file."""

    def IsAuthenticated(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_AuthenticationServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "IsAuthenticated": grpc.unary_unary_rpc_method_handler(
            servicer.IsAuthenticated,
            request_deserializer=meal_dot_meal__pb2.AuthReq.FromString,
            response_serializer=meal_dot_meal__pb2.AuthRes.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler("meal.Authentication", rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers("meal.Authentication", rpc_method_handlers)


# This class is part of an EXPERIMENTAL API.
class Authentication(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def IsAuthenticated(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/meal.Authentication/IsAuthenticated",
            meal_dot_meal__pb2.AuthReq.SerializeToString,
            meal_dot_meal__pb2.AuthRes.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True,
        )
