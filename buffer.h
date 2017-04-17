/*
  To reduce the number of writes we do to STDOUT (mainly
  for efficiency and speed) we use a 'dynamic' string buffer
  to hold a screen's state before a single write().

  This will avoid flickering and stuff between write() calls on
  slow terminals/ conenctions.
 */

struct abuf {
  char *b;
  int len;
};

#define ABUF_INIT {NULL, 0}

void abAppend(struct abuf *ab, const char *s, int len);
void abFree(struct abuf *ab);
