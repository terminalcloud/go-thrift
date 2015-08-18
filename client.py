#!/usr/bin/python

import sys, glob
sys.path.append('gen-py')

from example import Service
from example.ttypes import *

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

try:
    # Make socket
    transport = TSocket.TSocket('localhost', 9090)

    # Buffering is critical. Raw sockets are very slow
    transport = TTransport.TBufferedTransport(transport)

    # Wrap in a protocol
    protocol = TBinaryProtocol.TBinaryProtocol(transport)

    # Create a client to use the protocol encoder
    client = Service.Client(protocol)

    # Connect
    transport.open()

    # Test out operations
    f = flop()
    f.a = 1
    f.b = 2

    client.ping()
    print client.count()
    print client.count()
    print client.count()
    print client.echo("Hello, world!")
    print client.flip(f)
    client.fail()

    # Close
    transport.close()

except Thrift.TException, tx:
    print 'Failed: %s' % (tx.message)
