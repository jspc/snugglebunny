#define SNUGGLEBUNNY_VERSION "0.0.1"
#define SNUGGLEBUNNY_TAB_STOP 8
#define SNUGGLEBUNNY_QUIT_TIMES 3

#define DEFAULT_STATUS "hello snuggle bunny <3<3<3"

#define E_ERROR 1
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

/*
  Get some syntax highlighting going
*/

#define HL_HIGHLIGHT_NUMBERS (1<<0)
#define HL_HIGHLIGHT_STRINGS (1<<1)

enum editorHighlight {
  HL_NORMAL = 0,
  HL_STRING,
  HL_COMMENT,
  HL_MLCOMMENT,
  HL_KEYWORD1,
  HL_KEYWORD2,
  HL_NUMBER,
  HL_MATCH
};

struct editorSyntax {
  char *filetype;
  char **filematch;
  char **keywords;
  char *singleline_comment_start;
  char *multiline_comment_start;
  char *multiline_comment_end;
  int flags;
};

char *C_HL_extensions[] = { ".c", ".h", ".cpp", NULL };
char *C_HL_keywords[] = {
  "switch", "if", "while", "for", "break", "continue", "return", "else",
  "struct", "union", "typedef", "static", "enum", "class", "case",
  "int|", "long|", "double|", "float|", "char|", "unsigned|", "signed|",
  "void|", NULL
};
struct editorSyntax HLDB[] = {
  {
    "c",
    C_HL_extensions,
    C_HL_keywords,
    "//", "/*", "*/",
    HL_HIGHLIGHT_NUMBERS | HL_HIGHLIGHT_STRINGS
  },
};
#define HLDB_ENTRIES (sizeof(HLDB) / sizeof(HLDB[0]))
