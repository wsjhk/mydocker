package main

/*
void test() {
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/21218/ns/%s", namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system("/bin/sh");
	exit(0);
	return;
}
 */
import "C"

func main()  {
	C.test()
}