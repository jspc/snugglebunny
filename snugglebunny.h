#define SNUGGLEBUNNY_VERSION "0.0.1"
#define SNUGGLEBUNNY_TAB_STOP 8
#define SNUGGLEBUNNY_QUIT_TIMES 3

#define DEFAULT_STATUS "hello snuggle bunny <3<3<3"


#define CTRL_KEY(k) ((k) & 0x1f)

/* prototypes */
void editorSetStatusMessage(const char *fmt, ...);
void editorRefreshScreen();
char *editorPrompt(char *prompt, void (*callback)(char *, int));

/*
   Represent a row of stuff
 */
typedef struct erow {
  int idx;
  int size;
  int rsize;
  char *chars;
  char *render;
  unsigned char *hl;
  int hl_open_comment;
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
  struct editorSyntax *syntax;
  struct termios orig_termios;
};
