#define E_ERROR 1
#define CTRL_KEY(k) ((k) & 0x1f)

/*
  Hold some editor state
 */
struct editorConfig {
  int screencols;
  int screenrows;
  struct termios orig_termios;
};

/*
  Some unsutainably large numbers to help determine
  the maximum height and width of a terminal window.

  These numbers want replacing/ increasing when they
  become too low (on machines where ioctl can't do this
  for us: which seems super low).
*/
#define MASSIVE_WIDTH_HEIGHT "\x1b[9999C\x1b[9999B"
#define MASSIVE_WIDTH_HEIGHT_SIZE 14


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
