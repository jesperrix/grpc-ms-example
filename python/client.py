""" Python implementation of the gRPC my_model_api client"""

import logging

import grpc

from my_model_api_pb2 import MyModelInput, MyModelOutput
from my_model_api_pb2_grpc import MyModelApiStub

def run():
    # Connect to the gRPC channel and receive a prediction
    with grpc.insecure_channel('localhost:56789') as channel:
        stub = MyModelApiStub(channel)

        print("-------------- Predict --------------")
        for i in range(1002):

            # Generate input
            model_input = MyModelInput(x=[1,2,3,4], y=[5,6,7,8])
            # Send input
            res = stub.Predict(model_input)

            # Print received output
            model_id = res.model_id
            model_out = res.y
            if i % 100 == 0:
                print(f"recieved #{i}")
                print(f"received from model: model_id={model_id}, model_out={model_out}")



if __name__ == '__main__':
    logging.basicConfig()
    run()
