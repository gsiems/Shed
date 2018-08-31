# Shed

A Simple HEx Dumper

Reads a file, or STDIN, and writes out the hexedecimal representation of each byte read.

## Arguments

All arguments are optional.

 **-a**          Print ASCII letters, numbers, and symbols. For each byte
                 that represents an ASCII letter, number, or symbol (hex
                 values between 0x1e through 0x7e), print the character
                 instead of the hex value.

 **-c** *int*    The number of bytes to dump. The default is to dump the entire file.

 **-e**          Print codes for single character escape sequences. For hex
                 values in the table, print the code for the character
                 instead of the hex value.

                 000 nul  001 soh  002 stx  003 etx  004 eot  005 enq
                 006 ack  007 bel  008 bs   009 ht   00a lf   00b vt
                 00c ff   00d cr   00e so   00f si   010 dle  011 dc1
                 012 dc2  013 dc3  014 dc4  015 nak  016 syn  017 etb
                 018 can  019 em   01a sub  01b esc  01c fs   01d gs
                 01e rs   01f us   07f del

 **-i** *string* The input file to hexdump. If no file is specified then STDIN is read.

 **-o** *string* The output file to write to. If no file is specified then STDOUT is written to.

 **-w** *int*    The number of bytes to output per line. The default is wrap at 16 bytes.
                 If w is less than 1 then all output are written to one line.
