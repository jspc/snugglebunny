/*
  snugglebunny

  Copyright 2017 jspc <james@zero-internet.org.uk>

  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"),
  to deal in the Software without restriction, including without limitation
  the rights to use, copy, modify, merge, publish, distribute, sublicense,
  and/or sell copies of the Software, and to permit persons to whom the Software
  is furnished to do so, subject to the following conditions:

  The above copyright notice and this permission notice shall be included in all
  copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
  INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
  PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
  TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
  OR OTHER DEALINGS IN THE SOFTWARE.
 */

#include <unistd.h>
#include "syntax/c.h"
#include "syntax.h"

int editorSyntaxToColor(int hl) {
  switch (hl) {
  case HL_NUMBER:
    return 31;

  case HL_KEYWORD2:
    return 32;

  case HL_KEYWORD1:
    return 33;

  case HL_MATCH:
    return 34;

  case HL_STRING:
    return 35;

  case HL_MLCOMMENT:
  case HL_COMMENT:
    return 36;

  default:
    return 37;
  }
}
