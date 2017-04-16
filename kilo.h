#define KILO_VERSION "0.0.1"
#define KILO_TAB_STOP 8
#define KILO_QUIT_TIMES 3

#define DEFAULT_STATUS "hello snuggle bunny <3<3<3"

#define E_ERROR 1
#define CTRL_KEY(k) ((k) & 0x1f)


/* prototypes */
void editorSetStatusMessage(const char *fmt, ...);
void editorRefreshScreen();
char *editorPrompt(char *prompt);

/*
   Represent a row of stuff
 */
typedef struct erow {
  int size;
  int rsize;
  char *chars;
  char *render;
} erow;

/*
  Hold some editor state
 */
struct editorConfig {
  int cx, cy;
  int rx;
  int rowoff;
  int coloff;
  int screencols;
  int screenrows;
  int numrows;
  erow *row;
  int dirty;
  char *filename;
  char statusmsg[80];
  time_t statusmsg_time;
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

/*
  Do some arrow key stuff I guess I dunno
 */
enum editorKey {
  BACKSPACE = 127,
  ARROW_LEFT = 1000,
  ARROW_RIGHT,
  ARROW_UP,
  ARROW_DOWN,
  DEL_KEY,
  HOME_KEY,
  END_KEY,
  PAGE_UP,
  PAGE_DOWN
};
