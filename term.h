int getCursorPosition(int *rows, int *cols);
int getWindowSize(int *rows, int *cols);

/*
  Some unsutainably large numbers to help determine
  the maximum height and width of a terminal window.

  These numbers want replacing/ increasing when they
  become too low (on machines where ioctl can't do this
  for us: which seems super low).
 */
#define MASSIVE_WIDTH_HEIGHT "\x1b[9999C\x1b[9999B"
#define MASSIVE_WIDTH_HEIGHT_SIZE 14
