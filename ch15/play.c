// Compile and disassemble:
// gcc -c func.c
// otool -xv func.o
//
// Explaining the func's assembly:
// https://stackoverflow.com/questions/55711485/if-c-file-only-has-one-function-why-is-the-pushq-and-movq-still-exist-at-the-be
void func() {
    unsigned int x = 4294967295; // 0xffff_ffff;
    unsigned int y = 0xffffffff;
    // x = x + 3;
}
