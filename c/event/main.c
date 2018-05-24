#include <stdio.h>
#include <stdint.h>
#include <unistd.h>
#include <pthread.h>

enum events {
	NOEVENT,
	EVFIN,
	EV1,
	EV2,
	EVPOSTERSTOP
};

enum events e = NOEVENT;
void *eargs = NULL;
pthread_mutex_t emtx;
pthread_cond_t econd = PTHREAD_COND_INITIALIZER;

void evpost(enum events ev, void *args)
{
	pthread_mutex_lock(&emtx);
	while (e != NOEVENT)
		pthread_cond_wait(&econd, &emtx);
	eargs = args;
	e = ev;
	pthread_cond_signal(&econd);
	pthread_mutex_unlock(&emtx);
}

void ev1(uint8_t *arg)
{
	(void)arg;
	printf("ev1\n");
}

static unsigned int evposter_running = 1;

void evposterstop(void)
{
	evposter_running = 0;
}

void ev2(const char *s)
{
	printf("ev2 %s\n", s);
}

struct evitem {
	enum events id;
	void (*handler)(void *arg);
};
static const struct evitem evitems[] = {
	{
		.id = EV1,
		.handler = (void *)(void *)ev1
	},
	{
		.id = EV2,
		.handler = (void *)(void *)ev2
	},
	{
		.id = EVPOSTERSTOP,
		.handler = (void *)(void *)evposterstop
	}

};

void *evhandler(void *args)
{
	(void)args;

	while (1) {
		pthread_mutex_lock(&emtx);
		while (e == NOEVENT)
			pthread_cond_wait(&econd, &emtx);

		enum events ev = e;
		void *args = eargs;
		
		e = NOEVENT;
		eargs = NULL;

		pthread_cond_signal(&econd);
		pthread_mutex_unlock(&emtx);

		if (ev == EVFIN)
			break;

		for (size_t n = 0; n < sizeof(evitems)/sizeof(evitems[0]); n++) {
			if (evitems[n].id == ev) {
				evitems[n].handler(args);
				break;
			}
		}
	}

	pthread_exit(NULL);
}

void *evposter(void *args)
{
	(void)args;

	uint16_t i = 0;

	evpost(EV2, "evposter start");
	while (evposter_running) {
		evpost(EV2, "evposter");
		usleep(100 * 1000);
		i++;
	}

	sleep(1);
	evpost(EV2, "evposter finish");
	pthread_exit(NULL);
}

int main(void)
{
	pthread_t thd[3];

	pthread_mutex_init(&emtx, NULL);
	pthread_create(&thd[0], NULL, evhandler, NULL);
	pthread_create(&thd[1], NULL, evposter, NULL);

	uint16_t i = 0;

	evpost(EV2, "main start");
	while (1) {
		evpost(EV2, "main");
		usleep(100 * 1000);

		i++;
		if (i == 10)
			break;
	}
	evpost(EV2, "main finish");

	evpost(EVPOSTERSTOP, NULL);
	pthread_join(thd[1], NULL);
	evpost(EVFIN, NULL);
	pthread_join(thd[0], NULL);

	return 0;
}
