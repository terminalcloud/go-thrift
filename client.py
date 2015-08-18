#!/usr/bin/python

import sys, glob
sys.path.append('gen-py')

from example import Service
from example import OtherService
from example.ttypes import *

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

try:
    # Make socket
    serviceTransport = TSocket.TSocket('localhost', 9090)
    otherTransport = TSocket.TSocket('localhost', 9091)

    # Buffering is critical. Raw sockets are very slow
    serviceTransport = TTransport.TBufferedTransport(serviceTransport)
    otherTransport = TTransport.TBufferedTransport(otherTransport)

    # Wrap in a protocol
    serviceProtocol = TBinaryProtocol.TBinaryProtocol(serviceTransport)
    otherProtocol = TBinaryProtocol.TBinaryProtocol(otherTransport)

    # Create a client to use the protocol encoder
    serviceClient = Service.Client(serviceProtocol)
    otherClient = OtherService.Client(otherProtocol)

    # Connect
    serviceTransport.open()
    otherTransport.open()

    # Test out operations
    f = flop()
    f.a = 1
    f.b = 2

    serviceClient.ping()
    print serviceClient.count()
    print serviceClient.count()
    print serviceClient.count()
    print serviceClient.echo("Hello, world!")
    print serviceClient.flip(f)

    otherClient.noop()

    serviceClient.fail()

    # Close
    serviceTransport.close()
    otherTransport.close()

except Thrift.TException, tx:
    print 'Failed: %s' % (tx.message)
