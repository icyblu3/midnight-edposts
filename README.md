There exists a encrypted flag on the server. In order to decrypt this flag, the host needs to send a 32-byte key. 
If the key is wrong, the server will reply with "failure!"
If the key is correct, the server will reply with "success!" and print out the decrypted flag.

Run the server with ./server 

A sample host is provided and can be run with ./host