# xtensible-message-sender
This project aims to create byte message from a yaml structured message. This message content can be sent over a socket or can be printed to screen.

<pre>
input yaml file
  encoding: big-endian
 
  uint8: 24
  uint16: 5192
  double: 3.7
  float: 0.4
  
output
  181448400D99999999999A3ECCCCCD
</pre>
