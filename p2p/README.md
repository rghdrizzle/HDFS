## p2p package

The peer interface is a representation of a remote node. Lets say you are dialing to someone through the phone , the someone you are dialing to is the peer interface.

The transport interface handles the communication between multiple nodes in the network , it can be in the form of TCP , UDP , websockets etc.

The Tcp_transport has various functions such as listenAndAccept ( which btw is a function of the Transport interface), handleConn , acceptLoop etc. Read the comments written in the code to understand what each function or data type does.

The encoding module is responsible for decoding the encoded message. For now it decodes a Gob encoded message and the default message (normal strings/bytes). 

The handshake module is responsible for handling handshake process, for now there isnt a proper handshake process.

The Message data type represents any data that is sent over each transport between two nodes in the network