ASSUME DS:DATA,SS:STACK,CS:CODE
DATA SEGMENT
      str  DB '00012345','$'
      res  DB '0000','$'
DATA ENDS

STACK SEGMENT
            DB 100 DUP (0)
STACK ENDS

CODE SEGMENT                        ;功能：string转int(16位)
      start:
            mov  ax,data
            mov  ds,ax
            mov  ax,stack
            mov  ss,ax
            mov  si,4

            mov  bx,0
            mov  cx,8
            mov  ax,0
      for1: 
            mov  dx,ax
            shl  dx,1               ;左移一位
            shl  ax,1
            shl  ax,1
            shl  ax,1               ;左移3位
            add  ax,dx              ;相当于ax * 10

            add  al,str[bx]         ;字加字节
            adc  ah,0               ;字加字节
            sub  ax,30H             ;减去'0'

            inc  bx
            loop for1               ;ax = 3039

            mov  cx,4
      l:    
            mov  dx,ax
            and  dx,0FH             ;与000F得到个位
            add  dx,30H             ;如果是0-9加30h
            cmp  dx,3AH
            jb   s1
            add  dx,7H              ;如果是A-F加37h
            
      s1:   
            dec  si
            mov  ds:res[si],dl      ;付给res
            shr  ax,1               ;ax右移4次
            shr  ax,1
            shr  ax,1
            shr  ax,1
            loop l

            lea  dx,res             ;显示16位到屏幕 ’3039H‘
            mov  ah,9
            int  21H

            mov  ax,4c00H
            int  21H

CODE ENDS
end start