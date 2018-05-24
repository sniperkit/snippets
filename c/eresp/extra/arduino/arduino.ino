#include <SPI.h>
#include <Ethernet.h>
#include <eRESP.h>

// Enter a MAC address and IP address for your controller below.
// The IP address will be dependent on your local network.
// gateway and subnet are optional:
byte mac[] = {
	0x00, 0xAA, 0xBB, 0xCC, 0xDE, 0x02
};
IPAddress ip(192, 168, 1, 240);

// telnet defaults to port 23
EthernetServer server(6379);
eresp_ctx_t erespServer;
char erespServerBuffer[256];

static void pingCommand(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)ctx;

	Serial.println("PING");

	if (argc > 0) {
		Serial.print(argc);
		Serial.print(argv[0].data);
	}

	eresp_write_string(ctx, "PONG");
}

static void disconnectCommand(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)ctx;
	Serial.println("DISCONNECT");
	EthernetClient client = server.available();
	client.stop();
}

static const struct eresp_cmd_t erespServerCommandTable[] = {
	ERESP_CMD("PING",pingCommand,0),
	ERESP_CMD("DISCONNECT",disconnectCommand,0),
	ERESP_CMD_LAST
};

void erespWriterCallback(const eresp_ctx_t *ctx, const char c)
{
	(void)ctx;
	server.write(c);
}

void setup() {
	Serial.begin(9600);

	Serial.println("Trying to get an IP address using DHCP");
	if (Ethernet.begin(mac) == 0) {
		Serial.println("Failed to configure Ethernet using DHCP");
		exit(1);
	}

	Serial.print("My IP address: ");
	ip = Ethernet.localIP();
	for (byte thisByte = 0; thisByte < 4; thisByte++) {
		Serial.print(ip[thisByte], DEC);
		Serial.print(".");
	}
	Serial.println();

	eresp_init(&erespServer, erespServerBuffer, sizeof(erespServerBuffer));
	eresp_set_commands(&erespServer, erespServerCommandTable);
	eresp_set_writer(&erespServer, erespWriterCallback);

	server.begin();
}

void loop() {
	EthernetClient client = server.available();

	if (!client)
		return;

	eresp_read_char(&erespServer, client.read());
	Ethernet.maintain();
}
