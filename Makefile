MAKE=make

all:
	+$(MAKE) -C udpproto_client/client
	+$(MAKE) -C udpproto_server/server

clean:
	rm -f udpproto_client/client/client
	rm -f udpproto_server/server/server

