""" Python implementation of the gRPC my_model_api server"""

import logging
from concurrent import futures

import grpc

import my_model_api_pb2
import my_model_api_pb2_grpc

class MyModelApiServicer(my_model_api_pb2_grpc.MyModelApiServicer):
    """ provides methods that implement functionality of My Model API server. """
    c = 0

    def Predict(self, request, context):
        """Predict function """
        self.c += 1
        return my_model_api_pb2.MyModelOutput(model_id="model_xyz", y=[1,2,3,self.c])


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    my_model_api_pb2_grpc.add_MyModelApiServicer_to_server(
        MyModelApiServicer(), server)
    server.add_insecure_port('[::]:56789')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()
