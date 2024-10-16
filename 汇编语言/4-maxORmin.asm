ASSUME DS:DATA,SS:STACK,CS:CODE
DATA SEGMENT
      arr  DB "helloworld"
DATA ENDS

STACK SEGMENT
            DB 10 DUP (0)
STACK ENDS

CODE SEGMENT
      start:
            mov  ax,data
            mov  ds,ax

            mov  bx,0
            mov  cx,11
            mov  ah,arr[bx]
      s:    
            inc  bx
            mov  al,arr[bx]
            cmp  ah,al           ;求最大值
            jna  s1              ;如果ah不小于(not below)al:跳转到s1
            mov  ah,al           ;
      s1:   
            mov  arr[bx],ah
            loop s

      ;mov dx,offset arr
            lea  dx,arr
            mov  ah,9            ;将dx的内容显示到屏幕
            int  21H

            mov  ah,4cH
            int  21H

CODE ENDS
end start